package comp

type ForagerActivity uint8

const (
	ActivityLazy ForagerActivity = iota
)

type Milage struct {
	Today float32
	Total float32
}

type Age struct {
	DayOfBirth int
}

type Activity struct {
	Current ForagerActivity
}
