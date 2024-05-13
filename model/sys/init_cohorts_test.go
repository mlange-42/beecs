package sys

import (
	"testing"

	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/beecs/model/params"
	"github.com/stretchr/testify/assert"
)

func TestInitCohorts(t *testing.T) {
	world := ecs.NewWorld()
	ecs.AddResource(&world, &params.AgeFirstForagingParams{Max: 5})
	ecs.AddResource(&world, &params.WorkerDevelopment{
		EggTime:     2,
		LarvaeTime:  3,
		PupaeTime:   4,
		MaxLifespan: 100,
	})
	ecs.AddResource(&world, &params.DroneDevelopment{
		EggTime:     3,
		LarvaeTime:  4,
		PupaeTime:   5,
		MaxLifespan: 6,
	})

	s := InitCohorts{}
	s.Initialize(&world)

	assert.Equal(t, []int{0, 0}, s.eggs.Workers)
	assert.Equal(t, []int{0, 0, 0}, s.larvae.Workers)
	assert.Equal(t, []int{0, 0, 0, 0}, s.pupae.Workers)
	assert.Equal(t, []int{0, 0, 0, 0, 0, 0}, s.inHive.Workers)

	assert.Equal(t, []int{0, 0, 0}, s.eggs.Drones)
	assert.Equal(t, []int{0, 0, 0, 0}, s.larvae.Drones)
	assert.Equal(t, []int{0, 0, 0, 0, 0}, s.pupae.Drones)
	assert.Equal(t, []int{0, 0, 0, 0, 0, 0}, s.inHive.Drones)
}
