package sys

import (
	"math"

	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche-model/resource"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
	"github.com/mlange-42/beecs/model/activity"
	"github.com/mlange-42/beecs/model/comp"
	"github.com/mlange-42/beecs/model/globals"
	"github.com/mlange-42/beecs/model/params"
	"golang.org/x/exp/rand"
	"gonum.org/v1/gonum/stat/distuv"
)

type Foraging struct {
	PatchUpdater model.System
	rng          *rand.Rand

	foragerParams      *params.ForagerParams
	forageParams       *params.ForagingParams
	handlingTimeParams *params.HandlingTimeParams
	danceParams        *params.DanceParams
	energyParams       *params.EnergyParams
	storeParams        *params.StoreParams

	foragePeriod *globals.ForagingPeriod
	stores       *globals.Stores
	pop          *globals.PopulationStats

	patches  []patchCandidate
	toRemove []ecs.Entity
	resting  []ecs.Entity
	dances   []ecs.Entity

	patchResourceMapper  generic.Map1[comp.Resource]
	patchDanceMapper     generic.Map2[comp.Resource, comp.Dance]
	patchTripMapper      generic.Map1[comp.Trip]
	patchMortalityMapper generic.Map1[comp.Mortality]
	patchConfigMapper    generic.Map2[comp.PatchConfig, comp.Trip]
	foragerMapper        generic.Map2[comp.Activity, comp.KnownPatch]

	activityFilter      generic.Filter1[comp.Activity]
	loadFilter          generic.Filter2[comp.Activity, comp.NectarLoad]
	foragerFilter       generic.Filter3[comp.Activity, comp.KnownPatch, comp.Milage]
	foragerFilterLoad   generic.Filter4[comp.Activity, comp.KnownPatch, comp.Milage, comp.NectarLoad]
	foragerFilterSimple generic.Filter2[comp.Activity, comp.KnownPatch]
	patchFilter         generic.Filter2[comp.Resource, comp.PatchConfig]

	maxHoneyStore float64
}

func (s *Foraging) Initialize(w *ecs.World) {
	s.foragerParams = ecs.GetResource[params.ForagerParams](w)
	s.forageParams = ecs.GetResource[params.ForagingParams](w)
	s.handlingTimeParams = ecs.GetResource[params.HandlingTimeParams](w)
	s.danceParams = ecs.GetResource[params.DanceParams](w)
	s.energyParams = ecs.GetResource[params.EnergyParams](w)
	s.storeParams = ecs.GetResource[params.StoreParams](w)

	s.pop = ecs.GetResource[globals.PopulationStats](w)
	s.foragePeriod = ecs.GetResource[globals.ForagingPeriod](w)
	s.stores = ecs.GetResource[globals.Stores](w)

	s.activityFilter = *generic.NewFilter1[comp.Activity]()
	s.loadFilter = *generic.NewFilter2[comp.Activity, comp.NectarLoad]()
	s.foragerFilter = *generic.NewFilter3[comp.Activity, comp.KnownPatch, comp.Milage]()
	s.foragerFilterLoad = *generic.NewFilter4[comp.Activity, comp.KnownPatch, comp.Milage, comp.NectarLoad]()
	s.foragerFilterSimple = *generic.NewFilter2[comp.Activity, comp.KnownPatch]()
	s.patchFilter = *generic.NewFilter2[comp.Resource, comp.PatchConfig]()

	s.patchResourceMapper = generic.NewMap1[comp.Resource](w)
	s.patchDanceMapper = generic.NewMap2[comp.Resource, comp.Dance](w)
	s.patchTripMapper = generic.NewMap1[comp.Trip](w)
	s.patchMortalityMapper = generic.NewMap1[comp.Mortality](w)
	s.patchConfigMapper = generic.NewMap2[comp.PatchConfig, comp.Trip](w)
	s.foragerMapper = generic.NewMap2[comp.Activity, comp.KnownPatch](w)

	storeParams := ecs.GetResource[params.StoreParams](w)
	energyParams := ecs.GetResource[params.EnergyParams](w)

	s.maxHoneyStore = storeParams.MaxHoneyStoreKg * 1000.0 * energyParams.Honey
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

	hangAroundDuration := s.forageParams.SearchLength / s.foragerParams.FlightVelocity
	forageProb := s.calcForagingProb()

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

func (s *Foraging) calcForagingProb() float64 {
	if s.stores.Pollen/s.stores.IdealPollen > 0.5 && s.stores.Honey/s.stores.DecentHoney > 1 {
		return 0
	}
	prob := s.forageParams.ProbBase
	if s.stores.Pollen/s.stores.IdealPollen < 0.2 || s.stores.Honey/s.stores.DecentHoney < 0.5 {
		prob = s.forageParams.ProbHigh
	}
	if s.stores.Honey/s.stores.DecentHoney < 0.2 {
		prob = s.forageParams.ProbEmergency
	}
	return prob
}

func (s *Foraging) foragingRound(w *ecs.World, forageProb float64) (duration float64, foragers int) {
	probCollectPollen := (1.0 - s.stores.Pollen/s.stores.IdealPollen) * s.danceParams.MaxProportionPollenForagers

	if s.stores.Honey/s.stores.DecentHoney < 0.5 {
		probCollectPollen *= s.stores.Honey / s.stores.DecentHoney
	}

	s.PatchUpdater.Update(w)

	s.decisions(w, forageProb, probCollectPollen)
	s.searching(w)
	s.collecting(w)
	duration, foragers = s.flightCost(w)
	s.mortality(w)
	s.dancing(w)
	s.unloading(w)
	return
}

func (s *Foraging) decisions(w *ecs.World, probForage, probCollectPollen float64) {
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

		if !act.PollenForager && !patch.Nectar.IsZero() {
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

		if milage.Today > float32(s.foragerParams.MaxKmPerDay) {
			act.Current = activity.Resting
		}
	}
}

func (s *Foraging) searching(w *ecs.World) {
	cumProb := 0.0
	nonDetectionProb := 1.0

	sz := float64(s.foragerParams.SquadronSize)
	patchQuery := s.patchFilter.Query(w)
	for patchQuery.Next() {
		res, conf := patchQuery.Get()
		hasNectar := res.Nectar >= s.foragerParams.NectarLoad*sz
		hasPollen := res.Pollen >= s.foragerParams.PollenLoad*sz
		if !hasNectar && !hasPollen {
			continue
		}
		s.patches = append(s.patches, patchCandidate{
			Patch:       patchQuery.Entity(),
			Probability: conf.DetectionProbability,
			HasNectar:   hasNectar,
			HasPollen:   hasPollen,
		})

		cumProb += conf.DetectionProbability
		nonDetectionProb *= 1.0 - conf.DetectionProbability
	}
	detectionProb := 1.0 - nonDetectionProb

	// TODO: shuffle foragers
	foragerQuery := s.foragerFilterSimple.Query(w)
	for foragerQuery.Next() {
		act, patch := foragerQuery.Get()

		if act.Current == activity.Searching {
			if s.rng.Float64() >= detectionProb {
				continue
			}
			p := s.rng.Float64() * cumProb
			cum := 0.0
			var selected patchCandidate
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
					res.Pollen -= s.foragerParams.PollenLoad * sz
				} else {
					patch.Pollen = ecs.Entity{}
				}
			} else {
				if selected.HasNectar {
					patch.Nectar = selected.Patch
					act.Current = activity.BringNectar
					res := s.patchResourceMapper.Get(selected.Patch)
					res.Nectar -= s.foragerParams.NectarLoad * sz
				} else {
					patch.Nectar = ecs.Entity{}
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
				if res.Nectar >= s.foragerParams.NectarLoad*sz {
					res.Nectar -= s.foragerParams.NectarLoad * sz
					act.Current = activity.BringNectar
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
				if res.Pollen >= s.foragerParams.PollenLoad*sz {
					res.Pollen -= s.foragerParams.PollenLoad * sz
					act.Current = activity.BringPollen
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

func (s *Foraging) collecting(w *ecs.World) {
	sz := float64(s.foragerParams.SquadronSize)
	foragerQuery := s.foragerFilterLoad.Query(w)
	for foragerQuery.Next() {
		act, patch, milage, load := foragerQuery.Get()

		if act.Current == activity.Experienced {
			if act.PollenForager {
				if patch.Pollen.IsZero() {
					act.Current = activity.Resting
				} else {
					res := s.patchResourceMapper.Get(patch.Pollen)
					if res.Pollen >= s.foragerParams.PollenLoad*sz {
						act.Current = activity.BringPollen
						res.Pollen -= s.foragerParams.PollenLoad * sz
					} else {
						act.Current = activity.Searching
						patch.Pollen = ecs.Entity{}
					}
				}
			} else {
				if patch.Nectar.IsZero() {
					act.Current = activity.Resting
				} else {
					res := s.patchResourceMapper.Get(patch.Nectar)
					if res.Nectar >= s.foragerParams.NectarLoad*sz {
						act.Current = activity.BringNectar
						res.Nectar -= s.foragerParams.NectarLoad * sz
					} else {
						act.Current = activity.Searching
						patch.Nectar = ecs.Entity{}
					}
				}
			}
		}

		if act.Current == activity.BringNectar {
			conf, trip := s.patchConfigMapper.Get(patch.Nectar)
			load.Energy = conf.NectarConcentration * s.foragerParams.NectarLoad * s.energyParams.Scurose
			dist := trip.CostNectar / (s.foragerParams.FlightCostPerM * 1000)
			milage.Today += float32(dist)
			milage.Total += float32(dist)
		}

		if act.Current == activity.BringPollen {
			trip := s.patchTripMapper.Get(patch.Pollen)
			dist := trip.CostPollen / (s.foragerParams.FlightCostPerM * 1000)
			milage.Today += float32(dist)
			milage.Total += float32(dist)
		}
	}
}

func (s *Foraging) flightCost(w *ecs.World) (duration float64, foragers int) {
	duration = 0.0
	foragers = 0

	query := s.foragerFilter.Query(w)
	for query.Next() {
		act, patch, milage := query.Get()

		if act.Current == activity.Searching {
			dist := s.forageParams.SearchLength / 1000.0
			milage.Today += float32(dist)
			milage.Total += float32(dist)

			en := s.forageParams.SearchLength * s.foragerParams.FlightCostPerM
			s.stores.Honey -= en * float64(s.foragerParams.SquadronSize)

			duration += s.forageParams.SearchLength / s.foragerParams.FlightVelocity
			foragers += 1
		} else if act.Current == activity.BringNectar || act.Current == activity.BringPollen {
			en := 0.0
			if act.PollenForager {
				trip := s.patchTripMapper.Get(patch.Pollen)
				duration += trip.DurationPollen + s.handlingTimeParams.PollenUnloading
				en = trip.CostPollen
			} else {
				trip := s.patchTripMapper.Get(patch.Nectar)
				duration += trip.DurationNectar + s.handlingTimeParams.NectarUnloading
				en = trip.CostNectar
			}
			s.stores.Honey -= en * float64(s.foragerParams.SquadronSize)
			foragers += 1
		}
	}

	return
}

func (s *Foraging) mortality(w *ecs.World) {
	searchDuration := s.forageParams.SearchLength / s.foragerParams.FlightVelocity

	foragerQuery := s.foragerFilterSimple.Query(w)
	for foragerQuery.Next() {
		act, patch := foragerQuery.Get()

		if act.Current == activity.Searching {
			if s.rng.Float64() < 1-math.Pow(1-s.forageParams.MortalityPerSec, searchDuration) {
				s.toRemove = append(s.toRemove, foragerQuery.Entity())
			}
		} else if act.Current == activity.BringNectar {
			m := s.patchMortalityMapper.Get(patch.Nectar)
			if s.rng.Float64() < m.Nectar {
				s.toRemove = append(s.toRemove, foragerQuery.Entity())
			}
		} else if act.Current == activity.BringPollen {
			m := s.patchMortalityMapper.Get(patch.Pollen)
			if s.rng.Float64() < m.Pollen {
				s.toRemove = append(s.toRemove, foragerQuery.Entity())
			}
		}
	}

	for _, e := range s.toRemove {
		w.RemoveEntity(e)
	}
	s.toRemove = s.toRemove[:0]
}

func (s *Foraging) dancing(w *ecs.World) {
	activityQuery := s.activityFilter.Query(w)
	for activityQuery.Next() {
		act := activityQuery.Get()
		if act.Current == activity.Resting {
			s.resting = append(s.resting, activityQuery.Entity())
		} else if act.Current == activity.BringNectar || act.Current == activity.BringPollen {
			s.dances = append(s.dances, activityQuery.Entity())
		}
	}
	s.rng.Shuffle(len(s.resting), func(i, j int) { s.resting[i], s.resting[j] = s.resting[j], s.resting[i] })
	s.rng.Shuffle(len(s.dances), func(i, j int) { s.dances[i], s.dances[j] = s.dances[j], s.dances[i] })

	for _, e := range s.dances {
		act, patch := s.foragerMapper.Get(e)

		if act.Current != activity.BringNectar && act.Current != activity.BringPollen {
			continue
		}

		if act.Current == activity.BringNectar {
			patchRes, dance := s.patchDanceMapper.Get(patch.Nectar)
			danceEEF := patchRes.EnergyEfficiency

			rPoisson := distuv.Poisson{
				Src:    s.rng,
				Lambda: dance.Circuits * 0.05,
			}
			danceFollowers := int(rPoisson.Rand())

			if danceFollowers == 0 {
				continue
			}
			if len(s.resting) < danceFollowers {
				continue
			}

			for i := 0; i < danceFollowers; i++ {
				follower := s.resting[len(s.resting)-1]
				fAct, fPatch := s.foragerMapper.Get(follower)

				if fPatch.Nectar.IsZero() {
					fPatch.Nectar = patch.Nectar
					fAct.Current = activity.Recruited
					fAct.PollenForager = false
				} else {
					knownRes := s.patchResourceMapper.Get(fPatch.Nectar)
					if danceEEF > knownRes.EnergyEfficiency {
						fPatch.Nectar = patch.Nectar
						fAct.Current = activity.Recruited
						fAct.PollenForager = false
					} else {
						// TODO: really? was resting before!
						fAct.Current = activity.Experienced
					}
				}

				s.resting = s.resting[:len(s.resting)-1]
			}
		}

		if act.Current == activity.BringPollen {
			trip := s.patchTripMapper.Get(patch.Pollen)
			danceTripDuration := trip.DurationPollen

			danceFollowers := s.danceParams.PollenDanceFollowers

			if len(s.resting) < danceFollowers {
				continue
			}

			for i := 0; i < danceFollowers; i++ {
				follower := s.resting[len(s.resting)-1]
				fAct, fPatch := s.foragerMapper.Get(follower)

				if fPatch.Pollen.IsZero() {
					fPatch.Pollen = patch.Pollen
					fAct.Current = activity.Recruited
					fAct.PollenForager = true
				} else {
					knownTrip := s.patchTripMapper.Get(fPatch.Pollen)
					if danceTripDuration < knownTrip.DurationPollen {
						fPatch.Pollen = patch.Pollen
						fAct.Current = activity.Recruited
						fAct.PollenForager = true
					} else {
						// TODO: really? was resting before!
						fAct.Current = activity.Experienced
					}
				}

				s.resting = s.resting[:len(s.resting)-1]
			}
		}
	}

	s.resting = s.resting[:0]
	s.dances = s.dances[:0]
}

func (s *Foraging) unloading(w *ecs.World) {
	query := s.loadFilter.Query(w)
	for query.Next() {
		act, load := query.Get()
		if act.Current == activity.BringNectar {
			s.stores.Honey = math.Min(
				s.stores.Honey+load.Energy*float64(s.foragerParams.SquadronSize),
				s.maxHoneyStore,
			)
			load.Energy = 0
			act.Current = activity.Experienced
		} else if act.Current == activity.BringPollen {
			s.stores.Pollen += s.foragerParams.PollenLoad * float64(s.foragerParams.SquadronSize)
			act.Current = activity.Experienced
		}
	}
}

type patchCandidate struct {
	Patch       ecs.Entity
	Probability float64
	HasNectar   bool
	HasPollen   bool
}
