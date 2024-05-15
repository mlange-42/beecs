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
	"github.com/mlange-42/beecs/experiment"
	"github.com/mlange-42/beecs/model"
	"github.com/mlange-42/beecs/obs"
	"github.com/mlange-42/beecs/params"
	"golang.org/x/exp/rand"
)

func main() {
	// Define parameter variations.
	vars := []experiment.ParameterVariation{
		{
			Parameter: "params.InitialPopulation.Count",
			SequenceIntValues: &experiment.SequenceIntValues{
				Values: []int{1000, 2000, 5000, 10000}, // values to iterate
			},
		},
		{
			Parameter: "params.InitialStores.Honey",
			SequenceFloatRange: &experiment.SequenceFloatRange{
				Min:    0,
				Max:    100,
				Values: 11,
			},
		},
	}

	// Create an experiment.
	exp, err := experiment.New(vars, rand.New(rand.NewSource(0)))
	if err != nil {
		log.Fatal(err)
	}
	// Get the number of parameter sets in the experiment.
	sets := exp.ParameterSets()

	// Create an empty model.
	m := amodel.New()

	fmt.Printf("Running experiment with %d parameter sets\n", sets)

	for i := 0; i < sets; i++ {
		// Initialize and run the model.
		run(m, &exp, i)
	}
}

func run(m *amodel.Model, exp *experiment.Experiment, idx int) {
	// Get the default parameters.
	p := params.Default()
	// Create a model with the default sub-models.
	m = model.Default(&p, m)

	// Get parameter values for the current run.
	values := exp.Values(idx)
	fmt.Printf("Running set %v\n", values)
	// Set/overwrite parameters from the experiment.
	exp.ApplyValues(values, &m.World)

	// Add a CSV output system using observer [obs.WorkerCohorts].
	m.AddSystem(&reporter.CSV{
		Observer: &obs.WorkerCohorts{},
		File:     fmt.Sprintf("out/worker-cohorts-%04d.csv", idx),
	})

	// Run the model.
	m.Run()
}
