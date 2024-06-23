package kzkutils

import "strconv"

func IntToString(n int) string {
	return strconv.Itoa(n)
}

func Int32ToString(n int32) string {
	return strconv.FormatInt(int64(n), 10)
}

func Int64ToString(n int64) string {
	return strconv.FormatInt(n, 10)
}

func UInt64ToString(n uint64) string {
	return strconv.FormatInt(int64(n), 10)
}
