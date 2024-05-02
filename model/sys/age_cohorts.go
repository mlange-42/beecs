package sys

import (
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/beecs/model/res"
)

type AgeCohorts struct {
	eggs   *res.Eggs
	larvae *res.Larvae
	pupae  *res.Pupae
	inHive *res.InHive
}

func (s *AgeCohorts) Initialize(w *ecs.World) {
	s.eggs = ecs.GetResource[res.Eggs](w)
	s.larvae = ecs.GetResource[res.Larvae](w)
	s.pupae = ecs.GetResource[res.Pupae](w)
	s.inHive = ecs.GetResource[res.InHive](w)
}

func (s *AgeCohorts) Update(w *ecs.World) {
	newWorkerLarvae := s.eggs.Workers[len(s.eggs.Workers)-1]
	newDroneLarvae := s.eggs.Drones[len(s.eggs.Drones)-1]
	shiftCohorts(s.eggs.Workers, 0)
	shiftCohorts(s.eggs.Drones, 0)

	newWorkerPupae := s.larvae.Workers[len(s.larvae.Workers)-1]
	newDronePupae := s.larvae.Drones[len(s.larvae.Drones)-1]
	shiftCohorts(s.larvae.Workers, newWorkerLarvae)
	shiftCohorts(s.larvae.Drones, newDroneLarvae)

	newWorkers := s.pupae.Workers[len(s.pupae.Workers)-1]
	newDrones := s.pupae.Drones[len(s.pupae.Drones)-1]
	shiftCohorts(s.pupae.Workers, newWorkerPupae)
	shiftCohorts(s.pupae.Drones, newDronePupae)

	shiftCohorts(s.inHive.Workers, newWorkers)
	shiftCohorts(s.inHive.Drones, newDrones)
}

func (s *AgeCohorts) Finalize(w *ecs.World) {}

func shiftCohorts(coh []int, add int) {
	for i := len(coh) - 1; i > 0; i-- {
		coh[i] = coh[i-1]
	}
	coh[0] = add
}
