package sys

import (
	"github.com/mlange-42/arche-model/resource"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
	"github.com/mlange-42/beecs/model/comp"
	"github.com/mlange-42/beecs/model/res"
	"github.com/mlange-42/beecs/model/util"
	"golang.org/x/exp/rand"
	"gonum.org/v1/gonum/stat/distuv"
)

type MortalityCohorts struct {
	eggs   *res.Eggs
	larvae *res.Larvae
	pupae  *res.Pupae
	inHive *res.InHive
	rng    *resource.Rand

	workerMort *res.WorkerMortality
	droneMort  *res.DroneMortality

	toRemove      []ecs.Entity
	foragerFilter generic.Filter0
}

func (s *MortalityCohorts) Initialize(w *ecs.World) {
	s.eggs = ecs.GetResource[res.Eggs](w)
	s.larvae = ecs.GetResource[res.Larvae](w)
	s.pupae = ecs.GetResource[res.Pupae](w)
	s.inHive = ecs.GetResource[res.InHive](w)
	s.rng = ecs.GetResource[resource.Rand](w)

	s.workerMort = ecs.GetResource[res.WorkerMortality](w)
	s.droneMort = ecs.GetResource[res.DroneMortality](w)

	s.foragerFilter = *generic.NewFilter0().With(generic.T[comp.Milage]())
}

func (s *MortalityCohorts) Update(w *ecs.World) {
	applyMortality(s.eggs.Workers, s.workerMort.Eggs, s.rng)
	applyMortality(s.eggs.Drones, s.droneMort.Eggs, s.rng)

	applyMortality(s.larvae.Workers, s.workerMort.Larvae, s.rng)
	applyMortality(s.larvae.Drones, s.droneMort.Larvae, s.rng)

	applyMortality(s.pupae.Workers, s.workerMort.Pupae, s.rng)
	applyMortality(s.pupae.Drones, s.droneMort.Pupae, s.rng)

	applyMortality(s.inHive.Workers, s.workerMort.InHive, s.rng)
	applyMortality(s.inHive.Drones, s.droneMort.InHive, s.rng)

	r := rand.New(s.rng)
	query := s.foragerFilter.Query(w)
	for query.Next() {
		if r.Float64() < s.workerMort.InHive {
			s.toRemove = append(s.toRemove, query.Entity())
		}
	}
	for _, e := range s.toRemove {
		w.RemoveEntity(e)
	}
	s.toRemove = s.toRemove[:0]
}

func (s *MortalityCohorts) Finalize(w *ecs.World) {}

func applyMortality(coh []int, m float64, rng rand.Source) {
	for i := range coh {
		num := coh[i]
		rng := distuv.Poisson{
			Src:    rng,
			Lambda: m * float64(num),
		}
		toDie := int(rng.Rand())
		coh[i] = util.MaxInt(0, num-toDie)
	}
}
