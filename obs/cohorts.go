package obs

import (
	"github.com/mlange-42/ark/ecs"
	"github.com/mlange-42/beecs/globals"
)

// WorkerCohorts is a row observer for the colony structure by cohorts.
//
// Columns are the cohorts, either plain or cumulatively.
// Provides one row per model tick.
type WorkerCohorts struct {
	pop  *globals.PopulationStats
	data []float64

	Cumulative bool // Whether cohorts should be cumulative.
}

func (o *WorkerCohorts) Initialize(w *ecs.World) {
	o.pop = ecs.GetResource[globals.PopulationStats](w)
	o.data = make([]float64, len(o.Header()))
}
func (o *WorkerCohorts) Update(w *ecs.World) {}
func (o *WorkerCohorts) Header() []string {
	if o.Cumulative {
		return []string{"Eggs", "+Larvae", "+Pupae", "+InHive", "+Foragers"}
	}
	return []string{"Eggs", "Larvae", "Pupae", "InHive", "Foragers"}
}
func (o *WorkerCohorts) Values(w *ecs.World) []float64 {
	if o.Cumulative {
		o.data[0] = float64(o.pop.WorkerEggs)
		o.data[1] = o.data[0] + float64(o.pop.WorkerLarvae)
		o.data[2] = o.data[1] + float64(o.pop.WorkerPupae)
		o.data[3] = o.data[2] + float64(o.pop.WorkersInHive)
		o.data[4] = o.data[3] + float64(o.pop.WorkersForagers)
	} else {
		o.data[0] = float64(o.pop.WorkerEggs)
		o.data[1] = float64(o.pop.WorkerLarvae)
		o.data[2] = float64(o.pop.WorkerPupae)
		o.data[3] = float64(o.pop.WorkersInHive)
		o.data[4] = float64(o.pop.WorkersForagers)
	}

	return o.data
}
