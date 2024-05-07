package sys

import (
	"math"

	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
	"github.com/mlange-42/beecs/model/comp"
	"github.com/mlange-42/beecs/model/res"
)

type PollenConsumption struct {
	needs          *res.PollenNeeds
	stores         *res.Stores
	pop            *res.PopulationStats
	params         *res.Params
	workerDev      *res.WorkerDevelopment
	nurseParams    *res.NurseParams
	aff            *res.AgeFirstForaging
	foragersFilter *generic.Filter0
}

func (s *PollenConsumption) Initialize(w *ecs.World) {
	s.needs = ecs.GetResource[res.PollenNeeds](w)
	s.stores = ecs.GetResource[res.Stores](w)
	s.pop = ecs.GetResource[res.PopulationStats](w)
	s.params = ecs.GetResource[res.Params](w)
	s.workerDev = ecs.GetResource[res.WorkerDevelopment](w)
	s.nurseParams = ecs.GetResource[res.NurseParams](w)
	s.aff = ecs.GetResource[res.AgeFirstForaging](w)
	s.foragersFilter = generic.NewFilter0().With(generic.T[comp.Age]())
}

func (s *PollenConsumption) Update(w *ecs.World) {
	needLarva := s.needs.WorkerLarvaTotal / float64(s.workerDev.LarvaeTime)

	needAdult := float64(s.pop.WorkersInHive+s.pop.WorkersForagers)*s.needs.Worker + float64(s.pop.DronesInHive)*s.needs.Drone
	needLarvae := float64(s.pop.WorkerLarvae)*needLarva + float64(s.pop.DroneLarvae)*s.needs.DroneLarva

	consumption := (needAdult + needLarvae) / 1000.0
	s.stores.Pollen = math.Max(s.stores.Pollen-consumption, 0)

	s.stores.IdealPollen = math.Max(consumption*float64(s.needs.IdealStoreDays), s.needs.MinIdealStore)
}

func (s *PollenConsumption) Finalize(w *ecs.World) {}
