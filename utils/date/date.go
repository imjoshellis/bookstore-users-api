// Package date has date utils
package date

import "time"

const (
	apiDateLayout = "2006-01-02T15:04:05Z"
)

// GetNow returns the current UTC time.
func GetNow() time.Time {
	return time.Now().UTC()
}

// GetNowString returns a formatted string of the current UTC time.
func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}
