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
	store *database.Store
}

func NewAccountRepository(store *database.Store) *AccountRepository {
	db := database.GetDB()
	if db == nil {
		panic("database connection is nil")
	}

	queries := dbstore.New(db)
	if queries == nil {
		panic("failed to initialize dbstore.Queries")
	}

	return &AccountRepository{
		store: store,
	}
}

func (r *AccountRepository) CreateAccount(ctx context.Context, account *domain.Account) (*domain.Account, error) {
	var newAccount dbstore.Account
	var balance pgtype.Numeric

	balanceStr := account.Balance.RatString()
	if err := balance.Scan(balanceStr); err != nil {
		return nil, err
	}

	err := r.store.ExecTx(ctx, func(q *dbstore.Queries) error {
		var err error
		newAccount, err = q.CreateAccount(ctx, dbstore.CreateAccountParams{
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
		return err
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
	account, err := r.store.GetAccountByID(ctx, id)
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
	account, err := r.store.GetAccountByAccountNumber(ctx, accountNumber)
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
