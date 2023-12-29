package helpers

import "time"

func GetStartEndOfToday() (time.Time, time.Time) {
	today := time.Now()
	startOfDay := time.Date(today.Year(), today.Month(), today.Day(), 0, 0, 0, 0, today.Location())
	endOfDay := startOfDay.AddDate(0, 0, 1).Add(-time.Nanosecond)
	return startOfDay, endOfDay
}
