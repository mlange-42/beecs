package sys

import (
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
	"github.com/mlange-42/beecs/model/comp"
	"github.com/mlange-42/beecs/model/res"
)

type CountPopulation struct {
	eggs           *res.Eggs
	larvae         *res.Larvae
	pupae          *res.Pupae
	inHive         *res.InHive
	params         *res.Params
	stats          *res.PopulationStats
	foragersFilter *generic.Filter0
}

func (s *CountPopulation) Initialize(w *ecs.World) {
	s.eggs = ecs.GetResource[res.Eggs](w)
	s.larvae = ecs.GetResource[res.Larvae](w)
	s.pupae = ecs.GetResource[res.Pupae](w)
	s.inHive = ecs.GetResource[res.InHive](w)
	s.params = ecs.GetResource[res.Params](w)
	s.stats = ecs.GetResource[res.PopulationStats](w)
	s.foragersFilter = generic.NewFilter0().With(generic.T[comp.Age]())

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

	query := s.foragersFilter.Query(w)
	s.stats.WorkersForagers = query.Count() * s.params.SquadronSize
	query.Close()

	s.stats.TotalBrood =
		s.stats.WorkerEggs + s.stats.WorkerLarvae + s.stats.WorkerPupae +
			s.stats.DroneEggs + s.stats.DroneLarvae + s.stats.DronePupae
}

func countCohort(coh []int) int {
	sum := 0
	for _, cnt := range coh {
		sum += cnt
	}
	return sum
}