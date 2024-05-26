package util

import "time"

// TickToDate converts a model tick to a date, without leap years.
func TickToDate(tick int64) time.Time {

	year := tick / 365
	day := tick % 365
	date := time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC).Add(time.Hour * 24 * time.Duration(day))

	return time.Date(int(year), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC)
}
