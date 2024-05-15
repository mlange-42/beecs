package sys

import (
	"fmt"
	"log"

	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/beecs/globals"
	"github.com/mlange-42/beecs/params"
	"github.com/mlange-42/beecs/util"
)

type InitPatchesList struct{}

func (s *InitPatchesList) Initialize(w *ecs.World) {
	patches := ecs.GetResource[params.InitialPatches](w)
	fac := globals.NewPatchFactory(w)

	for _, p := range patches.Patches {
		_ = fac.CreatePatch(p)
	}

	if patches.File != "" {
		pch, err := util.PatchesFromFile(patches.File)
		if err != nil {
			log.Fatal(fmt.Errorf("error reading patches file '%s': %s", patches.File, err.Error()))
		}
		for _, p := range pch {
			_ = fac.CreatePatch(p)
		}
	}
}

func (s *InitPatchesList) Update(w *ecs.World) {}

func (s *InitPatchesList) Finalize(w *ecs.World) {}
