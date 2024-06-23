package kzkutils

import "math"

func SumCreditQuantity(creditRate float64, amount float64) int {
	sum := (1 / creditRate) * amount
	creditQuantity := int(sum)
	return creditQuantity
}

func SumMessageCredit(message string) float64 {
	textAmount := len(message)
	credit := float64(textAmount) / 1000
	res := math.Ceil(credit)
	return res
}
