package res

type Params struct {
	SquadronSize int
}

type AgeFirstForaging struct {
	Base    int
	Min     int
	Max     int
	Current int
}

type WorkerMortality struct {
	Eggs   float64
	Larvae float64
	Pupae  float64
	InHive float64

	MaxLifespan int
	MaxMilage   float32
}

type DroneMortality struct {
	Eggs   float64
	Larvae float64
	Pupae  float64
	InHive float64
}

type StoreThresholds struct {
	IdealPollenStore float64
	DecentHoneyStore float64
}
