package globals

import (
	"math"

	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
	"github.com/mlange-42/beecs/comp"
)

// PatchFactory is a helper resource for creating flower patch entities.
type PatchFactory struct {
	builder generic.Map9[
		comp.PatchProperties, comp.Coords,
		comp.PatchDistance, comp.Trip, comp.HandlingTime,
		comp.Resource, comp.Mortality, comp.Dance,
		comp.Visits]

	constantPatchMapper generic.Map1[comp.ConstantPatch]
	seasonalPatchMapper generic.Map1[comp.SeasonalPatch]
	scriptedPatchMapper generic.Map1[comp.ScriptedPatch]
}

// NewPatchFactory creates a new PatchFactory
func NewPatchFactory(world *ecs.World) PatchFactory {
	return PatchFactory{
		builder: generic.NewMap9[
			comp.PatchProperties, comp.Coords,
			comp.PatchDistance, comp.Trip, comp.HandlingTime,
			comp.Resource, comp.Mortality, comp.Dance,
			comp.Visits](world),

		constantPatchMapper: generic.NewMap1[comp.ConstantPatch](world),
		seasonalPatchMapper: generic.NewMap1[comp.SeasonalPatch](world),
		scriptedPatchMapper: generic.NewMap1[comp.ScriptedPatch](world),
	}
}

func (f *PatchFactory) CreatePatches(conf []comp.PatchConfig) {
	cnt := len(conf)
	ang := (2 * math.Pi) / float64(cnt)
	for i, p := range conf {
		var coords comp.Coords
		if p.Coords == nil {
			coords = comp.Coords{
				X: math.Sin(ang*float64(i)) * p.DistToColony,
				Y: math.Cos(ang*float64(i)) * p.DistToColony,
			}
		} else {
			coords = *p.Coords
		}
		f.createPatch(p, coords)
	}
}

func (f *PatchFactory) createPatch(conf comp.PatchConfig, coords comp.Coords) {
	e := f.builder.NewWith(
		&comp.PatchProperties{}, &coords,
		&comp.PatchDistance{DistToColony: conf.DistToColony}, &comp.Trip{}, &comp.HandlingTime{},
		&comp.Resource{}, &comp.Mortality{}, &comp.Dance{}, &comp.Visits{})

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
}
