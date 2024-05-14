package experiment

import (
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/beecs/model"
	"golang.org/x/exp/rand"
)

type Experiment struct {
	rng           *rand.Rand
	parameters    []string
	functions     []ParameterFunction
	parameterSets int
}

func New(vars []ParameterVariation, rng *rand.Rand) (Experiment, error) {
	pars := []string{}
	f := []ParameterFunction{}

	stride := 1
	for _, v := range vars {
		fn, err := NewParameterFunction(v, stride)
		if err != nil {
			return Experiment{}, err
		}
		stride *= fn.Stride()
		f = append(f, fn)
		pars = append(pars, v.Parameter)
	}

	return Experiment{
		parameters:    pars,
		functions:     f,
		parameterSets: stride,
		rng:           rng,
	}, nil
}

func (e *Experiment) ParameterSets() int {
	return e.parameterSets
}

func (e *Experiment) Parameters() []string {
	return e.parameters
}

func (e *Experiment) Seed(seed uint64) {
	e.rng.Seed(seed)
}

func (e *Experiment) Values(idx int) []ParameterValue {
	values := []ParameterValue{}
	for i, par := range e.parameters {
		fn := e.functions[i]
		values = append(values, ParameterValue{Parameter: par, Value: fn.Next(idx, e.rng)})
	}
	return values
}

func (e *Experiment) ApplyValues(idx int, world *ecs.World) error {
	values := e.Values(idx)
	for _, par := range values {
		if err := model.SetParameter(world, par.Parameter, par.Value); err != nil {
			return err
		}
	}
	return nil
}

type ParameterValue struct {
	Parameter string
	Value     any
}
