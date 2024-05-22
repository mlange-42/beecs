// Package activity provides an enumeration of forager activities.
package activity

// ForagerActivity type alias for use as enumeration.
type ForagerActivity uint8

// ForagerActivity values
const (
	// Lazy winter bees that will not forage this day.
	Lazy ForagerActivity = iota
	// Resting bees.
	Resting
	// Searching/scouting bees.
	Searching
	// Recruited bees, after waggle dance.
	Recruited
	// Experienced foragers with a known nectar or pollen patch.
	Experienced
	// Bees returning to the hive with nectar.
	BringNectar
	// Bees returning to the hive with pollen.
	BringPollen
)
