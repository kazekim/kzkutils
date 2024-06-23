package kzkutils

import (
	"fmt"
	"math"
)

func RoundedFloat2Decimal(value float64) float64 {
	return math.Round(value*100) / 100
}

func RoundedFloat4Decimal(value float64) float64 {
	return math.Round(value*10000) / 10000
}

func Float64ToString(value float64) string {
	return fmt.Sprintf("%.2f", value)
}
