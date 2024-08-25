package domain

import (
	"encoding/json"
	"errors"
	"fmt"

	dbstore "github.com/CaioDGallo/granite-identity/db"
)

type AccountStatus int

const (
	Active AccountStatus = iota
	Suspended
	Closed
	Pending
)

func (s AccountStatus) String() string {
	return [...]string{"Active", "Suspended", "Closed", "Pending"}[s]
}

func (t AccountStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func ParseAccountStatusFromString(status string) (AccountStatus, error) {
	switch status {
	case "Active":
		return Active, nil
	case "Suspended":
		return Suspended, nil
	case "Closed":
		return Closed, nil
	case "Pending":
		return Pending, nil
	default:
		return -1, fmt.Errorf("invalid account status: %s", status)
	}
}

func ParseAccountStatus(status dbstore.AccountStatus) (AccountStatus, error) {
	switch status {
	case dbstore.AccountStatusActive:
		return Active, nil
	case dbstore.AccountStatusSuspended:
		return Suspended, nil
	case dbstore.AccountStatusClosed:
		return Closed, nil
	case dbstore.AccountStatusPending:
		return Pending, nil
	default:
		return -1, errors.New("invalid account status")
	}
}
