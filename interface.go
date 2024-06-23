package kzkutils

import "reflect"

func IsNil(c interface{}) bool {
	if c == nil || (reflect.ValueOf(c).Kind() == reflect.Ptr && reflect.ValueOf(c).IsNil()) {
		return true
	}
	return false
}
