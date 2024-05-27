package sys

import (
	"github.com/mlange-42/arche-model/resource"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/beecs/globals"
	"github.com/mlange-42/beecs/params"
)

// FixedTermination terminates the simulation after the number of ticks
// given in [params.Termination.MaxTicks].
type FixedTermination struct {
	termRes    *resource.Termination
	termParams *params.Termination
	popStats   *globals.PopulationStats
	step       int64
}

func (s *FixedTermination) Initialize(w *ecs.World) {
	s.termRes = ecs.GetResource[resource.Termination](w)
	s.termParams = ecs.GetResource[params.Termination](w)
	s.popStats = ecs.GetResource[globals.PopulationStats](w)
	s.step = 0
}

func (s *FixedTermination) Update(w *ecs.World) {
	if s.termParams.OnExtinction && s.popStats.TotalPopulation == 0 {
		s.termRes.Terminate = true
	} else if s.termParams.MaxTicks > 0 && s.step+1 >= int64(s.termParams.MaxTicks) {
		s.termRes.Terminate = true
	}
	s.step++
}

// Finalize the system
func (s *FixedTermination) Finalize(w *ecs.World) {}
