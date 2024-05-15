package comp

import "github.com/mlange-42/beecs/enum/interp"

// PatchConfig for initialization of flower patches.
type PatchConfig struct {
	DistToColony  float64        // Distance to the colony [m].
	ConstantPatch *ConstantPatch `json:",omitempty"`
	SeasonalPatch *SeasonalPatch `json:",omitempty"`
	ScriptedPatch *ScriptedPatch `json:",omitempty"`
}

type ConstantPatch struct {
	Nectar               float64 // Maximum of available nectar [L].
	Pollen               float64 // Maximum of available pollen [kg].
	NectarConcentration  float64 // Sucrose concentration in the nectar [mol/L].
	DetectionProbability float64 // Detection probability, e.g. from BeeScout.
}

type SeasonalPatch struct {
	SeasonShift int     // Shift of the season [d]
	MaxNectar   float64 // Maximum of available nectar [L].
	MaxPollen   float64 // Maximum of available pollen [kg].

	NectarConcentration  float64 // Sucrose concentration in the nectar [mol/L].
	DetectionProbability float64 // Detection probability, e.g. from BeeScout.
}

type ScriptedPatch struct {
	Nectar               [][2]float64 // Maximum of available nectar [L].
	Pollen               [][2]float64 // Maximum of available pollen [kg].
	NectarConcentration  [][2]float64 // Sucrose concentration in the nectar [mol/L].
	DetectionProbability [][2]float64 // Detection probability, e.g. from BeeScout.
	Interpolation        interp.Interpolation
}

// PatchProperties component for flower patches.
type PatchProperties struct {
	MaxNectar            float64 // Maximum of available nectar [L].
	NectarConcentration  float64 // Sucrose concentration in the nectar [mol/L].
	MaxPollen            float64 // Maximum of available pollen [kg].
	DetectionProbability float64 // Detection probability, e.g. from BeeScout.
}

// PatchDistance component for flower patches.
type PatchDistance struct {
	DistToColony float64 // Distance to the colony [m].
}

// Resource component for flower patches.
//
// Holds information on available nectar and pollen resources.
type Resource struct {
	Nectar           float64 // Currently available nectar [muL].
	MaxNectar        float64 // Maximum currently available nectar (before any collecting) [muL].
	Pollen           float64 // Currently available pollen [g].
	MaxPollen        float64 // Maximum currently available pollen (before any collecting) [g].
	EnergyEfficiency float64 // Energy efficiency of nectar foraging.
}

// HandlingTime component for flower patches.
//
// Holds information on current nectar and pollen handling times.
type HandlingTime struct {
	Nectar float64 // Nectar handling time [s].
	Pollen float64 // Pollen handling time [s].
}

// Trip component for flower patches.
//
// Holds information on current nectar and pollen trip durations and costs.
type Trip struct {
	DurationNectar float64 // Current trip duration for nectar [s].
	DurationPollen float64 // Current trip duration for pollen [s].
	CostNectar     float64 // Current trip energy cost for nectar [kJ].
	CostPollen     float64 // Current trip energy cost for pollen [kJ].
}

// Mortality component for flower patches.
//
// Holds information on current mortality for foragers.
type Mortality struct {
	Nectar float64 // Current mortality for nectar foragers.
	Pollen float64 // Current mortality for pollen foragers.
}

// Dance component for flower patches.
//
// Holds information on current dance circuits for the flower patch.
type Dance struct {
	Circuits float64 // Current number of dance circuits.
}
