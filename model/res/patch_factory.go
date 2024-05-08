package res

import (
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
	"github.com/mlange-42/beecs/model/comp"
)

type PatchFactory struct {
	builder generic.Map5[comp.PatchConfig, comp.Trip, comp.HandlingTime, comp.Resource, comp.Mortality]
}

func NewPatchFactory(world *ecs.World) PatchFactory {
	return PatchFactory{
		builder: generic.NewMap5[comp.PatchConfig, comp.Trip, comp.HandlingTime, comp.Resource, comp.Mortality](world),
	}
}

func (f *PatchFactory) CreatePatch(conf comp.PatchConfig) ecs.Entity {
	return f.builder.NewWith(&conf, &comp.Trip{}, &comp.HandlingTime{}, &comp.Resource{}, &comp.Mortality{})
}
