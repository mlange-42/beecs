package experiment

import (
	"fmt"
	"math/rand/v2"
	"reflect"
)

// ParameterVariation definition.
//
// Only one of the pointer fields may be non-nil.
type ParameterVariation struct {
	Parameter string

	RandomFloatRange    *RandomFloatRange    `json:",omitempty"`
	RandomFloatValues   *RandomFloatValues   `json:",omitempty"`
	SequenceFloatRange  *SequenceFloatRange  `json:",omitempty"`
	SequenceFloatValues *SequenceFloatValues `json:",omitempty"`

	RandomIntRange    *RandomIntRange    `json:",omitempty"`
	RandomIntValues   *RandomIntValues   `json:",omitempty"`
	SequenceIntRange  *SequenceIntRange  `json:",omitempty"`
	SequenceIntValues *SequenceIntValues `json:",omitempty"`

	RandomBoolValues   *RandomBoolValues   `json:",omitempty"`
	SequenceBoolValues *SequenceBoolValues `json:",omitempty"`

	RandomStringValues   *RandomStringValues   `json:",omitempty"`
	SequenceStringValues *SequenceStringValues `json:",omitempty"`

	NoStride bool `json:",omitempty"`
}

// ParameterFunction interface for creating parameter values.
type ParameterFunction interface {
	Next(index int, rng *rand.Rand) any // Next returns the parameter value for the given run index.
	Stride() int                        // Stride returns the stride of the parameter function.
	SetRepetitions(rep int)             // SetRepetitions sets the number of repetitions.
}

// NewParameterFunction creates a new ParameterFunction.
func NewParameterFunction(v ParameterVariation, stride int) (ParameterFunction, error) {
	if v.NoStride {
		stride = 1
	}

	nonNin := 0
	var function ParameterFunction
	for _, ptr := range []ParameterFunction{
		v.RandomBoolValues, v.RandomFloatRange, v.RandomFloatValues, v.RandomIntRange, v.RandomIntValues, v.RandomStringValues,
		v.SequenceBoolValues, v.SequenceFloatRange, v.SequenceFloatValues, v.SequenceIntRange, v.SequenceIntValues, v.SequenceStringValues,
	} {
		if !reflect.ValueOf(ptr).IsNil() {
			nonNin++
			function = ptr
		}
	}
	if nonNin != 1 {
		return nil, fmt.Errorf("exactly one of RandomFloatValues, RandomIntRange, ...  must be given in a ParameterVariation")
	}

	function.SetRepetitions(stride)

	return function, nil
}

// RandomFloatRange generates random float values in the given range.
type RandomFloatRange struct {
	Min float64 // Lower limit of the range.
	Max float64 // Upper limit of the range.
}

// Next returns the parameter value for the given run index.
func (r *RandomFloatRange) Next(index int, rng *rand.Rand) any {
	return rng.Float64()*(r.Max-r.Min) + r.Min
}

// Stride returns the stride of the parameter function.
func (r *RandomFloatRange) Stride() int { return 1 }

// SetRepetitions sets the number of repetitions.
func (r *RandomFloatRange) SetRepetitions(rep int) {}

// RandomFloatValues generates float values by randomly drawing from the provided values.
type RandomFloatValues struct {
	Values []float64 // Values to draw from.
}

// Next returns the parameter value for the given run index.
func (r *RandomFloatValues) Next(index int, rng *rand.Rand) any {
	return r.Values[rng.IntN(len(r.Values))]
}

// Stride returns the stride of the parameter function.
func (r *RandomFloatValues) Stride() int { return 1 }

// SetRepetitions sets the number of repetitions.
func (r *RandomFloatValues) SetRepetitions(rep int) {}

// SequenceFloatRange generates float values by iterating a range.
type SequenceFloatRange struct {
	Min         float64 // Lowest value.
	Max         float64 // Highest value.
	Values      int     // Number of values.
	repetitions int     // Number of repetitions from strides of previous parameter functions.
}

// Next returns the parameter value for the given run index.
func (s *SequenceFloatRange) Next(index int, rng *rand.Rand) any {
	numSteps := s.Values - 1
	idx := (index / s.repetitions) % (numSteps + 1)
	step := (s.Max - s.Min) / float64(numSteps)
	return s.Min + float64(idx)*step
}

// Stride returns the stride of the parameter function.
func (s *SequenceFloatRange) Stride() int { return s.Values }

// SetRepetitions sets the number of repetitions.
func (s *SequenceFloatRange) SetRepetitions(rep int) {
	s.repetitions = rep
}

// SequenceFloatValues generates float values by iterating the given values.
type SequenceFloatValues struct {
	Values      []float64 // Values to iterate.
	repetitions int       // Number of repetitions from strides of previous parameter functions.
}

// Next returns the parameter value for the given run index.
func (s *SequenceFloatValues) Next(index int, rng *rand.Rand) any {
	numValues := len(s.Values)
	idx := (index / s.repetitions) % numValues
	return s.Values[idx]
}

// Stride returns the stride of the parameter function.
func (s *SequenceFloatValues) Stride() int { return len(s.Values) }

// SetRepetitions sets the number of repetitions.
func (s *SequenceFloatValues) SetRepetitions(rep int) {
	s.repetitions = rep
}

// RandomIntRange generates random int values in the given range.
type RandomIntRange struct {
	Min int // Lower limit of the range.
	Max int // Upper limit of the range (exclusive).
}

// Next returns the parameter value for the given run index.
func (r *RandomIntRange) Next(index int, rng *rand.Rand) any {
	return rng.IntN(r.Max-r.Min) + r.Min
}

// Stride returns the stride of the parameter function.
func (r *RandomIntRange) Stride() int { return 1 }

// SetRepetitions sets the number of repetitions.
func (r *RandomIntRange) SetRepetitions(rep int) {}

// RandomIntValues generates int values by randomly drawing from the provided values.
type RandomIntValues struct {
	Values []int // Values to draw from.
}

// Next returns the parameter value for the given run index.
func (r *RandomIntValues) Next(index int, rng *rand.Rand) any {
	return r.Values[rng.IntN(len(r.Values))]
}

// Stride returns the stride of the parameter function.
func (r *RandomIntValues) Stride() int { return 1 }

// SetRepetitions sets the number of repetitions.
func (r *RandomIntValues) SetRepetitions(rep int) {}

// SequenceIntRange generates int values by iterating a range.
type SequenceIntRange struct {
	Min         int // Lowest value.
	Step        int // Step size.
	Values      int // Number of values/steps.
	repetitions int // Number of repetitions from strides of previous parameter functions.
}

// Next returns the parameter value for the given run index.
func (s *SequenceIntRange) Next(index int, rng *rand.Rand) any {
	numValues := int(s.Values)
	idx := (index / s.repetitions) % numValues
	return s.Min + idx*s.Step
}

// Stride returns the stride of the parameter function.
func (s *SequenceIntRange) Stride() int { return int(s.Values) }

// SetRepetitions sets the number of repetitions.
func (s *SequenceIntRange) SetRepetitions(rep int) {
	s.repetitions = rep
}

// SequenceIntValues generates int values by iterating the given values.
type SequenceIntValues struct {
	Values      []int // Values to iterate.
	repetitions int   // Number of repetitions from strides of previous parameter functions.
}

// Next returns the parameter value for the given run index.
func (s *SequenceIntValues) Next(index int, rng *rand.Rand) any {
	numValues := len(s.Values)
	idx := (index / s.repetitions) % numValues
	return s.Values[idx]
}

// Stride returns the stride of the parameter function.
func (s *SequenceIntValues) Stride() int { return len(s.Values) }

// SetRepetitions sets the number of repetitions.
func (s *SequenceIntValues) SetRepetitions(rep int) {
	s.repetitions = rep
}

// RandomBoolValues generates bool values by randomly drawing from the provided values.
type RandomBoolValues struct {
	Values []bool // Values to draw from.
}

// Next returns the parameter value for the given run index.
func (r *RandomBoolValues) Next(index int, rng *rand.Rand) any {
	return r.Values[rng.IntN(len(r.Values))]
}

// Stride returns the stride of the parameter function.
func (r *RandomBoolValues) Stride() int { return 1 }

// SetRepetitions sets the number of repetitions.
func (r *RandomBoolValues) SetRepetitions(rep int) {}

// SequenceBoolValues generates bool values by iterating the given values.
type SequenceBoolValues struct {
	Values      []bool // Values to iterate.
	repetitions int    // Number of repetitions from strides of previous parameter functions.
}

// Next returns the parameter value for the given run index.
func (s *SequenceBoolValues) Next(index int, rng *rand.Rand) any {
	numValues := len(s.Values)
	idx := (index / s.repetitions) % numValues
	return s.Values[idx]
}

// Stride returns the stride of the parameter function.
func (s *SequenceBoolValues) Stride() int { return len(s.Values) }

// SetRepetitions sets the number of repetitions.
func (s *SequenceBoolValues) SetRepetitions(rep int) {
	s.repetitions = rep
}

// RandomStringValues generates string values by randomly drawing from the provided values.
type RandomStringValues struct {
	Values []string // Values to draw from.
}

// Next returns the parameter value for the given run index.
func (r *RandomStringValues) Next(index int, rng *rand.Rand) any {
	return r.Values[rng.IntN(len(r.Values))]
}

// Stride returns the stride of the parameter function.
func (r *RandomStringValues) Stride() int { return 1 }

// SetRepetitions sets the number of repetitions.
func (r *RandomStringValues) SetRepetitions(rep int) {}

// SequenceStringValues generates string values by iterating the given values.
type SequenceStringValues struct {
	Values      []string // Values to iterate.
	repetitions int      // Number of repetitions from strides of previous parameter functions.
}

// Next returns the parameter value for the given run index.
func (s *SequenceStringValues) Next(index int, rng *rand.Rand) any {
	numValues := len(s.Values)
	idx := (index / s.repetitions) % numValues
	return s.Values[idx]
}

// Stride returns the stride of the parameter function.
func (s *SequenceStringValues) Stride() int { return len(s.Values) }

// SetRepetitions sets the number of repetitions.
func (s *SequenceStringValues) SetRepetitions(rep int) {
	s.repetitions = rep
}
