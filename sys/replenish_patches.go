package sys

import (
	"github.com/mlange-42/arche-model/resource"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
	"github.com/mlange-42/beecs/comp"
	"github.com/mlange-42/beecs/util"
)

type ReplenishPatches struct {
	time   *resource.Tick
	filter generic.Filter2[comp.PatchProperties, comp.Resource]

	constantFilter generic.Filter2[comp.PatchProperties, comp.ConstantPatch]
	seasonalFilter generic.Filter2[comp.PatchProperties, comp.SeasonalPatch]
	scriptedFilter generic.Filter2[comp.PatchProperties, comp.ScriptedPatch]
}

func (s *ReplenishPatches) Initialize(w *ecs.World) {
	s.time = ecs.GetResource[resource.Tick](w)
	s.filter = *generic.NewFilter2[comp.PatchProperties, comp.Resource]()

	s.constantFilter = *generic.NewFilter2[comp.PatchProperties, comp.ConstantPatch]()
	s.seasonalFilter = *generic.NewFilter2[comp.PatchProperties, comp.SeasonalPatch]()
	s.scriptedFilter = *generic.NewFilter2[comp.PatchProperties, comp.ScriptedPatch]()
}

func (s *ReplenishPatches) Update(w *ecs.World) {
	constQuery := s.constantFilter.Query(w)
	for constQuery.Next() {
		props, con := constQuery.Get()

		props.MaxNectar = con.Nectar
		props.MaxPollen = con.Pollen
		props.DetectionProbability = con.DetectionProbability
		props.NectarConcentration = con.NectarConcentration
	}

	seasonalQuery := s.seasonalFilter.Query(w)
	for seasonalQuery.Next() {
		props, seas := seasonalQuery.Get()

		day := (s.time.Tick + int64(seas.SeasonShift)) % 365
		season := util.Season(day)

		props.MaxNectar = seas.MaxNectar * season
		props.MaxPollen = seas.MaxNectar * season

		props.DetectionProbability = seas.DetectionProbability
		props.NectarConcentration = seas.NectarConcentration
	}

	day := s.time.Tick % 365
	scriptedQuery := s.scriptedFilter.Query(w)
	for scriptedQuery.Next() {
		props, scr := scriptedQuery.Get()

		props.MaxNectar = util.Interpolate(scr.Nectar, float64(day), scr.Interpolation)
		props.MaxPollen = util.Interpolate(scr.Pollen, float64(day), scr.Interpolation)

		props.DetectionProbability = util.Interpolate(scr.DetectionProbability, float64(day), scr.Interpolation)
		props.NectarConcentration = util.Interpolate(scr.NectarConcentration, float64(day), scr.Interpolation)
	}

	query := s.filter.Query(w)
	for query.Next() {
		conf, res := query.Get()

		res.MaxNectar = conf.MaxNectar * 1000 * 1000
		res.MaxPollen = conf.MaxPollen * 1000

		res.Nectar = res.MaxNectar
		res.Pollen = res.MaxPollen
	}
}

func (s *ReplenishPatches) Finalize(w *ecs.World) {}
