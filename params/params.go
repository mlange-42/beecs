package params

import (
	"encoding/json"
	"os"
	"time"

	"github.com/mlange-42/arche-model/resource"
	"github.com/mlange-42/arche/ecs"
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
	Forager           Foragers
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
	decoder := json.NewDecoder(file)
	decoder.DisallowUnknownFields()
	return decoder.Decode(p)
}

func (p *DefaultParams) Apply(world *ecs.World) {
	// Random seed
	rand := ecs.GetResource[resource.Rand](world)
	if p.RandomSeed.Seed > 0 {
		rand.Seed(uint64(p.RandomSeed.Seed))
	} else {
		rand.Seed(uint64(time.Now().UnixNano()))
	}
	// Resources
	ecs.AddResource(world, &p.Termination)
	ecs.AddResource(world, &p.WorkerDevelopment)
	ecs.AddResource(world, &p.DroneDevelopment)
	ecs.AddResource(world, &p.WorkerMortality)
	ecs.AddResource(world, &p.DroneMortality)
	ecs.AddResource(world, &p.AgeFirstForaging)
	ecs.AddResource(world, &p.Forager)
	ecs.AddResource(world, &p.Foraging)
	ecs.AddResource(world, &p.HandlingTime)
	ecs.AddResource(world, &p.Dance)
	ecs.AddResource(world, &p.Energy)
	ecs.AddResource(world, &p.Stores)
	ecs.AddResource(world, &p.HoneyNeeds)
	ecs.AddResource(world, &p.PollenNeeds)
	ecs.AddResource(world, &p.Nursing)
	ecs.AddResource(world, &p.InitialPopulation)
	ecs.AddResource(world, &p.InitialStores)
	ecs.AddResource(world, &p.InitialPatches)
}
