package sys

import (
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/beecs/model/res"
)

type InitStore struct{}

func (s *InitStore) Initialize(w *ecs.World) {
	init := ecs.GetResource[res.InitialStores](w)
	energyParams := ecs.GetResource[res.EnergyParams](w)
	stores := res.Stores{
		Honey:               init.Honey * 1000.0 * energyParams.EnergyHoney,
		Pollen:              init.Pollen,
		ProteinFactorNurses: 1.0,
	}

	ecs.AddResource(w, &stores)
}

func (s *InitStore) Update(w *ecs.World) {}

func (s *InitStore) Finalize(w *ecs.World) {}
