package kzkutils

func Float64ToCentValue(value float64) int64 {

	v := value * 100
	cent := int64(v)
	return cent
}

func CentValueToFloat(value uint64) float64 {

	v := Uint64ToFloat64(value)
	v = v / 100
	return v
}
