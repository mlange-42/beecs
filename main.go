package main

import (
	"github.com/mlange-42/arche-model/reporter"
	"github.com/mlange-42/arche-model/system"
	"github.com/mlange-42/arche-pixel/plot"
	"github.com/mlange-42/arche-pixel/window"
	"github.com/mlange-42/beecs/model"
	"github.com/mlange-42/beecs/model/obs"
)

func main() {
	gui := true
	files := false
	ticks := 3650

	params := model.DefaultParams()
	m := model.Default(&params, nil)
	if gui {
		m.TPS = 30
		m.FPS = 30
	}

	//m.AddSystem(&sys.Pause{Steps: 91})
	m.AddSystem(&system.FixedTermination{Steps: int64(ticks)})

	// File output

	if files {
		m.AddSystem(&reporter.CSV{
			Observer: &obs.Debug{},
			File:     "out/out-beecs.csv",
			Sep:      ";",
		})
	}

	// Graphics

	if gui {
		m.AddUISystem((&window.Window{
			Title:        "Cohorts",
			Bounds:       window.B(1, 30, 600, 400),
			DrawInterval: 15,
		}).
			With(
				&plot.TimeSeries{
					Observer:       &obs.WorkerCohorts{Cumulative: true},
					UpdateInterval: 1,
					MaxRows:        2 * 365,
					Labels:         plot.Labels{Title: "Cohorts", X: "Time [days]", Y: "Count"},
				},
				&plot.Controls{}))

		m.AddUISystem((&window.Window{
			Title:  "Age distribution",
			Bounds: window.B(1, 465, 600, 400),
		}).
			With(
				&plot.Lines{
					Observer: &obs.AgeStructure{},
					YLim:     [...]float64{0, 2000},
					Labels:   plot.Labels{Title: "Age distribution", X: "Age [days]", Y: "Count"},
				},
				&plot.Controls{}))

		m.AddUISystem((&window.Window{
			Title:        "Stores",
			Bounds:       window.B(610, 30, 600, 400),
			DrawInterval: 15,
		}).
			With(
				&plot.TimeSeries{
					Observer:       &obs.Stores{PollenFactor: 20},
					Columns:        []string{"Honey", "Pollen x20"},
					UpdateInterval: 1,
					MaxRows:        2 * 365,
					Labels:         plot.Labels{Title: "Stores", X: "Time [days]", Y: "Mass [kg]"},
				},
				&plot.Controls{}))

		m.AddUISystem((&window.Window{
			Title:        "Foraging Period",
			Bounds:       window.B(610, 465, 600, 400),
			DrawInterval: 15,
		}).
			With(
				&plot.TimeSeries{
					Observer:       &obs.ForagingPeriod{},
					UpdateInterval: 1,
					MaxRows:        2 * 365,
					Labels:         plot.Labels{Title: "Foraging Period", X: "Time [days]", Y: "Foraging Period [h]"},
				},
				&plot.Controls{}))
	}

	// Run

	window.Run(m)
}
