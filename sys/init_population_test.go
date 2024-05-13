package sys

import (
	"testing"

	"github.com/mlange-42/arche-model/resource"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
	"github.com/mlange-42/beecs/comp"
	"github.com/mlange-42/beecs/globals"
	"github.com/mlange-42/beecs/params"
	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/rand"
)

func TestInitPopulation(t *testing.T) {
	world := ecs.NewWorld()

	fac := globals.NewForagerFactory(&world)
	ecs.AddResource(&world, &fac)
	ecs.AddResource(&world, &params.Foragers{SquadronSize: 100})
	ecs.AddResource(&world, &resource.Rand{Source: rand.NewSource(0)})
	ecs.AddResource(&world, &params.InitialPopulation{
		Count:     10_000,
		MinAge:    100,
		MaxAge:    160,
		MinMilage: 0,
		MaxMilage: 200,
	})

	s := InitPopulation{}
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
