package params

import (
	"fmt"

	"golang.org/x/exp/rand"
)

const (
	VaryRandomRange    = "random-range"
	VaryRandomValues   = "random-values"
	VarySequenceRange  = "sequence-range"
	VarySequenceValues = "sequence-values"
)

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

func (e Experiment) Values(idx int) map[string]any {
	for i, par := range e.parameters {
		fn := e.functions[i]
		e.values[par] = fn.Next(idx, e.rng)
	}
	return e.values
}

type ParameterVariation struct {
	Parameter   string
	Type        string
	FloatParams []float64
	IntParams   []int
	BoolParams  []bool
	NoStride    bool
}

type ParameterFunction interface {
	Next(index int, rng *rand.Rand) any
	Stride() int
}

func NewParameterFunction(v ParameterVariation, stride int) (ParameterFunction, error) {
	if v.NoStride {
		stride = 1
	}

	if len(v.FloatParams) > 0 {
		if len(v.IntParams) > 0 || len(v.BoolParams) > 0 {
			return nil, fmt.Errorf("only one of FloatParams, IntParams and BoolParams may be given in a ParameterVariation")
		}

		switch v.Type {
		case VaryRandomRange:
			return &randomFloatRange{
				values: v.FloatParams,
			}, nil
		case VaryRandomValues:
			return &randomFloatValues{
				values: v.FloatParams,
			}, nil
		case VarySequenceRange:
			return &sequenceFloatRange{
				values: v.FloatParams,
				stride: stride,
			}, nil
		case VarySequenceValues:
			return &sequenceFloatValues{
				values: v.FloatParams,
				stride: stride,
			}, nil
		default:
			return nil, fmt.Errorf("invalid random variation type '%s' for float values", v.Type)
		}
	}

	if len(v.IntParams) > 0 {
		if len(v.FloatParams) > 0 || len(v.BoolParams) > 0 {
			return nil, fmt.Errorf("only one of FloatParams, IntParams and BoolParams may be given in a ParameterVariation")
		}

		switch v.Type {
		case VaryRandomRange:
			return &randomIntRange{
				values: v.IntParams,
			}, nil
		case VaryRandomValues:
			return &randomIntValues{
				values: v.IntParams,
			}, nil
		case VarySequenceRange:
			return &sequenceIntRange{
				values: v.IntParams,
				stride: stride,
			}, nil
		case VarySequenceValues:
			return &sequenceIntValues{
				values: v.IntParams,
				stride: stride,
			}, nil
		default:
			return nil, fmt.Errorf("invalid random variation type '%s' for float values", v.Type)
		}
	}

	if len(v.BoolParams) > 0 {
		if len(v.FloatParams) > 0 || len(v.IntParams) > 0 {
			return nil, fmt.Errorf("only one of FloatParams, IntParams and BoolParams may be given in a ParameterVariation")
		}

		switch v.Type {
		case VaryRandomValues:
			return &randomBoolValues{
				values: v.BoolParams,
			}, nil
		case VarySequenceValues:
			return &sequenceBoolValues{
				values: v.BoolParams,
				stride: stride,
			}, nil
		default:
			return nil, fmt.Errorf("invalid random variation type '%s' for float values", v.Type)
		}
	}

	return nil, fmt.Errorf("exactly one of FloatParams, IntParams and BoolParams must be given in a ParameterVariation")
}

type randomFloatRange struct {
	values []float64
}

func (r *randomFloatRange) Next(index int, rng *rand.Rand) any {
	return rng.Float64()*(r.values[1]-r.values[0]) + r.values[0]
}

func (r *randomFloatRange) Stride() int { return 1 }

type randomFloatValues struct {
	values []float64
}

func (r *randomFloatValues) Next(index int, rng *rand.Rand) any {
	return r.values[rng.Intn(len(r.values))]
}

func (r *randomFloatValues) Stride() int { return 1 }

type sequenceFloatRange struct {
	values []float64
	stride int
}

func (s *sequenceFloatRange) Next(index int, rng *rand.Rand) any {
	numSteps := int(s.values[2]) - 1
	idx := (index / s.stride) % (numSteps + 1)
	step := (s.values[1] - s.values[0]) / float64(numSteps)
	return s.values[0] + float64(idx)*step
}

func (s *sequenceFloatRange) Stride() int { return int(s.values[2]) }

type sequenceFloatValues struct {
	values []float64
	stride int
}

func (s *sequenceFloatValues) Next(index int, rng *rand.Rand) any {
	numValues := len(s.values)
	idx := (index / s.stride) % numValues
	return s.values[idx]
}

func (s *sequenceFloatValues) Stride() int { return len(s.values) }

type randomIntRange struct {
	values []int
}

func (r *randomIntRange) Next(index int, rng *rand.Rand) any {
	return rng.Intn(r.values[1]-r.values[0]) + r.values[0]
}

func (r *randomIntRange) Stride() int { return 1 }

type randomIntValues struct {
	values []int
}

func (r *randomIntValues) Next(index int, rng *rand.Rand) any {
	return r.values[rng.Intn(len(r.values))]
}

func (r *randomIntValues) Stride() int { return 1 }

type sequenceIntRange struct {
	values []int
	stride int
}

func (s *sequenceIntRange) Next(index int, rng *rand.Rand) any {
	numValues := int(s.values[2])
	idx := (index / s.stride) % numValues
	step := s.values[1]
	return s.values[0] + idx*step
}

func (s *sequenceIntRange) Stride() int { return int(s.values[2]) }

type sequenceIntValues struct {
	values []int
	stride int
}

func (s *sequenceIntValues) Next(index int, rng *rand.Rand) any {
	numValues := len(s.values)
	idx := (index / s.stride) % numValues
	return s.values[idx]
}

func (s *sequenceIntValues) Stride() int { return len(s.values) }

type randomBoolValues struct {
	values []bool
}

func (r *randomBoolValues) Next(index int, rng *rand.Rand) any {
	return r.values[rng.Intn(len(r.values))]
}

func (r *randomBoolValues) Stride() int { return 1 }

type sequenceBoolValues struct {
	values []bool
	stride int
}

func (s *sequenceBoolValues) Next(index int, rng *rand.Rand) any {
	numValues := len(s.values)
	idx := (index / s.stride) % numValues
	return s.values[idx]
}

func (s *sequenceBoolValues) Stride() int { return len(s.values) }
