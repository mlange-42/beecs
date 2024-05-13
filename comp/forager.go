package comp

import (
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/beecs/activity"
)

// Milage component for forager squadrons.
type Milage struct {
	Today float32 // Milage today [km].
	Total float32 // Milage over total lifetime [km].
}

// Age component for forager squadrons.
type Age struct {
	DayOfBirth int // Date of birth for calculating the age from the current model tick.
}

// Activity component for forager squadrons.
type Activity struct {
	Current       activity.ForagerActivity // Current activity.
	PollenForager bool                     // Whether it is currently foraging for pollen.
}

// KnownPatch component for forager squadrons.
type KnownPatch struct {
	Nectar ecs.Entity // Known nectar patch.
	Pollen ecs.Entity // Known pollen patch.
}

// NectarLoad component for forager squadrons.
type NectarLoad struct {
	Energy float64 // Current nectar energy load per individual [kJ]
}
