package sys

import (
	"github.com/mlange-42/arche-model/resource"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/beecs/model/res"
)

type TransitionForagers struct {
	inHive  *res.InHive
	params  *res.Params
	aff     *res.AgeFirstForaging
	time    *resource.Tick
	factory *res.ForagerFactory
}

func (s *TransitionForagers) Initialize(w *ecs.World) {
	s.inHive = ecs.GetResource[res.InHive](w)
	s.params = ecs.GetResource[res.Params](w)
	s.aff = ecs.GetResource[res.AgeFirstForaging](w)
	s.time = ecs.GetResource[resource.Tick](w)
	s.factory = ecs.GetResource[res.ForagerFactory](w)
}

func (s *TransitionForagers) Update(w *ecs.World) {
	aff := s.aff.Aff
	newForagers := 0
	for i := aff; i < len(s.inHive.Workers); i++ {
		newForagers += s.inHive.Workers[i]
		s.inHive.Workers[i] = 0
	}

	squadrons := newForagers / s.params.SquadronSize
	remainder := newForagers % s.params.SquadronSize

	s.inHive.Workers[aff-1] += remainder

	if squadrons > 0 {
		s.factory.CreateSquadrons(squadrons, int(s.time.Tick)-aff)
	}
}

func (s *TransitionForagers) Finalize(w *ecs.World) {}
