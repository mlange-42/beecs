package res

import (
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
	"github.com/mlange-42/beecs/model/activity"
	"github.com/mlange-42/beecs/model/comp"
	"golang.org/x/exp/rand"
)

type ForagerFactory struct {
	builder generic.Map5[comp.Activity, comp.KnownPatch, comp.Age, comp.Milage, comp.NectarLoad]
}

func NewForagerFactory(world *ecs.World) ForagerFactory {
	return ForagerFactory{
		builder: generic.NewMap5[comp.Activity, comp.KnownPatch, comp.Age, comp.Milage, comp.NectarLoad](world),
	}
}

func (f *ForagerFactory) CreateSquadrons(count int, dayOfBirth int) {
	q := f.builder.NewBatchQ(count)
	for q.Next() {
		act, _, age, _, _ := q.Get()
		age.DayOfBirth = dayOfBirth
		act.Current = activity.Resting
	}
}

func (f *ForagerFactory) CreateInitialSquadrons(count int, minDayOfBirth, maxDayOfBirth int, minMilage, maxMilage float32, rnd rand.Source) {
	q := f.builder.NewBatchQ(count)
	rng := rand.New(rnd)
	for q.Next() {
		act, _, age, milage, _ := q.Get()
		age.DayOfBirth = rng.Intn(maxDayOfBirth-minDayOfBirth) + minDayOfBirth
		milage.Total = rng.Float32()*(maxMilage-minMilage) + minMilage
		act.Current = activity.Resting
	}
}
