// Demonstrates how to run experiments with parameter variations.
//
// When running experiments from Go code, parameters are better manipulated directly.
// The feature of setting parameters from their names demonstrated here is primarily meant
// for processing experiments defined in e.g. JSON files.
package main

import (
	"fmt"
	"log"

	amodel "github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche-model/reporter"
	"github.com/mlange-42/arche-model/system"
	"github.com/mlange-42/beecs/model"
	"github.com/mlange-42/beecs/obs"
	"github.com/mlange-42/beecs/params"
	"golang.org/x/exp/rand"
)

func main() {
	// Define parameter variations.
	vars := []params.ParameterVariation{
		{
			Parameter: "params.InitialPopulation.Count",
			Type:      "sequence-values",
			IntParams: []int{1000, 2000, 5000, 10000}, // values to iterate
		},
		{
			Parameter:   "params.InitialStores.Honey",
			Type:        "sequence-range",
			FloatParams: []float64{0, 100, 11}, // lower limit, upper limit, #values
		},
	}

	// Create an experiment.
	experiment, err := params.NewExperiment(vars, rand.New(nil))
	if err != nil {
		log.Fatal(err)
	}
	// Get the number of parameter sets in the experiment.
	sets := experiment.ParameterSets()

	// Create an empty model.
	m := amodel.New()

	fmt.Printf("Running experiment with %d parameter sets\n", sets)

	for i := 0; i < sets; i++ {
		// Get and print the parameter values for the current run.
		values := experiment.Values(i)
		fmt.Printf("Running set %v\n", values)

		// Initialize and run the model with the given parameters.
		run(m, values, i)
	}
}

func run(m *amodel.Model, values map[string]any, idx int) {
	// Get the default parameters.
	p := params.Default()
	// Create a model with the default sub-models.
	m = model.Default(&p, m)

	// Set/overwrite parameters from the experiment.
	for par, value := range values {
		if err := model.SetParameter(&m.World, par, value); err != nil {
			log.Fatal(err)
		}
	}

	// Add a system ("sub-model") to terminate after 365 ticks.
	m.AddSystem(&system.FixedTermination{Steps: 365})

	// Add a CSV output system using observer [obs.WorkerCohorts].
	m.AddSystem(&reporter.CSV{
		Observer: &obs.WorkerCohorts{},
		File:     fmt.Sprintf("out/worker-cohorts-%04d.csv", idx),
	})

	// Run the model.
	m.Run()
}
