package kzkutils

import (
	"math/rand"
)

var numberRunes = []rune("0123456789")
var stringRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func RandomNumber(digitAmount int) string {
	b := make([]rune, digitAmount)
	for i := range b {
		b[i] = numberRunes[rand.Intn(len(numberRunes))]
	}
	return string(b)
}

func RandomString(digitAmount int) string {
	b := make([]rune, digitAmount)
	for i := range b {
		b[i] = stringRunes[rand.Intn(len(numberRunes))]
	}
	return string(b)
}
