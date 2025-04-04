package sys

import (
	"math"

	"github.com/mlange-42/ark/ecs"
	"github.com/mlange-42/beecs/globals"
	"github.com/mlange-42/beecs/params"
)

// BroodCare calculates and applies brood mortality based on food or nursing shortage.
type BroodCare struct {
	nurseParams *params.Nursing

	pop    *globals.PopulationStats
	stores *globals.Stores

	eggs   *globals.Eggs
	larvae *globals.Larvae
	pupae  *globals.Pupae
}

func (s *BroodCare) Initialize(w *ecs.World) {
	s.nurseParams = ecs.GetResource[params.Nursing](w)

	s.pop = ecs.GetResource[globals.PopulationStats](w)
	s.stores = ecs.GetResource[globals.Stores](w)

	s.eggs = ecs.GetResource[globals.Eggs](w)
	s.larvae = ecs.GetResource[globals.Larvae](w)
	s.pupae = ecs.GetResource[globals.Pupae](w)
}

func (s *BroodCare) Update(w *ecs.World) {
	maxBrood := (float64(s.pop.WorkersInHive) + float64(s.pop.WorkersForagers)*s.nurseParams.ForagerNursingContribution) *
		s.nurseParams.MaxBroodNurseRatio

	excessBrood := int(math.Ceil(float64(s.pop.TotalBrood) - maxBrood))
	lacksNurses := excessBrood > 0

	starved := int(math.Ceil((float64(s.pop.WorkerLarvae+s.pop.DroneLarvae) * (1.0 - s.stores.ProteinFactorNurses))))

	if starved > excessBrood {
		excessBrood = starved
	}

	s.killBrood(excessBrood, lacksNurses)
}

func (s *BroodCare) Finalize(w *ecs.World) {}

func (s *BroodCare) killBrood(excess int, lacksNurses bool) {
	if excess <= 0 {
		return
	}

	if lacksNurses {
		if excess = reduceCohorts(s.eggs.Drones, excess); excess == 0 {
			return
		}
	}
	if excess = reduceCohorts(s.larvae.Drones, excess); excess == 0 {
		return
	}

	if lacksNurses {
		if excess = reduceCohorts(s.eggs.Workers, excess); excess == 0 {
			return
		}
	}
	if excess = reduceCohorts(s.larvae.Workers, excess); excess == 0 {
		return
	}

	if lacksNurses {
		if excess = reduceCohorts(s.pupae.Drones, excess); excess == 0 {
			return
		}
		if excess = reduceCohorts(s.pupae.Workers, excess); excess == 0 {
			return
		}
	}

	panic("still brood to kill - code should not be reachable")
}

func reduceCohorts(coh []int, excess int) int {
	for i, v := range coh {
		if v >= excess {
			coh[i] -= excess
			return 0
		}
		coh[i] = 0
		excess -= v
	}
	return excess
}
