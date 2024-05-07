package sys

import (
	"testing"

	"github.com/mlange-42/arche-model/resource"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
	"github.com/mlange-42/beecs/model/comp"
	"github.com/mlange-42/beecs/model/res"
	"github.com/stretchr/testify/assert"
)

func TestTransitionForagers(t *testing.T) {
	world := ecs.NewWorld()

	ecs.AddResource(&world, &resource.Tick{})
	ecs.AddResource(&world, &res.Params{SquadronSize: 100})
	ecs.AddResource(&world, &res.AgeFirstForaging{Current: 3, Max: 5})

	fac := res.NewForagerFactory(&world)
	ecs.AddResource(&world, &fac)

	init := InitCohorts{
		EggTimeWorker:    2,
		LarvaeTimeWorker: 3,
		PupaeTimeWorker:  4,
		EggTimeDrone:     3,
		LarvaeTimeDrone:  4,
		PupaeTimeDrone:   5,
		LifespanDrone:    6,
	}
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
