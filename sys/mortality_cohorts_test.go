package sys

import (
	"math/rand/v2"
	"testing"

	"github.com/mlange-42/ark-tools/resource"
	"github.com/mlange-42/ark/ecs"
	"github.com/mlange-42/beecs/params"
	"github.com/stretchr/testify/assert"
)

func TestMortalityCohorts(t *testing.T) {
	world := ecs.NewWorld()

	ecs.AddResource(&world, &resource.Rand{Source: rand.NewPCG(0, 0)})
	ecs.AddResource(&world, &params.AgeFirstForaging{Max: 5})
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
	ecs.AddResource(&world, &params.WorkerMortality{
		Eggs:      0.5,
		Larvae:    0.5,
		Pupae:     0.5,
		InHive:    0.5,
		MaxMilage: 200,
	})
	ecs.AddResource(&world, &params.DroneMortality{
		Eggs:   0.5,
		Larvae: 0.5,
		Pupae:  0.5,
		InHive: 0.5,
	})

	init := InitCohorts{}
	init.Initialize(&world)

	mort := MortalityCohorts{}
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
