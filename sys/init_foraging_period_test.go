package sys_test

import (
	"testing"

	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/beecs/globals"
	"github.com/mlange-42/beecs/params"
	"github.com/mlange-42/beecs/sys"
	"github.com/stretchr/testify/assert"
)

func TestInitForagingPeriod(t *testing.T) {
	world := ecs.NewWorld()
	ecs.AddResource(&world, &params.WorkingDirectory{})
	ecs.AddResource(&world, &params.ForagingPeriod{
		Years: [][]float64{
			make([]float64, 730),
		},
	})

	s := sys.InitForagingPeriod{}
	s.Initialize(&world)

	p := ecs.GetResource[globals.ForagingPeriodData](&world)
	assert.Equal(t, 2, len(p.Years))
}
