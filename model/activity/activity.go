package activity

type ForagerActivity uint8

const (
	Lazy ForagerActivity = iota
	Resting
	Searching
	Recruited
	Experienced
	BringNectar
	BringPollen
)
