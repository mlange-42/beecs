package res

import (
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
	"github.com/mlange-42/beecs/model/comp"
	"golang.org/x/exp/rand"
)

type ForagerFactory struct {
	builder generic.Map2[comp.Milage, comp.Age]
}

func NewForagerFactory(world *ecs.World) ForagerFactory {
	return ForagerFactory{
		builder: generic.NewMap2[comp.Milage, comp.Age](world),
	}
}

func (f *ForagerFactory) CreateSquadrons(count int, dayOfBirth int) {
	q := f.builder.NewBatchQ(count)
	for q.Next() {
		_, a := q.Get()
		a.DayOfBirth = dayOfBirth
	}
}

func (f *ForagerFactory) CreateInitialSquadrons(count int, minDayOfBirth, maxDayOfBirth int, minMilage, maxMilage float32, rnd rand.Source) {
	q := f.builder.NewBatchQ(count)
	rng := rand.New(rnd)
	for q.Next() {
		m, a := q.Get()
		a.DayOfBirth = rng.Intn(maxDayOfBirth-minDayOfBirth) + minDayOfBirth
		m.Total = rng.Float32()*(maxMilage-minMilage) + minMilage
	}
}
