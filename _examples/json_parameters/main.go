// Demonstrates how to read parameters from JSON.
package main

import (
	"fmt"
	"log"

	"github.com/mlange-42/arche-model/reporter"
	"github.com/mlange-42/arche-model/system"
	"github.com/mlange-42/beecs/model"
	"github.com/mlange-42/beecs/model/obs"
	"github.com/mlange-42/beecs/model/params"
)

func main() {
	// Get the default parameters.
	p := params.Default()
	// Read JSON to modify some parameters.
	err := p.FromJSON("_examples/json_parameters/params.json")
	if err != nil {
		log.Fatal(err)
	}
	// Print one of the modified sections of the parameters.
	fmt.Printf("%+v\n", p.Forager)

	// Create a model with the default sub-models.
	m := model.Default(&p, nil)

	// Add a system ("sub-model") to terminate after 365 ticks.
	m.AddSystem(&system.FixedTermination{Steps: 365})

	// Add a CSV output system using observer [obs.WorkerCohorts].
	m.AddSystem(&reporter.CSV{
		Observer: &obs.WorkerCohorts{},
		File:     "out/worker-cohorts.csv",
	})

	// Run the model.
	m.Run()
}
