package model

import (
	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/beecs/model/res"
	"github.com/mlange-42/beecs/model/sys"
)

// Default sets up the default beecs model with the standard sub-models.
//
// If the argument m is nil, a new model instance is created.
// If it is non-nil, the model is reset and re-used, saving some time for initialization and memory allocation.
func Default(params *Params, m *model.Model) *model.Model {
	if m == nil {
		m = model.New()
	} else {
		m.Reset()
	}

	// Resources

	ecs.AddResource(&m.World, &params.WorkerDevelopment)
	ecs.AddResource(&m.World, &params.DroneDevelopment)
	ecs.AddResource(&m.World, &params.WorkerMortality)
	ecs.AddResource(&m.World, &params.DroneMortality)
	ecs.AddResource(&m.World, &params.AgeFirstForaging)
	ecs.AddResource(&m.World, &params.Forager)
	ecs.AddResource(&m.World, &params.Foraging)
	ecs.AddResource(&m.World, &params.HandlingTime)
	ecs.AddResource(&m.World, &params.Dance)
	ecs.AddResource(&m.World, &params.Energy)
	ecs.AddResource(&m.World, &params.Stores)
	ecs.AddResource(&m.World, &params.HoneyNeeds)
	ecs.AddResource(&m.World, &params.PollenNeeds)
	ecs.AddResource(&m.World, &params.Nursing)
	ecs.AddResource(&m.World, &params.InitialPopulation)
	ecs.AddResource(&m.World, &params.InitialStores)

	factory := res.NewForagerFactory(&m.World)
	ecs.AddResource(&m.World, &factory)

	stats := res.PopulationStats{}
	ecs.AddResource(&m.World, &stats)

	consumptionStats := res.ConsumptionStats{}
	ecs.AddResource(&m.World, &consumptionStats)

	// Initialization

	m.AddSystem(&sys.InitStore{})

	m.AddSystem(&sys.InitCohorts{})

	m.AddSystem(&sys.InitPopulation{})

	m.AddSystem(&sys.InitPatchesList{
		Patches: params.Patches,
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
