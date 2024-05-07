package sys

import (
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/beecs/model/res"
)

type InitCohorts struct {
	eggs   res.Eggs
	larvae res.Larvae
	pupae  res.Pupae
	inHive res.InHive
}

func (s *InitCohorts) Initialize(w *ecs.World) {
	aff := ecs.GetResource[res.AgeFirstForagingParams](w)

	workerDev := ecs.GetResource[res.WorkerDevelopment](w)
	droneDev := ecs.GetResource[res.DroneDevelopment](w)

	s.eggs = res.Eggs{
		Workers: make([]int, workerDev.EggTime),
		Drones:  make([]int, droneDev.EggTime),
	}
	ecs.AddResource(w, &s.eggs)

	s.larvae = res.Larvae{
		Workers: make([]int, workerDev.LarvaeTime),
		Drones:  make([]int, droneDev.LarvaeTime),
	}
	ecs.AddResource(w, &s.larvae)

	s.pupae = res.Pupae{
		Workers: make([]int, workerDev.PupaeTime),
		Drones:  make([]int, droneDev.PupaeTime),
	}
	ecs.AddResource(w, &s.pupae)

	s.inHive = res.InHive{
		Workers: make([]int, aff.Max+1),
		Drones:  make([]int, droneDev.MaxLifespan),
	}
	ecs.AddResource(w, &s.inHive)
}

func (s *InitCohorts) Update(w *ecs.World) {}

func (s *InitCohorts) Finalize(w *ecs.World) {}
