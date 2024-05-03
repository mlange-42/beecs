package obs

import (
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/beecs/model/res"
)

type Cohorts struct {
	eggs   *res.Eggs
	larvae *res.Larvae
	pupae  *res.Pupae
	inHive *res.InHive

	data [][]float64
}

func (o *Cohorts) Initialize(w *ecs.World) {
	o.eggs = ecs.GetResource[res.Eggs](w)
	o.larvae = ecs.GetResource[res.Larvae](w)
	o.pupae = ecs.GetResource[res.Pupae](w)
	o.inHive = ecs.GetResource[res.InHive](w)

	ln := len(o.eggs.Workers) + len(o.larvae.Workers) + len(o.pupae.Workers) + len(o.inHive.Workers)

	o.data = make([][]float64, ln)
	for i := range o.data {
		o.data[i] = []float64{0}
	}
}
func (o *Cohorts) Update(w *ecs.World) {}
func (o *Cohorts) Header() []string {
	return []string{"Pop"}
}
func (o *Cohorts) Values(w *ecs.World) [][]float64 {
	idx := 0
	for _, v := range o.eggs.Workers {
		o.data[idx][0] = float64(v)
		idx++
	}
	for _, v := range o.larvae.Workers {
		o.data[idx][0] = float64(v)
		idx++
	}
	for _, v := range o.pupae.Workers {
		o.data[idx][0] = float64(v)
		idx++
	}
	for _, v := range o.inHive.Workers {
		o.data[idx][0] = float64(v)
		idx++
	}

	return o.data
}
