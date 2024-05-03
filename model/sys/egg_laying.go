package sys

import (
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/beecs/model/res"
)

type EggLaying struct {
	eggs *res.Eggs
	time *res.Time

	MaxEggsPerDay int
}

func (s *EggLaying) Initialize(w *ecs.World) {
	s.eggs = ecs.GetResource[res.Eggs](w)
	s.time = ecs.GetResource[res.Time](w)
}

func (s *EggLaying) Update(w *ecs.World) {
	if !s.time.IsDayTick {
		return
	}
}

func (s *EggLaying) Finalize(w *ecs.World) {}
