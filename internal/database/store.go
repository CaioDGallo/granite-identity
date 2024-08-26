package database

import (
	"context"
	"fmt"
	"log/slog"

	dbstore "github.com/CaioDGallo/granite-identity/db"
	"github.com/CaioDGallo/granite-identity/internal/logger"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Store struct {
	*dbstore.Queries
	connPool *pgxpool.Pool
}

var store *Store

func GetStore() *Store {
	if store == nil {
		return NewStore(GetDB())
	}

	return store
}

func NewStore(connPool *pgxpool.Pool) *Store {
	return &Store{
		connPool: connPool,
		Queries:  dbstore.New(connPool),
	}
}

func (s *Store) ExecTx(ctx context.Context, fn func(*dbstore.Queries) error) error {
	tx, err := s.connPool.Begin(ctx)

	txID := uuid.New().String()

	logger.GetLogger().Info("starting transaction", slog.String("tx_id", txID))

	if err != nil {
		logger.GetLogger().Error("failed to begin transaction", slog.String("tx_id", txID), slog.String("error", err.Error()))
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	q := dbstore.New(tx)

	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(ctx); rbErr != nil {
			logger.GetLogger().Error("failed to rollback transaction", slog.String("tx_id", txID), slog.String("error", rbErr.Error()))
			return fmt.Errorf("transaction rollback failed: %w", rbErr)
		}

		logger.GetLogger().Error("transaction failed", slog.String("tx_id", txID), slog.String("error", err.Error()))
		return fmt.Errorf("transaction failed: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		logger.GetLogger().Error("failed to commit transaction", slog.String("tx_id", txID), slog.String("error", err.Error()))
		return fmt.Errorf("transaction commit failed: %w", err)
	}

	logger.GetLogger().Info("transaction committed", slog.String("tx_id", txID))
	return nil
}
