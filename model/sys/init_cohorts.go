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

	EggTimeWorker       int
	LarvaeTimeWorker    int
	PupaeTimeWorker     int
	MaxInHiveTimeWorker int

	EggTimeDrone    int
	LarvaeTimeDrone int
	PupaeTimeDrone  int
	LifespanDrone   int
}

func (s *InitCohorts) Initialize(w *ecs.World) {
	s.eggs = res.Eggs{
		Workers: make([]int, s.EggTimeWorker),
		Drones:  make([]int, s.EggTimeDrone),
	}
	ecs.AddResource(w, &s.eggs)

	s.larvae = res.Larvae{
		Workers: make([]int, s.LarvaeTimeWorker),
		Drones:  make([]int, s.LarvaeTimeDrone),
	}
	ecs.AddResource(w, &s.larvae)

	s.pupae = res.Pupae{
		Workers: make([]int, s.PupaeTimeWorker),
		Drones:  make([]int, s.PupaeTimeDrone),
	}
	ecs.AddResource(w, &s.pupae)

	s.inHive = res.InHive{
		Workers: make([]int, s.MaxInHiveTimeWorker),
		Drones:  make([]int, s.LifespanDrone),
	}
	ecs.AddResource(w, &s.inHive)
}

func (s *InitCohorts) Update(w *ecs.World) {}

func (s *InitCohorts) Finalize(w *ecs.World) {}
