package obs

import (
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/beecs/model/globals"
)

// Debug is a row observer for several colony structure variables,
// using the same names as the original BEEHAVE implementation.
//
// Primarily meant for validation of beecs against BEEHAVE.
type Debug struct {
	pop      *globals.PopulationStats
	stores   *globals.Stores
	foraging *globals.ForagingPeriod
	data     []float64
}

func (o *Debug) Initialize(w *ecs.World) {
	o.pop = ecs.GetResource[globals.PopulationStats](w)
	o.stores = ecs.GetResource[globals.Stores](w)
	o.foraging = ecs.GetResource[globals.ForagingPeriod](w)
	o.data = make([]float64, len(o.Header()))
}
func (o *Debug) Update(w *ecs.World) {}
func (o *Debug) Header() []string {
	return []string{"DailyForagingPeriod", "HoneyEnergyStore", "PollenStore_g", "TotalEggs", "TotalLarvae", "TotalPupae", "TotalIHBees", "TotalForagers"}
}
func (o *Debug) Values(w *ecs.World) []float64 {
	o.data[0] = float64(o.foraging.SecondsToday)
	o.data[1] = o.stores.Honey
	o.data[2] = o.stores.Pollen

	o.data[3] = float64(o.pop.WorkerEggs)
	o.data[4] = float64(o.pop.WorkerLarvae)
	o.data[5] = float64(o.pop.WorkerPupae)
	o.data[6] = float64(o.pop.WorkersInHive)
	o.data[7] = float64(o.pop.WorkersForagers)

	return o.data
}
