package sys

import (
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/beecs/globals"
	"github.com/mlange-42/beecs/params"
)

// InitStore initializes and adds [globals.Stores]
// according to the settings in [params.InitialStores].
type InitStore struct{}

func (s *InitStore) Initialize(w *ecs.World) {
	init := ecs.GetResource[params.InitialStores](w)
	energyParams := ecs.GetResource[params.EnergyContent](w)
	stores := globals.Stores{
		Honey:               init.Honey * 1000.0 * energyParams.Honey,
		Pollen:              init.Pollen,
		ProteinFactorNurses: 1.0,
	}

	ecs.AddResource(w, &stores)
}

func (s *InitStore) Update(w *ecs.World) {}

func (s *InitStore) Finalize(w *ecs.World) {}
