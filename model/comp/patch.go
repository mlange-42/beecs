package comp

type TripDuration struct {
	Nectar float64 // [s]
	Pollen float64 // [s]
}

type PatchConfig struct {
	Nectar               float64 // [L]
	NectarConcentration  float64 // [mol/L]
	Pollen               float64 // [kg]
	DistToColony         float64 // [m]
	DetectionProbability float64
}
