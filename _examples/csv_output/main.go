// Demonstrates how to write CSV output.
package main

import (
	"github.com/mlange-42/ark-tools/reporter"
	"github.com/mlange-42/beecs/model"
	"github.com/mlange-42/beecs/obs"
	"github.com/mlange-42/beecs/params"
)

func main() {
	// Get the default parameters.
	p := params.Default()

	// Create a model with the default sub-models.
	m := model.Default(&p, nil)

	// Add a CSV output system using observer [obs.WorkerCohorts].
	m.AddSystem(&reporter.CSV{
		Observer: &obs.WorkerCohorts{},
		File:     "out/worker-cohorts.csv",
	})

	// Add a CSV snapshot output system using observer [obs.ForagingStats].
	m.AddSystem(&reporter.SnapshotCSV{
		Observer:    &obs.ForagingStats{},
		FilePattern: "out/foraging-%04d.csv",
	})

	// Run the model.
	m.Run()
}
