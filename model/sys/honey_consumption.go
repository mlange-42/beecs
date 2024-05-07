package sys

import (
	"math"

	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
	"github.com/mlange-42/beecs/model/comp"
	"github.com/mlange-42/beecs/model/res"
)

type HoneyConsumption struct {
	needs          *res.HoneyNeeds
	stores         *res.Stores
	pop            *res.PopulationStats
	params         *res.Params
	workerDev      *res.WorkerDevelopment
	nurseParams    *res.NurseParams
	energyParams   *res.EnergyParams
	aff            *res.AgeFirstForaging
	foragersFilter *generic.Filter0
}

func (s *HoneyConsumption) Initialize(w *ecs.World) {
	s.needs = ecs.GetResource[res.HoneyNeeds](w)
	s.stores = ecs.GetResource[res.Stores](w)
	s.pop = ecs.GetResource[res.PopulationStats](w)
	s.params = ecs.GetResource[res.Params](w)
	s.workerDev = ecs.GetResource[res.WorkerDevelopment](w)
	s.nurseParams = ecs.GetResource[res.NurseParams](w)
	s.energyParams = ecs.GetResource[res.EnergyParams](w)
	s.aff = ecs.GetResource[res.AgeFirstForaging](w)
	s.foragersFilter = generic.NewFilter0().With(generic.T[comp.Age]())
}

func (s *HoneyConsumption) Update(w *ecs.World) {
	thermoRegBrood := s.needs.WorkerNurse - s.needs.WorkerResting
	needLarva := s.needs.WorkerLarvaTotal / float64(s.workerDev.LarvaeTime)

	needAdult := float64(s.pop.WorkersInHive+s.pop.WorkersForagers)*s.needs.WorkerResting + float64(s.pop.DronesInHive)*s.needs.Drone
	needLarvae := float64(s.pop.WorkerLarvae)*needLarva + float64(s.pop.DroneLarvae)*s.needs.DroneLarva

	consumption := needAdult + needLarvae + float64(s.pop.TotalBrood)*thermoRegBrood

	s.stores.Honey = math.Max(s.stores.Honey-consumption, 0)
	s.stores.DecentHoney = float64(s.pop.WorkersInHive+s.pop.WorkersForagers) * 1.5 * s.energyParams.EnergyHoney
}

func (s *HoneyConsumption) Finalize(w *ecs.World) {}
