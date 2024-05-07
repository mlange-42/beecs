package sys

import (
	"testing"

	"github.com/mlange-42/arche-model/resource"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
	"github.com/mlange-42/beecs/model/comp"
	"github.com/mlange-42/beecs/model/res"
	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/rand"
)

func TestInitPopulation(t *testing.T) {
	world := ecs.NewWorld()

	fac := res.NewForagerFactory(&world)
	ecs.AddResource(&world, &fac)
	ecs.AddResource(&world, &res.Params{SquadronSize: 100})
	ecs.AddResource(&world, &resource.Rand{Source: rand.NewSource(0)})

	s := InitPopulation{
		InitialCount: 10_000,
		MinAge:       100,
		MaxAge:       160,
		MinMilage:    0,
		MaxMilage:    200,
	}
	s.Initialize(&world)

	filter := generic.NewFilter2[comp.Milage, comp.Age]()
	query := filter.Query(&world)
	assert.Equal(t, 100, query.Count())

	for query.Next() {
		m, a := query.Get()

		assert.GreaterOrEqual(t, m.Total, float32(0.0))
		assert.Less(t, m.Total, float32(200.0))

		assert.GreaterOrEqual(t, a.DayOfBirth, -160)
		assert.Less(t, a.DayOfBirth, -100)
	}
}
