package sys

import (
	"math"

	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/beecs/model/globals"
	"github.com/mlange-42/beecs/model/params"
	"github.com/mlange-42/beecs/model/util"
)

type PollenConsumption struct {
	needs       *params.PollenNeeds
	storeParams *params.StoreParams
	nurseParams *params.NursingParams
	workerDev   *params.WorkerDevelopment

	stores *globals.Stores
	pop    *globals.PopulationStats
}

func (s *PollenConsumption) Initialize(w *ecs.World) {
	s.needs = ecs.GetResource[params.PollenNeeds](w)
	s.storeParams = ecs.GetResource[params.StoreParams](w)
	s.nurseParams = ecs.GetResource[params.NursingParams](w)
	s.workerDev = ecs.GetResource[params.WorkerDevelopment](w)

	s.stores = ecs.GetResource[globals.Stores](w)
	s.pop = ecs.GetResource[globals.PopulationStats](w)
}

func (s *PollenConsumption) Update(w *ecs.World) {
	needLarva := s.needs.WorkerLarvaTotal / float64(s.workerDev.LarvaeTime)

	needAdult := float64(s.pop.WorkersInHive+s.pop.WorkersForagers)*s.needs.Worker + float64(s.pop.DronesInHive)*s.needs.Drone
	needLarvae := float64(s.pop.WorkerLarvae)*needLarva + float64(s.pop.DroneLarvae)*s.needs.DroneLarva

	consumption := (needAdult + needLarvae) / 1000.0
	s.stores.Pollen = math.Max(s.stores.Pollen-consumption, 0)

	s.stores.IdealPollen = math.Max(consumption*float64(s.storeParams.IdealPollenStoreDays), s.storeParams.MinIdealPollenStore)

	if s.stores.Pollen > 0 {
		s.stores.ProteinFactorNurses = s.stores.ProteinFactorNurses + 1.0/s.storeParams.ProteinStoreNurse
	} else {
		maxBrood := (float64(s.pop.WorkersInHive) + float64(s.pop.WorkersForagers)*s.nurseParams.ForagerNursingContribution) *
			s.nurseParams.MaxBroodNurseRatio
		workLoad := 0.0
		if maxBrood > 0 {
			workLoad = float64(s.pop.TotalBrood) / maxBrood
		}
		s.stores.ProteinFactorNurses = s.stores.ProteinFactorNurses - workLoad/s.storeParams.ProteinStoreNurse
	}
	s.stores.ProteinFactorNurses = util.Clamp(s.stores.ProteinFactorNurses, 0.0, 1.0)
}

func (s *PollenConsumption) Finalize(w *ecs.World) {}
