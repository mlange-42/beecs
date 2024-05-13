package params

import (
	"encoding/json"
	"os"

	"github.com/mlange-42/beecs/model/comp"
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

func Default() Params {
	return Params{
		WorkerDevelopment: WorkerDevelopment{
			EggTime:     3,
			LarvaeTime:  6,
			PupaeTime:   12,
			MaxLifespan: 290,
		},
		DroneDevelopment: DroneDevelopment{
			EggTime:     3,
			LarvaeTime:  7,
			PupaeTime:   14,
			MaxLifespan: 37,
		},
		WorkerMortality: WorkerMortality{
			Eggs:      0.03,
			Larvae:    0.01,
			Pupae:     0.001,
			InHive:    0.004,
			MaxMilage: 800,
		},
		DroneMortality: DroneMortality{
			Eggs:   0.064,
			Larvae: 0.044,
			Pupae:  0.005,
			InHive: 0.05,
		},
		AgeFirstForaging: AgeFirstForaging{
			Base: 21,
			Min:  7,
			Max:  50,
		},
		Forager: Foragers{
			FlightVelocity: 6.5,      // [m/s]
			FlightCostPerM: 0.000006, //[kJ/m]
			NectarLoad:     50,       // [muL]
			PollenLoad:     0.015,    // [g]
			MaxKmPerDay:    7299,     // ???
			SquadronSize:   100,
		},
		Foraging: Foraging{
			ProbBase:      0.01,
			ProbHigh:      0.05,
			ProbEmergency: 0.2,

			SearchLength: 6.5 * 60 * 17, // [m] (6630m, 17 min)

			EnergyOnFlower:  0.2,
			MortalityPerSec: 0.00001,

			StopProbability:     0.3,
			AbandonPollenPerSec: 0.00002,
		},
		HandlingTime: HandlingTime{
			NectarGathering:      1200,
			PollenGathering:      600,
			NectarUnloading:      116,
			PollenUnloading:      210,
			ConstantHandlingTime: false,
		},
		Dance: Dance{
			Slope:                       1.16,
			Intercept:                   0.0,
			MaxCircuits:                 117,
			FindProbability:             0.5,
			PollenDanceFollowers:        2,
			MaxProportionPollenForagers: 0.8,
		},
		Energy: EnergyContent{
			Honey:   12.78,
			Scurose: 0.00582,
		},
		Stores: Stores{
			IdealPollenStoreDays: 7,
			MinIdealPollenStore:  250.0,
			MaxHoneyStoreKg:      50.0,
			ProteinStoreNurse:    7, // [d]
		},
		HoneyNeeds: HoneyNeeds{
			WorkerResting:    11.0,  // [mg/d]
			WorkerNurse:      53.42, // [mg/d]
			WorkerLarvaTotal: 65.4,  // [mg]
			DroneLarva:       19.2,  // [mg/d]
			Drone:            10.0,  // [mg/d]
		},
		PollenNeeds: PollenNeeds{
			WorkerLarvaTotal: 142.0, // [mg]
			DroneLarva:       50.0,  // [mg/d]
			Worker:           1.5,   // [mg/d]
			Drone:            2.0,   // [mg/d]
		},
		Nursing: Nursing{
			MaxBroodNurseRatio:         3.0,
			ForagerNursingContribution: 0.2,
			MaxEggsPerDay:              1600,
			DroneEggsProportion:        0.04,
			EggNursingLimit:            true,
			MaxBroodCells:              200_000,
			DroneEggLayingSeasonStart:  115,
			DroneEggLayingSeasonEnd:    240,
		},
		InitialPopulation: InitialPopulation{
			Count:     10_000,
			MinAge:    100,
			MaxAge:    160,
			MinMilage: 0,
			MaxMilage: 200,
		},
		InitialStores: InitialStores{
			Honey:  25,  // [kg]
			Pollen: 100, // [g]
		},

		Patches: []comp.PatchConfig{
			{
				Nectar:               20,
				NectarConcentration:  1.5,
				Pollen:               1,
				DistToColony:         1500,
				DetectionProbability: 0.2,
			},
			{
				Nectar:               20,
				NectarConcentration:  1.5,
				Pollen:               1,
				DistToColony:         500,
				DetectionProbability: 0.2,
			},
		},
	}
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
