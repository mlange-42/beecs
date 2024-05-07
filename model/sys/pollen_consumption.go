package sys

import (
	"math"

	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/beecs/model/res"
)

type PollenConsumption struct {
	needs       *res.PollenNeeds
	storeParams *res.StoreParams
	stores      *res.Stores
	pop         *res.PopulationStats
	workerDev   *res.WorkerDevelopment
}

func (s *PollenConsumption) Initialize(w *ecs.World) {
	s.needs = ecs.GetResource[res.PollenNeeds](w)
	s.storeParams = ecs.GetResource[res.StoreParams](w)
	s.stores = ecs.GetResource[res.Stores](w)
	s.pop = ecs.GetResource[res.PopulationStats](w)
	s.workerDev = ecs.GetResource[res.WorkerDevelopment](w)
}

func (s *PollenConsumption) Update(w *ecs.World) {
	needLarva := s.needs.WorkerLarvaTotal / float64(s.workerDev.LarvaeTime)

	needAdult := float64(s.pop.WorkersInHive+s.pop.WorkersForagers)*s.needs.Worker + float64(s.pop.DronesInHive)*s.needs.Drone
	needLarvae := float64(s.pop.WorkerLarvae)*needLarva + float64(s.pop.DroneLarvae)*s.needs.DroneLarva

	consumption := (needAdult + needLarvae) / 1000.0
	s.stores.Pollen = math.Max(s.stores.Pollen-consumption, 0)

	s.stores.IdealPollen = math.Max(consumption*float64(s.storeParams.IdealPollenStoreDays), s.storeParams.MinIdealPollenStore)
}

func (s *PollenConsumption) Finalize(w *ecs.World) {}
