package globals

import (
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
	"github.com/mlange-42/beecs/comp"
)

type PatchFactory struct {
	builder generic.Map7[
		comp.PatchProperties,
		comp.PatchDistance, comp.Trip, comp.HandlingTime,
		comp.Resource, comp.Mortality, comp.Dance]

	constantPatchMapper generic.Map1[comp.ConstantPatch]
	seasonalPatchMapper generic.Map1[comp.SeasonalPatch]
	scriptedPatchMapper generic.Map1[comp.ScriptedPatch]
}

func NewPatchFactory(world *ecs.World) PatchFactory {
	return PatchFactory{
		builder: generic.NewMap7[
			comp.PatchProperties,
			comp.PatchDistance, comp.Trip, comp.HandlingTime,
			comp.Resource, comp.Mortality, comp.Dance](world),

		constantPatchMapper: generic.NewMap1[comp.ConstantPatch](world),
		seasonalPatchMapper: generic.NewMap1[comp.SeasonalPatch](world),
		scriptedPatchMapper: generic.NewMap1[comp.ScriptedPatch](world),
	}
}

func (f *PatchFactory) CreatePatch(conf comp.PatchConfig) ecs.Entity {
	e := f.builder.NewWith(
		&comp.PatchProperties{},
		&comp.PatchDistance{DistToColony: conf.DistToColony}, &comp.Trip{}, &comp.HandlingTime{},
		&comp.Resource{}, &comp.Mortality{}, &comp.Dance{})

	anyPatch := false

	if conf.ConstantPatch != nil {
		f.constantPatchMapper.Assign(e, conf.ConstantPatch)
		anyPatch = true
	}

	if conf.SeasonalPatch != nil {
		if anyPatch {
			panic("each patch must have exactly one of ConstantPatch, SeasonalPatch or ScriptedPatch, has multiple")
		}
		f.seasonalPatchMapper.Assign(e, conf.SeasonalPatch)
		anyPatch = true
	}

	if conf.ScriptedPatch != nil {
		if anyPatch {
			panic("each patch must have exactly one of ConstantPatch, SeasonalPatch or ScriptedPatch, has multiple")
		}
		f.scriptedPatchMapper.Assign(e, conf.ScriptedPatch)
		anyPatch = true
	}

	if !anyPatch {
		panic("each patch must have exactly one of ConstantPatch, SeasonalPatch or ScriptedPatch, has none")
	}

	return e
}
