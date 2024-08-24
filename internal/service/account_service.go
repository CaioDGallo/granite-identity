package service

import (
	"errors"

	dbstore "github.com/CaioDGallo/granite-identity/db"
	"github.com/CaioDGallo/granite-identity/internal/database"
	"github.com/google/uuid"
)

var (
	ErrAccountNotFound   = errors.New("account not found")
	ErrInsufficientFunds = errors.New("insufficient funds")
)

type Account struct {
	ID uuid.UUID
}

type CreateAccountRequest struct {
	UserID uuid.UUID
}

type AccountService struct {
	q *dbstore.Queries
}

func NewAccountService() *AccountService {
	return &AccountService{q: dbstore.New(database.GetDB())}
}

func (s *AccountService) CreateAccount(req CreateAccountRequest) (Account, error) {
	account := Account{
		ID: uuid.New(),
	}

	return account, nil
}

func (s *AccountService) GetAccount(accountID string) (Account, error) {
	account := Account{
		ID: uuid.New(),
	}

	return account, nil
}
