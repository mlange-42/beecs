package util

import "math/rand/v2"

// RandWrapper is a wrapper to allow for using a math/rand source for distributions.
type RandWrapper struct {
	Src rand.Source
}

// Uint64 random number.
func (r *RandWrapper) Uint64() uint64 {
	return r.Src.Uint64()
}

// Seed does nothing and is just there to fulfill the interface.
func (r *RandWrapper) Seed(seed uint64) {}
