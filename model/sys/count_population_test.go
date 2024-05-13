package sys

import (
	"testing"

	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/beecs/model/globals"
	"github.com/mlange-42/beecs/model/params"
	"github.com/stretchr/testify/assert"
)

func TestCountPopulation(t *testing.T) {
	world := ecs.NewWorld()

	ecs.AddResource(&world, &params.ForagerParams{SquadronSize: 100})
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

	stats := globals.PopulationStats{}
	ecs.AddResource(&world, &stats)

	fac := globals.NewForagerFactory(&world)

	init := InitCohorts{}
	init.Initialize(&world)

	pop := CountPopulation{}
	pop.Initialize(&world)

	assert.Equal(t, globals.PopulationStats{}, stats)

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
	assert.Equal(t, globals.PopulationStats{
		WorkerEggs:      10,
		WorkerLarvae:    30,
		WorkerPupae:     50,
		WorkersInHive:   70,
		WorkersForagers: 900,

		DroneEggs:    20,
		DroneLarvae:  40,
		DronePupae:   60,
		DronesInHive: 80,

		TotalBrood:      90 + 120,
		TotalAdults:     900 + 70 + 80,
		TotalPopulation: 900 + 70 + 80 + 90 + 120,
	}, stats)
}
