package utilities

import "time"

// GetToday get todays date at midnight
func GetToday() time.Time {
	return getDay(time.Now())
}

// GetYesterday get yesterdays date at midnight
func GetYesterday() time.Time {
	return getDay(time.Now().Add(-24 * time.Hour))
}

func getDay(t time.Time) time.Time {
	year, month, day := t.Date()

	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}
