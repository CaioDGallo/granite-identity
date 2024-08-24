package repository

import (
	"context"

	dbstore "github.com/CaioDGallo/granite-identity/db"
	"github.com/CaioDGallo/granite-identity/internal/database"
	"github.com/CaioDGallo/granite-identity/internal/domain"
	utils "github.com/CaioDGallo/granite-identity/internal/util"
	"github.com/jackc/pgx/v5/pgtype"
)

type AccountRepository struct {
	queries *dbstore.Queries
}

func NewAccountRepository() *AccountRepository {
	return &AccountRepository{
		queries: dbstore.New(database.GetDB()),
	}
}

func (r *AccountRepository) CreateAccount(ctx context.Context, account *domain.Account) (*domain.Account, error) {
	var balance pgtype.Numeric
	if err := balance.Scan(account.Balance); err != nil {
		return nil, err
	}

	newAccount, err := r.queries.CreateAccount(ctx, dbstore.CreateAccountParams{
		ID:            account.ID,
		UserID:        account.UserID,
		Balance:       balance,
		Currency:      account.Currency,
		Status:        account.Status,
		CreatedAt:     pgtype.Timestamp{Time: account.CreatedAt, Valid: true},
		UpdatedAt:     pgtype.Timestamp{Time: account.UpdatedAt, Valid: true},
		AccountType:   account.AccountType,
		AccountNumber: account.AccountNumber,
	})
	if err != nil {
		return nil, err
	}

	nb, err := utils.NumericToBigRat(newAccount.Balance)
	if err != nil {
		return nil, err
	}

	return &domain.Account{
		ID:            newAccount.ID,
		UserID:        newAccount.UserID,
		Balance:       *nb,
		Currency:      newAccount.Currency,
		Status:        newAccount.Status,
		CreatedAt:     newAccount.CreatedAt.Time,
		UpdatedAt:     newAccount.UpdatedAt.Time,
		AccountType:   newAccount.AccountType,
		AccountNumber: newAccount.AccountNumber,
		LastActivity:  newAccount.LastActivity.Time,
	}, nil
}
