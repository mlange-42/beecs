package sys

import (
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/beecs/globals"
	"github.com/mlange-42/beecs/params"
)

type InitPatchesList struct{}

func (s *InitPatchesList) Initialize(w *ecs.World) {
	patches := ecs.GetResource[params.InitialPatches](w)
	fac := globals.NewPatchFactory(w)

	for _, p := range patches.Patches {
		_ = fac.CreatePatch(p)
	}
}

func (s *InitPatchesList) Update(w *ecs.World) {}

func (s *InitPatchesList) Finalize(w *ecs.World) {}
