package sys

import (
	"math"

	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/beecs/model/res"
)

type HoneyConsumption struct {
	needs        *res.HoneyNeeds
	stores       *res.Stores
	pop          *res.PopulationStats
	cons         *res.ConsumptionStats
	workerDev    *res.WorkerDevelopment
	nurseParams  *res.NursingParams
	energyParams *res.EnergyParams
}

func (s *HoneyConsumption) Initialize(w *ecs.World) {
	s.needs = ecs.GetResource[res.HoneyNeeds](w)
	s.stores = ecs.GetResource[res.Stores](w)
	s.pop = ecs.GetResource[res.PopulationStats](w)
	s.cons = ecs.GetResource[res.ConsumptionStats](w)
	s.workerDev = ecs.GetResource[res.WorkerDevelopment](w)
	s.nurseParams = ecs.GetResource[res.NursingParams](w)
	s.energyParams = ecs.GetResource[res.EnergyParams](w)
}

func (s *HoneyConsumption) Update(w *ecs.World) {
	thermoRegBrood := (s.needs.WorkerNurse - s.needs.WorkerResting) / s.nurseParams.MaxBroodNurseRatio
	needLarva := s.needs.WorkerLarvaTotal / float64(s.workerDev.LarvaeTime)

	needAdult := float64(s.pop.WorkersInHive+s.pop.WorkersForagers)*s.needs.WorkerResting + float64(s.pop.DronesInHive)*s.needs.Drone
	needLarvae := float64(s.pop.WorkerLarvae)*needLarva + float64(s.pop.DroneLarvae)*s.needs.DroneLarva

	consumption := needAdult + needLarvae + float64(s.pop.TotalBrood)*thermoRegBrood
	consumptionEnergy := 0.001 * consumption * s.energyParams.Honey

	s.stores.Honey = math.Max(s.stores.Honey-consumptionEnergy, 0)
	s.stores.DecentHoney = math.Max(float64(s.pop.WorkersInHive+s.pop.WorkersForagers), 1) * 1.5 * s.energyParams.Honey
	s.cons.HoneyDaily = consumption
}

func (s *HoneyConsumption) Finalize(w *ecs.World) {}
