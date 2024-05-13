// Demonstrates how to implement observers.
package main

import (
	"github.com/mlange-42/arche-model/reporter"
	"github.com/mlange-42/arche-model/system"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/beecs/model"
	"github.com/mlange-42/beecs/model/globals"
	"github.com/mlange-42/beecs/model/params"
)

type StoresObserver struct {
	stores *globals.Stores // Access to some global variables.
	data   []float64       // We use a persistent data slice for the best performance.
}

func (o *StoresObserver) Initialize(w *ecs.World) {
	// Get the global variables.
	o.stores = ecs.GetResource[globals.Stores](w)
	// Create the data slice from the header length.
	o.data = make([]float64, len(o.Header()))
}

func (o *StoresObserver) Update(w *ecs.World) {}

func (o *StoresObserver) Header() []string {
	// The "table columns".
	return []string{"HoneyStore", "PollenStore"}
}

func (o *StoresObserver) Values(w *ecs.World) []float64 {
	// Fill the data slice.
	o.data[0] = o.stores.Honey
	o.data[1] = o.stores.Pollen
	// Return it.
	return o.data
}

func main() {
	// Get the default parameters.
	p := params.Default()

	// Create a model with the default sub-models.
	m := model.Default(&p, nil)

	// Add a system ("sub-model") to terminate after 365 ticks.
	m.AddSystem(&system.FixedTermination{Steps: 365})

	// Add a CSV output system using the observer defined above.
	m.AddSystem(&reporter.CSV{
		Observer: &StoresObserver{},
		File:     "out/stores.csv",
	})

	// Run the model.
	m.Run()
}
