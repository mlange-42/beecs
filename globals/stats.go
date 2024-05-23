package globals

// PopulationStats contains summarized population numbers for development stages.
//
// PopulationStats is updated at the end of each simulation step.
// Thus, it contains stats of the previous step.
type PopulationStats struct {
	WorkerEggs      int
	WorkerLarvae    int
	WorkerPupae     int
	WorkersInHive   int
	WorkersForagers int

	DroneEggs    int
	DroneLarvae  int
	DronePupae   int
	DronesInHive int

	TotalBrood      int
	TotalAdults     int
	TotalPopulation int
}

// Reset all stats to zero.
func (s *PopulationStats) Reset() {
	s.WorkerEggs = 0
	s.WorkerLarvae = 0
	s.WorkerPupae = 0
	s.WorkersInHive = 0
	s.WorkersForagers = 0

	s.DroneEggs = 0
	s.DroneLarvae = 0
	s.DronePupae = 0
	s.DronesInHive = 0

	s.TotalBrood = 0
	s.TotalAdults = 0
	s.TotalPopulation = 0
}

// ConsumptionStats contains statistics on daily consumption.
type ConsumptionStats struct {
	HoneyDaily float64 // Today's honey consumption [mg].
}

// Reset all stats to zero.
func (s *ConsumptionStats) Reset() {
	s.HoneyDaily = 0
}

// ForagingStats contains statistics on foraging per foraging round.
type ForagingStats struct {
	Rounds []ForagingRound
}

// Reset all stats.
func (s *ForagingStats) Reset() {
	s.Rounds = s.Rounds[:0]
}

// ForagingRound contains statistics for a single foraging round.
// Not used as an ECS resource directly!
type ForagingRound struct {
	Lazy      int
	Resting   int
	Searching int
	Recruited int
	Nectar    int
	Pollen    int
}
