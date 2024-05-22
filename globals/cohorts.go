package globals

// Eggs contains worker and drone egg age cohorts.
type Eggs struct {
	Workers []int // Worker eggs per day of age.
	Drones  []int // Drone eggs per day of age.
}

// Larvae contains worker and drone larvae age cohorts.
type Larvae struct {
	Workers []int // Worker larvae per day since hatching.
	Drones  []int // Drone larvae per day since hatching.
}

// Pupae contains worker and drone pupae age cohorts.
type Pupae struct {
	Workers []int // Worker pupae per day of since pupation.
	Drones  []int // Drone pupae per day of since pupation.
}

// InHive contains in-hive worker and drone age cohorts.
type InHive struct {
	Workers []int // In-hive workers per day of age since emergence.
	Drones  []int // Drones per day of age since emergence.
}
