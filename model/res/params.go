package res

type Params struct {
	SquadronSize int
}

type AgeFirstForagingParams struct {
	Base int
	Min  int
	Max  int
}

type ForagingProbabilityParams struct {
	Base      float64
	High      float64
	Emergency float64
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
	EnergyHoney   float64
	EnergyScurose float64
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

	IdealStoreDays int
	MinIdealStore  float64
}

type NurseParams struct {
	MaxBroodNurseRatio         float64
	ForagerNursingContribution float64
}
