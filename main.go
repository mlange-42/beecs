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

	params := res.Params{
		SquadronSize: 100,
	}
	ecs.AddResource(&m.World, &params)

	aff := res.AgeFirstForaging{
		Base: 21,
		Min:  7,
		Max:  50,
	}
	ecs.AddResource(&m.World, &aff)

	m.AddSystem(&sys.InitCohorts{
		EggTimeWorker:    3,
		LarvaeTimeWorker: 6,
		PupaeTimeWorker:  12,

		EggTimeDrone:    3,
		LarvaeTimeDrone: 7,
		PupaeTimeDrone:  14,
		LifespanDrone:   37,
	})

	m.AddSystem(&sys.Time{
		TicksPerDay: 1, //6 * 24,
	})

	m.AddSystem(&sys.CalcAff{})

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
	m.AddSystem(&sys.TransitionForagers{})

	m.AddSystem(&sys.EggLaying{
		MaxEggsPerDay: 1600,
	})

	m.AddSystem(&system.FixedTermination{Steps: 100000})

	m.AddUISystem((&window.Window{
		DrawInterval: 1,
	}).
		With(&plot.Lines{
			Observer: &obs.Cohorts{},
			YLim:     [...]float64{0, 4}, // Optional Y axis limits.
		}))

	window.Run(m)
}
