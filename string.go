package kzkutils

import (
	"regexp"
	"strconv"
	"strings"
)

func StringToUInt64(value string) (uint64, error) {
	if s, err := strconv.ParseUint(value, 10, 64); err == nil {
		return s, nil
	} else {
		return 0, err
	}
}

func StringToFloat64(value string) (float64, error) {
	v, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, err
	}
	return v, nil
}

func ToSnakeCase(str string) string {
	var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func StringToInt(value string) (int, error) {
	v, err := strconv.Atoi(value)
	if err != nil {
		return 0, err
	}
	return v, nil
}

func StringToUInt(value string) (uint, error) {
	v, err := strconv.Atoi(value)
	if err != nil {
		return 0, err
	}
	return uint(v), nil
}

func StringToInt64(value string) (int64, error) {
	v, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0, err
	}
	return v, nil
}

func StringReverse(value string) string {
	runes := []rune(value)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func ConvertString(value interface{}) string {

	switch value.(type) {
	case int64:
		return Int64ToString(value.(int64))
	case float64:
		return Float64ToString(value.(float64))
	case string:
		return value.(string)
	}
	return ""
}

func StripNonCharacter(s string) string {
	var result strings.Builder
	for i := 0; i < len(s); i++ {
		b := s[i]
		if ('a' <= b && b <= 'z') ||
			('A' <= b && b <= 'Z') ||
			('0' <= b && b <= '9') ||
			b == ' ' {
			result.WriteByte(b)
		}
	}
	return result.String()
}
