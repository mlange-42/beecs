package sys

import (
	"testing"

	"github.com/mlange-42/arche/ecs"
	"github.com/stretchr/testify/assert"
)

func TestAgeCohorts(t *testing.T) {
	world := ecs.NewWorld()

	init := InitCohorts{
		EggTimeWorker:       2,
		LarvaeTimeWorker:    3,
		PupaeTimeWorker:     4,
		MaxInHiveTimeWorker: 5,
		EggTimeDrone:        3,
		LarvaeTimeDrone:     4,
		PupaeTimeDrone:      5,
		LifespanDrone:       6,
	}
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
	assert.Equal(t, []int{0, 0, 0, 0, 10}, init.inHive.Workers)

	assert.Equal(t, []int{0, 0, 0}, init.eggs.Drones)
	assert.Equal(t, []int{0, 0, 0, 0}, init.larvae.Drones)
	assert.Equal(t, []int{0, 0, 0, 0, 0}, init.pupae.Drones)
	assert.Equal(t, []int{0, 20, 0, 0, 0, 0}, init.inHive.Drones)
}
