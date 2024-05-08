package sys

import (
	"math"

	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
	"github.com/mlange-42/beecs/model/comp"
	"github.com/mlange-42/beecs/model/res"
	"github.com/mlange-42/beecs/model/util"
)

type UpdatePatchesForaging struct {
	params       *res.ForagingParams
	energyParams *res.EnergyParams
	danceParams  *res.DanceParams

	filter generic.Filter6[comp.PatchConfig, comp.Resource, comp.HandlingTime, comp.Trip, comp.Mortality, comp.Dance]
}

func (s *UpdatePatchesForaging) Initialize(w *ecs.World) {
	s.params = ecs.GetResource[res.ForagingParams](w)
	s.energyParams = ecs.GetResource[res.EnergyParams](w)
	s.danceParams = ecs.GetResource[res.DanceParams](w)

	s.filter = *generic.NewFilter6[comp.PatchConfig, comp.Resource, comp.HandlingTime, comp.Trip, comp.Mortality, comp.Dance]()
}

func (s *UpdatePatchesForaging) Update(w *ecs.World) {
	query := s.filter.Query(w)
	for query.Next() {
		conf, r, ht, trip, mort, dance := query.Get()

		if s.params.ConstantHandlingTime {
			ht.Pollen = s.params.TimePollenGathering
			ht.Nectar = s.params.TimeNectarGathering
		} else {
			ht.Pollen = s.params.TimePollenGathering * r.MaxPollen / r.Pollen
			ht.Nectar = s.params.TimeNectarGathering * r.MaxNectar / r.Nectar
		}

		trip.CostNectar = (2 * conf.DistToColony * s.params.FlightCostPerM) +
			(s.params.FlightCostPerM * ht.Nectar *
				s.params.FlightVelocity * s.params.EnergyOnFlower) // [kJ] = [m*kJ/m + kJ/m * s * m/s]

		trip.CostPollen = (2 * conf.DistToColony * s.params.FlightCostPerM) +
			(s.params.FlightCostPerM * ht.Pollen *
				s.params.FlightVelocity * s.params.EnergyOnFlower) // [kJ] = [m*kJ/m + kJ/m * s * m/s]

		r.EnergyEfficiency = (conf.NectarConcentration*s.params.VolumeCarried*s.energyParams.EnergyScurose - trip.CostNectar) / trip.CostNectar

		trip.DurationNectar = 2*conf.DistToColony/s.params.FlightVelocity + ht.Nectar
		trip.DurationPollen = 2*conf.DistToColony/s.params.FlightVelocity + ht.Pollen

		mort.Nectar = 1.0 - (math.Pow(1.0-s.params.MortalityPerSec, trip.DurationNectar))
		mort.Pollen = 1.0 - (math.Pow(1.0-s.params.MortalityPerSec, trip.DurationPollen))

		circ := r.EnergyEfficiency*s.danceParams.Slope + s.danceParams.Intercept
		dance.Circuits = util.Clamp(circ, 0, float64(s.danceParams.MaxCircuits))
	}
}

func (s *UpdatePatchesForaging) Finalize(w *ecs.World) {}
