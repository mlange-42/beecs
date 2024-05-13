package sys

import (
	"github.com/mlange-42/arche-model/resource"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/beecs/model/globals"
)

var foragingHoursBerlin2000 = []float64{
	0, // TODO: Added this value as there is a shift by one day in the original BEEHAVE
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 7.2,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2.5, 0, 0, 0, 0,
	0, 0, 0, 10.7, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 7, 0, 7.9, 6.8, 4.7, 10.8, 11.2, 11.8,
	11.2, 9.9, 0, 10.7, 10.4, 4.2, 10.6, 8.7, 5.7, 13.3, 13.2, 12, 14, 14.1, 13.9,
	13.1, 10.7, 7.1, 13.7, 14.6, 15, 15.1, 15, 13.5, 10.3, 2.6, 5.9, 0, 6, 0, 8.4, 2.4,
	0.7, 12.1, 5.8, 6.8, 8.7, 6, 10, 8.7, 14.2, 12.3, 7.4, 3.4, 0.2, 7.2, 13.2, 15.8,
	13.9, 9.5, 11, 15.3, 4.1, 2.1, 6, 12.7, 10.4, 15.4, 15.1, 11.4, 8.5, 8, 1.5, 1.5,
	2.4, 2.6, 1.1, 0.1, 0, 9.5, 4.5, 2.4, 3.9, 1.3, 2.2, 8.3, 1.1, 3.4, 2.8, 5.1, 0.2,
	6.4, 0.5, 3.4, 5.2, 5.4, 0.1, 0, 1.5, 0, 0.5, 7.9, 9.8, 4.4, 1.6, 3.8, 2.1, 0.6, 1,
	1.5, 10.7, 3.8, 8.3, 7.1, 9.3, 12.7, 6.9, 3.6, 10.3, 3.3, 0.2, 5.7, 11.7, 13.4,
	7.8, 5.2, 9.5, 5, 4.2, 5.4, 2, 7.3, 8.5, 9, 4.7, 13.1, 10.5, 0, 7.5, 8.6, 4.3, 8,
	2.5, 0, 2.2, 1.2, 8.1, 2.8, 0, 0.4, 5.1, 1.2, 6.2, 2.1, 0.1, 5.1, 0.3, 0, 11.7, 0,
	0, 10.4, 6.5, 11.1, 11.3, 8.5, 1.2, 8.8, 5.6, 10.6, 10.3, 8.1, 3.7, 9.4, 2.2, 0.2,
	0, 0, 0, 0, 0, 2.2, 2.9, 2.7, 6.9, 0, 6, 3.3, 0, 0, 0, 7.4, 9.1, 8.9, 1.7, 0, 0, 0, 0, 4.1,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
}

type CalcForagingPeriod struct {
	time   *resource.Tick
	period globals.ForagingPeriod
}

func (s *CalcForagingPeriod) Initialize(w *ecs.World) {
	s.time = ecs.GetResource[resource.Tick](w)
	s.period = globals.ForagingPeriod{}
	ecs.AddResource(w, &s.period)
}

func (s *CalcForagingPeriod) Update(w *ecs.World) {
	dayOfYear := int(s.time.Tick) % len(foragingHoursBerlin2000)
	s.period.SecondsToday = int(foragingHoursBerlin2000[dayOfYear] * 3600)
}

func (s *CalcForagingPeriod) Finalize(w *ecs.World) {}
