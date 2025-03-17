package globals

import (
	"math/rand/v2"

	"github.com/mlange-42/ark/ecs"
	"github.com/mlange-42/beecs/comp"
	"github.com/mlange-42/beecs/enum/activity"
)

// ForagerFactory is a helper resource for creating forager squadron entities.
type ForagerFactory struct {
	builder *ecs.Map5[comp.Activity, comp.KnownPatch, comp.Age, comp.Milage, comp.NectarLoad]
}

// NewForagerFactory creates a new ForagerFactory
func NewForagerFactory(world *ecs.World) ForagerFactory {
	return ForagerFactory{
		builder: ecs.NewMap5[comp.Activity, comp.KnownPatch, comp.Age, comp.Milage, comp.NectarLoad](world),
	}
}

// CreateSquadrons creates the given number of squadrons with the given day of birth
// (usually the current model tick).
func (f *ForagerFactory) CreateSquadrons(count int, dayOfBirth int) {
	f.builder.NewBatchFn(count, func(entity ecs.Entity, act *comp.Activity, _ *comp.KnownPatch, age *comp.Age, _ *comp.Milage, _ *comp.NectarLoad) {
		age.DayOfBirth = dayOfBirth
		act.Current = activity.Resting
	})
}

// CreateInitialSquadrons creates the given number of squadrons with random day of birth and milage as given by the ranges.
//
// Used to create initial foragers.
func (f *ForagerFactory) CreateInitialSquadrons(count int, minDayOfBirth, maxDayOfBirth int, minMilage, maxMilage float32, rnd rand.Source) {
	rng := rand.New(rnd)
	f.builder.NewBatchFn(count, func(entity ecs.Entity, act *comp.Activity, _ *comp.KnownPatch, age *comp.Age, milage *comp.Milage, _ *comp.NectarLoad) {
		age.DayOfBirth = rng.IntN(maxDayOfBirth-minDayOfBirth) + minDayOfBirth
		milage.Total = rng.Float32()*(maxMilage-minMilage) + minMilage
		act.Current = activity.Resting
	})
}
