package main

import (
	"fmt"
	"time"

	amodel "github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche-model/reporter"
	"github.com/mlange-42/arche-model/system"
	"github.com/mlange-42/beecs/model"
	"github.com/mlange-42/beecs/model/obs"
)

func main() {
	m := amodel.New()

	start := time.Now()

	for i := 0; i < 100; i++ {
		run(m, i)
	}

	dur := time.Now().Sub(start)
	fmt.Println(dur)
}

func run(m *amodel.Model, idx int) {
	ticks := 365

	m = model.Default(m)

	m.AddSystem(&system.FixedTermination{Steps: int64(ticks)})

	m.AddSystem(&reporter.CSV{
		Observer: &obs.Debug{},
		File:     fmt.Sprintf("out/beecs-%04d.csv", idx),
		Sep:      ";",
	})

	m.Run()
}
