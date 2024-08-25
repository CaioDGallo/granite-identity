package database

import (
	"context"
	"fmt"
	"log"

	"github.com/CaioDGallo/granite-identity/internal/config"
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
		return nil, err
	}

	if err := pool.Ping(ctx); err != nil {
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
