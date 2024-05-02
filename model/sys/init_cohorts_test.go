package sys

import (
	"testing"

	"github.com/mlange-42/arche/ecs"
	"github.com/stretchr/testify/assert"
)

func TestInitCohorts(t *testing.T) {
	world := ecs.NewWorld()

	s := InitCohorts{
		EggTimeWorker:       2,
		LarvaeTimeWorker:    3,
		PupaeTimeWorker:     4,
		MaxInHiveTimeWorker: 5,
		EggTimeDrone:        3,
		LarvaeTimeDrone:     4,
		PupaeTimeDrone:      5,
		LifespanDrone:       6,
	}
	s.Initialize(&world)

	assert.Equal(t, []int{0, 0}, s.eggs.Workers)
	assert.Equal(t, []int{0, 0, 0}, s.larvae.Workers)
	assert.Equal(t, []int{0, 0, 0, 0}, s.pupae.Workers)
	assert.Equal(t, []int{0, 0, 0, 0, 0}, s.inHive.Workers)

	assert.Equal(t, []int{0, 0, 0}, s.eggs.Drones)
	assert.Equal(t, []int{0, 0, 0, 0}, s.larvae.Drones)
	assert.Equal(t, []int{0, 0, 0, 0, 0}, s.pupae.Drones)
	assert.Equal(t, []int{0, 0, 0, 0, 0, 0}, s.inHive.Drones)
}
