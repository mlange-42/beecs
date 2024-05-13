package sys

import (
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/beecs/comp"
	"github.com/mlange-42/beecs/globals"
)

type InitPatchesList struct {
	Patches []comp.PatchConfig
}

func (s *InitPatchesList) Initialize(w *ecs.World) {
	fac := globals.NewPatchFactory(w)

	for _, p := range s.Patches {
		_ = fac.CreatePatch(p)
	}
}

func (s *InitPatchesList) Update(w *ecs.World) {}

func (s *InitPatchesList) Finalize(w *ecs.World) {}
