package experiment

import (
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/beecs/model"
	"golang.org/x/exp/rand"
)

// Experiment definition.
type Experiment struct {
	rng           *rand.Rand
	parameters    []string
	functions     []ParameterFunction
	parameterSets int
}

// New creates a new Experiment with the given parameter variations and PRNG instance.
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

// ParameterSets returns the number of unique parameter sets.
// Random variations do not count towards the number of sets.
func (e *Experiment) ParameterSets() int {
	return e.parameterSets
}

// Parameters returns the names of the parameters varied in the experiment.
func (e *Experiment) Parameters() []string {
	return e.parameters
}

// Re-seeds the experiment's PRNG.
func (e *Experiment) Seed(seed uint64) {
	e.rng.Seed(seed)
}

// Values returns the parameter values for the given run index.
func (e *Experiment) Values(idx int) []ParameterValue {
	values := []ParameterValue{}
	for i, par := range e.parameters {
		fn := e.functions[i]
		values = append(values, ParameterValue{Parameter: par, Value: fn.Next(idx, e.rng)})
	}
	return values
}

// ApplyValues applies the given parameter values to a model.
func (e *Experiment) ApplyValues(values []ParameterValue, world *ecs.World) error {
	for _, par := range values {
		if err := model.SetParameter(world, par.Parameter, par.Value); err != nil {
			return err
		}
	}
	return nil
}

// ParameterValue pair.
type ParameterValue struct {
	Parameter string
	Value     any
}
