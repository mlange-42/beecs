package sys

import (
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
	"github.com/mlange-42/beecs/comp"
)

type ReplenishPatches struct {
	filter generic.Filter2[comp.PatchConfig, comp.Resource]
}

func (s *ReplenishPatches) Initialize(w *ecs.World) {
	s.filter = *generic.NewFilter2[comp.PatchConfig, comp.Resource]()
}

func (s *ReplenishPatches) Update(w *ecs.World) {
	query := s.filter.Query(w)
	for query.Next() {
		conf, res := query.Get()

		res.MaxNectar = conf.Nectar * 1000 * 1000
		res.MaxPollen = conf.Pollen * 1000

		res.Nectar = res.MaxNectar
		res.Pollen = res.MaxPollen
	}
}

func (s *ReplenishPatches) Finalize(w *ecs.World) {}
