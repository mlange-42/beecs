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

	m.AddSystem(&sys.AgeCohorts{})

	m.AddSystem(&system.FixedTermination{Steps: 1000})

	m.Run()
}
