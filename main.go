package main

import (
	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche-model/system"
	"github.com/mlange-42/beecs/model/sys"
)

func main() {
	m := model.New()

	m.AddSystem(&sys.InitCohorts{
		EggTimeWorker:       3,
		LarvaeTimeWorker:    6,
		PupaeTimeWorker:     12,
		MaxInHiveTimeWorker: 50,

		EggTimeDrone:    3,
		LarvaeTimeDrone: 7,
		PupaeTimeDrone:  14,
		LifespanDrone:   37,
	})

	m.AddSystem(&sys.MortalityCohorts{
		EggMortalityWorker:    0.03,
		LarvaeMortalityWorker: 0.01,
		PupaeMortalityWorker:  0.001,
		InHiveMortalityWorker: 0.004,

		EggMortalityDrone:    0.064,
		LarvaeMortalityDrone: 0.044,
		PupaeMortalityDrone:  0.005,
		InHiveMortalityDrone: 0.05,
	})

	m.AddSystem(&sys.AgeCohorts{})

	m.AddSystem(&sys.EggLaying{
		MaxEggsPerDay: 1600,
	})

	m.AddSystem(&sys.Time{
		TicksPerDay: 6 * 24,
	})

	m.AddSystem(&system.FixedTermination{Steps: 1000})

	m.Run()
}
