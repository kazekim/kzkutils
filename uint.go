package kzkutils

import "strconv"

func Uint64ToFloat64(value uint64) float64 {
	return float64(value)
}

func UIntToString(n uint) string {
	return strconv.Itoa(int(n))
}
