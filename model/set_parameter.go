package model

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/mlange-42/ark/ecs"
)

// SetParameter sets a parameter of the model from it's string identifier.
func SetParameter(world *ecs.World, param string, value any) error {
	parts := strings.Split(param, ".")
	if len(parts) < 2 {
		return fmt.Errorf("invalid parameter name '%s', should be pkg.Type.Field", param)
	}
	par := parts[len(parts)-1]
	group := strings.Join(parts[:len(parts)-1], ".")

	resID, resType, err := findResource(world, group)
	if err != nil {
		return err
	}

	res := world.Resources().Get(resID)
	if res == nil {
		return fmt.Errorf("resource type registered but nil: %s", resType.String())
	}

	rValue := reflect.ValueOf(res).Elem()

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
		case string:
			vv, err := strconv.Atoi(v)
			if err != nil {
				return err
			}
			f.SetInt(int64(vv))
		default:
			return fmt.Errorf("unsupported parameter type %s for integer field", reflect.TypeOf(value).String())
		}
	} else if f.Kind() == reflect.Float32 || f.Kind() == reflect.Float64 {
		switch v := value.(type) {
		case float32:
			f.SetFloat(float64(v))
		case float64:
			f.SetFloat(v)
		case string:
			vv, err := strconv.ParseFloat(v, 64)
			if err != nil {
				return err
			}
			f.SetFloat(vv)
		default:
			return fmt.Errorf("unsupported parameter type %s for float field", reflect.TypeOf(value).String())
		}
	} else if f.Kind() == reflect.Bool {
		switch v := value.(type) {
		case bool:
			f.SetBool(v)
		case string:
			vv, err := strconv.ParseBool(v)
			if err != nil {
				return err
			}
			f.SetBool(vv)
		default:
			return fmt.Errorf("unsupported parameter type %s for bool field", reflect.TypeOf(value).String())
		}
	} else if f.Kind() == reflect.String {
		switch v := value.(type) {
		case string:
			f.SetString(v)
		default:
			return fmt.Errorf("unsupported parameter type %s for string field", reflect.TypeOf(value).String())
		}
	} else {
		return fmt.Errorf("unsupported field kind %s", f.Kind().String())
	}

	return nil
}

// GetParameter gets a parameter of the model from it's string identifier.
func GetParameter(world *ecs.World, param string) (any, error) {
	parts := strings.Split(param, ".")
	if len(parts) < 2 {
		return nil, fmt.Errorf("invalid parameter name '%s', should be pkg.Type.Field", param)
	}
	par := parts[len(parts)-1]
	group := strings.Join(parts[:len(parts)-1], ".")

	resID, resType, err := findResource(world, group)
	if err != nil {
		return nil, err
	}

	res := world.Resources().Get(resID)
	if res == nil {
		return nil, fmt.Errorf("resource type registered but nil: %s", resType.String())
	}

	rValue := reflect.ValueOf(res).Elem()

	f := rValue.FieldByName(par)
	if !f.IsValid() {
		return nil, fmt.Errorf("%s is not a valid field of %s", par, resType.String())
	}

	if f.Kind() == reflect.Int || f.Kind() == reflect.Int32 || f.Kind() == reflect.Int64 {
		return f.Int(), nil
	} else if f.Kind() == reflect.Float32 || f.Kind() == reflect.Float64 {
		return f.Float(), nil
	} else if f.Kind() == reflect.Bool {
		return f.Bool(), nil
	}

	return nil, fmt.Errorf("unsupported parameter type %s", f.Kind())
}

func findResource(world *ecs.World, group string) (ecs.ResID, reflect.Type, error) {
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
		return resID, resType, fmt.Errorf("could not find resource '%s'", group)
	}
	return resID, resType, nil
}
