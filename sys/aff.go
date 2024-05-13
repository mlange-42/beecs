package sys

import (
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/beecs/globals"
	"github.com/mlange-42/beecs/params"
	"github.com/mlange-42/beecs/util"
)

type CalcAff struct {
	affParams    *params.AgeFirstForaging
	energyParams *params.EnergyContent
	nurseParams  *params.Nursing
	consStats    *globals.ConsumptionStats
	stores       *globals.Stores
	pop          *globals.PopulationStats

	aff globals.AgeFirstForaging
}

func (s *CalcAff) Initialize(w *ecs.World) {
	s.affParams = ecs.GetResource[params.AgeFirstForaging](w)
	s.energyParams = ecs.GetResource[params.EnergyContent](w)
	s.nurseParams = ecs.GetResource[params.Nursing](w)
	s.consStats = ecs.GetResource[globals.ConsumptionStats](w)
	s.stores = ecs.GetResource[globals.Stores](w)
	s.pop = ecs.GetResource[globals.PopulationStats](w)

	s.aff = globals.AgeFirstForaging{
		Aff: s.affParams.Base,
	}
	ecs.AddResource(w, &s.aff)
}

func (s *CalcAff) Update(w *ecs.World) {
	pollenTH := 0.5
	proteinTH := 1.0
	honeyTH := 35.0 * (s.consStats.HoneyDaily / 1000) * s.energyParams.Honey
	broodTH := 0.1
	foragerToWorkerTH := 0.3

	aff := s.aff.Aff

	if s.stores.Pollen/s.stores.IdealPollen < pollenTH {
		aff--
	}
	if s.stores.ProteinFactorNurses < proteinTH {
		aff--
	}
	if s.stores.Honey < honeyTH {
		aff -= 2
	}
	if s.pop.WorkersInHive > 0 &&
		float64(s.pop.WorkersForagers)/float64(s.pop.WorkersInHive) < foragerToWorkerTH {
		aff--
	}
	maxBrood := (float64(s.pop.WorkersInHive) + float64(s.pop.WorkersForagers)*s.nurseParams.ForagerNursingContribution) *
		s.nurseParams.MaxBroodNurseRatio
	if maxBrood > 0 && float64(s.pop.TotalBrood)/maxBrood > broodTH {
		aff += 2
	}

	if s.aff.Aff < s.affParams.Base-7 {
		aff++
	} else if s.aff.Aff > s.affParams.Base+7 {
		aff--
	}

	if aff < s.aff.Aff {
		aff = s.aff.Aff - 1
	} else if aff > s.aff.Aff {
		aff = s.aff.Aff + 1
	}
	s.aff.Aff = util.Clamp(aff, s.affParams.Min, s.affParams.Max)
}

func (s *CalcAff) Finalize(w *ecs.World) {}
