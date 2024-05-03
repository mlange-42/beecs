package sys

import (
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/beecs/model/res"
)

type Time struct {
	time res.Time

	TicksPerDay int
}

func (s *Time) Initialize(w *ecs.World) {
	s.time = res.Time{
		Tick: -1,
	}
	ecs.AddResource(w, &s.time)
}

func (s *Time) Update(w *ecs.World) {
	s.time.Tick++
	s.time.Day = s.time.Tick / s.TicksPerDay
	s.time.IsDayTick = s.time.Tick%s.TicksPerDay == 0
}

func (s *Time) Finalize(w *ecs.World) {}
