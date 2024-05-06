package obs

import (
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
	"github.com/mlange-42/beecs/model/comp"
	"github.com/mlange-42/beecs/model/res"
)

type Cohorts struct {
	eggs   *res.Eggs
	larvae *res.Larvae
	pupae  *res.Pupae
	inHive *res.InHive
	aff    *res.AgeFirstForaging
	time   *res.Time
	params *res.Params

	filter generic.Filter1[comp.Age]

	data [][]float64
}

func (o *Cohorts) Initialize(w *ecs.World) {
	o.eggs = ecs.GetResource[res.Eggs](w)
	o.larvae = ecs.GetResource[res.Larvae](w)
	o.pupae = ecs.GetResource[res.Pupae](w)
	o.inHive = ecs.GetResource[res.InHive](w)
	o.aff = ecs.GetResource[res.AgeFirstForaging](w)
	o.time = ecs.GetResource[res.Time](w)
	o.params = ecs.GetResource[res.Params](w)

	o.filter = *generic.NewFilter1[comp.Age]()

	// TODO: make x limits depend on parameters
	ln := len(o.eggs.Workers) + len(o.larvae.Workers) + len(o.pupae.Workers) + len(o.inHive.Workers) + 50

	o.data = make([][]float64, ln)
	for i := range o.data {
		o.data[i] = []float64{0, 0, 0, 0, 0}
	}
}
func (o *Cohorts) Update(w *ecs.World) {}
func (o *Cohorts) Header() []string {
	return []string{"Eggs", "Larvae", "Pupae", "InHive", "Foragers"}
}
func (o *Cohorts) Values(w *ecs.World) [][]float64 {
	for i := 0; i < len(o.data); i++ {
		o.data[i][3] = 0
		o.data[i][4] = 0
	}

	idx := 0
	for _, v := range o.eggs.Workers {
		o.data[idx][0] = float64(v)
		idx++
	}
	o.data[idx][0] = float64(o.larvae.Workers[0])

	for _, v := range o.larvae.Workers {
		o.data[idx][1] = float64(v)
		idx++
	}
	o.data[idx][1] = float64(o.pupae.Workers[0])

	for _, v := range o.pupae.Workers {
		o.data[idx][2] = float64(v)
		idx++
	}
	o.data[idx][2] = float64(o.inHive.Workers[0])

	offset := idx

	aff := o.aff.Current
	for i := 0; i < aff; i++ {
		o.data[idx][3] = float64(o.inHive.Workers[i])
		idx++
	}

	query := o.filter.Query(w)
	for query.Next() {
		a := query.Get()
		x := offset + o.time.Day - a.DayOfBirth
		if x >= len(o.data) {
			continue
		}
		o.data[x][4] += float64(o.params.SquadronSize)
	}
	o.data[idx][3] = float64(o.data[offset+aff][4])

	return o.data
}
