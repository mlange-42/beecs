package sys

import (
	"github.com/mlange-42/ark/ecs"
	"github.com/mlange-42/beecs/globals"
)

// AgeCohorts moves all cohort-based age classes to the next day's cohort.
// It also handles transition from eggs to larvae, larvae to pupae and pupae to in-hive bees.
// It does not handle transition from in-hive bees to foragers.
type AgeCohorts struct {
	eggs       *globals.Eggs
	larvae     *globals.Larvae
	pupae      *globals.Pupae
	inHive     *globals.InHive
	newCohorts *globals.NewCohorts
}

func (s *AgeCohorts) Initialize(w *ecs.World) {
	s.eggs = ecs.GetResource[globals.Eggs](w)
	s.larvae = ecs.GetResource[globals.Larvae](w)
	s.pupae = ecs.GetResource[globals.Pupae](w)
	s.inHive = ecs.GetResource[globals.InHive](w)
	s.newCohorts = ecs.GetResource[globals.NewCohorts](w)
}

func (s *AgeCohorts) Update(w *ecs.World) {
	s.newCohorts.IHbees = s.pupae.Workers[len(s.pupae.Workers)-1] // new IHbees now get "saved" here until they get initialized later
	s.newCohorts.Drones = s.pupae.Drones[len(s.pupae.Drones)-1]   // new Drones now get "saved" here until they get initialized later

	shiftCohorts(s.inHive.Workers, 0) // first cohort now 0 until it gets initialized in the new subprocess; similar to egglaying
	shiftCohorts(s.inHive.Drones, 0)  // first cohort now 0 until it gets initialized in the new subprocess; similar to egglaying

	shiftCohorts(s.pupae.Workers, s.larvae.Workers[len(s.larvae.Workers)-1])
	shiftCohorts(s.pupae.Drones, s.larvae.Drones[len(s.larvae.Drones)-1])

	shiftCohorts(s.larvae.Workers, s.eggs.Workers[len(s.eggs.Workers)-1])
	shiftCohorts(s.larvae.Drones, s.eggs.Drones[len(s.eggs.Drones)-1])

	shiftCohorts(s.eggs.Workers, 0)
	shiftCohorts(s.eggs.Drones, 0)
}

func (s *AgeCohorts) Finalize(w *ecs.World) {}

func shiftCohorts(coh []int, add int) {
	for i := len(coh) - 1; i > 0; i-- {
		coh[i] = coh[i-1]
	}
	coh[0] = add
}
