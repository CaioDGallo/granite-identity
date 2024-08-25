package domain

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func GenerateAccountNumber() (string, error) {
	const length = 10

	accountNumber, err := randomNumericString(length - 1)
	if err != nil {
		return "", err
	}

	checkDigit := calculateCheckDigit(accountNumber)

	accountNumber += checkDigit

	return accountNumber, nil
}

func randomNumericString(length int) (string, error) {
	const digits = "0123456789"
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(digits))))
		if err != nil {
			return "", fmt.Errorf("failed to generate random number: %w", err)
		}
		result[i] = digits[n.Int64()]
	}
	return string(result), nil
}

func calculateCheckDigit(accountNumber string) string {
	sum := 0
	double := false

	for i := len(accountNumber) - 1; i >= 0; i-- {
		digit := int(accountNumber[i] - '0')

		if double {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}

		sum += digit
		double = !double
	}

	checkDigit := (10 - (sum % 10)) % 10
	return fmt.Sprintf("%d", checkDigit)
}
