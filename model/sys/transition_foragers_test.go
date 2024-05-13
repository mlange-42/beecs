package sys

import (
	"testing"

	"github.com/mlange-42/arche-model/resource"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
	"github.com/mlange-42/beecs/model/comp"
	"github.com/mlange-42/beecs/model/globals"
	"github.com/mlange-42/beecs/model/params"
	"github.com/stretchr/testify/assert"
)

func TestTransitionForagers(t *testing.T) {
	world := ecs.NewWorld()

	ecs.AddResource(&world, &resource.Tick{})
	ecs.AddResource(&world, &params.ForagerParams{SquadronSize: 100})
	ecs.AddResource(&world, &params.AgeFirstForagingParams{Max: 5})
	ecs.AddResource(&world, &globals.AgeFirstForaging{Aff: 3})
	ecs.AddResource(&world, &params.WorkerDevelopment{
		EggTime:     2,
		LarvaeTime:  3,
		PupaeTime:   4,
		MaxLifespan: 390,
	})
	ecs.AddResource(&world, &params.DroneDevelopment{
		EggTime:     3,
		LarvaeTime:  4,
		PupaeTime:   5,
		MaxLifespan: 6,
	})

	fac := globals.NewForagerFactory(&world)
	ecs.AddResource(&world, &fac)

	init := InitCohorts{}
	init.Initialize(&world)

	transition := TransitionForagers{}
	transition.Initialize(&world)

	assert.Equal(t, 6, len(init.inHive.Workers))
	init.inHive.Workers = []int{0, 0, 2000, 1000, 100, 25}

	transition.Update(&world)

	assert.Equal(t, []int{0, 0, 2025, 0, 0, 0}, init.inHive.Workers)

	filter := *generic.NewFilter1[comp.Milage]()
	query := filter.Query(&world)
	assert.Equal(t, 11, query.Count())

	query.Close()
}
