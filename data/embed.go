package data

import "embed"

// Embedded data for daily foraging hours.
//
// # Available data
//
//   - foraging-period/berlin2000.txt
//   - foraging-period/berlin2001.txt
//   - foraging-period/berlin2002.txt
//   - foraging-period/berlin2003.txt
//   - foraging-period/berlin2004.txt
//   - foraging-period/berlin2005.txt
//   - foraging-period/berlin2006.txt
//
//go:embed foraging-period
var ForagingPeriod embed.FS
