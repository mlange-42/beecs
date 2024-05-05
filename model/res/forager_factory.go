package res

import (
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
	"github.com/mlange-42/beecs/model/comp"
)

type ForagerFactory struct {
	builder generic.Map1[comp.Milage]
}

func NewForagerFactory(world *ecs.World) ForagerFactory {
	return ForagerFactory{
		builder: generic.NewMap1[comp.Milage](world),
	}
}

func (f *ForagerFactory) CreateSquadrons(count int) {
	f.builder.NewBatch(count)
}
