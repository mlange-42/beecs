package res

type Params struct {
	SquadronSize int
}

type AgeFirstForagingParams struct {
	Base int
	Min  int
	Max  int
}

type ForagingParams struct {
	ProbBase      float64
	ProbHigh      float64
	ProbEmergency float64

	FlightVelocity              float64 // [m/s]
	SearchLength                float64 // [m]
	MaxProportionPollenForagers float64

	EnergyOnFlower  float64
	MortalityPerSec float64
	FlightCostPerM  float64 // [kJ/m]

	NectarLoad           float64 // [muL]
	PollenLoad           float64 // [g]
	TimeNectarGathering  float64 // [s]
	TimePollenGathering  float64 // [s]
	ConstantHandlingTime bool
	StopProbability      float64
	AbandonPollenPerSec  float64
	MaxKmPerDay          float64

	TimeUnloadingNectar float64
	TimeUnloadingPollen float64
}

type DanceParams struct {
	Slope                float64
	Intercept            float64
	MaxCircuits          int
	FindProbability      float64
	PollenDanceFollowers int
}

type WorkerDevelopment struct {
	EggTime     int
	LarvaeTime  int
	PupaeTime   int
	MaxLifespan int
}

type DroneDevelopment struct {
	EggTime     int
	LarvaeTime  int
	PupaeTime   int
	MaxLifespan int
}

type WorkerMortality struct {
	Eggs   float64
	Larvae float64
	Pupae  float64
	InHive float64

	MaxMilage float32
}

type DroneMortality struct {
	Eggs   float64
	Larvae float64
	Pupae  float64
	InHive float64

	MaxLifespan int
}

type EnergyParams struct {
	EnergyHoney   float64 // [kJ/g]
	EnergyScurose float64 // [kJ/micromol]
}

type HoneyNeeds struct {
	WorkerResting float64
	WorkerNurse   float64

	WorkerLarvaTotal float64
	DroneLarva       float64

	Drone float64
}

type PollenNeeds struct {
	WorkerLarvaTotal float64
	DroneLarva       float64

	Worker float64
	Drone  float64
}

type NurseParams struct {
	MaxBroodNurseRatio         float64
	ForagerNursingContribution float64
	MaxEggsPerDay              int
	DroneEggsProportion        float64
	EggNursingLimit            bool
	MaxBroodCells              int
}

type StoreParams struct {
	IdealPollenStoreDays int
	MinIdealPollenStore  float64
	MaxHoneyStoreKg      float64
	ProteinStoreNurse    float64 // [d]
}
