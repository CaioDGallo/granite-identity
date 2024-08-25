package repository

import (
	"context"

	dbstore "github.com/CaioDGallo/granite-identity/db"
	"github.com/CaioDGallo/granite-identity/internal/database"
	"github.com/CaioDGallo/granite-identity/internal/domain"
	utils "github.com/CaioDGallo/granite-identity/internal/util"
	"github.com/google/uuid"
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
		Status:        dbstore.AccountStatus(account.Status.String()),
		CreatedAt:     pgtype.Timestamp{Time: account.CreatedAt, Valid: true},
		UpdatedAt:     pgtype.Timestamp{Time: account.UpdatedAt, Valid: true},
		AccountType:   dbstore.AccountType(account.AccountType.String()),
		AccountNumber: account.AccountNumber,
	})
	if err != nil {
		return nil, err
	}

	nb, err := utils.NumericToBigRat(newAccount.Balance)
	if err != nil {
		return nil, err
	}

	status, err := domain.ParseAccountStatus(newAccount.Status)
	if err != nil {
		return nil, err
	}

	accountType, err := domain.ParseAccountType(newAccount.AccountType)
	if err != nil {
		return nil, err
	}

	return &domain.Account{
		ID:            newAccount.ID,
		UserID:        newAccount.UserID,
		Balance:       *nb,
		Currency:      newAccount.Currency,
		Status:        status,
		CreatedAt:     newAccount.CreatedAt.Time,
		UpdatedAt:     newAccount.UpdatedAt.Time,
		AccountType:   accountType,
		AccountNumber: newAccount.AccountNumber,
		LastActivity:  newAccount.LastActivity.Time,
	}, nil
}

func (r *AccountRepository) GetAccountByID(ctx context.Context, id uuid.UUID) (*domain.Account, error) {
	account, err := r.queries.GetAccountByID(ctx, id)
	if err != nil {
		return nil, err
	}

	nb, err := utils.NumericToBigRat(account.Balance)
	if err != nil {
		return nil, err
	}

	status, err := domain.ParseAccountStatus(account.Status)
	if err != nil {
		return nil, err
	}

	accountType, err := domain.ParseAccountType(account.AccountType)
	if err != nil {
		return nil, err
	}

	return &domain.Account{
		ID:            account.ID,
		UserID:        account.UserID,
		Balance:       *nb,
		Currency:      account.Currency,
		Status:        status,
		CreatedAt:     account.CreatedAt.Time,
		UpdatedAt:     account.UpdatedAt.Time,
		AccountType:   accountType,
		AccountNumber: account.AccountNumber,
		LastActivity:  account.LastActivity.Time,
	}, nil
}

func (r *AccountRepository) GetAccountByAccountNumber(ctx context.Context, accountNumber string) (*domain.Account, error) {
	account, err := r.queries.GetAccountByAccountNumber(ctx, accountNumber)
	if err != nil {
		return nil, err
	}

	nb, err := utils.NumericToBigRat(account.Balance)
	if err != nil {
		return nil, err
	}

	status, err := domain.ParseAccountStatus(account.Status)
	if err != nil {
		return nil, err
	}

	accountType, err := domain.ParseAccountType(account.AccountType)
	if err != nil {
		return nil, err
	}

	return &domain.Account{
		ID:            account.ID,
		UserID:        account.UserID,
		Balance:       *nb,
		Currency:      account.Currency,
		Status:        status,
		CreatedAt:     account.CreatedAt.Time,
		UpdatedAt:     account.UpdatedAt.Time,
		AccountType:   accountType,
		AccountNumber: account.AccountNumber,
		LastActivity:  account.LastActivity.Time,
	}, nil
}
