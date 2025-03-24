package model

import (
	"github.com/mlange-42/ark-tools/app"
	"github.com/mlange-42/ark/ecs"
	"github.com/mlange-42/beecs/globals"
	"github.com/mlange-42/beecs/params"
	"github.com/mlange-42/beecs/sys"
)

// Default sets up the default beecs model with the standard sub-models.
//
// If the argument m is nil, a new model instance is created.
// If it is non-nil, the model is reset and re-used, saving some time for initialization and memory allocation.
func Default(p params.Params, app *app.App) *app.App {

	// Add parameters and other resources

	app = initializeModel(p, app)

	// Initialization

	app.AddSystem(&sys.InitStore{})
	app.AddSystem(&sys.InitCohorts{})
	app.AddSystem(&sys.InitPopulation{})
	app.AddSystem(&sys.InitPatchesList{})
	app.AddSystem(&sys.InitForagingPeriod{})

	// Sub-models

	app.AddSystem(&sys.CalcAff{})
	app.AddSystem(&sys.CalcForagingPeriod{})
	app.AddSystem(&sys.ReplenishPatches{})

	app.AddSystem(&sys.BroodCare{}) // Moved before any other population changes for now, to avoid counting more than once.
	app.AddSystem(&sys.AgeCohorts{})
	app.AddSystem(&sys.TransitionForagers{})
	app.AddSystem(&sys.EggLaying{})

	app.AddSystem(&sys.MortalityCohorts{})
	app.AddSystem(&sys.MortalityForagers{})

	app.AddSystem(&sys.Foraging{})
	app.AddSystem(&sys.HoneyConsumption{})
	app.AddSystem(&sys.PollenConsumption{})

	app.AddSystem(&sys.CountPopulation{})

	app.AddSystem(&sys.FixedTermination{})

	return app
}

// WithSystems sets up a beecs model with the given systems instead of the default ones.
//
// If the argument m is nil, a new model instance is created.
// If it is non-nil, the model is reset and re-used, saving some time for initialization and memory allocation.
func WithSystems(p params.Params, sys []app.System, app *app.App) *app.App {

	app = initializeModel(p, app)

	for _, s := range sys {
		app.AddSystem(s)
	}

	return app
}

func initializeModel(p params.Params, a *app.App) *app.App {
	if a == nil {
		a = app.New()
	} else {
		a.Reset()
	}

	p.Apply(&a.World)

	factory := globals.NewForagerFactory(&a.World)
	ecs.AddResource(&a.World, &factory)

	stats := globals.PopulationStats{}
	ecs.AddResource(&a.World, &stats)

	consumptionStats := globals.ConsumptionStats{}
	ecs.AddResource(&a.World, &consumptionStats)

	foragingStats := globals.ForagingStats{}
	ecs.AddResource(&a.World, &foragingStats)

	return a
}
