package experiment

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/rand"
)

func TestExperiment(t *testing.T) {
	rng := rand.New(nil)

	vars := []ParameterVariation{
		{
			Parameter:   "a",
			Type:        VarySequenceRange,
			FloatParams: []float64{0, 1, 11},
		},
		{
			Parameter: "b",
			Type:      VarySequenceRange,
			IntParams: []int{1, 1, 3},
		},
	}

	e, err := New(vars, rng)

	assert.Nil(t, err)

	assert.Equal(t, 33, e.ParameterSets())

	res := e.Values(0)
	assert.Equal(t, map[string]any{
		"a": 0.0,
		"b": 1,
	}, res)

	res = e.Values(5)
	assert.Equal(t, map[string]any{
		"a": 0.5,
		"b": 1,
	}, res)

	res = e.Values(10)
	assert.Equal(t, map[string]any{
		"a": 1.0,
		"b": 1,
	}, res)

	res = e.Values(11)
	assert.Equal(t, map[string]any{
		"a": 0.0,
		"b": 2,
	}, res)

	res = e.Values(12)
	assert.Equal(t, map[string]any{
		"a": 0.1,
		"b": 2,
	}, res)

	res = e.Values(32)
	assert.Equal(t, map[string]any{
		"a": 1.0,
		"b": 3,
	}, res)

	res = e.Values(33)
	assert.Equal(t, map[string]any{
		"a": 0.0,
		"b": 1,
	}, res)
}
