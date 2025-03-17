package sys

import (
	"math/rand/v2"

	"github.com/mlange-42/ark-tools/resource"
	"github.com/mlange-42/ark/ecs"
	"github.com/mlange-42/beecs/comp"
	"github.com/mlange-42/beecs/params"
)

// MortalityForagers applies worker mortality, including
//   - background mortality from [params.WorkerMortality.InHive]
//   - removal of squadrons reaching [params.WorkerDevelopment.MaxLifespan]
//   - removal of squadrons exceeding [params.WorkerMortality.MaxMilage]
type MortalityForagers struct {
	rng           *resource.Rand
	time          *resource.Tick
	workerMort    *params.WorkerMortality
	workerDev     *params.WorkerDevelopment
	toRemove      []ecs.Entity
	foragerFilter *ecs.Filter2[comp.Age, comp.Milage]
}

func (s *MortalityForagers) Initialize(w *ecs.World) {
	s.rng = ecs.GetResource[resource.Rand](w)
	s.time = ecs.GetResource[resource.Tick](w)
	s.workerMort = ecs.GetResource[params.WorkerMortality](w)
	s.workerDev = ecs.GetResource[params.WorkerDevelopment](w)
	s.foragerFilter = s.foragerFilter.New(w)
}

func (s *MortalityForagers) Update(w *ecs.World) {
	r := rand.New(s.rng)
	query := s.foragerFilter.Query()
	for query.Next() {
		a, m := query.Get()
		if int(s.time.Tick)-a.DayOfBirth >= s.workerDev.MaxLifespan ||
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
