package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"sync"

	middleware "github.com/CaioDGallo/granite-identity/internal/api/middlewares"
	v1 "github.com/CaioDGallo/granite-identity/internal/api/v1"
	"github.com/CaioDGallo/granite-identity/internal/config"
	database "github.com/CaioDGallo/granite-identity/internal/database"
	grpc "github.com/CaioDGallo/granite-identity/internal/grpc/server"
	"github.com/CaioDGallo/granite-identity/internal/security/keymanager"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	var keyLoader keymanager.KeyLoader

	switch cfg.KeyType {
	case "file":
		keyLoader = &keymanager.FileKeyLoader{Path: cfg.KeySource}
	case "env":
		keyLoader = &keymanager.EnvKeyLoader{EnvVar: cfg.KeySource}
	case "kms":
		keyLoader = &keymanager.KMSKeyLoader{KeyID: cfg.KeySource}
	default:
		fmt.Println("No valid KeyType specified")
		return
	}

	_, err = keymanager.LoadKey(keyLoader)
	if err != nil {
		log.Fatalf("Failed to load key: %v", err)
	}

	pool, err := database.Connect(context.Background(), *cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	defer pool.Close()

	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func(cfg *config.Config, wg *sync.WaitGroup) {
		graniteGRPCServer := grpc.GraniteGRPCServer{}
		graniteGRPCServer.StartListening(cfg, wg)
	}(cfg, wg)

	go startGinServer(cfg, wg)

	wg.Wait()

	slog.Info("Application started", slog.String("env", cfg.Environment), slog.String("version", cfg.Version))
}

func startGinServer(cfg *config.Config, wg *sync.WaitGroup) {
	defer wg.Done()
	router := gin.Default()

	router.Use(middleware.Logging())

	v1.RegisterRoutes(router)

	if err := router.Run(cfg.ServerAddress); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
