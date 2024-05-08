package comp

import (
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/beecs/model/activity"
)

type Milage struct {
	Today float32
	Total float32
}

type Age struct {
	DayOfBirth int
}

type Activity struct {
	Current       activity.ForagerActivity
	PollenForager bool
}

type KnownPatch struct {
	Nectar ecs.Entity
	Pollen ecs.Entity
}
