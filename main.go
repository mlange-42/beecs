package main

import (
	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche-model/system"
)

func main() {
	m := model.New()

	m.AddSystem(&system.FixedTermination{Steps: 1000})

	m.Run()
}
