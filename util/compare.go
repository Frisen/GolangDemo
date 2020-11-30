package util

import (
	"errors"
	"reflect"
)

func Compare(a interface{}, b interface{}) (int, error) {
	aType := reflect.TypeOf(a).String()
	bType := reflect.TypeOf(b).String()
	if aType != bType {
		return 0, errors.New("type is not same,can't compare")
	}

	switch a.(type) {
	case int:
		if a.(int) > b.(int) {
			return 1, nil
		} else if a.(int) < b.(int) {
			return -1, nil
		}
		return 0, nil
	case string:
		if a.(string) > b.(string) {
			return 1, nil
		} else if a.(string) < b.(string) {
			return -1, nil
		}
		return 0, nil
	default:
		panic("type not contain")
	}
}
