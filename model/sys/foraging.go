package sys

import (
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/beecs/model/res"
)

type Foraging struct {
	pop        *res.PopulationStats
	probParams *res.ForagingProbabilityParams
	stores     *res.Stores
}

func (s *Foraging) Initialize(w *ecs.World) {
	s.pop = ecs.GetResource[res.PopulationStats](w)
	s.probParams = ecs.GetResource[res.ForagingProbabilityParams](w)
	s.stores = ecs.GetResource[res.Stores](w)
}

func (s *Foraging) Update(w *ecs.World) {
	forageProb := s.calcForagingProb(s.stores.DecentHoney, s.stores.IdealPollen)
	_ = forageProb
}

func (s *Foraging) Finalize(w *ecs.World) {}

func (s *Foraging) calcForagingProb(decentHoney, idealPollen float64) float64 {
	if s.stores.Pollen/idealPollen > 0.5 && s.stores.Honey/decentHoney > 1 {
		return 0
	}
	prob := s.probParams.Base
	if s.stores.Pollen/idealPollen < 0.2 || s.stores.Honey/decentHoney < 0.5 {
		prob = s.probParams.High
	}
	if s.stores.Honey/decentHoney < 0.2 {
		prob = s.probParams.Emergency
	}
	return prob
}
