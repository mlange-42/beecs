package sys

import (
	"github.com/mlange-42/arche-model/resource"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/beecs/globals"
	"github.com/mlange-42/beecs/params"
	"github.com/mlange-42/beecs/util"
	"golang.org/x/exp/rand"
	"gonum.org/v1/gonum/stat/distuv"
)

// MortalityCohorts applies background mortality to all cohort-based development stages
// (i.e. all except foragers).
type MortalityCohorts struct {
	workerMort *params.WorkerMortality
	droneMort  *params.DroneMortality
	rng        *resource.Rand

	eggs   *globals.Eggs
	larvae *globals.Larvae
	pupae  *globals.Pupae
	inHive *globals.InHive
}

func (s *MortalityCohorts) Initialize(w *ecs.World) {
	s.workerMort = ecs.GetResource[params.WorkerMortality](w)
	s.droneMort = ecs.GetResource[params.DroneMortality](w)
	s.rng = ecs.GetResource[resource.Rand](w)

	s.eggs = ecs.GetResource[globals.Eggs](w)
	s.larvae = ecs.GetResource[globals.Larvae](w)
	s.pupae = ecs.GetResource[globals.Pupae](w)
	s.inHive = ecs.GetResource[globals.InHive](w)
}

func (s *MortalityCohorts) Update(w *ecs.World) {
	applyMortality(s.eggs.Workers, s.workerMort.Eggs, s.rng)
	applyMortality(s.eggs.Drones, s.droneMort.Eggs, s.rng)

	applyMortality(s.larvae.Workers, s.workerMort.Larvae, s.rng)
	applyMortality(s.larvae.Drones, s.droneMort.Larvae, s.rng)

	applyMortality(s.pupae.Workers, s.workerMort.Pupae, s.rng)
	applyMortality(s.pupae.Drones, s.droneMort.Pupae, s.rng)

	applyMortality(s.inHive.Workers, s.workerMort.InHive, s.rng)
	applyMortality(s.inHive.Drones, s.droneMort.InHive, s.rng)
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
