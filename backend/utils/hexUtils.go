package utils

import (
	"math"
	"math/big"
	"strconv"
	"strings"
)

func HexToDecimal(hex string) (*big.Int, bool) {
	formatted := strings.TrimPrefix(hex, "0x")
	decimalNum := new(big.Int)
	decimalNum, success := decimalNum.SetString(formatted, 16)
	return decimalNum, success
}

func HexToFloat64(hex string) (float64, error) {
	formatted := strings.TrimLeft(hex, "0x")
	i, err := strconv.ParseUint(formatted, 16, 64)
	if err != nil {
		return 0, err
	}
	return math.Float64frombits(i), nil
}
