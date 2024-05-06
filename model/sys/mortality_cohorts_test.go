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

func TestMortalityCohorts(t *testing.T) {
	world := ecs.NewWorld()

	fac := res.NewForagerFactory(&world)

	rng := resource.Rand{Source: rand.NewSource(0)}
	ecs.AddResource(&world, &rng)
	ecs.AddResource(&world, &res.AgeFirstForaging{Max: 5})
	ecs.AddResource(&world, &res.WorkerMortality{
		Eggs:   0.5,
		Larvae: 0.5,
		Pupae:  0.5,
		InHive: 0.5,
	})
	ecs.AddResource(&world, &res.DroneMortality{
		Eggs:   0.5,
		Larvae: 0.5,
		Pupae:  0.5,
		InHive: 0.5,
	})

	time := Time{TicksPerDay: 1}
	time.Initialize(&world)

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

	mort := MortalityCohorts{}
	mort.Initialize(&world)

	fac.CreateSquadrons(100)

	fillCohorts(init.eggs.Workers, 10000)
	fillCohorts(init.eggs.Drones, 10000)

	fillCohorts(init.larvae.Workers, 10000)
	fillCohorts(init.larvae.Drones, 10000)

	fillCohorts(init.pupae.Workers, 10000)
	fillCohorts(init.pupae.Drones, 10000)

	fillCohorts(init.inHive.Workers, 10000)
	fillCohorts(init.inHive.Drones, 10000)

	time.Update(&world)
	mort.Update(&world)

	checkCohorts(t, init.eggs.Workers, 0, 10000)
	checkCohorts(t, init.eggs.Drones, 0, 10000)

	checkCohorts(t, init.larvae.Workers, 0, 10000)
	checkCohorts(t, init.larvae.Drones, 0, 10000)

	checkCohorts(t, init.pupae.Workers, 0, 10000)
	checkCohorts(t, init.pupae.Drones, 0, 10000)

	checkCohorts(t, init.inHive.Workers, 0, 10000)
	checkCohorts(t, init.inHive.Drones, 0, 10000)

	f := generic.NewFilter1[comp.Milage]()
	q := f.Query(&world)
	cnt := q.Count()
	q.Close()

	assert.Greater(t, cnt, 0)
	assert.Less(t, cnt, 100)
}

func fillCohorts(coh []int, count int) {
	for i := range coh {
		coh[i] = count
	}
}

func checkCohorts(t *testing.T, coh []int, min int, max int) {
	for _, n := range coh {
		assert.Less(t, n, max)
		assert.Greater(t, n, min)
	}
}
