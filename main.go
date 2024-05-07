package main

import (
	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche-model/system"
	"github.com/mlange-42/arche-pixel/plot"
	"github.com/mlange-42/arche-pixel/window"
	"github.com/mlange-42/arche/ecs"
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

	workerMort := res.WorkerMortality{
		Eggs:   0.03,
		Larvae: 0.01,
		Pupae:  0.001,
		InHive: 0.004,

		MaxLifespan: 390,
		MaxMilage:   200,
	}
	ecs.AddResource(&m.World, &workerMort)

	droneMort := res.DroneMortality{
		Eggs:   0.064,
		Larvae: 0.044,
		Pupae:  0.005,
		InHive: 0.05,
	}
	ecs.AddResource(&m.World, &droneMort)

	aff := res.AgeFirstForaging{
		Base: 21,
		Min:  7,
		Max:  50,
	}
	ecs.AddResource(&m.World, &aff)

	factory := res.NewForagerFactory(&m.World)
	ecs.AddResource(&m.World, &factory)

	// Initialization

	m.AddSystem(&sys.InitCohorts{
		EggTimeWorker:    3,
		LarvaeTimeWorker: 6,
		PupaeTimeWorker:  12,

		EggTimeDrone:    3,
		LarvaeTimeDrone: 7,
		PupaeTimeDrone:  14,
		LifespanDrone:   37,
	})

	m.AddSystem(&sys.InitPopulation{
		InitialCount: 10_000,
		MinAge:       100,
		MaxAge:       160,
		MinMilage:    0,
		MaxMilage:    200,
	})

	// Sub-models

	m.AddSystem(&sys.CalcAff{})

	m.AddSystem(&sys.MortalityCohorts{})

	m.AddSystem(&sys.AgeCohorts{})
	m.AddSystem(&sys.TransitionForagers{})

	m.AddSystem(&sys.EggLaying{
		MaxEggsPerDay: 1600,
	})

	m.AddSystem(&system.FixedTermination{Steps: 100000})

	// Graphics

	m.AddUISystem((&window.Window{
		Title: "Age distribution",
	}).
		With(&plot.Lines{
			Observer: &obs.Cohorts{
				MaxAge: 400,
			},
			YLim: [...]float64{0, 1600},
		}))

	// Run

	window.Run(m)
}
