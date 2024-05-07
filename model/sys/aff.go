package sys

import (
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/beecs/model/res"
)

type CalcAff struct {
	affParams *res.AgeFirstForagingParams
	aff       res.AgeFirstForaging
}

func (s *CalcAff) Initialize(w *ecs.World) {
	s.affParams = ecs.GetResource[res.AgeFirstForagingParams](w)
	s.aff = res.AgeFirstForaging{}
	ecs.AddResource(w, &s.aff)

	s.aff.Aff = s.affParams.Base
}

func (s *CalcAff) Update(w *ecs.World) {
	/*pollenTH := 0.5
	proteinTH := 1.0
	honeyTH := 35.0 * (DailyHoneyConsumption / 1000) * ENERGY_HONEY_per_g
	broodTH := 0.1
	foragerToWorkerTH := 0.3

	lastAff := s.aff.Current*/
	s.aff.Aff = s.affParams.Base
}

func (s *CalcAff) Finalize(w *ecs.World) {}
