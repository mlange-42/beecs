// Demonstrates how to read parameters from JSON.
package main

import (
	"fmt"
	"log"

	"github.com/mlange-42/ark-tools/reporter"
	"github.com/mlange-42/beecs/model"
	"github.com/mlange-42/beecs/obs"
	"github.com/mlange-42/beecs/params"
)

func main() {
	// Get the default parameters.
	p := params.Default()
	// Read JSON to modify some parameters.
	err := p.FromJSONFile("_examples/json_parameters/params.json")
	if err != nil {
		log.Fatal(err)
	}
	// Print one of the modified sections of the parameters.
	fmt.Printf("%+v\n", p.Foragers)

	// Create a model with the default sub-models.
	m := model.Default(&p, nil)

	// Add a CSV output system using observer [obs.WorkerCohorts].
	m.AddSystem(&reporter.CSV{
		Observer: &obs.WorkerCohorts{},
		File:     "out/worker-cohorts.csv",
	})

	// Run the model.
	m.Run()
}
