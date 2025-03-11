package sys

import (
	"fmt"
	"io/fs"
	"log"
	"os"

	"github.com/mlange-42/ark/ecs"
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
		wd := ecs.GetResource[params.WorkingDirectory](w).Path
		fileSys = os.DirFS(wd)
	}

	// fill from data provided directly
	for _, arr := range periodParams.Years {
		if len(arr)%365 != 0 {
			log.Fatal(fmt.Errorf("foraging period entries requires multiple of 365 values, parameters have %d", len(arr)))
		}
		years := len(arr) / 365
		for year := 0; year < years; year++ {
			s.periodData.Years = append(s.periodData.Years, arr[year*365:(year+1)*365])
		}
	}

	// fill from files
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
