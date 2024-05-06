package res

import (
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
	"github.com/mlange-42/beecs/model/comp"
)

type ForagerFactory struct {
	builder generic.Map2[comp.Milage, comp.Age]
}

func NewForagerFactory(world *ecs.World) ForagerFactory {
	return ForagerFactory{
		builder: generic.NewMap2[comp.Milage, comp.Age](world),
	}
}

func (f *ForagerFactory) CreateSquadrons(count int, dayOfBirth int) {
	q := f.builder.NewBatchQ(count)
	for q.Next() {
		_, a := q.Get()
		a.DayOfBirth = dayOfBirth
	}
}
