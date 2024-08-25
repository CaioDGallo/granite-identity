package service

import (
	"context"
	"errors"
	"math/big"
	"time"

	"github.com/CaioDGallo/granite-identity/internal/domain"
	"github.com/CaioDGallo/granite-identity/internal/repository"
	"github.com/google/uuid"
)

var (
	ErrAccountNotFound   = errors.New("account not found")
	ErrInsufficientFunds = errors.New("insufficient funds")
)

type CreateAccountRequest struct {
	Currency    string
	AccountType string
	UserID      uuid.UUID
}

type AccountService struct {
	r *repository.AccountRepository
}

func NewAccountService() *AccountService {
	return &AccountService{
		r: repository.NewAccountRepository(),
	}
}

func (s *AccountService) CreateAccount(req CreateAccountRequest) (*domain.Account, error) {
	ctx := context.Background()

	balance := *new(big.Rat).SetInt64(0)

	var accountNumber string
	var err error

	for {
		accountNumber, err = domain.GenerateAccountNumber()
		if err != nil {
			return nil, err
		}

		found, _ := s.r.GetAccountByAccountNumber(ctx, accountNumber)

		if found == nil {
			break
		}
	}

	accountType, err := domain.ParseAccountTypeFromString(req.AccountType)
	if err != nil {
		return nil, err
	}

	account := &domain.Account{
		ID:            uuid.New(),
		Currency:      req.Currency,
		AccountType:   accountType,
		UserID:        req.UserID,
		Status:        domain.Active,
		Balance:       balance,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		LastActivity:  time.Now(),
		AccountNumber: accountNumber,
	}

	newAccount, err := s.r.CreateAccount(ctx, account)
	if err != nil {
		return nil, err
	}

	return newAccount, nil
}

func (s *AccountService) GetAccountByID(accountID string) (*domain.Account, error) {
	accountUUID, err := uuid.Parse(accountID)
	if err != nil {
		return nil, err
	}

	ctx := context.Background()

	account, err := s.r.GetAccountByID(ctx, accountUUID)
	if err != nil {
		return nil, err
	}

	return account, nil
}
