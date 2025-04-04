package experiment

import (
	"math/rand/v2"

	"github.com/mlange-42/ark/ecs"
	"github.com/mlange-42/beecs/model"
)

// Experiment definition.
type Experiment struct {
	parameters    []string
	values        [][]any
	parameterSets int
	runsPerSet    int
}

// New creates a new Experiment with the given parameter variations,
// PRNG instance and number of runs per parameter set.
func New(vars []ParameterVariation, rng *rand.Rand, runs int) (Experiment, error) {
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

	values := make([][]any, stride*runs)
	for i := range values {
		values[i] = make([]any, len(pars))
		for j := range pars {
			fn := f[j]
			values[i][j] = fn.Next(i, rng)
		}
	}

	return Experiment{
		parameters:    pars,
		parameterSets: stride,
		runsPerSet:    runs,
		values:        values,
	}, nil
}

// ParameterSets returns the number of unique parameter sets.
// Random variations do not count towards the number of sets.
func (e *Experiment) ParameterSets() int {
	return e.parameterSets
}

// RunsPerSet returns the number of runs to perform per parameter sets.
// Random variations do not count towards the number of sets.
func (e *Experiment) RunsPerSet() int {
	return e.runsPerSet
}

// TotalRuns returns the total number of runs of this experiment.
func (e *Experiment) TotalRuns() int {
	return e.parameterSets * e.runsPerSet
}

// Parameters returns the names of the parameters varied in the experiment.
func (e *Experiment) Parameters() []string {
	return e.parameters
}

// Values returns the parameter values for the given run index.
func (e *Experiment) Values(idx int) []ParameterValue {
	values := []ParameterValue{}
	for i, par := range e.parameters {
		values = append(values, ParameterValue{Parameter: par, Value: e.values[idx][i]})
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
