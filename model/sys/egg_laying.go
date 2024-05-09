package sys

import (
	"math"

	"github.com/mlange-42/arche-model/resource"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/beecs/model/res"
	"github.com/mlange-42/beecs/model/util"
)

type EggLaying struct {
	time        *resource.Tick
	eggs        *res.Eggs
	pop         *res.PopulationStats
	nurseParams *res.NurseParams
	workerDev   *res.WorkerDevelopment
}

func (s *EggLaying) Initialize(w *ecs.World) {
	s.time = ecs.GetResource[resource.Tick](w)
	s.eggs = ecs.GetResource[res.Eggs](w)
	s.pop = ecs.GetResource[res.PopulationStats](w)
	s.nurseParams = ecs.GetResource[res.NurseParams](w)
	s.workerDev = ecs.GetResource[res.WorkerDevelopment](w)
}

func (s *EggLaying) Update(w *ecs.World) {
	// TODO: limit number of eggs
	eggs := int(float64(s.nurseParams.MaxEggsPerDay) * season(s.time.Tick))

	if s.nurseParams.EggNursingLimit {
		emergingAge := float64(s.workerDev.EggTime + s.workerDev.LarvaeTime + s.workerDev.PupaeTime)
		eggsNurse := int((float64(s.pop.WorkersInHive) + float64(s.pop.WorkersForagers)*s.nurseParams.ForagerNursingContribution) *
			s.nurseParams.MaxBroodNurseRatio / emergingAge)

		if eggsNurse < eggs {
			eggs = eggsNurse
		}
	}
	if eggs > s.nurseParams.MaxEggsPerDay {
		eggs = s.nurseParams.MaxEggsPerDay
	}
	if s.pop.TotalBrood+int(eggs) > s.nurseParams.MaxBroodCells {
		eggs = s.nurseParams.MaxBroodCells - s.pop.TotalBrood
	}

	droneEggs := 0
	dayOfYear := int(s.time.Tick % 365)
	if dayOfYear >= s.nurseParams.DroneEggLayingSeason[0] && dayOfYear <= s.nurseParams.DroneEggLayingSeason[1] {
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

	return 1 - math.Max(s1, s2)
}
