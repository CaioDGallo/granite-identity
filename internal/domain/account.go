package domain

import (
	"math/big"
	"time"

	"github.com/google/uuid"
)

type Account struct {
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	LastActivity  time.Time `json:"last_activity"`
	AccountType   string    `json:"account_type"`
	Currency      string    `json:"currency"`
	Status        string    `json:"status"`
	AccountNumber string    `json:"account_number"`
	Balance       big.Rat   `json:"balance"`
	ID            uuid.UUID `json:"id"`
	UserID        uuid.UUID `json:"user_id"`
}
