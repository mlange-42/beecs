package params

import (
	"encoding/json"
	"os"

	"github.com/mlange-42/arche-model/resource"
	"github.com/mlange-42/arche/ecs"
	"golang.org/x/exp/rand"
)

type Params interface {
	Apply(world *ecs.World)
	FromJSON(path string) error
}

type DefaultParams struct {
	Termination       Termination
	InitialPatches    InitialPatches
	Nursing           Nursing
	Foraging          Foraging
	Foragers          Foragers
	ForagingPeriod    ForagingPeriod
	Dance             Dance
	HandlingTime      HandlingTime
	WorkerMortality   WorkerMortality
	DroneMortality    DroneMortality
	HoneyNeeds        HoneyNeeds
	PollenNeeds       PollenNeeds
	Stores            Stores
	WorkerDevelopment WorkerDevelopment
	DroneDevelopment  DroneDevelopment
	InitialPopulation InitialPopulation
	AgeFirstForaging  AgeFirstForaging
	Energy            EnergyContent
	InitialStores     InitialStores
	RandomSeed        RandomSeed
}

func (p *DefaultParams) FromJSON(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	decoder.DisallowUnknownFields()
	return decoder.Decode(p)
}

func (p *DefaultParams) Apply(world *ecs.World) {
	// Random seed
	seed := p.RandomSeed
	if seed.Seed <= 0 {
		seed.Seed = int(rand.Int31())
	}
	rng := ecs.GetResource[resource.Rand](world)
	rng.Seed(uint64(seed.Seed))

	pCopy := *p

	// Resources
	ecs.AddResource(world, &seed)
	ecs.AddResource(world, &pCopy.Termination)
	ecs.AddResource(world, &pCopy.WorkerDevelopment)
	ecs.AddResource(world, &pCopy.DroneDevelopment)
	ecs.AddResource(world, &pCopy.WorkerMortality)
	ecs.AddResource(world, &pCopy.DroneMortality)
	ecs.AddResource(world, &pCopy.AgeFirstForaging)
	ecs.AddResource(world, &pCopy.Foragers)
	ecs.AddResource(world, &pCopy.Foraging)
	ecs.AddResource(world, &pCopy.ForagingPeriod)
	ecs.AddResource(world, &pCopy.HandlingTime)
	ecs.AddResource(world, &pCopy.Dance)
	ecs.AddResource(world, &pCopy.Energy)
	ecs.AddResource(world, &pCopy.Stores)
	ecs.AddResource(world, &pCopy.HoneyNeeds)
	ecs.AddResource(world, &pCopy.PollenNeeds)
	ecs.AddResource(world, &pCopy.Nursing)
	ecs.AddResource(world, &pCopy.InitialPopulation)
	ecs.AddResource(world, &pCopy.InitialStores)
	ecs.AddResource(world, &pCopy.InitialPatches)
}
