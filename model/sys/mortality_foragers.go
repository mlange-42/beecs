package sys

import (
	"github.com/mlange-42/arche-model/resource"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
	"github.com/mlange-42/beecs/model/comp"
	"github.com/mlange-42/beecs/model/res"
	"golang.org/x/exp/rand"
)

type MortalityForagers struct {
	rng           *resource.Rand
	time          *resource.Tick
	workerMort    *res.WorkerMortality
	toRemove      []ecs.Entity
	foragerFilter generic.Filter2[comp.Age, comp.Milage]
}

func (s *MortalityForagers) Initialize(w *ecs.World) {
	s.rng = ecs.GetResource[resource.Rand](w)
	s.time = ecs.GetResource[resource.Tick](w)
	s.workerMort = ecs.GetResource[res.WorkerMortality](w)
	s.foragerFilter = *generic.NewFilter2[comp.Age, comp.Milage]()
}

func (s *MortalityForagers) Update(w *ecs.World) {
	r := rand.New(s.rng)
	query := s.foragerFilter.Query(w)
	for query.Next() {
		a, m := query.Get()
		if int(s.time.Tick)-a.DayOfBirth >= s.workerMort.MaxLifespan ||
			m.Total >= s.workerMort.MaxMilage ||
			r.Float64() < s.workerMort.InHive {
			s.toRemove = append(s.toRemove, query.Entity())
		}
	}
	for _, e := range s.toRemove {
		w.RemoveEntity(e)
	}
	s.toRemove = s.toRemove[:0]
}

func (s *MortalityForagers) Finalize(w *ecs.World) {}
