package params

import "golang.org/x/exp/rand"

type Experiment struct {
	values        map[string]any
	rng           *rand.Rand
	parameters    []string
	functions     []ParameterFunction
	parameterSets int
}

func NewExperiment(vars []ParameterVariation, rng *rand.Rand) (Experiment, error) {
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
		values:        map[string]any{},
	}, nil
}

func (e *Experiment) ParameterSets() int {
	return e.parameterSets
}

func (e *Experiment) Parameters() []string {
	return e.parameters
}

func (e Experiment) Values(idx int) map[string]any {
	for i, par := range e.parameters {
		fn := e.functions[i]
		e.values[par] = fn.Next(idx, e.rng)
	}
	return e.values
}
