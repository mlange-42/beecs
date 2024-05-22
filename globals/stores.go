package globals

// Stores of the hive.
type Stores struct {
	Honey  float64 // Stored honey [kJ].
	Pollen float64 // Stored pollen [g].

	DecentHoney float64 // Amount of honey currently considered decent [kJ].
	IdealPollen float64 // Amount of pollen currently considered ideal [g].

	ProteinFactorNurses float64 // Current protein store of nurse bees, as fraction of the maximum.
}
