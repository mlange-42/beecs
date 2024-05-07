package res

import (
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
	"github.com/mlange-42/beecs/model/comp"
)

type PatchFactory struct {
	builder generic.Map2[comp.TripDuration, comp.DistanceToColony]
}

func NewPatchFactory(world *ecs.World) PatchFactory {
	return PatchFactory{
		builder: generic.NewMap2[comp.TripDuration, comp.DistanceToColony](world),
	}
}

func (f *PatchFactory) CreatePatch(dist float64) {
	e := f.builder.New()

	_, d := f.builder.Get(e)
	d.Dist = dist
}
