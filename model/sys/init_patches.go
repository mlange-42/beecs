package sys

import (
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/beecs/model/comp"
	"github.com/mlange-42/beecs/model/res"
)

type InitPatchesList struct {
	Patches []comp.PatchConfig
}

func (s *InitPatchesList) Initialize(w *ecs.World) {
	fac := res.NewPatchFactory(w)

	for _, p := range s.Patches {
		_ = fac.CreatePatch(p)
	}
}

func (s *InitPatchesList) Update(w *ecs.World) {}

func (s *InitPatchesList) Finalize(w *ecs.World) {}
