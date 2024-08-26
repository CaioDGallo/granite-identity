package database

import (
	"context"
	"fmt"
	"log"
	"log/slog"

	"github.com/CaioDGallo/granite-identity/internal/config"
	"github.com/CaioDGallo/granite-identity/internal/logger"
	"github.com/jackc/pgx/v5/pgxpool"
)

var dbPool *pgxpool.Pool

func Connect(ctx context.Context, cfg config.Config) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(
		ctx,
		fmt.Sprintf(
			"user=%s password=%s host=%s port=%s dbname=%s",
			cfg.DBUser,
			cfg.DBPassword,
			cfg.DBHost,
			cfg.DBPort,
			cfg.DBName,
		))
	if err != nil {
		logger.GetLogger().Error("failed to connect to database", slog.String("error", err.Error()))
		return nil, err
	}

	if err := pool.Ping(ctx); err != nil {
		logger.GetLogger().Error("failed to ping database", slog.String("error", err.Error()))
		return nil, err
	}

	dbPool = pool

	return pool, nil
}

func GetDB() *pgxpool.Pool {
	return dbPool
}

func Close() {
	dbPool.Close()
	log.Println("Database connection closed")
}
