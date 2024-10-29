package params

import (
	"bytes"
	"encoding/json"

	"github.com/mlange-42/beecs/comp"
)

type WorkingDirectory struct {
	Path string
}

// RandomSeed for the model run.
type RandomSeed struct {
	Seed int // The seed. A value <= 0 forces random seeding.
}

// Termination criteria.
type Termination struct {
	MaxTicks     int  // Maximum number of ticks to run [d].
	OnExtinction bool // Whether to terminate when there are no bees anymore.
}

// AgeFirstForaging (AFF) parameters.
type AgeFirstForaging struct {
	Base int // Base AFF [d].
	Min  int // Minimum AFF [d].
	Max  int // Maximum AFF [d].
}

// Foragers parameters.
type Foragers struct {
	FlightVelocity float64 // Flight velocity [m/s].
	FlightCostPerM float64 // Flight energy cost [kJ/m].
	MaxKmPerDay    float64 // Maximum distance to fly per day [km].
	NectarLoad     float64 // Maximum nectar load of a single forager [muL].
	PollenLoad     float64 // Maximum pollen load of a single forager [g].
	SquadronSize   int     // Size of forager squadrons.
}

// HandlingTime parameters.
type HandlingTime struct {
	// Time required for gathering nectar (minimum) [s].
	NectarGathering float64
	// Time required for gathering pollen (minimum) [s].
	PollenGathering float64
	// Time required for unloading nectar [s].
	NectarUnloading float64
	// Time required for unloading pollen [s].
	PollenUnloading float64
	// Whether a constant handling time should be used.
	// Otherwise, handling time depends on patch resource depletion.
	ConstantHandlingTime bool
}

// Foraging parameters.
type Foraging struct {
	ProbBase      float64 // Base probability to start foraging.
	ProbHigh      float64 // High probability to start foraging.
	ProbEmergency float64 // Emergency probability to start foraging.

	SearchLength float64 // Search length for scouts [m].

	EnergyOnFlower  float64 // Fraction of energy usage when on a flower.
	MortalityPerSec float64 // Mortality of foragers, per second.

	StopProbability     float64 // Probability to stop foraging.
	AbandonPollenPerSec float64 // Probability to abandon a pollen patch, per second.
}

// Dance parameters.
type Dance struct {
	Slope                       float64 // Slope for calculating the number of dance followers.
	Intercept                   float64 // Intercept for calculating the number of dance followers.
	MaxCircuits                 int     // Maximum number of dance circuits.
	FindProbability             float64 // Probability to find a patch that was learned from a dance.
	PollenDanceFollowers        int     // Fixed number of dance followers for advertised pollen patches.
	MaxProportionPollenForagers float64 // Maximum proportion of foragers that can forage for pollen.
}

// WorkerDevelopment parameters.
type WorkerDevelopment struct {
	EggTime     int // Time spent as eggs [d].
	LarvaeTime  int // Time spent as larvae [d].
	PupaeTime   int // Time spent as pupae [d].
	MaxLifespan int // Maximum lifespan of adult bees [d].
}

type DroneDevelopment struct {
	EggTime     int // Time spent as eggs [d].
	LarvaeTime  int // Time spent as larvae [d].
	PupaeTime   int // Time spent as pupae [d].
	MaxLifespan int // Maximum lifespan of adult drones [d].
}

// WorkerMortality parameters.
type WorkerMortality struct {
	Eggs   float64 // Daily mortality of eggs.
	Larvae float64 // Daily mortality of larvae.
	Pupae  float64 // Daily mortality of pupae.
	InHive float64 // Daily mortality of in-hive bees and foragers.

	MaxMilage float32 // Maximum milage foragers [km].
}

// DroneMortality parameters.
type DroneMortality struct {
	Eggs   float64 // Daily mortality of eggs.
	Larvae float64 // Daily mortality of larvae.
	Pupae  float64 // Daily mortality of pupae.
	InHive float64 // Daily mortality of adult drones.
}

// EnergyContent parameters.
type EnergyContent struct {
	Honey   float64 // Energy content of honey [kJ/g].
	Scurose float64 // Energy content of sucrose [kJ/micromol].
}

// HoneyNeeds parameters.
type HoneyNeeds struct {
	WorkerResting float64 // Daily honey needs of resting adults [mg/d].
	WorkerNurse   float64 // Daily honey needs of nursing adults [mg/d].

	WorkerLarvaTotal float64 // Total honey need for worker larvae development [mg].
	DroneLarva       float64 // Daily honey needs of drone larvae [mg/d].

	Drone float64 // [mg/d].
}

// PollenNeeds parameters.
type PollenNeeds struct {
	WorkerLarvaTotal float64 // Total pollen need for worker larvae development [mg].
	DroneLarva       float64 // Daily pollen needs of drone larvae [mg/d].

	Worker float64 // Daily pollen needs of workers [mg/d].
	Drone  float64 // Daily pollen needs of drones [mg/d].
}

// Nursing parameters.
type Nursing struct {
	MaxBroodNurseRatio         float64 // Maximum brood per nurse.
	ForagerNursingContribution float64 // Contribution fraction of foragers to nursing.
	MaxEggsPerDay              int     // Maximum eggs laid by a queen per day.
	DroneEggsProportion        float64 // Proportion of drone eggs.
	EggNursingLimit            bool    // Whether to limit egg laying by the number of available nurses.
	MaxBroodCells              int     // Maximum number of brood cells in the hive.
	DroneEggLayingSeasonStart  int     // Fist day of year of the drone egg laying season.
	DroneEggLayingSeasonEnd    int     // Last day of year of the drone egg laying season.
}

// Stores parameters.
type Stores struct {
	IdealPollenStoreDays int     // Number of days the pollen store should ideally last for [d].
	MinIdealPollenStore  float64 // Minimum pollen store to consider ideal [g].
	MaxHoneyStoreKg      float64 // Maximum honey store [kg].
	DecentHoneyPerWorker float64 // Honey needed per worker to consider stores decent [g].
	ProteinStoreNurse    float64 // Number of days nurse protein stores lasts [d].
}

// InitialPopulation parameters.
type InitialPopulation struct {
	Count     int     // Number of initial foragers.
	MinAge    int     // Minimum age of initial foragers [d].
	MaxAge    int     // Maximum age of initial foragers [d].
	MinMilage float32 // Minimum milage of initial foragers [km].
	MaxMilage float32 // Maximum milage of initial foragers [km].
}

// InitialStores parameters.
type InitialStores struct {
	Honey  float64 // Initial honey store [kg].
	Pollen float64 // Initial pollen store [g].
}

// InitialPatches parameters.
type InitialPatches struct {
	Patches []comp.PatchConfig // Initial patches. Optional.
	File    string             // File to read patches from. Applied after creating Patches.
}

// initialPatchesHelper is used to unmarshal the InitialPatches struct from JSON,
// properly overwriting the default patches.
type initialPatchesHelper struct {
	Patches []comp.PatchConfig // Initial patches. Optional.
	File    string             // File to read patches from. Applied after creating Patches.
}

func (p *InitialPatches) UnmarshalJSON(jsonData []byte) error {
	helper := initialPatchesHelper{}
	reader := bytes.NewReader(jsonData)
	decoder := json.NewDecoder(reader)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&helper)
	if err != nil {
		return err
	}

	p.File = helper.File
	p.Patches = helper.Patches

	return nil
}

// ForagingPeriod parameters.
//
// Data read from files (field Files) is appended to data provided directly (field Years).
type ForagingPeriod struct {
	Years       [][]float64 // Foraging period per day [h] as raw data. Each row must have a whole-numbered multiple of 365 entries.
	Files       []string    // Files with daily foraging period data to use.
	Builtin     bool        // Whether the used files are built-in. Use local files otherwise.
	RandomYears bool        // Whether to randomize years.
}
