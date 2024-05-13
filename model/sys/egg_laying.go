package sys

import (
	"math"

	"github.com/mlange-42/arche-model/resource"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/beecs/model/globals"
	"github.com/mlange-42/beecs/model/params"
	"github.com/mlange-42/beecs/model/util"
)

type EggLaying struct {
	time        *resource.Tick
	eggs        *globals.Eggs
	pop         *globals.PopulationStats
	nurseParams *params.Nursing
	workerDev   *params.WorkerDevelopment
}

func (s *EggLaying) Initialize(w *ecs.World) {
	s.time = ecs.GetResource[resource.Tick](w)
	s.eggs = ecs.GetResource[globals.Eggs](w)
	s.pop = ecs.GetResource[globals.PopulationStats](w)
	s.nurseParams = ecs.GetResource[params.Nursing](w)
	s.workerDev = ecs.GetResource[params.WorkerDevelopment](w)
}

func (s *EggLaying) Update(w *ecs.World) {
	elr := float64(s.nurseParams.MaxEggsPerDay) * (1.0 - season(s.time.Tick))

	if s.nurseParams.EggNursingLimit {
		emergingAge := float64(s.workerDev.EggTime + s.workerDev.LarvaeTime + s.workerDev.PupaeTime)
		elrNurse := (float64(s.pop.WorkersInHive) + float64(s.pop.WorkersForagers)*s.nurseParams.ForagerNursingContribution) *
			s.nurseParams.MaxBroodNurseRatio / emergingAge

		if elrNurse < elr {
			elr = elrNurse
		}
	}
	if elr > float64(s.nurseParams.MaxEggsPerDay) {
		elr = float64(s.nurseParams.MaxEggsPerDay)
	}
	eggs := int(math.Round(elr))
	if s.pop.TotalBrood+eggs > s.nurseParams.MaxBroodCells {
		eggs = s.nurseParams.MaxBroodCells - s.pop.TotalBrood
	}

	droneEggs := 0
	dayOfYear := int(s.time.Tick % 365)
	if dayOfYear >= s.nurseParams.DroneEggLayingSeasonStart && dayOfYear <= s.nurseParams.DroneEggLayingSeasonEnd {
		droneEggs = int(math.Max(s.nurseParams.DroneEggsProportion*float64(eggs), 0))
	}
	eggs = util.MaxInt(eggs-droneEggs, 0)

	// TODO: queen age

	s.eggs.Workers[0] = eggs
	s.eggs.Drones[0] = droneEggs
}

func (s *EggLaying) Finalize(w *ecs.World) {}

func season(t int64) float64 {
	d := float64(t % 365)
	x1, x2, x3, x4, x5 := 385.0, 25.0, 36.0, 155.0, 30.0
	s1 := (1 - (1 / (1 + x1*math.Exp(-2*d/x2))))
	s2 := (1 / (1 + x3*math.Exp(-2*(d-x4)/x5)))

	return math.Max(s1, s2)
}
