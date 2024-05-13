package sys

import (
	"math"

	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
	"github.com/mlange-42/beecs/comp"
	"github.com/mlange-42/beecs/params"
	"github.com/mlange-42/beecs/util"
)

type UpdatePatchesForaging struct {
	forageParams      *params.Foraging
	foragerParams     *params.Foragers
	handingTimeParams *params.HandlingTime
	energyParams      *params.EnergyContent
	danceParams       *params.Dance

	filter generic.Filter6[comp.PatchConfig, comp.Resource, comp.HandlingTime, comp.Trip, comp.Mortality, comp.Dance]
}

func (s *UpdatePatchesForaging) Initialize(w *ecs.World) {
	s.forageParams = ecs.GetResource[params.Foraging](w)
	s.foragerParams = ecs.GetResource[params.Foragers](w)
	s.handingTimeParams = ecs.GetResource[params.HandlingTime](w)
	s.energyParams = ecs.GetResource[params.EnergyContent](w)
	s.danceParams = ecs.GetResource[params.Dance](w)

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
