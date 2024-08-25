// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package dbstore

import (
	"database/sql/driver"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type AccountStatus string

const (
	AccountStatusActive    AccountStatus = "Active"
	AccountStatusSuspended AccountStatus = "Suspended"
	AccountStatusClosed    AccountStatus = "Closed"
	AccountStatusPending   AccountStatus = "Pending"
)

func (e *AccountStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = AccountStatus(s)
	case string:
		*e = AccountStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for AccountStatus: %T", src)
	}
	return nil
}

type NullAccountStatus struct {
	AccountStatus AccountStatus
	Valid         bool // Valid is true if AccountStatus is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullAccountStatus) Scan(value interface{}) error {
	if value == nil {
		ns.AccountStatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.AccountStatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullAccountStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.AccountStatus), nil
}

type AccountType string

const (
	AccountTypeSavings    AccountType = "Savings"
	AccountTypeChecking   AccountType = "Checking"
	AccountTypeBusiness   AccountType = "Business"
	AccountTypeCredit     AccountType = "Credit"
	AccountTypeInvestment AccountType = "Investment"
)

func (e *AccountType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = AccountType(s)
	case string:
		*e = AccountType(s)
	default:
		return fmt.Errorf("unsupported scan type for AccountType: %T", src)
	}
	return nil
}

type NullAccountType struct {
	AccountType AccountType
	Valid       bool // Valid is true if AccountType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullAccountType) Scan(value interface{}) error {
	if value == nil {
		ns.AccountType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.AccountType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullAccountType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.AccountType), nil
}

type Account struct {
	ID            uuid.UUID
	UserID        uuid.UUID
	Balance       pgtype.Numeric
	Currency      string
	Status        AccountStatus
	CreatedAt     pgtype.Timestamp
	UpdatedAt     pgtype.Timestamp
	AccountType   AccountType
	AccountNumber string
	LastActivity  pgtype.Timestamp
}

type Transaction struct {
	ID          uuid.UUID
	AccountID   uuid.UUID
	Type        string
	Amount      pgtype.Numeric
	Currency    string
	CreatedAt   pgtype.Timestamp
	ReferenceID string
	Description pgtype.Text
}
