package util

type numbers interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Clamp[T numbers](v, low, high T) T {
	if v < low {
		return low
	}
	if v > high {
		return high
	}
	return v
}
