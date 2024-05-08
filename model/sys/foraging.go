package sys

import (
	"math"

	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche-model/resource"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
	"github.com/mlange-42/beecs/model/activity"
	"github.com/mlange-42/beecs/model/comp"
	"github.com/mlange-42/beecs/model/res"
	"golang.org/x/exp/rand"
)

type Foraging struct {
	PatchUpdater model.System
	rng          *rand.Rand
	params       *res.Params
	foragePeriod *res.ForagingPeriod
	forageParams *res.ForagingParams
	danceParams  *res.DanceParams
	stores       *res.Stores
	pop          *res.PopulationStats

	patches []PatchCandidate

	patchResourceMapper generic.Map1[comp.Resource]
	patchTripMapper     generic.Map1[comp.Trip]
	foragerFilter       generic.Filter3[comp.Activity, comp.KnownPatch, comp.Milage]
	patchFilter         generic.Filter2[comp.Resource, comp.PatchConfig]

	maxHoneyStore float64
}

func (s *Foraging) Initialize(w *ecs.World) {
	s.pop = ecs.GetResource[res.PopulationStats](w)
	s.params = ecs.GetResource[res.Params](w)
	s.foragePeriod = ecs.GetResource[res.ForagingPeriod](w)
	s.forageParams = ecs.GetResource[res.ForagingParams](w)
	s.danceParams = ecs.GetResource[res.DanceParams](w)
	s.stores = ecs.GetResource[res.Stores](w)

	s.foragerFilter = *generic.NewFilter3[comp.Activity, comp.KnownPatch, comp.Milage]()
	s.patchFilter = *generic.NewFilter2[comp.Resource, comp.PatchConfig]()
	s.patchResourceMapper = generic.NewMap1[comp.Resource](w)
	s.patchTripMapper = generic.NewMap1[comp.Trip](w)

	storeParams := ecs.GetResource[res.StoreParams](w)
	energyParams := ecs.GetResource[res.EnergyParams](w)

	s.maxHoneyStore = storeParams.MaxHoneyStoreKg * 1000.0 * energyParams.EnergyHoney
	s.rng = rand.New(ecs.GetResource[resource.Rand](w))

	s.PatchUpdater.Initialize(w)
}

func (s *Foraging) Update(w *ecs.World) {
	if s.foragePeriod.SecondsToday <= 0 ||
		(s.stores.Honey >= 0.95*s.maxHoneyStore && s.stores.Pollen >= s.stores.IdealPollen) {
		return
	}

	query := s.foragerFilter.Query(w)
	for query.Next() {
		_, _, milage := query.Get()
		milage.Today = 0
	}

	hangAroundDuration := s.forageParams.SearchLength / s.forageParams.FlightVelocity
	forageProb := s.calcForagingProb(s.stores.DecentHoney, s.stores.IdealPollen)

	round := 0
	totalDuration := 0.0
	for {
		duration, foragers := s.foragingRound(w, forageProb)
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

	query = s.foragerFilter.Query(w)
	for query.Next() {
		act, _, _ := query.Get()
		act.Current = activity.Resting
	}
}

func (s *Foraging) Finalize(w *ecs.World) {
	s.PatchUpdater.Finalize(w)
}

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

func (s *Foraging) foragingRound(w *ecs.World, forageProb float64) (duration float64, foragers int) {
	//duration, foragers = 0.0, 0

	probCollectPollen := (1.0 - s.stores.Pollen/s.stores.IdealPollen) * s.forageParams.MaxProportionPollenForagers

	if s.stores.Honey/s.stores.DecentHoney < 0.5 {
		probCollectPollen *= s.stores.Honey / s.stores.DecentHoney
	}

	s.PatchUpdater.Update(w)

	s.foragerDecisions(w, forageProb, probCollectPollen)
	s.searching(w)

	_ = forageProb
	return 17 * 60, 1 // TODO!
}

func (s *Foraging) foragerDecisions(w *ecs.World, probForage, probCollectPollen float64) {
	query := s.foragerFilter.Query(w)
	for query.Next() {
		act, patch, milage := query.Get()

		if act.Current != activity.Recruited {
			act.PollenForager = s.rng.Float64() < probCollectPollen
		}

		if act.Current != activity.Recruited &&
			act.Current != activity.Resting &&
			act.Current != activity.Lazy {
			if s.rng.Float64() < s.forageParams.StopProbability {
				act.Current = activity.Resting
			}
		}

		if !patch.Nectar.IsZero() && !act.PollenForager {
			res := s.patchResourceMapper.Get(patch.Nectar)
			if s.rng.Float64() < 1.0/res.EnergyEfficiency &&
				s.rng.Float64() < s.stores.Honey/s.stores.DecentHoney {

				patch.Nectar = ecs.Entity{}
				if act.Current != activity.Resting && act.Current != activity.Lazy {
					act.Current = activity.Searching
				}
			}
		}

		if !patch.Pollen.IsZero() && act.PollenForager {
			trip := s.patchTripMapper.Get(patch.Pollen)
			if s.rng.Float64() < 1-math.Pow(1-s.forageParams.AbandonPollenPerSec, trip.DurationPollen) {

				patch.Nectar = ecs.Entity{}
				if act.Current != activity.Resting && act.Current != activity.Lazy {
					act.Current = activity.Searching
				}
			}
		}

		if act.Current == activity.Resting {
			if s.rng.Float64() < probForage {
				if act.PollenForager {
					if patch.Pollen.IsZero() {
						act.Current = activity.Searching
					} else {
						act.Current = activity.Experienced
					}
				} else {
					if patch.Nectar.IsZero() {
						act.Current = activity.Searching
					} else {
						act.Current = activity.Experienced
					}
				}
			}
		}

		if milage.Today > float32(s.forageParams.MaxKmPerDay) {
			act.Current = activity.Resting
		}
	}
}

func (s *Foraging) searching(w *ecs.World) {
	cumProb := 0.0
	nonDetectionProb := 1.0

	sz := float64(s.params.SquadronSize)
	patchQuery := s.patchFilter.Query(w)
	for patchQuery.Next() {
		res, conf := patchQuery.Get()
		hasNectar := res.Nectar >= s.forageParams.NectarLoad*sz
		hasPollen := res.Pollen >= s.forageParams.PollenLoad*sz
		if !hasNectar && !hasPollen {
			continue
		}
		s.patches = append(s.patches, PatchCandidate{
			Patch:       patchQuery.Entity(),
			Probability: conf.DetectionProbability,
			HasNectar:   hasNectar,
			HasPollen:   hasPollen,
		})

		cumProb += conf.DetectionProbability
		nonDetectionProb *= 1.0 - conf.DetectionProbability
	}
	detectionProb := 1.0 - nonDetectionProb

	foragerQuery := s.foragerFilter.Query(w)
	for foragerQuery.Next() {
		act, patch, _ := foragerQuery.Get()

		if act.Current == activity.Searching {
			if s.rng.Float64() < detectionProb {
				p := s.rng.Float64() * cumProb
				cum := 0.0
				var selected PatchCandidate
				for _, pch := range s.patches {
					cum += pch.Probability
					if cum >= p {
						selected = pch
						break
					}
				}
				if act.PollenForager {
					if selected.HasPollen {
						patch.Pollen = selected.Patch
						act.Current = activity.BringPollen
						res := s.patchResourceMapper.Get(selected.Patch)
						res.Pollen -= s.forageParams.PollenLoad * sz
					} else {
						patch.Pollen = ecs.Entity{}
					}
				} else {
					if selected.HasNectar {
						patch.Nectar = selected.Patch
						act.Current = activity.BringNectar
						res := s.patchResourceMapper.Get(selected.Patch)
						res.Nectar -= s.forageParams.NectarLoad * sz
					} else {
						patch.Nectar = ecs.Entity{}
					}
				}
			}
		}

		if act.Current != activity.Recruited {
			continue
		}

		if !act.PollenForager && !patch.Nectar.IsZero() {
			success := false
			if s.rng.Float64() < s.danceParams.FindProbability {
				res := s.patchResourceMapper.Get(patch.Nectar)
				if res.Nectar >= s.forageParams.NectarLoad*sz {
					res.Nectar -= s.forageParams.NectarLoad * sz
					success = true
				}
			}
			if !success {
				act.Current = activity.Searching
				patch.Nectar = ecs.Entity{}
			}
		}

		if act.PollenForager && !patch.Pollen.IsZero() {
			success := false
			if s.rng.Float64() < s.danceParams.FindProbability {
				res := s.patchResourceMapper.Get(patch.Pollen)
				if res.Pollen >= s.forageParams.PollenLoad*sz {
					res.Pollen -= s.forageParams.PollenLoad * sz
					success = true
				}
			}
			if !success {
				act.Current = activity.Searching
				patch.Pollen = ecs.Entity{}
			}
		}
	}

	s.patches = s.patches[:0]
}

type PatchCandidate struct {
	Patch       ecs.Entity
	Probability float64
	HasNectar   bool
	HasPollen   bool
}
