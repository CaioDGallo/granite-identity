package utils

import (
	"fmt"
	"math/big"

	"github.com/jackc/pgx/v5/pgtype"
)

// Convert pgtype.Numeric to *big.Rat
func NumericToBigRat(numeric pgtype.Numeric) (*big.Rat, error) {
	// Check if the numeric value is NaN
	if numeric.NaN {
		return nil, fmt.Errorf("numeric value is NaN")
	}

	// Convert pgtype.Numeric.Int (big.Int) to big.Rat
	rat := new(big.Rat).SetInt(numeric.Int)

	// Adjust the scale using Exp
	if numeric.Exp < 0 {
		// Scale down: divide by 10^(-Exp)
		scale := new(big.Rat).SetFloat64(1.0 / float64Pow10(-numeric.Exp))
		rat.Mul(rat, scale)
	} else if numeric.Exp > 0 {
		// Scale up: multiply by 10^Exp
		scale := new(big.Rat).SetFloat64(float64Pow10(numeric.Exp))
		rat.Mul(rat, scale)
	}

	return rat, nil
}

// Utility function to compute 10^exp
func float64Pow10(exp int32) float64 {
	result := 1.0
	for i := int32(0); i < exp; i++ {
		result *= 10
	}
	return result
}

func BigRatToNumeric(rat *big.Rat) (pgtype.Numeric, error) {
	// Convert big.Rat to string
	ratStr := rat.FloatString(10) // Convert to string with 10 decimal places

	// Create a new pgtype.Numeric
	var numeric pgtype.Numeric

	// Use Scan to set the value
	err := numeric.Scan(ratStr)
	if err != nil {
		return pgtype.Numeric{}, err
	}

	return numeric, nil
}
