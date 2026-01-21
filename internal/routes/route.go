package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rseigha/goecomapi/internal/handler"
	"github.com/rseigha/goecomapi/internal/middleware"
	jwtpkg "github.com/rseigha/goecomapi/pkg/jwt"
	"go.uber.org/zap"
)

type RouterConfig struct {
	AuthHandler    *handler.AuthHandler
	UserHandler    *handler.UserHandler
	ProductHandler *handler.ProductHandler
	OrderHandler   *handler.OrderHandler
	JWT            *jwtpkg.JWT
	Logger         *zap.Logger
}

func NewRouter(cfg *RouterConfig) *mux.Router {
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()

	// public
	api.HandleFunc("/auth/register", cfg.AuthHandler.Register).Methods("POST")
	api.HandleFunc("/auth/login", cfg.AuthHandler.Login).Methods("POST")

	// products: list and get are public
	api.HandleFunc("/products", cfg.ProductHandler.List).Methods("GET")
	api.HandleFunc("/products/{id}", cfg.ProductHandler.Get).Methods("GET")

	// protected routes
	authMiddleware := middleware.JWTAuth(cfg.JWT, cfg.Logger)

	userRouter := api.PathPrefix("/users").Subrouter()
	userRouter.HandleFunc("/{id}", cfg.UserHandler.GetByID).Methods("GET")
	userRouter.HandleFunc("/{id}", cfg.UserHandler.Update).Methods("PUT")
	userRouter.HandleFunc("/{id}", cfg.UserHandler.Delete).Methods("DELETE")
	userRouter.Use(authMiddleware)

	orderRouter := api.PathPrefix("/orders").Subrouter()
	orderRouter.HandleFunc("", cfg.OrderHandler.Create).Methods("POST")
	orderRouter.HandleFunc("", cfg.OrderHandler.ListByUser).Methods("GET")
	orderRouter.Use(authMiddleware)

	// admin product routes
	adminRouter := api.PathPrefix("/products").Subrouter()
	adminRouter.HandleFunc("", cfg.ProductHandler.Create).Methods("POST")
	adminRouter.HandleFunc("/{id}", cfg.ProductHandler.Update).Methods("PUT")
	adminRouter.HandleFunc("/{id}", cfg.ProductHandler.Delete).Methods("DELETE")
	adminRouter.Use(authMiddleware, middleware.RequireRole("admin"))

	// health
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	}).Methods("GET")

	return r
}