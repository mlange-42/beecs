package sys

import (
	"testing"

	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/beecs/model/res"
	"github.com/stretchr/testify/assert"
)

func TestAgeCohorts(t *testing.T) {
	world := ecs.NewWorld()

	ecs.AddResource(&world, &res.AgeFirstForagingParams{Max: 5})
	ecs.AddResource(&world, &res.WorkerDevelopment{
		EggTime:     2,
		LarvaeTime:  3,
		PupaeTime:   4,
		MaxLifespan: 100,
	})
	ecs.AddResource(&world, &res.DroneDevelopment{
		EggTime:     3,
		LarvaeTime:  4,
		PupaeTime:   5,
		MaxLifespan: 6,
	})

	init := InitCohorts{}
	init.Initialize(&world)

	age := AgeCohorts{}
	age.Initialize(&world)

	init.eggs.Workers[0] = 10
	init.eggs.Drones[0] = 20

	for i := 0; i < 13; i++ {
		age.Update(&world)
	}

	assert.Equal(t, []int{0, 0}, init.eggs.Workers)
	assert.Equal(t, []int{0, 0, 0}, init.larvae.Workers)
	assert.Equal(t, []int{0, 0, 0, 0}, init.pupae.Workers)
	assert.Equal(t, []int{0, 0, 0, 0, 10, 0}, init.inHive.Workers)

	assert.Equal(t, []int{0, 0, 0}, init.eggs.Drones)
	assert.Equal(t, []int{0, 0, 0, 0}, init.larvae.Drones)
	assert.Equal(t, []int{0, 0, 0, 0, 0}, init.pupae.Drones)
	assert.Equal(t, []int{0, 20, 0, 0, 0, 0}, init.inHive.Drones)
}
