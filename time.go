package utilities

import "time"

// DaysIn return number of days depends on year and month
func DaysIn(month, year int) int {
	// This is equivalent to time.daysIn(m, year).
	return time.Date(year, time.Month(month)+1, 0, 0, 0, 0, 0, time.UTC).Day()
}
