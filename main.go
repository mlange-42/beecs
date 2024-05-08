package main

import (
	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche-model/system"
	"github.com/mlange-42/arche-pixel/plot"
	"github.com/mlange-42/arche-pixel/window"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/beecs/model/comp"
	"github.com/mlange-42/beecs/model/obs"
	"github.com/mlange-42/beecs/model/res"
	"github.com/mlange-42/beecs/model/sys"
)

func main() {
	m := model.New()
	m.TPS = 30
	m.FPS = 30

	// Resources

	params := res.Params{
		SquadronSize: 100,
	}
	ecs.AddResource(&m.World, &params)

	workerDev := res.WorkerDevelopment{
		EggTime:     3,
		LarvaeTime:  6,
		PupaeTime:   12,
		MaxLifespan: 390,
	}
	ecs.AddResource(&m.World, &workerDev)

	droneDev := res.DroneDevelopment{
		EggTime:     3,
		LarvaeTime:  7,
		PupaeTime:   14,
		MaxLifespan: 37,
	}
	ecs.AddResource(&m.World, &droneDev)

	workerMort := res.WorkerMortality{
		Eggs:      0.03,
		Larvae:    0.01,
		Pupae:     0.001,
		InHive:    0.004,
		MaxMilage: 200,
	}
	ecs.AddResource(&m.World, &workerMort)

	droneMort := res.DroneMortality{
		Eggs:   0.064,
		Larvae: 0.044,
		Pupae:  0.005,
		InHive: 0.05,
	}
	ecs.AddResource(&m.World, &droneMort)

	aff := res.AgeFirstForagingParams{
		Base: 21,
		Min:  7,
		Max:  50,
	}
	ecs.AddResource(&m.World, &aff)

	forageParams := res.ForagingParams{
		ProbBase:      0.01,
		ProbHigh:      0.05,
		ProbEmergency: 0.2,

		FlightVelocity:              6.5,           // [m/s]
		SearchLength:                6.5 * 60 * 17, // [s] (17 min)
		MaxProportionPollenForagers: 0.8,

		EnergyOnFlower:  0.2,
		MortalityPerSec: 0.00001,
		FlightCostPerM:  0.000006, //[kJ/m]

		NectarLoad:           50,    // [muL]
		PollenLoad:           0.015, // [g]
		TimeNectarGathering:  1200,
		TimePollenGathering:  600,
		ConstantHandlingTime: false,
		StopProbability:      0.3,
		AbandonPollenPerSec:  0.00002,
		MaxKmPerDay:          7299, // ???
	}
	ecs.AddResource(&m.World, &forageParams)

	danceParams := res.DanceParams{
		Slope:       1.16,
		Intercept:   0.0,
		MaxCircuits: 117,
	}
	ecs.AddResource(&m.World, &danceParams)

	energy := res.EnergyParams{
		EnergyHoney:   12.78,
		EnergyScurose: 0.00582,
	}
	ecs.AddResource(&m.World, &energy)

	storeParams := res.StoreParams{
		IdealPollenStoreDays: 7,
		MinIdealPollenStore:  250.0,
		MaxHoneyStoreKg:      50.0,
	}
	ecs.AddResource(&m.World, &storeParams)

	honeyNeeds := res.HoneyNeeds{
		WorkerResting:    11.0,
		WorkerNurse:      53.42,
		WorkerLarvaTotal: 65.4,
		DroneLarva:       19.2,
		Drone:            10.0,
	}
	ecs.AddResource(&m.World, &honeyNeeds)

	pollenNeeds := res.PollenNeeds{
		WorkerLarvaTotal: 142.0,
		DroneLarva:       50.0,
		Worker:           1.5,
		Drone:            2.0,
	}
	ecs.AddResource(&m.World, &pollenNeeds)

	nurseParams := res.NurseParams{
		MaxBroodNurseRatio:         3.0,
		ForagerNursingContribution: 0.2,
	}
	ecs.AddResource(&m.World, &nurseParams)

	factory := res.NewForagerFactory(&m.World)
	ecs.AddResource(&m.World, &factory)

	stores := res.Stores{}
	ecs.AddResource(&m.World, &stores)

	stats := res.PopulationStats{}
	ecs.AddResource(&m.World, &stats)

	// Initialization

	m.AddSystem(&sys.InitCohorts{})

	m.AddSystem(&sys.InitPopulation{
		InitialCount: 10_000,
		MinAge:       100,
		MaxAge:       160,
		MinMilage:    0,
		MaxMilage:    200,
	})

	m.AddSystem(&sys.InitPatchesList{
		Patches: []comp.PatchConfig{
			{
				Nectar:               20,
				NectarConcentration:  1.5,
				Pollen:               1,
				DistToColony:         1500,
				DetectionProbability: 0.2,
			},
			{
				Nectar:               20,
				NectarConcentration:  1.5,
				Pollen:               1,
				DistToColony:         500,
				DetectionProbability: 0.2,
			},
		},
	})

	// Sub-models

	m.AddSystem(&sys.CalcAff{})
	m.AddSystem(&sys.CalcForagingPeriod{})
	m.AddSystem(&sys.ReplenishPatches{})

	m.AddSystem(&sys.MortalityCohorts{})
	m.AddSystem(&sys.MortalityForagers{})

	m.AddSystem(&sys.AgeCohorts{})
	m.AddSystem(&sys.TransitionForagers{})

	m.AddSystem(&sys.EggLaying{
		MaxEggsPerDay: 1600,
	})

	m.AddSystem(&sys.Foraging{
		PatchUpdater: &sys.UpdatePatchesForaging{},
	})
	m.AddSystem(&sys.HoneyConsumption{})
	m.AddSystem(&sys.PollenConsumption{})

	m.AddSystem(&sys.CountPopulation{})

	m.AddSystem(&system.FixedTermination{Steps: 100000})

	// Graphics

	m.AddUISystem((&window.Window{
		Title:        "Cohorts",
		Bounds:       window.B(1, 30, 600, 400),
		DrawInterval: 15,
	}).
		With(&plot.TimeSeries{
			Observer:       &obs.WorkerCohorts{},
			UpdateInterval: 1,
			Labels:         plot.Labels{Title: "Cohorts", X: "Time [days]", Y: "Count"},
		}))

	m.AddUISystem((&window.Window{
		Title:  "Age distribution",
		Bounds: window.B(1, 460, 600, 400),
	}).
		With(&plot.Lines{
			Observer: &obs.AgeStructure{
				MaxAge: 400,
			},
			YLim:   [...]float64{0, 1600},
			Labels: plot.Labels{Title: "Age distribution", X: "Age [days]", Y: "Count"},
		}))

	// Run

	window.Run(m)
}
