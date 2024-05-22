package util

import "math"

// Season from HoPoMo. Inverted using y=1-x.
func Season(t int64) float64 {
	d := float64(t % 365)
	x1, x2, x3, x4, x5 := 385.0, 25.0, 36.0, 155.0, 30.0
	s1 := (1 - (1 / (1 + x1*math.Exp(-2*d/x2))))
	s2 := (1 / (1 + x3*math.Exp(-2*(d-x4)/x5)))

	return 1.0 - math.Max(s1, s2)
}
