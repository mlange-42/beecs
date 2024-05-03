package sys

import (
	"testing"

	"github.com/mlange-42/arche-model/resource"
	"github.com/mlange-42/arche/ecs"
	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/rand"
)

func TestMortalityCohorts(t *testing.T) {
	world := ecs.NewWorld()
	rng := resource.Rand{Source: rand.NewSource(0)}
	ecs.AddResource(&world, &rng)

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

	mort := MortalityCohorts{
		EggMortalityWorker:    0.5,
		LarvaeMortalityWorker: 0.5,
		PupaeMortalityWorker:  0.5,
		InHiveMortalityWorker: 0.5,

		EggMortalityDrone:    0.5,
		LarvaeMortalityDrone: 0.5,
		PupaeMortalityDrone:  0.5,
		InHiveMortalityDrone: 0.5,
	}
	mort.Initialize(&world)

	fillCohorts(init.eggs.Workers, 10000)
	fillCohorts(init.eggs.Drones, 10000)

	fillCohorts(init.larvae.Workers, 10000)
	fillCohorts(init.larvae.Drones, 10000)

	fillCohorts(init.pupae.Workers, 10000)
	fillCohorts(init.pupae.Drones, 10000)

	fillCohorts(init.inHive.Workers, 10000)
	fillCohorts(init.inHive.Drones, 10000)

	mort.Update(&world)

	checkCohorts(t, init.eggs.Workers, 0, 10000)
	checkCohorts(t, init.eggs.Drones, 0, 10000)

	checkCohorts(t, init.larvae.Workers, 0, 10000)
	checkCohorts(t, init.larvae.Drones, 0, 10000)

	checkCohorts(t, init.pupae.Workers, 0, 10000)
	checkCohorts(t, init.pupae.Drones, 0, 10000)

	checkCohorts(t, init.inHive.Workers, 0, 10000)
	checkCohorts(t, init.inHive.Drones, 0, 10000)
}

func fillCohorts(coh []int, count int) {
	for i := range coh {
		coh[i] = count
	}
}

func checkCohorts(t *testing.T, coh []int, min int, max int) {
	for _, n := range coh {
		assert.Less(t, n, max)
		assert.Greater(t, n, min)
	}
}
