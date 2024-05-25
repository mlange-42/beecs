package util

import "reflect"

// CopyInterface copies the underlying value of an interface pointer using reflection.
func CopyInterface[T any](value any) T {
	v := reflect.Indirect(reflect.ValueOf(value))
	b := v.Interface()

	val := reflect.ValueOf(b)
	ptr := reflect.New(val.Type()).Elem()
	ptr.Set(val)
	return ptr.Addr().Interface().(T)
}
