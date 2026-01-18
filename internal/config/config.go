package config

import (
	"errors"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port             int
	MongoURI         string
	JWTSecret        string
	MongoDBName      string
	JWTExpiryMinutes int
}

func Load() (*Config, error) {
	_ = godotenv.Load()

	portStr := os.Getenv("PORT")
	if portStr == "" {
		portStr = "8080"
	}
	port, _ := strconv.Atoi(portStr)

	jwtExpiry, _ := strconv.Atoi(os.Getenv("JWT_EXPIRY_MINUTES"))
	if jwtExpiry == 0 {
		jwtExpiry = 60 // default to 60 minutes
	}

	cfg := &Config{
		Port:             port,
		MongoURI:         os.ExpandEnv("MONGODB_URI"),
		MongoDBName:      os.Getenv("MONGO_DB"),
		JWTSecret:        os.Getenv("JWT_SECRET"),
		JWTExpiryMinutes: jwtExpiry,
	}

	if cfg.MongoURI == "" || cfg.JWTSecret == "" || cfg.MongoDBName == "" {
		return nil, errors.New("missing required environment variables")
	}

	return cfg, nil
}
