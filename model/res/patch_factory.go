package res

import (
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
	"github.com/mlange-42/beecs/model/comp"
)

type PatchFactory struct {
	builder generic.Map7[
		comp.PatchConfig, comp.Trip, comp.HandlingTime,
		comp.Resource, comp.Mortality, comp.Dance,
		comp.KnownPatch]
}

func NewPatchFactory(world *ecs.World) PatchFactory {
	return PatchFactory{
		builder: generic.NewMap7[
			comp.PatchConfig, comp.Trip, comp.HandlingTime,
			comp.Resource, comp.Mortality, comp.Dance,
			comp.KnownPatch](world),
	}
}

func (f *PatchFactory) CreatePatch(conf comp.PatchConfig) ecs.Entity {
	return f.builder.NewWith(
		&conf, &comp.Trip{}, &comp.HandlingTime{},
		&comp.Resource{}, &comp.Mortality{}, &comp.Dance{},
		&comp.KnownPatch{})
}
