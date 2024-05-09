package sys

import (
	"math"
	"testing"

	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/beecs/model/res"
	"github.com/stretchr/testify/assert"
)

func TestHoneyConsumption(t *testing.T) {
	world := ecs.NewWorld()

	ecs.AddResource(&world, &res.Params{SquadronSize: 10})
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

	ecs.AddResource(&world, &res.HoneyNeeds{
		WorkerResting:    3,
		WorkerNurse:      5,
		WorkerLarvaTotal: 21, // -> 7/d
		DroneLarva:       7,
		Drone:            9,
	})

	stores := res.Stores{
		Honey:  100_000,
		Pollen: 100_000,
	}
	ecs.AddResource(&world, &stores)

	ecs.AddResource(&world, &res.EnergyParams{
		EnergyHoney:   12.78,
		EnergyScurose: 0.00582,
	})
	ecs.AddResource(&world, &res.NurseParams{
		MaxBroodNurseRatio: 3.0,
	})
	ecs.AddResource(&world, &res.ConsumptionStats{})

	stats := res.PopulationStats{}
	ecs.AddResource(&world, &stats)

	fac := res.NewForagerFactory(&world)

	init := InitCohorts{}
	init.Initialize(&world)

	pop := CountPopulation{}
	pop.Initialize(&world)

	cons := HoneyConsumption{}
	cons.Initialize(&world)

	init.eggs.Workers[1] = 10
	init.eggs.Drones[1] = 20

	init.larvae.Workers[1] = 30
	init.larvae.Drones[1] = 40

	init.pupae.Workers[1] = 50
	init.pupae.Drones[1] = 60

	init.inHive.Workers[1] = 70
	init.inHive.Drones[1] = 80

	fac.CreateSquadrons(9, 0)

	pop.Update(&world)
	cons.Update(&world)

	expected := float64(30*7+40*7+80*9+160*3+(210/3)*2) * 0.001 * 12.78
	actual := 100_000.0 - stores.Honey
	assert.Less(t, math.Abs(expected-actual), 0.0001)
	assert.Greater(t, stores.DecentHoney, 0.0)
}
