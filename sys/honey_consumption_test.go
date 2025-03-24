package sys

import (
	"math"
	"testing"

	"github.com/mlange-42/ark/ecs"
	"github.com/mlange-42/beecs/globals"
	"github.com/mlange-42/beecs/params"
	"github.com/stretchr/testify/assert"
)

func TestHoneyConsumption(t *testing.T) {
	world := ecs.NewWorld()

	ecs.AddResource(&world, &params.Foragers{SquadronSize: 10})
	ecs.AddResource(&world, &params.AgeFirstForaging{Max: 5})
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

	ecs.AddResource(&world, &params.HoneyNeeds{
		WorkerResting:    3,
		WorkerNurse:      5,
		WorkerLarvaTotal: 21, // -> 7/d
		DroneLarva:       7,
		Drone:            9,
	})

	ecs.AddResource(&world, &params.Stores{
		IdealPollenStoreDays: 7,
		MinIdealPollenStore:  250.0,
		MaxHoneyStoreKg:      50.0,
		DecentHoneyPerWorker: 1.5,
		ProteinStoreNurse:    7,
	})

	stores := globals.Stores{
		Honey:  100_000,
		Pollen: 100_000,
	}
	ecs.AddResource(&world, &stores)

	ecs.AddResource(&world, &params.EnergyContent{
		Honey:   12.78,
		Sucrose: 0.00582,
	})
	ecs.AddResource(&world, &params.Nursing{
		MaxBroodNurseRatio: 3.0,
	})
	ecs.AddResource(&world, &globals.ConsumptionStats{})

	stats := globals.PopulationStats{}
	ecs.AddResource(&world, &stats)

	fac := globals.NewForagerFactory(&world)

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
