package params

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"reflect"

	"github.com/mlange-42/ark/ecs"
	"github.com/mlange-42/beecs/registry"
	"github.com/mlange-42/beecs/util"
)

type entry struct {
	Bytes []byte
}

func (e *entry) UnmarshalJSON(jsonData []byte) error {
	e.Bytes = jsonData
	return nil
}

func (e entry) MarshalJSON() ([]byte, error) {
	return e.Bytes, nil
}

// CustomParams contains all default parameters of BEEHAVE and a map of additional custom parameter resources.
//
// CustomParams implements [Params].
type CustomParams struct {
	Parameters DefaultParams
	Custom     map[reflect.Type]any
}

type customParamsJs struct {
	Parameters DefaultParams
	Custom     map[string]entry
}

// FromJSONFile fills the parameter set with values from a JSON file.
//
// Only values present in the file are overwritten,
// all other values remain unchanged.
func (p *CustomParams) FromJSONFile(path string) error {
	content, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	return p.FromJSON(content)
}

// FromJSON fills the parameter set with values from JSON.
//
// Only values present in the file are overwritten,
// all other values remain unchanged.
func (p *CustomParams) FromJSON(data []byte) error {
	reader := bytes.NewReader(data)
	decoder := json.NewDecoder(reader)
	decoder.DisallowUnknownFields()

	pars := customParamsJs{
		Parameters: p.Parameters,
	}
	err := decoder.Decode(&pars)
	if err != nil {
		return err
	}

	p.Parameters = pars.Parameters
	if p.Custom == nil {
		p.Custom = map[reflect.Type]any{}
	}

	for tpName, entry := range pars.Custom {
		tp, ok := registry.GetResource(tpName)
		if !ok {
			return fmt.Errorf("resource type '%s' is not registered", tpName)
		}
		resourceVal := reflect.New(tp).Interface()

		decoder := json.NewDecoder(bytes.NewReader(entry.Bytes))
		decoder.DisallowUnknownFields()
		if err := decoder.Decode(resourceVal); err != nil {
			return err
		}

		p.Custom[tp] = resourceVal
	}
	return nil
}

// ToJSON marshals all parameters to JSON format.
func (p *CustomParams) ToJSON() ([]byte, error) {
	par := customParamsJs{
		Parameters: p.Parameters,
		Custom:     map[string]entry{},
	}

	for k, v := range p.Custom {
		js, err := json.MarshalIndent(&v, "", "    ")
		if err != nil {
			return []byte{}, err
		}
		par.Custom[k.String()] = entry{Bytes: js}
	}

	js, err := json.MarshalIndent(&par, "", "    ")
	if err != nil {
		return []byte{}, err
	}
	return js, nil
}

// Apply the parameters to a world by adding them as resources.
func (p *CustomParams) Apply(world *ecs.World) {
	p.Parameters.Apply(world)

	for tp, res := range p.Custom {
		id := ecs.ResourceTypeID(world, tp)
		world.Resources().Add(id, util.CopyInterface[any](res))
	}
}
