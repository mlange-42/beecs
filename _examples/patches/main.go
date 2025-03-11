// Demonstrates how to parametrize patches.
package main

import (
	"github.com/mlange-42/ark-tools/reporter"
	"github.com/mlange-42/beecs/comp"
	"github.com/mlange-42/beecs/enum/interp"
	"github.com/mlange-42/beecs/model"
	"github.com/mlange-42/beecs/obs"
	"github.com/mlange-42/beecs/params"
)

func main() {
	// Get the default parameters.
	p := params.Default()

	// Change initial patches
	p.InitialPatches = params.InitialPatches{
		Patches: []comp.PatchConfig{
			// A patch with constant resource availability.
			{
				DistToColony: 1000,
				ConstantPatch: &comp.ConstantPatch{
					Nectar:               20,  // [L]
					Pollen:               1,   // [kg]
					NectarConcentration:  1.5, // [mumol/L]
					DetectionProbability: 0.2,
				},
			},
			// A patch with seasonal resource availability.
			{
				DistToColony: 1000,
				SeasonalPatch: &comp.SeasonalPatch{
					MaxNectar:            20,  // [L]
					MaxPollen:            1,   // [kg]
					NectarConcentration:  1.5, // [mumol/L]
					DetectionProbability: 0.2,

					SeasonShift: 20, // [d]
				},
			},
			// A patch with scripted resource availability.
			{
				DistToColony: 1000,
				ScriptedPatch: &comp.ScriptedPatch{
					Nectar: [][2]float64{
						{0, 0},
						{100, 20},
						{250, 0},
					},
					Pollen: [][2]float64{
						{0, 0},
						{100, 1},
						{250, 0},
					},
					NectarConcentration: [][2]float64{
						{0, 1.5},
					},
					DetectionProbability: [][2]float64{
						{0, 0.2},
					},
					Interpolation: interp.Step,
				},
			},
		},
	}

	// Create a model with the default sub-models.
	m := model.Default(&p, nil)

	// Add a CSV outputs for patch nectar and pollen.
	m.AddSystem(&reporter.CSV{
		Observer: &obs.PatchNectar{},
		File:     "out/nectar.csv",
	})
	m.AddSystem(&reporter.CSV{
		Observer: &obs.PatchPollen{},
		File:     "out/pollen.csv",
	})

	// Add a CSV outputs for patch nectar and pollen visits.
	m.AddSystem(&reporter.CSV{
		Observer: &obs.NectarVisits{},
		File:     "out/nectar-visits.csv",
	})
	m.AddSystem(&reporter.CSV{
		Observer: &obs.PollenVisits{},
		File:     "out/pollen-visits.csv",
	})

	// Run the model.
	m.Run()
}
