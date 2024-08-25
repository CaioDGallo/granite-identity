package domain

import (
	"math/big"
	"time"

	"github.com/google/uuid"
)

type Account struct {
	CreatedAt     time.Time     `json:"created_at"`
	UpdatedAt     time.Time     `json:"updated_at"`
	LastActivity  time.Time     `json:"last_activity"`
	Currency      string        `json:"currency"`
	AccountNumber string        `json:"account_number"`
	Balance       big.Rat       `json:"balance"`
	AccountType   AccountType   `json:"account_type"`
	Status        AccountStatus `json:"status"`
	ID            uuid.UUID     `json:"id"`
	UserID        uuid.UUID     `json:"user_id"`
}
