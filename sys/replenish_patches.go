package sys

import (
	"github.com/mlange-42/ark-tools/resource"
	"github.com/mlange-42/ark/ecs"
	"github.com/mlange-42/beecs/comp"
	"github.com/mlange-42/beecs/util"
)

// ReplenishPatches resets and replenishes flower patches to their current maximum nectar and pollen.
type ReplenishPatches struct {
	time   *resource.Tick
	filter ecs.Filter3[comp.PatchProperties, comp.Resource, comp.Visits]

	constantFilter ecs.Filter2[comp.PatchProperties, comp.ConstantPatch]
	seasonalFilter ecs.Filter2[comp.PatchProperties, comp.SeasonalPatch]
	scriptedFilter ecs.Filter2[comp.PatchProperties, comp.ScriptedPatch]
}

func (s *ReplenishPatches) Initialize(w *ecs.World) {
	s.time = ecs.GetResource[resource.Tick](w)
	s.filter = *ecs.NewFilter3[comp.PatchProperties, comp.Resource, comp.Visits](w)

	s.constantFilter = *ecs.NewFilter2[comp.PatchProperties, comp.ConstantPatch](w)
	s.seasonalFilter = *ecs.NewFilter2[comp.PatchProperties, comp.SeasonalPatch](w)
	s.scriptedFilter = *ecs.NewFilter2[comp.PatchProperties, comp.ScriptedPatch](w)
}

func (s *ReplenishPatches) Update(w *ecs.World) {
	constQuery := s.constantFilter.Query()
	for constQuery.Next() {
		props, con := constQuery.Get()

		props.MaxNectar = con.Nectar
		props.MaxPollen = con.Pollen
		props.DetectionProbability = con.DetectionProbability
		props.NectarConcentration = con.NectarConcentration
	}

	seasonalQuery := s.seasonalFilter.Query()
	for seasonalQuery.Next() {
		props, seas := seasonalQuery.Get()

		day := (s.time.Tick + int64(seas.SeasonShift)) % 365
		season := util.Season(day)

		props.MaxNectar = seas.MaxNectar * season
		props.MaxPollen = seas.MaxPollen * season

		props.DetectionProbability = seas.DetectionProbability
		props.NectarConcentration = seas.NectarConcentration
	}

	day := s.time.Tick % 365
	scriptedQuery := s.scriptedFilter.Query()
	for scriptedQuery.Next() {
		props, scr := scriptedQuery.Get()

		props.MaxNectar = util.Interpolate(scr.Nectar, float64(day), scr.Interpolation)
		props.MaxPollen = util.Interpolate(scr.Pollen, float64(day), scr.Interpolation)

		props.DetectionProbability = util.Interpolate(scr.DetectionProbability, float64(day), scr.Interpolation)
		props.NectarConcentration = util.Interpolate(scr.NectarConcentration, float64(day), scr.Interpolation)
	}

	query := s.filter.Query()
	for query.Next() {
		conf, res, visits := query.Get()

		res.MaxNectar = conf.MaxNectar * 1000 * 1000
		res.MaxPollen = conf.MaxPollen * 1000

		res.Nectar = res.MaxNectar
		res.Pollen = res.MaxPollen

		visits.Nectar = 0
		visits.Pollen = 0
	}
}

func (s *ReplenishPatches) Finalize(w *ecs.World) {}
