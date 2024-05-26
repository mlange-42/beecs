package registry

import (
	"fmt"
	"reflect"
)

var observerRegistry = map[string]reflect.Type{}
var resourcesRegistry = map[string]reflect.Type{}
var systemsRegistry = map[string]reflect.Type{}

// RegisterObserver registers a type as an observer.
func RegisterObserver[T any]() {
	tp := reflect.TypeOf((*T)(nil)).Elem()
	if _, ok := observerRegistry[tp.String()]; ok {
		panic(fmt.Sprintf("there is already an observer with type name '%s' registered", tp.String()))
	}
	observerRegistry[tp.String()] = tp
}

// RegisterResource registers a type as a resource.
func RegisterResource[T any]() {
	tp := reflect.TypeOf((*T)(nil)).Elem()
	if _, ok := resourcesRegistry[tp.String()]; ok {
		panic(fmt.Sprintf("there is already a resource with type name '%s' registered", tp.String()))
	}
	resourcesRegistry[tp.String()] = tp
}

// RegisterSystem registers a type as a system.
func RegisterSystem[T any]() {
	tp := reflect.TypeOf((*T)(nil)).Elem()
	if _, ok := systemsRegistry[tp.String()]; ok {
		panic(fmt.Sprintf("there is already a system with type name '%s' registered", tp.String()))
	}
	systemsRegistry[tp.String()] = tp
}

// GetObserver gets a registered observer type by type name ([reflect.Type.String]).
func GetObserver(name string) (reflect.Type, bool) {
	t, ok := observerRegistry[name]
	return t, ok
}

// GetResource gets a registered resource type by type name ([reflect.Type.String]).
func GetResource(name string) (reflect.Type, bool) {
	t, ok := resourcesRegistry[name]
	return t, ok
}

// GetSystem gets a registered system type by type name ([reflect.Type.String]).
func GetSystem(name string) (reflect.Type, bool) {
	t, ok := systemsRegistry[name]
	return t, ok
}
