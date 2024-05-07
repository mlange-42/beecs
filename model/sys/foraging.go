package sys

import (
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
	"github.com/mlange-42/beecs/model/comp"
	"github.com/mlange-42/beecs/model/res"
)

type Foraging struct {
	inHive         *res.InHive
	params         *res.Params
	aff            *res.AgeFirstForaging
	probParams     *res.ForagingProbabilityParams
	energyParams   *res.EnergyParams
	stores         *res.Stores
	foragersFilter *generic.Filter0
}

func (s *Foraging) Initialize(w *ecs.World) {
	s.inHive = ecs.GetResource[res.InHive](w)
	s.params = ecs.GetResource[res.Params](w)
	s.aff = ecs.GetResource[res.AgeFirstForaging](w)
	s.probParams = ecs.GetResource[res.ForagingProbabilityParams](w)
	s.energyParams = ecs.GetResource[res.EnergyParams](w)
	s.stores = ecs.GetResource[res.Stores](w)
	s.foragersFilter = generic.NewFilter0().With(generic.T[comp.Age]())
}

func (s *Foraging) Update(w *ecs.World) {
	totalInHive, totalForagers := s.countBees(w)

	decentHoney := float64(totalInHive+totalForagers) * 1.5 * s.energyParams.EnergyHoney
	forageProb := s.calcForagingProb(w, decentHoney, 0)
	_ = forageProb
}

func (s *Foraging) Finalize(w *ecs.World) {}

func (s *Foraging) countBees(w *ecs.World) (inHive, foragers int) {
	inHive = 0
	for i := 0; i < s.aff.Aff; i++ {
		inHive += s.inHive.Workers[i]
	}
	query := s.foragersFilter.Query(w)
	foragers = query.Count() * s.params.SquadronSize
	query.Close()
	return
}

func (s *Foraging) calcForagingProb(w *ecs.World, decentHoney, decentPollen float64) float64 {
	if s.stores.Pollen/decentPollen > 0.5 && s.stores.Honey/decentHoney > 1 {
		return 0
	}
	prob := s.probParams.Base
	if s.stores.Pollen/decentPollen < 0.2 || s.stores.Honey/decentHoney < 0.5 {
		prob = s.probParams.High
	}
	if s.stores.Honey/decentHoney < 0.2 {
		prob = s.probParams.Emergency
	}
	return prob
}
