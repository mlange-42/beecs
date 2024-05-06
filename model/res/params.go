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
}

type DroneMortality struct {
	Eggs   float64
	Larvae float64
	Pupae  float64
	InHive float64
}