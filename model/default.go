package model

import (
	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/beecs/model/comp"
	"github.com/mlange-42/beecs/model/res"
	"github.com/mlange-42/beecs/model/sys"
)

// Default sets up the default beecs model with the standard BEEHAVE parameters.
//
// If the argument m is nil, a new model instance is created.
// If it is non-nil, the model is reset and re-used, saving some time for initialization and memory allocation.
func Default(m *model.Model) *model.Model {
	if m == nil {
		m = model.New()
	} else {
		m.Reset()
	}

	// Resources

	params := res.Params{
		SquadronSize: 100,
	}
	ecs.AddResource(&m.World, &params)

	workerDev := res.WorkerDevelopment{
		EggTime:     3,
		LarvaeTime:  6,
		PupaeTime:   12,
		MaxLifespan: 290,
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
		MaxMilage: 800,
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
		SearchLength:                6.5 * 60 * 17, // [m] (6630m, 17 min)
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

		TimeNectarUnloading: 116,
		TimePollenUnloading: 210,
	}
	ecs.AddResource(&m.World, &forageParams)

	danceParams := res.DanceParams{
		Slope:                1.16,
		Intercept:            0.0,
		MaxCircuits:          117,
		FindProbability:      0.5,
		PollenDanceFollowers: 2,
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
		ProteinStoreNurse:    7, // [d]
	}
	ecs.AddResource(&m.World, &storeParams)

	honeyNeeds := res.HoneyNeeds{
		WorkerResting:    11.0,  // [mg/d]
		WorkerNurse:      53.42, // [mg/d]
		WorkerLarvaTotal: 65.4,  // [mg]
		DroneLarva:       19.2,  // [mg/d]
		Drone:            10.0,  // [mg/d]
	}
	ecs.AddResource(&m.World, &honeyNeeds)

	pollenNeeds := res.PollenNeeds{
		WorkerLarvaTotal: 142.0, // [mg]
		DroneLarva:       50.0,  // [mg/d]
		Worker:           1.5,   // [mg/d]
		Drone:            2.0,   // [mg/d]
	}
	ecs.AddResource(&m.World, &pollenNeeds)

	nurseParams := res.NurseParams{
		MaxBroodNurseRatio:         3.0,
		ForagerNursingContribution: 0.2,
		MaxEggsPerDay:              1600,
		DroneEggsProportion:        0.04,
		EggNursingLimit:            true,
		MaxBroodCells:              200_000,
		DroneEggLayingSeason:       [2]int{115, 240},
	}
	ecs.AddResource(&m.World, &nurseParams)

	factory := res.NewForagerFactory(&m.World)
	ecs.AddResource(&m.World, &factory)

	stats := res.PopulationStats{}
	ecs.AddResource(&m.World, &stats)

	consumptionStats := res.ConsumptionStats{}
	ecs.AddResource(&m.World, &consumptionStats)

	// Initialization

	m.AddSystem(&sys.InitStore{
		InitialPollen: 100,
		InitialHoney:  25,
	})

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

	m.AddSystem(&sys.BroodCare{}) // Moved before any other population changes for now, to avoid counting more than once.
	m.AddSystem(&sys.AgeCohorts{})
	m.AddSystem(&sys.TransitionForagers{})
	m.AddSystem(&sys.EggLaying{})

	m.AddSystem(&sys.MortalityCohorts{})
	m.AddSystem(&sys.MortalityForagers{})

	m.AddSystem(&sys.Foraging{
		PatchUpdater: &sys.UpdatePatchesForaging{},
	})
	m.AddSystem(&sys.HoneyConsumption{})
	m.AddSystem(&sys.PollenConsumption{})

	m.AddSystem(&sys.CountPopulation{})

	return m
}
