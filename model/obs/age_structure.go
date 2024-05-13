package obs

import (
	"math"

	"github.com/mlange-42/arche-model/resource"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
	"github.com/mlange-42/beecs/model/comp"
	"github.com/mlange-42/beecs/model/globals"
	"github.com/mlange-42/beecs/model/params"
)

// AgeStructure is a table observer for the age and cohort structure of the colony.
//
// Rows are age in days, columns are cohorts ("Eggs", "Larvae", "Pupae", "InHive", "Foragers"),
// and values are the number of individuals with the given cohort and age.
//
// NaN values are used for invalid age/cohort combinations, like foragers aged 1 day (which can only be eggs).
type AgeStructure struct {
	eggs   *globals.Eggs
	larvae *globals.Larvae
	pupae  *globals.Pupae
	inHive *globals.InHive
	aff    *globals.AgeFirstForaging
	time   *resource.Tick
	params *params.ForagerParams

	data   [][]float64
	filter generic.Filter1[comp.Age]
}

func (o *AgeStructure) Initialize(w *ecs.World) {
	o.eggs = ecs.GetResource[globals.Eggs](w)
	o.larvae = ecs.GetResource[globals.Larvae](w)
	o.pupae = ecs.GetResource[globals.Pupae](w)
	o.inHive = ecs.GetResource[globals.InHive](w)
	o.aff = ecs.GetResource[globals.AgeFirstForaging](w)
	o.time = ecs.GetResource[resource.Tick](w)
	o.params = ecs.GetResource[params.ForagerParams](w)

	o.filter = *generic.NewFilter1[comp.Age]()

	maxAge := ecs.GetResource[params.WorkerDevelopment](w).MaxLifespan
	ln := len(o.eggs.Workers) + len(o.larvae.Workers) + len(o.pupae.Workers) + maxAge + 1

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
