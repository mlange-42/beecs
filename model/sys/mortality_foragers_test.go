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

func TestMortalityForagers(t *testing.T) {
	world := ecs.NewWorld()

	fac := res.NewForagerFactory(&world)

	time := resource.Tick{}
	ecs.AddResource(&world, &time)
	ecs.AddResource(&world, &resource.Rand{Source: rand.NewSource(0)})
	ecs.AddResource(&world, &res.AgeFirstForagingParams{Max: 5})
	ecs.AddResource(&world, &res.WorkerMortality{
		Eggs:        0.5,
		Larvae:      0.5,
		Pupae:       0.5,
		InHive:      0.5,
		MaxLifespan: 390,
		MaxMilage:   200,
	})
	ecs.AddResource(&world, &res.DroneMortality{
		Eggs:   0.5,
		Larvae: 0.5,
		Pupae:  0.5,
		InHive: 0.5,
	})

	init := InitCohorts{
		EggTimeWorker:    2,
		LarvaeTimeWorker: 3,
		PupaeTimeWorker:  4,
		EggTimeDrone:     3,
		LarvaeTimeDrone:  4,
		PupaeTimeDrone:   5,
		LifespanDrone:    6,
	}
	init.Initialize(&world)

	mort := MortalityForagers{}
	mort.Initialize(&world)

	fac.CreateSquadrons(100, -100)

	mort.Update(&world)

	f := generic.NewFilter1[comp.Milage]()
	q := f.Query(&world)
	cnt := q.Count()
	q.Close()

	assert.Greater(t, cnt, 0)
	assert.Less(t, cnt, 100)

	time.Tick = 400
	mort.Update(&world)

	q = f.Query(&world)
	cnt = q.Count()
	q.Close()

	assert.Equal(t, 0, cnt)
}