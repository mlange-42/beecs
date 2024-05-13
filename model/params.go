package model

import (
	"encoding/json"
	"os"

	"github.com/mlange-42/beecs/model/comp"
	"github.com/mlange-42/beecs/model/res"
)

type Params struct {
	Patches           []comp.PatchConfig
	Nursing           res.NursingParams
	Foraging          res.ForagingParams
	Forager           res.ForagerParams
	Dance             res.DanceParams
	HoneyNeeds        res.HoneyNeeds
	WorkerMortality   res.WorkerMortality
	DroneMortality    res.DroneMortality
	HandlingTime      res.HandlingTimeParams
	PollenNeeds       res.PollenNeeds
	Stores            res.StoreParams
	WorkerDevelopment res.WorkerDevelopment
	InitialPopulation res.InitialPopulation
	DroneDevelopment  res.DroneDevelopment
	AgeFirstForaging  res.AgeFirstForagingParams
	Energy            res.EnergyParams
	InitialStores     res.InitialStores
}

func DefaultParams() Params {
	return Params{
		WorkerDevelopment: res.WorkerDevelopment{
			EggTime:     3,
			LarvaeTime:  6,
			PupaeTime:   12,
			MaxLifespan: 290,
		},
		DroneDevelopment: res.DroneDevelopment{
			EggTime:     3,
			LarvaeTime:  7,
			PupaeTime:   14,
			MaxLifespan: 37,
		},
		WorkerMortality: res.WorkerMortality{
			Eggs:      0.03,
			Larvae:    0.01,
			Pupae:     0.001,
			InHive:    0.004,
			MaxMilage: 800,
		},
		DroneMortality: res.DroneMortality{
			Eggs:   0.064,
			Larvae: 0.044,
			Pupae:  0.005,
			InHive: 0.05,
		},
		AgeFirstForaging: res.AgeFirstForagingParams{
			Base: 21,
			Min:  7,
			Max:  50,
		},
		Forager: res.ForagerParams{
			FlightVelocity: 6.5,      // [m/s]
			FlightCostPerM: 0.000006, //[kJ/m]
			NectarLoad:     50,       // [muL]
			PollenLoad:     0.015,    // [g]
			MaxKmPerDay:    7299,     // ???
			SquadronSize:   100,
		},
		Foraging: res.ForagingParams{
			ProbBase:      0.01,
			ProbHigh:      0.05,
			ProbEmergency: 0.2,

			SearchLength: 6.5 * 60 * 17, // [m] (6630m, 17 min)

			EnergyOnFlower:  0.2,
			MortalityPerSec: 0.00001,

			StopProbability:     0.3,
			AbandonPollenPerSec: 0.00002,
		},
		HandlingTime: res.HandlingTimeParams{
			NectarGathering:      1200,
			PollenGathering:      600,
			NectarUnloading:      116,
			PollenUnloading:      210,
			ConstantHandlingTime: false,
		},
		Dance: res.DanceParams{
			Slope:                       1.16,
			Intercept:                   0.0,
			MaxCircuits:                 117,
			FindProbability:             0.5,
			PollenDanceFollowers:        2,
			MaxProportionPollenForagers: 0.8,
		},
		Energy: res.EnergyParams{
			Honey:   12.78,
			Scurose: 0.00582,
		},
		Stores: res.StoreParams{
			IdealPollenStoreDays: 7,
			MinIdealPollenStore:  250.0,
			MaxHoneyStoreKg:      50.0,
			ProteinStoreNurse:    7, // [d]
		},
		HoneyNeeds: res.HoneyNeeds{
			WorkerResting:    11.0,  // [mg/d]
			WorkerNurse:      53.42, // [mg/d]
			WorkerLarvaTotal: 65.4,  // [mg]
			DroneLarva:       19.2,  // [mg/d]
			Drone:            10.0,  // [mg/d]
		},
		PollenNeeds: res.PollenNeeds{
			WorkerLarvaTotal: 142.0, // [mg]
			DroneLarva:       50.0,  // [mg/d]
			Worker:           1.5,   // [mg/d]
			Drone:            2.0,   // [mg/d]
		},
		Nursing: res.NursingParams{
			MaxBroodNurseRatio:         3.0,
			ForagerNursingContribution: 0.2,
			MaxEggsPerDay:              1600,
			DroneEggsProportion:        0.04,
			EggNursingLimit:            true,
			MaxBroodCells:              200_000,
			DroneEggLayingSeasonStart:  115,
			DroneEggLayingSeasonEnd:    240,
		},
		InitialPopulation: res.InitialPopulation{
			Count:     10_000,
			MinAge:    100,
			MaxAge:    160,
			MinMilage: 0,
			MaxMilage: 200,
		},
		InitialStores: res.InitialStores{
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
