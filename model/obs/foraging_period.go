package obs

import (
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/beecs/model/res"
)

// ForagingPeriod is a row observer for the foraging period of the current day, in hours.
//
// Has a single column "Foraging Period [h]", and reports one row/value per model tick.
type ForagingPeriod struct {
	period *res.ForagingPeriod
	data   []float64
}

func (o *ForagingPeriod) Initialize(w *ecs.World) {
	o.period = ecs.GetResource[res.ForagingPeriod](w)
	o.data = make([]float64, len(o.Header()))
}
func (o *ForagingPeriod) Update(w *ecs.World) {}
func (o *ForagingPeriod) Header() []string {
	return []string{"Foraging Period [h]"}
}
func (o *ForagingPeriod) Values(w *ecs.World) []float64 {
	o.data[0] = float64(o.period.SecondsToday) / 3600.0

	return o.data
}
