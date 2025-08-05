package util

type numbers interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

// MinInt returns the smaller one of two int values.
func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// MaxInt returns the larger one of two int values.
func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Clamp a numeric value to the range low...high, inclusively.
func Clamp[T numbers](v, low, high T) T {
	if v < low {
		return low
	}
	if v > high {
		return high
	}
	return v
}
