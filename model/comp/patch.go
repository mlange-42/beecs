package comp

type Trip struct {
	DurationNectar float64 // [s]
	DurationPollen float64 // [s]
	CostNectar     float64
	CostPollen     float64
}

type HandlingTime struct {
	Nectar float64 // [s]
	Pollen float64 // [s]
}

type Mortality struct {
	Nectar float64
	Pollen float64
}

type Resource struct {
	Nectar           float64 // [muL]
	MaxNectar        float64 // [muL]
	Pollen           float64 // [g]
	MaxPollen        float64 // [g]
	EnergyEfficiency float64
}

type PatchConfig struct {
	Nectar               float64 // [L]
	NectarConcentration  float64 // [mol/L]
	Pollen               float64 // [kg]
	DistToColony         float64 // [m]
	DetectionProbability float64
}
