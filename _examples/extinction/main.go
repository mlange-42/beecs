// Demonstrates how to measure time to extinction using a callback reporter.
package main

import (
	"fmt"

	"github.com/mlange-42/ark-tools/reporter"
	"github.com/mlange-42/beecs/model"
	"github.com/mlange-42/beecs/obs"
	"github.com/mlange-42/beecs/params"
)

func main() {
	// Get the default parameters.
	p := params.Default()

	// Run for 100 years, but only until extinction.
	p.Termination.MaxTicks = 100 * 365
	p.Termination.OnExtinction = true

	// Increase handling time to provoke extinction.
	p.HandlingTime.NectarGathering = 3600

	// Create a model with the default sub-models.
	m := model.Default(&p, nil)

	// Add a CSV output system using observer [obs.WorkerCohorts].
	m.AddSystem(&reporter.RowCallback{
		Observer: &obs.Extinction{},
		Callback: func(step int, row []float64) {
			if row[0] > 0 {
				fmt.Printf("Extinction on tick %d (%02f years)\n", int(row[1]), row[1]/365)
			} else {
				fmt.Println("No extinction")
			}
		},
		Final: true,
	})

	// Run the model.
	m.Run()
}
