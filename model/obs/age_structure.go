package obs

import (
	"math"

	"github.com/mlange-42/arche-model/resource"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
	"github.com/mlange-42/beecs/model/comp"
	"github.com/mlange-42/beecs/model/res"
)

type AgeStructure struct {
	eggs   *res.Eggs
	larvae *res.Larvae
	pupae  *res.Pupae
	inHive *res.InHive
	aff    *res.AgeFirstForaging
	time   *resource.Tick
	params *res.Params

	data   [][]float64
	filter generic.Filter1[comp.Age]

	MaxAge int
}

func (o *AgeStructure) Initialize(w *ecs.World) {
	o.eggs = ecs.GetResource[res.Eggs](w)
	o.larvae = ecs.GetResource[res.Larvae](w)
	o.pupae = ecs.GetResource[res.Pupae](w)
	o.inHive = ecs.GetResource[res.InHive](w)
	o.aff = ecs.GetResource[res.AgeFirstForaging](w)
	o.time = ecs.GetResource[resource.Tick](w)
	o.params = ecs.GetResource[res.Params](w)

	o.filter = *generic.NewFilter1[comp.Age]()

	ln := len(o.eggs.Workers) + len(o.larvae.Workers) + len(o.pupae.Workers) + o.MaxAge

	o.data = make([][]float64, ln)
	for i := range o.data {
		o.data[i] = []float64{0, 0, 0, 0, 0}
	}
}
func (o *AgeStructure) Update(w *ecs.World) {}
func (o *AgeStructure) Header() []string {
	return []string{"Eggs", "Larvae", "Pupae", "InHive", "Foragers"}
}
func (o *AgeStructure) Values(w *ecs.World) [][]float64 {
	for i := 0; i < len(o.data); i++ {
		o.data[i][0] = math.NaN()
		o.data[i][1] = math.NaN()
		o.data[i][2] = math.NaN()
		o.data[i][3] = math.NaN()
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

	aff := o.aff.Aff
	for i := 0; i < aff; i++ {
		o.data[idx][3] = float64(o.inHive.Workers[i])
		idx++
	}

	for i := 0; i < offset+aff; i++ {
		o.data[i][4] = math.NaN()
	}

	query := o.filter.Query(w)
	for query.Next() {
		a := query.Get()
		x := offset + int(o.time.Tick) - a.DayOfBirth
		if x >= len(o.data) {
			continue
		}
		o.data[x][4] += float64(o.params.SquadronSize)
	}
	o.data[idx][3] = float64(o.data[offset+aff][4])

	return o.data
}
