package utils

import (
	"math/big"
	"strconv"
)

func ParseFloat(str string) float64 {
	value, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0
	}

	return value
}

func ParseInt(str string) *big.Int {
	value, ok := new(big.Int).SetString(str, 10)
	if !ok {
		return big.NewInt(0)
	}

	return value
}

func ToWei(value int64) (result *big.Int) {
	result = new(big.Int)
	result.Mul(big.NewInt(value), big.NewInt(1e18))
	return
}
