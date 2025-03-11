package sys

import (
	"github.com/mlange-42/ark-tools/resource"
	"github.com/mlange-42/ark/ecs"
	"github.com/mlange-42/beecs/globals"
	"github.com/mlange-42/beecs/params"
)

// TransitionForagers transitions all in-hive bees of age equal or above [globals.AgeFirstForaging.Aff]
// to forager squadrons.
type TransitionForagers struct {
	time    *resource.Tick
	params  *params.Foragers
	factory *globals.ForagerFactory
	inHive  *globals.InHive
	aff     *globals.AgeFirstForaging
}

func (s *TransitionForagers) Initialize(w *ecs.World) {
	s.time = ecs.GetResource[resource.Tick](w)
	s.params = ecs.GetResource[params.Foragers](w)
	s.factory = ecs.GetResource[globals.ForagerFactory](w)
	s.inHive = ecs.GetResource[globals.InHive](w)
	s.aff = ecs.GetResource[globals.AgeFirstForaging](w)
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
