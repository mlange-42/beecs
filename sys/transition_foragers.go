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
	time       *resource.Tick
	params     *params.Foragers
	factory    *globals.ForagerFactory
	inHive     *globals.InHive
	aff        *globals.AgeFirstForaging
	newCohorts *globals.NewCohorts
}

func (s *TransitionForagers) Initialize(w *ecs.World) {
	s.time = ecs.GetResource[resource.Tick](w)
	s.params = ecs.GetResource[params.Foragers](w)
	s.factory = ecs.GetResource[globals.ForagerFactory](w)
	s.inHive = ecs.GetResource[globals.InHive](w)
	s.aff = ecs.GetResource[globals.AgeFirstForaging](w)
	s.newCohorts = ecs.GetResource[globals.NewCohorts](w)
}

// TransitionForagers now only calculates the foragers that are to be transitioned and "kills" those cohorts.
// These new squadrons are saved and will be initialized in the foraging module before foraging activity starts as to not appear in the countingstats beforehand.
// To more closely mimic BEEHAVE dynamics squadons and remainder now also get calculated per cohort; this can make a difference if aff decreased this timestep, because
// then 2 cohorts would have been added up in NewForagers to then calculate squadrons/remainder. The old way seems to make more sense biologically, but can overestimate
// the amount of squadrons to be created by 1 compared to original BEEHAVE behavior.
func (s *TransitionForagers) Update(w *ecs.World) {
	aff := s.aff.Aff
	squadrons := 0
	remainder := 0

	for i := aff; i < len(s.inHive.Workers); i++ {
		newForagers := s.inHive.Workers[i]
		s.inHive.Workers[i] = 0
		squadrons += newForagers / s.params.SquadronSize // in case aff gets lowered there are 2 IHbee-cohorts getting transformed to foragers.
		remainder += newForagers % s.params.SquadronSize // BEEHAVE calculates squadrons and remainder per cohort. This can make a difference compared to first adding newForagers from both cohorts.
	}

	s.inHive.Workers[aff-1] += remainder
	s.newCohorts.Foragers = squadrons // new foragers now get "saved" here until they get initialized later
}

func (s *TransitionForagers) Finalize(w *ecs.World) {}
