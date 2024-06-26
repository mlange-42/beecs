package params_test

import (
	"reflect"
	"testing"

	"github.com/mlange-42/beecs/params"
	"github.com/mlange-42/beecs/registry"
	"github.com/stretchr/testify/assert"
)

type CustomParam struct {
	Value int
}

func TestCustomParamsFromJSON(t *testing.T) {
	registry.RegisterResource[CustomParam]()

	p := params.CustomParams{}

	js := `{
    "Parameters": {
        "Termination": {
            "MaxTicks": 3650,
            "OnExtinction": true
        }
    },
	"Custom": {
		"params_test.CustomParam": {
			"Value": 100
		}
	}
}`

	err := p.FromJSON([]byte(js))
	assert.Nil(t, err)

	assert.Equal(t, 3650, p.Parameters.Termination.MaxTicks)

	obj, ok := p.Custom[reflect.TypeOf(CustomParam{})]
	assert.True(t, ok)

	par, ok := obj.(*CustomParam)
	assert.True(t, ok)

	assert.Equal(t, 100, par.Value)
}
