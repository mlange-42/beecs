package globals

import (
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
	"github.com/mlange-42/beecs/comp"
	"github.com/mlange-42/beecs/enum/activity"
	"golang.org/x/exp/rand"
)

// ForagerFactory is a helper resource for creating forager squadron entities.
type ForagerFactory struct {
	builder generic.Map5[comp.Activity, comp.KnownPatch, comp.Age, comp.Milage, comp.NectarLoad]
}

// NewForagerFactory creates a new ForagerFactory
func NewForagerFactory(world *ecs.World) ForagerFactory {
	return ForagerFactory{
		builder: generic.NewMap5[comp.Activity, comp.KnownPatch, comp.Age, comp.Milage, comp.NectarLoad](world),
	}
}

// CreateSquadrons creates the given number of squadrons with the given day of birth
// (usually the current model tick).
func (f *ForagerFactory) CreateSquadrons(count int, dayOfBirth int) {
	q := f.builder.NewBatchQ(count)
	for q.Next() {
		act, _, age, _, _ := q.Get()
		age.DayOfBirth = dayOfBirth
		act.Current = activity.Resting
	}
}

// CreateInitialSquadrons creates the given number of squadrons with random day of birth and milage as given by the ranges.
//
// Used to create initial foragers.
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
