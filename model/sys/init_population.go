package sys

import (
	"github.com/mlange-42/arche-model/resource"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/beecs/model/res"
)

type InitPopulation struct {
	InitialCount int
	MinAge       int
	MaxAge       int
	MinMilage    float32
	MaxMilage    float32
}

func (s *InitPopulation) Initialize(w *ecs.World) {
	params := ecs.GetResource[res.Params](w)
	factory := ecs.GetResource[res.ForagerFactory](w)
	rand := ecs.GetResource[resource.Rand](w)

	squadrons := s.InitialCount / params.SquadronSize
	factory.CreateInitialSquadrons(squadrons, -s.MaxAge, -s.MinAge, s.MinMilage, s.MaxMilage, rand)
}

func (s *InitPopulation) Update(w *ecs.World) {}

func (s *InitPopulation) Finalize(w *ecs.World) {}
