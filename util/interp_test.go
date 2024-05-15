package util_test

import (
	"testing"

	"github.com/mlange-42/beecs/util"
	"github.com/stretchr/testify/assert"
)

func TestInterpolateStep(t *testing.T) {
	values := [][2]float64{
		{4, 1},
		{8, 2},
		{12, 3},
		{13, 4},
	}
	assert.Equal(t, 4.0, util.InterpolateStep(values, 1))
	assert.Equal(t, 1.0, util.InterpolateStep(values, 4))
	assert.Equal(t, 1.0, util.InterpolateStep(values, 7))
	assert.Equal(t, 2.0, util.InterpolateStep(values, 8))
	assert.Equal(t, 2.0, util.InterpolateStep(values, 11))
	assert.Equal(t, 3.0, util.InterpolateStep(values, 12))
	assert.Equal(t, 4.0, util.InterpolateStep(values, 13))
	assert.Equal(t, 4.0, util.InterpolateStep(values, 364))
}

func TestInterpolateLinear(t *testing.T) {
	values := [][2]float64{
		{4, 1},
		{8, 2},
		{12, 3},
		{13, 4},
		{361, -1},
	}
	assert.Equal(t, 0.0, util.InterpolateLinear(values, 0))
	assert.Equal(t, 0.5, util.InterpolateLinear(values, 2))

	assert.Equal(t, 1.0, util.InterpolateLinear(values, 4))
	assert.Equal(t, 1.25, util.InterpolateLinear(values, 5))
	assert.Equal(t, 1.5, util.InterpolateLinear(values, 6))
	assert.Equal(t, 1.75, util.InterpolateLinear(values, 7))
	assert.Equal(t, 2.0, util.InterpolateLinear(values, 8))
	assert.Equal(t, 2.5, util.InterpolateLinear(values, 10))
	assert.Equal(t, 3.0, util.InterpolateLinear(values, 12))
	assert.Equal(t, 4.0, util.InterpolateLinear(values, 13))

	assert.Equal(t, -1.0, util.InterpolateLinear(values, 361))
	assert.Equal(t, -0.5, util.InterpolateLinear(values, 363))
}
