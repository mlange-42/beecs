package globals

// ForagingPeriodData contains data on daily foraging hours.
type ForagingPeriodData struct {
	// Foraging period per day [h].
	// First index: year, second index: day of year.
	Years [][]float64
	// The currently selected year.
	CurrentYear int
}

// ForagingPeriod of the current day.
type ForagingPeriod struct {
	SecondsToday int // Today's foraging period [s].
}
