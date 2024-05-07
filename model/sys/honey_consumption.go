package sys

import (
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
	"github.com/mlange-42/beecs/model/comp"
	"github.com/mlange-42/beecs/model/res"
)

type HoneyConsumption struct {
	needs *res.HoneyNeeds

	stores         *res.Stores
	foragersFilter *generic.Filter0
}

func (s *HoneyConsumption) Initialize(w *ecs.World) {
	s.stores = ecs.GetResource[res.Stores](w)
	s.foragersFilter = generic.NewFilter0().With(generic.T[comp.Age]())
}

func (s *HoneyConsumption) Update(w *ecs.World) {
}

func (s *HoneyConsumption) Finalize(w *ecs.World) {}
