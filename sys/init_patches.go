package sys

import (
	"fmt"
	"log"
	"path"

	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/beecs/comp"
	"github.com/mlange-42/beecs/globals"
	"github.com/mlange-42/beecs/params"
	"github.com/mlange-42/beecs/util"
)

// InitPatchesList initializes flower patches based on the information
// in [params.InitialPatches].
type InitPatchesList struct{}

func (s *InitPatchesList) Initialize(w *ecs.World) {
	patches := ecs.GetResource[params.InitialPatches](w)
	fac := globals.NewPatchFactory(w)

	allPatches := append([]comp.PatchConfig{}, patches.Patches...)
	if patches.File != "" {
		wd := ecs.GetResource[params.WorkingDirectory](w).Path
		path := path.Join(wd, patches.File)
		pch, err := util.PatchesFromFile(path)
		if err != nil {
			log.Fatal(fmt.Errorf("error reading patches file '%s': %s", patches.File, err.Error()))
		}
		allPatches = append(allPatches, pch...)
	}

	fac.CreatePatches(allPatches)
}

func (s *InitPatchesList) Update(w *ecs.World) {}

func (s *InitPatchesList) Finalize(w *ecs.World) {}
