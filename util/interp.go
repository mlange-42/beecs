package util

import (
	"fmt"

	"github.com/mlange-42/beecs/enum/interp"
)

func Interpolate(data [][2]float64, day float64, inter interp.Interpolation) float64 {
	if inter == interp.Step {
		return InterpolateStep(data, day)
	}
	return InterpolateLinear(data, day)
}

func InterpolateLinear(data [][2]float64, day float64) float64 {
	dpy := 365.0
	last := len(data) - 1
	if day <= data[0][0] {
		before := (dpy - data[last][0]) + day
		after := data[0][0] - day
		frac := before / (before + after)
		return data[0][1]*frac + data[last][1]*(1-frac)
	}
	if day >= data[last][0] {
		before := day - data[last][0]
		after := (dpy - day) + data[0][0]
		frac := before / (before + after)
		return data[0][1]*frac + data[last][1]*(1-frac)
	}
	for i := 0; i < last; i++ {
		t1 := data[i][0]
		if day == t1 {
			return data[i][1]
		}
		t2 := data[i+1][0]
		if day == t2 {
			return data[i+1][1]
		}
		if t2 < day {
			continue
		}
		frac := (day - t1) / (t2 - t1)
		return data[i+1][1]*frac + data[i][1]*(1-frac)
	}
	panic(fmt.Sprintf("unable to do linear interpolation - code should not be reachable\n%v at %f", data, day))
}

func InterpolateStep(data [][2]float64, day float64) float64 {
	if day < data[0][0] {
		return data[len(data)-1][1]
	}
	v := 0.0
	for _, tv := range data {
		if tv[0] > day {
			break
		}
		v = tv[1]
	}
	return v
}
