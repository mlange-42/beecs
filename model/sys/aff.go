package sys

import (
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/beecs/model/res"
)

type CalcAff struct {
	aff  *res.AgeFirstForaging
	time *res.Time
}

func (s *CalcAff) Initialize(w *ecs.World) {
	s.aff = ecs.GetResource[res.AgeFirstForaging](w)
	s.time = ecs.GetResource[res.Time](w)

	s.aff.Current = s.aff.Base
}

func (s *CalcAff) Update(w *ecs.World) {
	if !s.time.IsDayTick {
		return
	}

	/*pollenTH := 0.5
	proteinTH := 1.0
	honeyTH := 35.0 * (DailyHoneyConsumption / 1000) * ENERGY_HONEY_per_g
	broodTH := 0.1
	foragerToWorkerTH := 0.3

	lastAff := s.aff.Current*/
	s.aff.Current = s.aff.Base
}

func (s *CalcAff) Finalize(w *ecs.World) {}
