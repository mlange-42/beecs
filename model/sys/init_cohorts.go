package sys

import (
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/beecs/model/globals"
	"github.com/mlange-42/beecs/model/params"
)

type InitCohorts struct {
	eggs   globals.Eggs
	larvae globals.Larvae
	pupae  globals.Pupae
	inHive globals.InHive
}

func (s *InitCohorts) Initialize(w *ecs.World) {
	aff := ecs.GetResource[params.AgeFirstForaging](w)

	workerDev := ecs.GetResource[params.WorkerDevelopment](w)
	droneDev := ecs.GetResource[params.DroneDevelopment](w)

	s.eggs = globals.Eggs{
		Workers: make([]int, workerDev.EggTime),
		Drones:  make([]int, droneDev.EggTime),
	}
	ecs.AddResource(w, &s.eggs)

	s.larvae = globals.Larvae{
		Workers: make([]int, workerDev.LarvaeTime),
		Drones:  make([]int, droneDev.LarvaeTime),
	}
	ecs.AddResource(w, &s.larvae)

	s.pupae = globals.Pupae{
		Workers: make([]int, workerDev.PupaeTime),
		Drones:  make([]int, droneDev.PupaeTime),
	}
	ecs.AddResource(w, &s.pupae)

	s.inHive = globals.InHive{
		Workers: make([]int, aff.Max+1),
		Drones:  make([]int, droneDev.MaxLifespan),
	}
	ecs.AddResource(w, &s.inHive)
}

func (s *InitCohorts) Update(w *ecs.World) {}

func (s *InitCohorts) Finalize(w *ecs.World) {}
