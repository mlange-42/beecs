package sys

import (
	"fmt"
	"io/fs"
	"log"
	"os"

	"github.com/mlange-42/arche-model/resource"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/beecs"
	"github.com/mlange-42/beecs/globals"
	"github.com/mlange-42/beecs/params"
	"github.com/mlange-42/beecs/util"
	"golang.org/x/exp/rand"
)

type CalcForagingPeriod struct {
	time   *resource.Tick
	period globals.ForagingPeriod

	years      [][]float64
	year       int
	rng        rand.Rand
	randomized bool
}

func (s *CalcForagingPeriod) Initialize(w *ecs.World) {
	s.time = ecs.GetResource[resource.Tick](w)

	s.period = globals.ForagingPeriod{}
	ecs.AddResource(w, &s.period)

	src := ecs.GetResource[resource.Rand](w)
	s.rng = *rand.New(src)

	periodParams := ecs.GetResource[params.ForagingPeriod](w)
	s.randomized = periodParams.RandomYears
	var fileSys fs.FS = beecs.Data
	if !periodParams.Builtin {
		fileSys = os.DirFS(".")
	}

	for _, f := range periodParams.Files {
		arr, err := util.FloatArrayFromFile(fileSys, f)
		if err != nil {
			log.Fatal(fmt.Errorf("error reading foraging period file '%s': %s", f, err.Error()))
		}
		if len(arr) != 365 {
			log.Fatal(fmt.Errorf("foraging period file require 365 values, '%s' has %d", f, len(arr)))
		}
		s.years = append(s.years, arr)
	}
}

func (s *CalcForagingPeriod) Update(w *ecs.World) {
	dayOfYear := int(s.time.Tick % 365)
	if dayOfYear == 0 {
		if s.randomized {
			s.year = s.rng.Intn(len(s.years))
		} else {
			s.year = int(s.time.Tick/365) % len(s.years)
		}
	}
	year := s.years[s.year]
	s.period.SecondsToday = int(year[dayOfYear] * 3600)
}

func (s *CalcForagingPeriod) Finalize(w *ecs.World) {}
