package obs

import (
	"github.com/mlange-42/ark/ecs"
	"github.com/mlange-42/beecs/globals"
)

// ForagingStats is a table observer for foraging activity over the day.
type ForagingStats struct {
	Relative bool
	stats    *globals.ForagingStats
	pop      *globals.PopulationStats
	data     [][]float64
}

func (o *ForagingStats) Initialize(w *ecs.World) {
	o.stats = ecs.GetResource[globals.ForagingStats](w)
	o.pop = ecs.GetResource[globals.PopulationStats](w)
	o.data = [][]float64{}
}
func (o *ForagingStats) Update(w *ecs.World) {}
func (o *ForagingStats) Header() []string {
	return []string{"Round", "Lazy", "Resting", "Searching", "Recruited", "Nectar", "Pollen"}
}
func (o *ForagingStats) Values(w *ecs.World) [][]float64 {
	o.data = o.data[:0]

	total := o.pop.WorkersForagers

	if o.Relative && total > 0 {
		for i, round := range o.stats.Rounds {
			row := make([]float64, 7)

			row[0] = float64(i)
			row[1] = float64(round.Lazy) / float64(total)
			row[2] = float64(round.Resting) / float64(total)
			row[3] = float64(round.Searching) / float64(total)
			row[4] = float64(round.Recruited) / float64(total)
			row[5] = float64(round.Nectar) / float64(total)
			row[6] = float64(round.Pollen) / float64(total)

			o.data = append(o.data, row)
		}
	} else {
		for i, round := range o.stats.Rounds {
			row := make([]float64, 7)

			row[0] = float64(i)
			row[1] = float64(round.Lazy)
			row[2] = float64(round.Resting)
			row[3] = float64(round.Searching)
			row[4] = float64(round.Recruited)
			row[5] = float64(round.Nectar)
			row[6] = float64(round.Pollen)

			o.data = append(o.data, row)
		}
	}

	return o.data
}
