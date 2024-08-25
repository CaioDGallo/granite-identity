package main

import (
	"context"
	"log"
	"log/slog"

	middleware "github.com/CaioDGallo/granite-identity/internal/api/middlewares"
	v1 "github.com/CaioDGallo/granite-identity/internal/api/v1"
	"github.com/CaioDGallo/granite-identity/internal/config"
	database "github.com/CaioDGallo/granite-identity/internal/database"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	pool, err := database.Connect(context.Background(), *cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	defer pool.Close()

	router := gin.Default()

	router.Use(middleware.RequestID())

	v1.RegisterRoutes(router)

	if err := router.Run(cfg.ServerAddress); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	slog.Info("Application started", slog.String("env", cfg.Environment), slog.String("version", cfg.Version))
}
