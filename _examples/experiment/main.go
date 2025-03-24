// Demonstrates how to run experiments with parameter variations.
//
// When running experiments from Go code, parameters are better manipulated directly.
// The feature of setting parameters from their names demonstrated here is primarily meant
// for processing experiments defined in e.g. JSON files.
package main

import (
	"fmt"
	"log"
	"math/rand/v2"

	"github.com/mlange-42/ark-tools/app"
	"github.com/mlange-42/ark-tools/reporter"
	"github.com/mlange-42/beecs/experiment"
	"github.com/mlange-42/beecs/model"
	"github.com/mlange-42/beecs/obs"
	"github.com/mlange-42/beecs/params"
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

	// Create an RNG.
	rng := rand.New(rand.NewPCG(0, 0))
	// Create an experiment for one run per parameter set.
	exp, err := experiment.New(vars, rng, 1)
	if err != nil {
		log.Fatal(err)
	}

	// Create an empty model.
	app := app.New()

	fmt.Printf("Running experiment with %d parameter sets\n", exp.ParameterSets())

	runs := exp.TotalRuns()
	for i := 0; i < runs; i++ {
		// Initialize and run the model.
		run(app, &exp, i)
	}
}

func run(app *app.App, exp *experiment.Experiment, idx int) {
	// Get the default parameters.
	p := params.Default()
	// Create a model with the default sub-models.
	m := model.Default(&p, app)

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
	app.Run()
}
