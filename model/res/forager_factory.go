package res

import (
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
	"github.com/mlange-42/beecs/model/activity"
	"github.com/mlange-42/beecs/model/comp"
	"golang.org/x/exp/rand"
)

type ForagerFactory struct {
	builder generic.Map4[comp.Milage, comp.Age, comp.Activity, comp.NectarLoad]
}

func NewForagerFactory(world *ecs.World) ForagerFactory {
	return ForagerFactory{
		builder: generic.NewMap4[comp.Milage, comp.Age, comp.Activity, comp.NectarLoad](world),
	}
}

func (f *ForagerFactory) CreateSquadrons(count int, dayOfBirth int) {
	q := f.builder.NewBatchQ(count)
	for q.Next() {
		_, a, act, _ := q.Get()
		a.DayOfBirth = dayOfBirth
		act.Current = activity.Resting
	}
}

func (f *ForagerFactory) CreateInitialSquadrons(count int, minDayOfBirth, maxDayOfBirth int, minMilage, maxMilage float32, rnd rand.Source) {
	q := f.builder.NewBatchQ(count)
	rng := rand.New(rnd)
	for q.Next() {
		m, a, act, _ := q.Get()
		a.DayOfBirth = rng.Intn(maxDayOfBirth-minDayOfBirth) + minDayOfBirth
		m.Total = rng.Float32()*(maxMilage-minMilage) + minMilage
		act.Current = activity.Resting
	}
}
