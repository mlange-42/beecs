package sys

import (
	"math"

	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/beecs/globals"
	"github.com/mlange-42/beecs/params"
)

type HoneyConsumption struct {
	needs        *params.HoneyNeeds
	workerDev    *params.WorkerDevelopment
	nurseParams  *params.Nursing
	energyParams *params.EnergyContent

	stores *globals.Stores
	pop    *globals.PopulationStats
	cons   *globals.ConsumptionStats
}

func (s *HoneyConsumption) Initialize(w *ecs.World) {
	s.needs = ecs.GetResource[params.HoneyNeeds](w)
	s.workerDev = ecs.GetResource[params.WorkerDevelopment](w)
	s.nurseParams = ecs.GetResource[params.Nursing](w)
	s.energyParams = ecs.GetResource[params.EnergyContent](w)

	s.stores = ecs.GetResource[globals.Stores](w)
	s.pop = ecs.GetResource[globals.PopulationStats](w)
	s.cons = ecs.GetResource[globals.ConsumptionStats](w)
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
