package sys

import (
	"github.com/mlange-42/arche-model/resource"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/beecs/model/globals"
	"github.com/mlange-42/beecs/model/params"
)

type InitPopulation struct{}

func (s *InitPopulation) Initialize(w *ecs.World) {
	init := ecs.GetResource[params.InitialPopulation](w)
	params := ecs.GetResource[params.Foragers](w)
	factory := ecs.GetResource[globals.ForagerFactory](w)
	rand := ecs.GetResource[resource.Rand](w)

	squadrons := init.Count / params.SquadronSize
	factory.CreateInitialSquadrons(squadrons, -init.MaxAge, -init.MinAge, init.MinMilage, init.MaxMilage, rand)
}

func (s *InitPopulation) Update(w *ecs.World) {}

func (s *InitPopulation) Finalize(w *ecs.World) {}
