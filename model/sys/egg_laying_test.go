package sys

import (
	"testing"

	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/beecs/model/res"
	"github.com/stretchr/testify/assert"
)

func TestEggLaying(t *testing.T) {
	world := ecs.NewWorld()

	ecs.AddResource(&world, &res.AgeFirstForagingParams{Max: 1})

	ecs.AddResource(&world, &res.WorkerDevelopment{
		EggTime: 1,
	})
	ecs.AddResource(&world, &res.DroneDevelopment{
		EggTime: 1,
	})

	init := InitCohorts{}
	init.Initialize(&world)

	lay := EggLaying{
		MaxEggsPerDay:       100,
		DroneEggsProportion: 0.2,
	}
	lay.Initialize(&world)

	lay.Update(&world)

	assert.Equal(t, []int{100}, init.eggs.Workers)
	assert.Equal(t, []int{20}, init.eggs.Drones)
}
