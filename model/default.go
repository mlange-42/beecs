package model

import (
	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/beecs/globals"
	"github.com/mlange-42/beecs/params"
	"github.com/mlange-42/beecs/sys"
)

// Default sets up the default beecs model with the standard sub-models.
//
// If the argument m is nil, a new model instance is created.
// If it is non-nil, the model is reset and re-used, saving some time for initialization and memory allocation.
func Default(p params.Params, m *model.Model) *model.Model {
	if m == nil {
		m = model.New()
	} else {
		m.Reset()
	}

	p.Apply(&m.World)

	factory := globals.NewForagerFactory(&m.World)
	ecs.AddResource(&m.World, &factory)

	stats := globals.PopulationStats{}
	ecs.AddResource(&m.World, &stats)

	consumptionStats := globals.ConsumptionStats{}
	ecs.AddResource(&m.World, &consumptionStats)

	// Initialization

	m.AddSystem(&sys.InitStore{})

	m.AddSystem(&sys.InitCohorts{})

	m.AddSystem(&sys.InitPopulation{})

	m.AddSystem(&sys.InitPatchesList{})

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

	m.AddSystem(&sys.FixedTermination{})

	return m
}
