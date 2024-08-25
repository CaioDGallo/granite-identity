package domain

import (
	"fmt"

	dbstore "github.com/CaioDGallo/granite-identity/db"
)

type AccountType int

const (
	Savings AccountType = iota
	Checking
	Business
	Credit
	Investment
)

func (t AccountType) String() string {
	return [...]string{"Savings", "Checking", "Business", "Credit", "Investment"}[t]
}

func ParseAccountTypeFromString(accountType string) (AccountType, error) {
	switch accountType {
	case "Savings":
		return Savings, nil
	case "Checking":
		return Checking, nil
	case "Business":
		return Business, nil
	case "Credit":
		return Credit, nil
	case "Investment":
		return Investment, nil
	default:
		return -1, fmt.Errorf("invalid account type: %s", accountType)
	}
}

func ParseAccountType(accountType dbstore.AccountType) (AccountType, error) {
	switch accountType {
	case dbstore.AccountTypeSavings:
		return Savings, nil
	case dbstore.AccountTypeChecking:
		return Checking, nil
	case dbstore.AccountTypeBusiness:
		return Business, nil
	case dbstore.AccountTypeCredit:
		return Credit, nil
	case dbstore.AccountTypeInvestment:
		return Investment, nil
	default:
		return -1, fmt.Errorf("invalid account type: %s", accountType)
	}
}
