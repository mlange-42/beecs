package experiment

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/rand"
)

func TestExperiment(t *testing.T) {
	rng := rand.New(rand.NewSource(0))

	vars := []ParameterVariation{
		{
			Parameter: "a",
			SequenceFloatRange: &SequenceFloatRange{
				Min:    0,
				Max:    1,
				Values: 11,
			},
		},
		{
			Parameter: "b",
			SequenceIntRange: &SequenceIntRange{
				Min:    1,
				Step:   1,
				Values: 3,
			},
		},
	}

	e, err := New(vars, rng)

	assert.Nil(t, err)

	assert.Equal(t, 33, e.ParameterSets())

	res := e.Values(0)
	assert.Equal(t, []ParameterValue{
		{Parameter: "a", Value: 0.0},
		{Parameter: "b", Value: 1},
	}, res)

	res = e.Values(5)
	assert.Equal(t, []ParameterValue{
		{Parameter: "a", Value: 0.5},
		{Parameter: "b", Value: 1},
	}, res)

	res = e.Values(10)
	assert.Equal(t, []ParameterValue{
		{Parameter: "a", Value: 1.0},
		{Parameter: "b", Value: 1},
	}, res)

	res = e.Values(11)
	assert.Equal(t, []ParameterValue{
		{Parameter: "a", Value: 0.0},
		{Parameter: "b", Value: 2},
	}, res)

	res = e.Values(12)
	assert.Equal(t, []ParameterValue{
		{Parameter: "a", Value: 0.1},
		{Parameter: "b", Value: 2},
	}, res)

	res = e.Values(32)
	assert.Equal(t, []ParameterValue{
		{Parameter: "a", Value: 1.0},
		{Parameter: "b", Value: 3},
	}, res)

	res = e.Values(33)
	assert.Equal(t, []ParameterValue{
		{Parameter: "a", Value: 0.0},
		{Parameter: "b", Value: 1},
	}, res)
}
