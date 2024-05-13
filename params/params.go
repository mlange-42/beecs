package params

import (
	"encoding/json"
	"os"

	"github.com/mlange-42/beecs/comp"
)

type Params struct {
	Patches           []comp.PatchConfig
	Nursing           Nursing
	Foraging          Foraging
	Forager           Foragers
	Dance             Dance
	HoneyNeeds        HoneyNeeds
	WorkerMortality   WorkerMortality
	DroneMortality    DroneMortality
	HandlingTime      HandlingTime
	PollenNeeds       PollenNeeds
	Stores            Stores
	WorkerDevelopment WorkerDevelopment
	InitialPopulation InitialPopulation
	DroneDevelopment  DroneDevelopment
	AgeFirstForaging  AgeFirstForaging
	Energy            EnergyContent
	InitialStores     InitialStores
}

func (p *Params) FromJSON(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	decoder := json.NewDecoder(file)
	decoder.DisallowUnknownFields()
	return decoder.Decode(p)
}
