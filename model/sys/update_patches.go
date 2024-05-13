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
	forageParams      *res.ForagingParams
	foragerParams     *res.ForagerParams
	handingTimeParams *res.HandlingTimeParams
	energyParams      *res.EnergyParams
	danceParams       *res.DanceParams

	filter generic.Filter6[comp.PatchConfig, comp.Resource, comp.HandlingTime, comp.Trip, comp.Mortality, comp.Dance]
}

func (s *UpdatePatchesForaging) Initialize(w *ecs.World) {
	s.forageParams = ecs.GetResource[res.ForagingParams](w)
	s.foragerParams = ecs.GetResource[res.ForagerParams](w)
	s.handingTimeParams = ecs.GetResource[res.HandlingTimeParams](w)
	s.energyParams = ecs.GetResource[res.EnergyParams](w)
	s.danceParams = ecs.GetResource[res.DanceParams](w)

	s.filter = *generic.NewFilter6[comp.PatchConfig, comp.Resource, comp.HandlingTime, comp.Trip, comp.Mortality, comp.Dance]()
}

func (s *UpdatePatchesForaging) Update(w *ecs.World) {
	query := s.filter.Query(w)
	for query.Next() {
		conf, r, ht, trip, mort, dance := query.Get()

		if s.handingTimeParams.ConstantHandlingTime {
			ht.Pollen = s.handingTimeParams.PollenGathering
			ht.Nectar = s.handingTimeParams.NectarGathering
		} else {
			ht.Pollen = s.handingTimeParams.PollenGathering * r.MaxPollen / r.Pollen
			ht.Nectar = s.handingTimeParams.NectarGathering * r.MaxNectar / r.Nectar
		}

		trip.CostNectar = (2 * conf.DistToColony * s.foragerParams.FlightCostPerM) +
			(s.foragerParams.FlightCostPerM * ht.Nectar *
				s.foragerParams.FlightVelocity * s.forageParams.EnergyOnFlower) // [kJ] = [m*kJ/m + kJ/m * s * m/s]

		trip.CostPollen = (2 * conf.DistToColony * s.foragerParams.FlightCostPerM) +
			(s.foragerParams.FlightCostPerM * ht.Pollen *
				s.foragerParams.FlightVelocity * s.forageParams.EnergyOnFlower) // [kJ] = [m*kJ/m + kJ/m * s * m/s]

		r.EnergyEfficiency = (conf.NectarConcentration*s.foragerParams.NectarLoad*s.energyParams.Scurose - trip.CostNectar) / trip.CostNectar

		trip.DurationNectar = 2*conf.DistToColony/s.foragerParams.FlightVelocity + ht.Nectar
		trip.DurationPollen = 2*conf.DistToColony/s.foragerParams.FlightVelocity + ht.Pollen

		mort.Nectar = 1.0 - (math.Pow(1.0-s.forageParams.MortalityPerSec, trip.DurationNectar))
		mort.Pollen = 1.0 - (math.Pow(1.0-s.forageParams.MortalityPerSec, trip.DurationPollen))

		circ := r.EnergyEfficiency*s.danceParams.Slope + s.danceParams.Intercept
		dance.Circuits = util.Clamp(circ, 0, float64(s.danceParams.MaxCircuits))
	}
}

func (s *UpdatePatchesForaging) Finalize(w *ecs.World) {}
