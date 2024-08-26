package service

import (
	"context"
	"errors"
	"math/big"
	"time"

	"github.com/CaioDGallo/granite-identity/internal/database"
	"github.com/CaioDGallo/granite-identity/internal/domain"
	"github.com/CaioDGallo/granite-identity/internal/logger"
	"github.com/CaioDGallo/granite-identity/internal/repository"
	"github.com/CaioDGallo/granite-identity/internal/security/encryption"
	"github.com/CaioDGallo/granite-identity/internal/security/keymanager"
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
		r: repository.NewAccountRepository(database.GetStore()),
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
			logger.GetLogger().Error("failed to generate account number", "error", err.Error())
			return nil, err
		}

		found, _ := s.r.GetAccountByAccountNumber(ctx, accountNumber)

		if found == nil {
			break
		}
	}

	encryptedAccountNumber, err := encryption.Encrypt(accountNumber, keymanager.GetKey())
	if err != nil {
		logger.GetLogger().Error("failed to encrypt account number", "error", err.Error())
		return nil, err
	}

	accountType, err := domain.ParseAccountTypeFromString(req.AccountType)
	if err != nil {
		logger.GetLogger().Error("failed to parse account type", "error", err.Error())
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
		AccountNumber: encryptedAccountNumber,
	}

	newAccount, err := s.r.CreateAccount(ctx, account)
	if err != nil {
		logger.GetLogger().Error("failed to create account", "error", err.Error())
		return nil, err
	}

	decryptedAccountNumber, err := encryption.Decrypt(newAccount.AccountNumber, keymanager.GetKey())
	if err != nil {
		logger.GetLogger().Error("failed to decrypt account number", "error", err.Error())
		return nil, err
	}

	newAccount.AccountNumber = decryptedAccountNumber

	return newAccount, nil
}

func (s *AccountService) GetAccountByID(accountID string) (*domain.Account, error) {
	accountUUID, err := uuid.Parse(accountID)
	if err != nil {
		logger.GetLogger().Error("failed to parse account ID", "error", err.Error())
		return nil, err
	}

	ctx := context.Background()

	account, err := s.r.GetAccountByID(ctx, accountUUID)
	if err != nil {
		logger.GetLogger().Error("failed to get account by ID", "error", err.Error())
		return nil, err
	}

	decryptedAccountNumber, err := encryption.Decrypt(account.AccountNumber, keymanager.GetKey())
	if err != nil {
		logger.GetLogger().Error("failed to decrypt account number", "error", err.Error())
		return nil, err
	}

	account.AccountNumber = decryptedAccountNumber

	return account, nil
}
