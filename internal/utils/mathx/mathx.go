package mathx

import (
	"github.com/shopspring/decimal"
	"math"
)

// Round 四捨五入
func Round(val float64, precision int) float64 {
	p := math.Pow10(precision)
	return math.Floor(val*p+0.5) / p
}

// Roundx 四捨六入 五不動
func Roundx(val float64, precision int) float64 {
	p := math.Pow10(precision)
	val *= p
	// 提取整數和小數
	intPart := math.Floor(val)
	decimalPart, _ := decimal.NewFromFloat(val).
		Sub(decimal.NewFromFloat(intPart)).
		Mul(decimal.NewFromInt(10)).
		Floor().
		Div(decimal.NewFromInt(10)).Float64()

	// 處理小數
	switch {
	case decimalPart < 0.5:
		return intPart / p
	case decimalPart > 0.5:
		return (intPart + 1) / p
	default:
		return (intPart + decimalPart) / p
	}
}
