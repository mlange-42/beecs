package obs

import (
	"fmt"

	"github.com/mlange-42/ark/ecs"
	"github.com/mlange-42/beecs/comp"
)

// PatchNectar is a row observer for the nectar availability of all patches in L (liters).
type PatchNectar struct {
	patchMapper *ecs.Map1[comp.Resource]
	data        []float64
	patches     []ecs.Entity
	header      []string
}

func (o *PatchNectar) Initialize(w *ecs.World) {
	o.patchMapper = o.patchMapper.New(w)

	patchFilter := *ecs.NewFilter1[comp.Resource](w)
	query := patchFilter.Query()
	for query.Next() {
		e := query.Entity()
		o.patches = append(o.patches, e)
	}

	for i := range o.patches {
		o.header = append(o.header, fmt.Sprintf("Nectar_%d", i))
	}

	o.data = make([]float64, len(o.patches))
}
func (o *PatchNectar) Update(w *ecs.World) {}
func (o *PatchNectar) Header() []string {
	return o.header
}
func (o *PatchNectar) Values(w *ecs.World) []float64 {
	for i, e := range o.patches {
		res := o.patchMapper.Get(e)
		o.data[i] = res.Nectar * 0.000_001
	}
	return o.data
}

// PatchPollen is a row observer for the pollen availability of all patches, in g (grams).
type PatchPollen struct {
	patchMapper *ecs.Map1[comp.Resource]
	data        []float64
	patches     []ecs.Entity
	header      []string
}

func (o *PatchPollen) Initialize(w *ecs.World) {
	o.patchMapper = o.patchMapper.New(w)

	patchFilter := ecs.NewFilter1[comp.Resource](w)
	query := patchFilter.Query()
	for query.Next() {
		e := query.Entity()
		o.patches = append(o.patches, e)
	}

	for i := range o.patches {
		o.header = append(o.header, fmt.Sprintf("Pollen_%d", i))
	}

	o.data = make([]float64, len(o.patches))
}
func (o *PatchPollen) Update(w *ecs.World) {}
func (o *PatchPollen) Header() []string {
	return o.header
}
func (o *PatchPollen) Values(w *ecs.World) []float64 {
	for i, e := range o.patches {
		res := o.patchMapper.Get(e)
		o.data[i] = res.Pollen
	}
	return o.data
}
