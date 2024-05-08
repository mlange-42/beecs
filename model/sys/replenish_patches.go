package sys

import (
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
	"github.com/mlange-42/beecs/model/comp"
	"github.com/mlange-42/beecs/model/res"
)

type ReplenishPatches struct {
	params *res.ForagingParams

	filter generic.Filter2[comp.PatchConfig, comp.Resource]
}

func (s *ReplenishPatches) Initialize(w *ecs.World) {
	s.params = ecs.GetResource[res.ForagingParams](w)

	s.filter = *generic.NewFilter2[comp.PatchConfig, comp.Resource]()
}

func (s *ReplenishPatches) Update(w *ecs.World) {
	query := s.filter.Query(w)
	for query.Next() {
		conf, res := query.Get()

		res.MaxNectar = conf.Nectar * 1000 * 1000
		res.MaxPollen = conf.Pollen * 1000
		res.Nectar = res.MaxNectar
		res.Pollen = res.MaxNectar
	}
}

func (s *ReplenishPatches) Finalize(w *ecs.World) {}
