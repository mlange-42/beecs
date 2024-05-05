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
	aff    *res.AgeFirstForaging
	time   *res.Time
}

func (s *AgeCohorts) Initialize(w *ecs.World) {
	s.eggs = ecs.GetResource[res.Eggs](w)
	s.larvae = ecs.GetResource[res.Larvae](w)
	s.pupae = ecs.GetResource[res.Pupae](w)
	s.inHive = ecs.GetResource[res.InHive](w)
	s.aff = ecs.GetResource[res.AgeFirstForaging](w)
	s.time = ecs.GetResource[res.Time](w)
}

func (s *AgeCohorts) Update(w *ecs.World) {
	if !s.time.IsDayTick {
		return
	}

	shiftCohorts(s.inHive.Workers, s.pupae.Workers[len(s.pupae.Workers)-1])
	shiftCohorts(s.inHive.Drones, s.pupae.Drones[len(s.pupae.Drones)-1])

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
