package utils

import (
	"fmt"
	"math/big"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

func NumericToBigRat(numeric pgtype.Numeric) (*big.Rat, error) {
	if numeric.NaN {
		return nil, fmt.Errorf("numeric value is NaN")
	}

	rat := new(big.Rat).SetInt(numeric.Int)

	if numeric.Exp < 0 {

		scale := new(big.Rat).SetFloat64(1.0 / float64Pow10(-numeric.Exp))
		rat.Mul(rat, scale)
	} else if numeric.Exp > 0 {

		scale := new(big.Rat).SetFloat64(float64Pow10(numeric.Exp))
		rat.Mul(rat, scale)
	}

	return rat, nil
}

func float64Pow10(exp int32) float64 {
	result := 1.0
	for i := int32(0); i < exp; i++ {
		result *= 10
	}
	return result
}

func BigRatToNumeric(rat *big.Rat) (pgtype.Numeric, error) {
	ratStr := rat.FloatString(10)

	var numeric pgtype.Numeric

	err := numeric.Scan(ratStr)
	if err != nil {
		return pgtype.Numeric{}, err
	}

	return numeric, nil
}

func GetRequestID(c *gin.Context) (string, bool) {
	requestID, exists := c.Get("RequestID")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Request ID not found"})
		return "", false
	}

	return requestID.(string), true
}
