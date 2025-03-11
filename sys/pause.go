package sys

import (
	"github.com/mlange-42/ark-tools/app"
	"github.com/mlange-42/ark/ecs"
)

// Pause the simulation at the given simulation step.
type Pause struct {
	sys   *app.Systems
	Steps int64
	step  int64
}

// Initialize the system
func (s *Pause) Initialize(w *ecs.World) {
	s.sys = ecs.GetResource[app.Systems](w)
	s.step = 0
}

// Update the system
func (s *Pause) Update(w *ecs.World) {
	if s.step+1 >= s.Steps {
		s.sys.Paused = true
	}
	s.step++
}

// Finalize the system
func (s *Pause) Finalize(w *ecs.World) {}
