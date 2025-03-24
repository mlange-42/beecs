// Demonstrates how to parametrize the foraging period.
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

	// Use randomized builtin weather data.
	p.ForagingPeriod = params.ForagingPeriod{
		Years: [][]float64{}, // Can be used to provide daily foraging hours directly.
		Files: []string{ // Alternatively (or additionally), read data from text files.
			"foraging-period/berlin2000.txt",
			"foraging-period/berlin2001.txt",
			"foraging-period/berlin2002.txt",
			"foraging-period/berlin2003.txt",
			"foraging-period/berlin2004.txt",
			"foraging-period/berlin2005.txt",
			"foraging-period/berlin2006.txt",
		},
		Builtin:     true,
		RandomYears: true,
	}

	// Run for 10 years
	p.Termination.MaxTicks = 3650

	// Create a model with the default sub-models.
	m := model.Default(&p, nil)

	// Add a CSV output system using the observer defined above.
	m.AddSystem(&reporter.CSV{
		Observer: &obs.ForagingPeriod{},
		File:     "out/foraging-period.csv",
	})

	// Run the model.
	m.Run()
}
