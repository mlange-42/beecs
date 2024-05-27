package sys

import (
	"testing"

	"github.com/mlange-42/arche-model/resource"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/beecs/globals"
	"github.com/mlange-42/beecs/params"
	"github.com/stretchr/testify/assert"
)

func TestFixedTermination(t *testing.T) {
	world := ecs.NewWorld()

	ecs.AddResource(&world, &params.Termination{MaxTicks: 10})
	ecs.AddResource(&world, &globals.PopulationStats{TotalPopulation: 100})

	termRes := resource.Termination{}
	ecs.AddResource(&world, &termRes)

	term := FixedTermination{}

	term.Initialize(&world)

	for i := 0; i < 9; i++ {
		term.Update(&world)
	}
	assert.False(t, termRes.Terminate)

	term.Update(&world)
	assert.True(t, termRes.Terminate)
}

func TestExtinctionTermination(t *testing.T) {
	world := ecs.NewWorld()

	ecs.AddResource(&world, &params.Termination{OnExtinction: true})

	pop := globals.PopulationStats{TotalPopulation: 100}
	ecs.AddResource(&world, &pop)

	termRes := resource.Termination{}
	ecs.AddResource(&world, &termRes)

	term := FixedTermination{}

	term.Initialize(&world)

	for i := 0; i < 10; i++ {
		term.Update(&world)
	}
	assert.False(t, termRes.Terminate)

	pop.TotalPopulation = 0
	term.Update(&world)
	assert.True(t, termRes.Terminate)
}
