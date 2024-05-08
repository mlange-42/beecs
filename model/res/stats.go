package res

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
}
