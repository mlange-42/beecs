package sys

import (
	"github.com/mlange-42/arche-model/resource"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/beecs/model/res"
	"github.com/mlange-42/beecs/model/util"
	"golang.org/x/exp/rand"
	"gonum.org/v1/gonum/stat/distuv"
)

type MortalityCohorts struct {
	eggs   *res.Eggs
	larvae *res.Larvae
	pupae  *res.Pupae
	inHive *res.InHive
	rng    *resource.Rand

	EggMortalityWorker    float64
	LarvaeMortalityWorker float64
	PupaeMortalityWorker  float64
	InHiveMortalityWorker float64

	EggMortalityDrone    float64
	LarvaeMortalityDrone float64
	PupaeMortalityDrone  float64
	InHiveMortalityDrone float64
}

func (s *MortalityCohorts) Initialize(w *ecs.World) {
	s.eggs = ecs.GetResource[res.Eggs](w)
	s.larvae = ecs.GetResource[res.Larvae](w)
	s.pupae = ecs.GetResource[res.Pupae](w)
	s.inHive = ecs.GetResource[res.InHive](w)
	s.rng = ecs.GetResource[resource.Rand](w)
}

func (s *MortalityCohorts) Update(w *ecs.World) {
	applyMortality(s.eggs.Workers, s.EggMortalityWorker, s.rng)
	applyMortality(s.eggs.Drones, s.EggMortalityDrone, s.rng)

	applyMortality(s.larvae.Workers, s.LarvaeMortalityWorker, s.rng)
	applyMortality(s.larvae.Drones, s.LarvaeMortalityDrone, s.rng)

	applyMortality(s.pupae.Workers, s.PupaeMortalityWorker, s.rng)
	applyMortality(s.pupae.Drones, s.PupaeMortalityDrone, s.rng)

	applyMortality(s.inHive.Workers, s.InHiveMortalityWorker, s.rng)
	applyMortality(s.inHive.Drones, s.InHiveMortalityDrone, s.rng)
}

func (s *MortalityCohorts) Finalize(w *ecs.World) {}

func applyMortality(coh []int, m float64, rng rand.Source) {
	for i := range coh {
		num := coh[i]
		rng := distuv.Poisson{
			Src:    rng,
			Lambda: m * float64(num),
		}
		toDie := int(rng.Rand())
		coh[i] = util.MaxInt(0, num-toDie)
	}
}
