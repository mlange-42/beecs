package obs

import (
	"fmt"

	"github.com/mlange-42/ark/ecs"
	"github.com/mlange-42/beecs/comp"
)

// NectarVisits is a row observer for the number of nectar visits of all patches.
type NectarVisits struct {
	patchMapper *ecs.Map1[comp.Visits]
	data        []float64
	patches     []ecs.Entity
	header      []string
}

func (o *NectarVisits) Initialize(w *ecs.World) {
	o.patchMapper = o.patchMapper.New(w)

	patchFilter := *ecs.NewFilter1[comp.Visits](w)
	query := patchFilter.Query()
	for query.Next() {
		e := query.Entity()
		o.patches = append(o.patches, e)
	}

	for i := range o.patches {
		o.header = append(o.header, fmt.Sprintf("NectarVisits_%d", i))
	}

	o.data = make([]float64, len(o.patches))
}
func (o *NectarVisits) Update(w *ecs.World) {}
func (o *NectarVisits) Header() []string {
	return o.header
}
func (o *NectarVisits) Values(w *ecs.World) []float64 {
	for i, e := range o.patches {
		vis := o.patchMapper.Get(e)
		o.data[i] = float64(vis.Nectar)
	}
	return o.data
}

// NectarVisits is a row observer for the number of pollen visits of all patches.
type PollenVisits struct {
	patchMapper *ecs.Map1[comp.Visits]
	data        []float64
	patches     []ecs.Entity
	header      []string
}

func (o *PollenVisits) Initialize(w *ecs.World) {
	o.patchMapper = o.patchMapper.New(w)

	patchFilter := *ecs.NewFilter1[comp.Visits](w)
	query := patchFilter.Query()
	for query.Next() {
		e := query.Entity()
		o.patches = append(o.patches, e)
	}

	for i := range o.patches {
		o.header = append(o.header, fmt.Sprintf("PollenVisits_%d", i))
	}

	o.data = make([]float64, len(o.patches))
}
func (o *PollenVisits) Update(w *ecs.World) {}
func (o *PollenVisits) Header() []string {
	return o.header
}
func (o *PollenVisits) Values(w *ecs.World) []float64 {
	for i, e := range o.patches {
		vis := o.patchMapper.Get(e)
		o.data[i] = float64(vis.Pollen)
	}
	return o.data
}
