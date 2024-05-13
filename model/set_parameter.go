package model

import (
	"fmt"
	"reflect"

	"github.com/mlange-42/arche/ecs"
)

func SetParameter(world *ecs.World, group, par string, value any) error {
	var resID ecs.ResID
	var resType reflect.Type = nil
	allRes := ecs.ResourceIDs(world)
	for _, id := range allRes {
		if tp, ok := ecs.ResourceType(world, id); ok {
			if tp.String() == group {
				resID = id
				resType = tp
				break
			}
		}
	}
	if resType == nil {
		return fmt.Errorf("could not find resource '%s'", group)
	}

	resLoc := world.Resources().Get(resID)
	if resLoc == nil {
		return fmt.Errorf("resource type registered but nil: %s", resType.String())
	}

	rValue := reflect.ValueOf(resLoc).Elem()

	f := rValue.FieldByName(par)
	if !f.IsValid() {
		return fmt.Errorf("%s is not a valid field of %s", par, resType.String())
	}

	if f.Kind() == reflect.Int || f.Kind() == reflect.Int32 || f.Kind() == reflect.Int64 {
		switch v := value.(type) {
		case int:
			f.SetInt(int64(v))
		case int32:
			f.SetInt(int64(v))
		case int64:
			f.SetInt(v)
		default:
			return fmt.Errorf("unsupported parameter type %s for integer field", reflect.TypeOf(value).String())
		}
	} else if f.Kind() == reflect.Float32 || f.Kind() == reflect.Float64 {
		switch v := value.(type) {
		case float32:
			f.SetFloat(float64(v))
		case float64:
			f.SetFloat(v)
		default:
			return fmt.Errorf("unsupported parameter type %s for float field", reflect.TypeOf(value).String())
		}
	} else if f.Kind() == reflect.Bool {
		switch v := value.(type) {
		case bool:
			f.SetBool(v)
		default:
			return fmt.Errorf("unsupported parameter type %s for bool field", reflect.TypeOf(value).String())
		}
	}

	return nil
}
