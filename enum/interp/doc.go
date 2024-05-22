// Package interp provides an enumeration of interpolation methods.
package interp

// Interpolation type alias for use as enumeration.
type Interpolation uint8

const (
	// Step-wise interpolation.
	Step Interpolation = iota
	// Linear interpolation.
	Linear
)
