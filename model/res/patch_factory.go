package res

import (
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
	"github.com/mlange-42/beecs/model/comp"
)

type PatchFactory struct {
	builder generic.Map2[comp.PatchConfig, comp.TripDuration]
}

func NewPatchFactory(world *ecs.World) PatchFactory {
	return PatchFactory{
		builder: generic.NewMap2[comp.PatchConfig, comp.TripDuration](world),
	}
}

func (f *PatchFactory) CreatePatch(conf comp.PatchConfig) ecs.Entity {
	return f.builder.NewWith(&conf, &comp.TripDuration{})
}
