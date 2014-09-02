// Package is provides an is.Zero function
package is

import (
	"errors"
	"reflect"
)

// Zero tells whether a given value is its "zero value". It returns an error if given an array.
func Zero(val interface{}) (bool, error) {
	value := reflect.ValueOf(val)
	switch value.Kind() {
	case reflect.Bool:
		return !value.Bool(), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return value.Int() == 0, nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return value.Uint() == 0, nil
	case reflect.Uintptr:
		return value.Pointer() == 0, nil
	case reflect.Float32, reflect.Float64:
		return value.Float() == 0.0, nil
	case reflect.Complex64, reflect.Complex128:
		return value.Complex() == (0 + 0i), nil
	case reflect.Array:
		return false, errors.New("Cannot find the zero value of the given array")
	case reflect.Slice:
		return value.Len() == 0, nil
	case reflect.String:
		return value.String() == "", nil
	case reflect.Struct:
		numberOfFields := value.NumField()
		for i := 0; i < numberOfFields; i++ {
			if !Zero(value.Field(i).Interface()) {
				return false, nil
			}
		}
		return true, nil
	default:
		return value.IsNil(), nil
	}
}
