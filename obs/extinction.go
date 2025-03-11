package obs

import (
	"github.com/mlange-42/ark-tools/resource"
	"github.com/mlange-42/ark/ecs"
	"github.com/mlange-42/beecs/globals"
)

// Extinction is a row observer for the tick of extinction of all bees.
//
// Best use as a finalization observer.
// Columns are "Extinct" and "ExtinctionTick".
type Extinction struct {
	time *resource.Tick
	pop  *globals.PopulationStats
	data []float64
}

func (o *Extinction) Initialize(w *ecs.World) {
	o.time = ecs.GetResource[resource.Tick](w)
	o.pop = ecs.GetResource[globals.PopulationStats](w)
	o.data = make([]float64, len(o.Header()))
}
func (o *Extinction) Update(w *ecs.World) {
	if o.data[0] == 0 && o.pop.TotalPopulation == 0 {
		o.data[0] = 1
		o.data[1] = float64(o.time.Tick)
	}
}
func (o *Extinction) Header() []string {
	return []string{"Extinct", "ExtinctionTick"}
}
func (o *Extinction) Values(w *ecs.World) []float64 {
	return o.data
}
