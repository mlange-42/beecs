package sys

import (
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/beecs/model/res"
)

type EggLaying struct {
	eggs *res.Eggs

	MaxEggsPerDay       int
	DroneEggsProportion float64
}

func (s *EggLaying) Initialize(w *ecs.World) {
	s.eggs = ecs.GetResource[res.Eggs](w)
}

func (s *EggLaying) Update(w *ecs.World) {
	// TODO: limit number of eggs
	eggs := s.MaxEggsPerDay
	// TODO: limit to drone eggs period
	droneEggs := int(s.DroneEggsProportion * float64(eggs))
	s.eggs.Workers[0] = eggs
	s.eggs.Drones[0] = droneEggs
}

func (s *EggLaying) Finalize(w *ecs.World) {}
