package registry

import (
	"fmt"
	"reflect"
)

var observerRegistry = map[string]reflect.Type{}
var resourcesRegistry = map[string]reflect.Type{}
var systemsRegistry = map[string]reflect.Type{}

func RegisterObserver[T any]() {
	tp := reflect.TypeOf((*T)(nil)).Elem()
	if _, ok := observerRegistry[tp.String()]; ok {
		panic(fmt.Sprintf("there is already an observer with type name '%s' registered", tp.String()))
	}
	observerRegistry[tp.String()] = tp
}

func RegisterResource[T any]() {
	tp := reflect.TypeOf((*T)(nil)).Elem()
	if _, ok := resourcesRegistry[tp.String()]; ok {
		panic(fmt.Sprintf("there is already a resource with type name '%s' registered", tp.String()))
	}
	resourcesRegistry[tp.String()] = tp
}

func RegisterSystem[T any]() {
	tp := reflect.TypeOf((*T)(nil)).Elem()
	if _, ok := systemsRegistry[tp.String()]; ok {
		panic(fmt.Sprintf("there is already a system with type name '%s' registered", tp.String()))
	}
	systemsRegistry[tp.String()] = tp
}

func GetObserver(name string) (reflect.Type, bool) {
	t, ok := observerRegistry[name]
	return t, ok
}

func GetResource(name string) (reflect.Type, bool) {
	t, ok := resourcesRegistry[name]
	return t, ok
}

func GetSystem(name string) (reflect.Type, bool) {
	t, ok := systemsRegistry[name]
	return t, ok
}
