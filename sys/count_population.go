package sys

import (
	"github.com/mlange-42/ark/ecs"
	"github.com/mlange-42/beecs/comp"
	"github.com/mlange-42/beecs/globals"
	"github.com/mlange-42/beecs/params"
)

// CountPopulation counts the population per developmental stage.
// Results are written to a [globals.PopulationStats] resource.
type CountPopulation struct {
	params *params.Foragers

	eggs   *globals.Eggs
	larvae *globals.Larvae
	pupae  *globals.Pupae
	inHive *globals.InHive
	stats  *globals.PopulationStats

	foragersFilter *ecs.Filter0
}

func (s *CountPopulation) Initialize(w *ecs.World) {
	s.params = ecs.GetResource[params.Foragers](w)

	s.eggs = ecs.GetResource[globals.Eggs](w)
	s.larvae = ecs.GetResource[globals.Larvae](w)
	s.pupae = ecs.GetResource[globals.Pupae](w)
	s.inHive = ecs.GetResource[globals.InHive](w)
	s.stats = ecs.GetResource[globals.PopulationStats](w)

	s.foragersFilter = ecs.NewFilter0(w).With(ecs.C[comp.Age]())

	s.count(w)
}

func (s *CountPopulation) Update(w *ecs.World) {
	s.count(w)
}

func (s *CountPopulation) Finalize(w *ecs.World) {}

func (s *CountPopulation) count(w *ecs.World) {
	s.stats.Reset()

	s.stats.WorkerEggs = countCohort(s.eggs.Workers)
	s.stats.WorkerLarvae = countCohort(s.larvae.Workers)
	s.stats.WorkerPupae = countCohort(s.pupae.Workers)
	s.stats.WorkersInHive = countCohort(s.inHive.Workers)

	s.stats.DroneEggs = countCohort(s.eggs.Drones)
	s.stats.DroneLarvae = countCohort(s.larvae.Drones)
	s.stats.DronePupae = countCohort(s.pupae.Drones)
	s.stats.DronesInHive = countCohort(s.inHive.Drones)

	query := s.foragersFilter.Query()
	s.stats.WorkersForagers = query.Count() * s.params.SquadronSize
	query.Close()

	s.stats.TotalBrood =
		s.stats.WorkerEggs + s.stats.WorkerLarvae + s.stats.WorkerPupae +
			s.stats.DroneEggs + s.stats.DroneLarvae + s.stats.DronePupae
	s.stats.TotalAdults = s.stats.WorkersInHive + s.stats.WorkersForagers + s.stats.DronesInHive
	s.stats.TotalPopulation = s.stats.TotalBrood + s.stats.TotalAdults
}

func countCohort(coh []int) int {
	sum := 0
	for _, cnt := range coh {
		sum += cnt
	}
	return sum
}
