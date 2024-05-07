package sys

import (
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/beecs/model/res"
)

type CalcAff struct {
	aff *res.AgeFirstForaging
}

func (s *CalcAff) Initialize(w *ecs.World) {
	s.aff = ecs.GetResource[res.AgeFirstForaging](w)

	s.aff.Current = s.aff.Base
}

func (s *CalcAff) Update(w *ecs.World) {
	/*pollenTH := 0.5
	proteinTH := 1.0
	honeyTH := 35.0 * (DailyHoneyConsumption / 1000) * ENERGY_HONEY_per_g
	broodTH := 0.1
	foragerToWorkerTH := 0.3

	lastAff := s.aff.Current*/
	s.aff.Current = s.aff.Base
}

func (s *CalcAff) Finalize(w *ecs.World) {}
