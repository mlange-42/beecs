package sys

import (
	"fmt"
	"io/fs"
	"log"
	"os"

	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/beecs/data"
	"github.com/mlange-42/beecs/globals"
	"github.com/mlange-42/beecs/params"
	"github.com/mlange-42/beecs/util"
)

type InitForagingPeriod struct {
	periodData globals.ForagingPeriodData
}

func (s *InitForagingPeriod) Initialize(w *ecs.World) {
	s.periodData = globals.ForagingPeriodData{}
	ecs.AddResource(w, &s.periodData)

	periodParams := ecs.GetResource[params.ForagingPeriod](w)
	var fileSys fs.FS = data.ForagingPeriod
	if !periodParams.Builtin {
		fileSys = os.DirFS(".")
	}

	for _, f := range periodParams.Files {
		arr, err := util.FloatArrayFromFile(fileSys, f)
		if err != nil {
			log.Fatal(fmt.Errorf("error reading foraging period file '%s': %s", f, err.Error()))
		}
		if len(arr)%365 != 0 {
			log.Fatal(fmt.Errorf("foraging period file requires multiple of 365 values, '%s' has %d", f, len(arr)))
		}
		years := len(arr) / 365
		for year := 0; year < years; year++ {
			s.periodData.Years = append(s.periodData.Years, arr[year*365:(year+1)*365])
		}
	}
}

func (s *InitForagingPeriod) Update(w *ecs.World) {}

func (s *InitForagingPeriod) Finalize(w *ecs.World) {}
