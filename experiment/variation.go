package experiment

import (
	"fmt"
	"reflect"

	"golang.org/x/exp/rand"
)

// ParameterVariation definition.
//
// Only one of the pointer fields may be non-nil.
type ParameterVariation struct {
	Parameter string

	RandomFloatRange    *RandomFloatRange
	RandomFloatValues   *RandomFloatValues
	SequenceFloatRange  *SequenceFloatRange
	SequenceFloatValues *SequenceFloatValues

	RandomIntRange    *RandomIntRange
	RandomIntValues   *RandomIntValues
	SequenceIntRange  *SequenceIntRange
	SequenceIntValues *SequenceIntValues

	RandomBoolValues   *RandomBoolValues
	SequenceBoolValues *SequenceBoolValues

	NoStride bool
}

// ParameterFunction interface for creating parameter values.
type ParameterFunction interface {
	Next(index int, rng *rand.Rand) any
	Stride() int
}

// NewParameterFunction creates a new ParameterFunction.
func NewParameterFunction(v ParameterVariation, stride int) (ParameterFunction, error) {
	if v.NoStride {
		stride = 1
	}

	nonNin := 0
	var function ParameterFunction
	for _, ptr := range []ParameterFunction{
		v.RandomBoolValues, v.RandomFloatRange, v.RandomFloatValues, v.RandomIntRange, v.RandomIntValues,
		v.SequenceBoolValues, v.SequenceFloatRange, v.SequenceFloatValues, v.SequenceIntRange, v.SequenceIntValues,
	} {
		if !reflect.ValueOf(ptr).IsNil() {
			nonNin++
			function = ptr
		}
	}
	if nonNin != 1 {
		return nil, fmt.Errorf("exactly one of RandomFloatValues, RandomIntRange, ...  must be given in a ParameterVariation")
	}
	switch f := function.(type) {
	case *SequenceFloatRange:
		f.stride = stride
	case *SequenceFloatValues:
		f.stride = stride
	case *SequenceIntRange:
		f.stride = stride
	case *SequenceIntValues:
		f.stride = stride
	case *SequenceBoolValues:
		f.stride = stride
	}

	return function, nil
}

type RandomFloatRange struct {
	Min float64
	Max float64
}

func (r *RandomFloatRange) Next(index int, rng *rand.Rand) any {
	return rng.Float64()*(r.Max-r.Min) + r.Min
}

func (r *RandomFloatRange) Stride() int { return 1 }

type RandomFloatValues struct {
	Values []float64
}

func (r *RandomFloatValues) Next(index int, rng *rand.Rand) any {
	return r.Values[rng.Intn(len(r.Values))]
}

func (r *RandomFloatValues) Stride() int { return 1 }

type SequenceFloatRange struct {
	Min    float64
	Max    float64
	Values int
	stride int
}

func (s *SequenceFloatRange) Next(index int, rng *rand.Rand) any {
	numSteps := s.Values - 1
	idx := (index / s.stride) % (numSteps + 1)
	step := (s.Max - s.Min) / float64(numSteps)
	return s.Min + float64(idx)*step
}

func (s *SequenceFloatRange) Stride() int { return s.Values }

type SequenceFloatValues struct {
	Values []float64
	stride int
}

func (s *SequenceFloatValues) Next(index int, rng *rand.Rand) any {
	numValues := len(s.Values)
	idx := (index / s.stride) % numValues
	return s.Values[idx]
}

func (s *SequenceFloatValues) Stride() int { return len(s.Values) }

type RandomIntRange struct {
	Min int
	Max int
}

func (r *RandomIntRange) Next(index int, rng *rand.Rand) any {
	return rng.Intn(r.Max-r.Min) + r.Min
}

func (r *RandomIntRange) Stride() int { return 1 }

type RandomIntValues struct {
	Values []int
}

func (r *RandomIntValues) Next(index int, rng *rand.Rand) any {
	return r.Values[rng.Intn(len(r.Values))]
}

func (r *RandomIntValues) Stride() int { return 1 }

type SequenceIntRange struct {
	Min    int
	Step   int
	Values int
	stride int
}

func (s *SequenceIntRange) Next(index int, rng *rand.Rand) any {
	numValues := int(s.Values)
	idx := (index / s.stride) % numValues
	return s.Min + idx*s.Step
}

func (s *SequenceIntRange) Stride() int { return int(s.Values) }

type SequenceIntValues struct {
	Values []int
	stride int
}

func (s *SequenceIntValues) Next(index int, rng *rand.Rand) any {
	numValues := len(s.Values)
	idx := (index / s.stride) % numValues
	return s.Values[idx]
}

func (s *SequenceIntValues) Stride() int { return len(s.Values) }

type RandomBoolValues struct {
	Values []bool
}

func (r *RandomBoolValues) Next(index int, rng *rand.Rand) any {
	return r.Values[rng.Intn(len(r.Values))]
}

func (r *RandomBoolValues) Stride() int { return 1 }

type SequenceBoolValues struct {
	Values []bool
	stride int
}

func (s *SequenceBoolValues) Next(index int, rng *rand.Rand) any {
	numValues := len(s.Values)
	idx := (index / s.stride) % numValues
	return s.Values[idx]
}

func (s *SequenceBoolValues) Stride() int { return len(s.Values) }
