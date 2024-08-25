package database

import (
	"context"
	"fmt"

	dbstore "github.com/CaioDGallo/granite-identity/db"
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
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	q := dbstore.New(tx)

	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(ctx); rbErr != nil {
			return fmt.Errorf("transaction rollback failed: %w", rbErr)
		}
		return fmt.Errorf("transaction failed: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("transaction commit failed: %w", err)
	}

	return nil
}
