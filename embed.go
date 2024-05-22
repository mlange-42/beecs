package beecs

import "embed"

// Embedded data like time series of daily foraging hours.
//
//go:embed data
var Data embed.FS
