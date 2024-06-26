package params_test

import (
	"testing"

	"github.com/mlange-42/beecs/params"
	"github.com/stretchr/testify/assert"
)

func TestDefaultParamsFromJSON(t *testing.T) {
	p := params.Default()

	js := `{
    "Termination": {
        "MaxTicks": 3650,
        "OnExtinction": true
    }
}`

	err := p.FromJSON([]byte(js))
	assert.Nil(t, err)

	assert.Equal(t, 3650, p.Termination.MaxTicks)
}
