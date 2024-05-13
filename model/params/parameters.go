package params

type AgeFirstForagingParams struct {
	Base int
	Min  int
	Max  int
}

type ForagerParams struct {
	FlightVelocity float64 // [m/s]
	FlightCostPerM float64 // [kJ/m]
	MaxKmPerDay    float64
	NectarLoad     float64 // [muL]
	PollenLoad     float64 // [g]
	SquadronSize   int
}

type HandlingTimeParams struct {
	NectarGathering      float64 // [s]
	PollenGathering      float64 // [s]
	NectarUnloading      float64 // [s]
	PollenUnloading      float64 // [s]
	ConstantHandlingTime bool
}

type ForagingParams struct {
	ProbBase      float64
	ProbHigh      float64
	ProbEmergency float64

	SearchLength float64 // [m]

	EnergyOnFlower  float64
	MortalityPerSec float64

	StopProbability     float64
	AbandonPollenPerSec float64
}

type DanceParams struct {
	Slope                       float64
	Intercept                   float64
	MaxCircuits                 int
	FindProbability             float64
	PollenDanceFollowers        int
	MaxProportionPollenForagers float64
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
	Honey   float64 // [kJ/g]
	Scurose float64 // [kJ/micromol]
}

type HoneyNeeds struct {
	WorkerResting float64 // [mg/d]
	WorkerNurse   float64 // [mg/d]

	WorkerLarvaTotal float64 // [mg]
	DroneLarva       float64 // [mg/d]

	Drone float64 // [mg/d]
}

type PollenNeeds struct {
	WorkerLarvaTotal float64 // [mg]
	DroneLarva       float64 // [mg/d]

	Worker float64 // [mg/d]
	Drone  float64 // [mg/d]
}

type NursingParams struct {
	MaxBroodNurseRatio         float64
	ForagerNursingContribution float64
	MaxEggsPerDay              int
	DroneEggsProportion        float64
	EggNursingLimit            bool
	MaxBroodCells              int
	DroneEggLayingSeasonStart  int
	DroneEggLayingSeasonEnd    int
}

type StoreParams struct {
	IdealPollenStoreDays int
	MinIdealPollenStore  float64
	MaxHoneyStoreKg      float64
	ProteinStoreNurse    float64 // [d]
}

type InitialPopulation struct {
	Count     int
	MinAge    int
	MaxAge    int
	MinMilage float32
	MaxMilage float32
}

type InitialStores struct {
	Honey  float64 // [kg]
	Pollen float64 // [g]
}
