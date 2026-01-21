package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rseigha/goecomapi/internal/config"
	"github.com/rseigha/goecomapi/internal/database"
	"github.com/rseigha/goecomapi/internal/handler"
	"github.com/rseigha/goecomapi/internal/repository"
	"github.com/rseigha/goecomapi/internal/routes"
	"github.com/rseigha/goecomapi/internal/service"
	jwtpkg "github.com/rseigha/goecomapi/pkg/jwt"
	"go.uber.org/zap"
)

func main() {
	// Logger
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	// load config from environment
	cfg, err := config.Load()

	if err != nil {
		logger.Fatal("failed to load config", zap.Error(err))
	}


	ctx := context.Background()
	mongoDB, err := database.NewMongo(ctx, cfg.MongoURI, cfg.MongoDBName, logger)
	if err != nil {
		logger.Fatal("failed to connect to MongoDB", zap.Error(err))
	}
	defer mongoDB.Close(ctx, logger)

	// Repositories
	userRepo := repository.NewUserRepository(mongoDB, logger)
	productRepo := repository.NewProductRepository(mongoDB, logger)
	orderRepo := repository.NewOrderRepository(mongoDB, logger)

	// JWT
	jwt := jwtpkg.NewJWT(cfg.JWTSecret, time.Duration(cfg.JWTExpiryMinutes)*time.Minute)

	// Services
	authSvc := service.NewAuthService(userRepo, jwt, logger)
	userSvc := service.NewUserService(userRepo, logger)
	productSvc := service.NewProductService(productRepo)
	orderSvc := service.NewOrderService(orderRepo)

	// Handlers
	authHandler := handler.NewAuthHandler(authSvc, logger)
	userHandler := handler.NewUserHandler(userSvc)
	productHandler := handler.NewProductHandler(productSvc)
	orderHandler := handler.NewOrderHandler(orderSvc)


	// Router
	router := routes.NewRouter(&routes.RouterConfig{
		AuthHandler:    authHandler,
		UserHandler:    userHandler,
		ProductHandler: productHandler,
		OrderHandler:   orderHandler,
		JWT:            jwt,
		Logger:         logger,
	})


	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}


	// Run server in goroutine
	go func() {
		logger.Info("starting server", zap.Int("port", cfg.Port))
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Fatal("server error", zap.Error(err))
		}
	}()


		// Wait for interrupt to gracefully shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	logger.Info("shutting down server")

	ctxShutdown, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctxShutdown); err != nil {
		logger.Fatal("server forced to shutdown", zap.Error(err))
	}
	logger.Info("server exiting")
}