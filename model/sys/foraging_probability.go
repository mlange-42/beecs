package sys

import (
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/beecs/model/res"
)

type CalcForagingProbability struct {
	prob *res.ForagingProbability
}

func (s *CalcForagingProbability) Initialize(w *ecs.World) {
	s.prob = ecs.GetResource[res.ForagingProbability](w)
}

func (s *CalcForagingProbability) Update(w *ecs.World) {
	s.prob.Prob = 0.01
}

func (s *CalcForagingProbability) Finalize(w *ecs.World) {}
