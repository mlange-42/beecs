package main

import (
	"fmt"
	"time"

	amodel "github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche-model/reporter"
	"github.com/mlange-42/beecs/model"
	"github.com/mlange-42/beecs/obs"
	"github.com/mlange-42/beecs/params"
)

func main() {
	m := amodel.New()

	p := params.Default()
	p.Termination.MaxTicks = 365

	start := time.Now()

	for i := 0; i < 100; i++ {
		run(m, i, &p)
	}

	dur := time.Since(start)
	fmt.Println(dur)
}

func run(m *amodel.Model, idx int, params params.Params) {
	m = model.Default(params, m)

	m.AddSystem(&reporter.CSV{
		Observer: &obs.Debug{},
		File:     fmt.Sprintf("out/beecs-%04d.csv", idx),
		Sep:      ";",
	})

	m.Run()
}
