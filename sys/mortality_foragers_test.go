package sys

import (
	"math/rand/v2"
	"testing"

	"github.com/mlange-42/ark-tools/resource"
	"github.com/mlange-42/ark/ecs"
	"github.com/mlange-42/beecs/comp"
	"github.com/mlange-42/beecs/globals"
	"github.com/mlange-42/beecs/params"
	"github.com/stretchr/testify/assert"
)

func TestMortalityForagers(t *testing.T) {
	world := ecs.NewWorld()

	fac := globals.NewForagerFactory(&world)

	time := resource.Tick{}
	ecs.AddResource(&world, &time)
	ecs.AddResource(&world, &resource.Rand{Source: rand.NewPCG(0, 0)})
	ecs.AddResource(&world, &params.AgeFirstForaging{Max: 5})
	ecs.AddResource(&world, &params.WorkerDevelopment{
		EggTime:     2,
		LarvaeTime:  3,
		PupaeTime:   4,
		MaxLifespan: 390,
	})
	ecs.AddResource(&world, &params.DroneDevelopment{
		EggTime:     3,
		LarvaeTime:  4,
		PupaeTime:   5,
		MaxLifespan: 6,
	})
	ecs.AddResource(&world, &params.WorkerMortality{
		Eggs:      0.5,
		Larvae:    0.5,
		Pupae:     0.5,
		InHive:    0.5,
		MaxMilage: 200,
	})
	ecs.AddResource(&world, &params.DroneMortality{
		Eggs:   0.5,
		Larvae: 0.5,
		Pupae:  0.5,
		InHive: 0.5,
	})

	init := InitCohorts{}
	init.Initialize(&world)

	mort := MortalityForagers{}
	mort.Initialize(&world)

	fac.CreateSquadrons(100, -100)

	mort.Update(&world)

	f := ecs.NewFilter1[comp.Milage](&world)
	q := f.Query()
	cnt := q.Count()
	q.Close()

	assert.Greater(t, cnt, 0)
	assert.Less(t, cnt, 100)

	time.Tick = 400
	mort.Update(&world)

	q = f.Query()
	cnt = q.Count()
	q.Close()

	assert.Equal(t, 0, cnt)
}
