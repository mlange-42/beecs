package sys

import (
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/beecs/model/res"
)

type InitStore struct {
	InitialPollen float64 // [g]
	InitialHoney  float64 // [kg]
}

func (s *InitStore) Initialize(w *ecs.World) {
	energyParams := ecs.GetResource[res.EnergyParams](w)
	stores := res.Stores{
		Pollen:              s.InitialPollen,
		Honey:               s.InitialHoney * 1000.0 * energyParams.EnergyHoney,
		ProteinFactorNurses: 1.0,
	}

	ecs.AddResource(w, &stores)
}

func (s *InitStore) Update(w *ecs.World) {}

func (s *InitStore) Finalize(w *ecs.World) {}
