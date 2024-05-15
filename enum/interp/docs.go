// Package interp provides an enumeration of interpolation methods.
package interp

type Interpolation uint8

const (
	Step Interpolation = iota
	Linear
)
