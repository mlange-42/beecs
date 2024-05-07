package sys

import (
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
	"github.com/mlange-42/beecs/model/activity"
	"github.com/mlange-42/beecs/model/comp"
	"github.com/mlange-42/beecs/model/res"
)

type Foraging struct {
	pop          *res.PopulationStats
	foragePeriod *res.ForagingPeriod
	forageParams *res.ForagingParams
	stores       *res.Stores

	foragerFilter generic.Filter1[comp.Activity]

	maxHoneyStore float64
}

func (s *Foraging) Initialize(w *ecs.World) {
	s.pop = ecs.GetResource[res.PopulationStats](w)
	s.foragePeriod = ecs.GetResource[res.ForagingPeriod](w)
	s.forageParams = ecs.GetResource[res.ForagingParams](w)
	s.stores = ecs.GetResource[res.Stores](w)

	s.foragerFilter = *generic.NewFilter1[comp.Activity]()

	storeParams := ecs.GetResource[res.StoreParams](w)
	energyParams := ecs.GetResource[res.EnergyParams](w)

	s.maxHoneyStore = storeParams.MaxHoneyStoreKg * 1000.0 * energyParams.EnergyHoney
}

func (s *Foraging) Update(w *ecs.World) {
	if s.foragePeriod.SecondsToday <= 0 ||
		(s.stores.Honey >= 0.95*s.maxHoneyStore && s.stores.Pollen >= s.stores.IdealPollen) {
		return
	}

	hangAroundDuration := s.forageParams.SearchLength / s.forageParams.FlightVelocity
	forageProb := s.calcForagingProb(s.stores.DecentHoney, s.stores.IdealPollen)

	round := 0
	totalDuration := 0.0
	for {
		duration, foragers := s.foragingRound(forageProb)
		meanDuration := 0.0
		if foragers > 0 {
			meanDuration = duration / float64(foragers)
		} else {
			meanDuration = hangAroundDuration
		}
		totalDuration += meanDuration

		if totalDuration > float64(s.foragePeriod.SecondsToday) {
			break
		}

		round++
	}

	query := s.foragerFilter.Query(w)
	for query.Next() {
		act := query.Get()
		act.Current = activity.Resting
	}
}

func (s *Foraging) Finalize(w *ecs.World) {}

func (s *Foraging) calcForagingProb(decentHoney, idealPollen float64) float64 {
	if s.stores.Pollen/idealPollen > 0.5 && s.stores.Honey/decentHoney > 1 {
		return 0
	}
	prob := s.forageParams.ProbBase
	if s.stores.Pollen/idealPollen < 0.2 || s.stores.Honey/decentHoney < 0.5 {
		prob = s.forageParams.ProbHigh
	}
	if s.stores.Honey/decentHoney < 0.2 {
		prob = s.forageParams.ProbEmergency
	}
	return prob
}

func (s *Foraging) foragingRound(forageProb float64) (duration float64, foragers int) {
	duration, foragers = 0.0, 0

	probCollectPollen := (1.0 - s.stores.Pollen/s.stores.IdealPollen) * s.forageParams.MaxProportionPollenForagers

	if s.stores.Honey/s.stores.DecentHoney < 0.5 {
		probCollectPollen *= s.stores.Honey / s.stores.DecentHoney
	}

	return 17 * 60, 1 // TODO!
}
