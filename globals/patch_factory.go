package globals

import (
	"math"

	"github.com/mlange-42/ark/ecs"
	"github.com/mlange-42/beecs/comp"
)

// PatchFactory is a helper resource for creating flower patch entities.
type PatchFactory struct {
	builder *ecs.Map9[
		comp.PatchProperties, comp.Coords,
		comp.PatchDistance, comp.Trip, comp.HandlingTime,
		comp.Resource, comp.Mortality, comp.Dance,
		comp.Visits]
	initMapper *ecs.Map2[comp.Coords, comp.PatchDistance]

	constantPatchMapper *ecs.Map1[comp.ConstantPatch]
	seasonalPatchMapper *ecs.Map1[comp.SeasonalPatch]
	scriptedPatchMapper *ecs.Map1[comp.ScriptedPatch]
}

// NewPatchFactory creates a new PatchFactory
func NewPatchFactory(world *ecs.World) PatchFactory {
	return PatchFactory{
		builder: ecs.NewMap9[
			comp.PatchProperties, comp.Coords,
			comp.PatchDistance, comp.Trip, comp.HandlingTime,
			comp.Resource, comp.Mortality, comp.Dance,
			comp.Visits](world),

		initMapper: ecs.NewMap2[comp.Coords, comp.PatchDistance](world),

		constantPatchMapper: ecs.NewMap1[comp.ConstantPatch](world),
		seasonalPatchMapper: ecs.NewMap1[comp.SeasonalPatch](world),
		scriptedPatchMapper: ecs.NewMap1[comp.ScriptedPatch](world),
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
	e := f.builder.NewEntityFn(nil)

	co, dist := f.initMapper.Get(e)
	*co = coords
	dist.DistToColony = conf.DistToColony

	anyPatch := false

	if conf.ConstantPatch != nil {
		f.constantPatchMapper.Add(e, conf.ConstantPatch)
		anyPatch = true
	}

	if conf.SeasonalPatch != nil {
		if anyPatch {
			panic("each patch must have exactly one of ConstantPatch, SeasonalPatch or ScriptedPatch, has multiple")
		}
		f.seasonalPatchMapper.Add(e, conf.SeasonalPatch)
		anyPatch = true
	}

	if conf.ScriptedPatch != nil {
		if anyPatch {
			panic("each patch must have exactly one of ConstantPatch, SeasonalPatch or ScriptedPatch, has multiple")
		}
		f.scriptedPatchMapper.Add(e, conf.ScriptedPatch)
		anyPatch = true
	}

	if !anyPatch {
		panic("each patch must have exactly one of ConstantPatch, SeasonalPatch or ScriptedPatch, has none")
	}
}
