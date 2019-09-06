// Package datetime contains various functions and types that related to date and time
package datetime

import (
	"time"
)

// DurationString is string that contains duration string
type DurationString string

// Duration returns time.Duration value
func (t DurationString) Duration() (time.Duration, error) {
	return time.ParseDuration(string(t))
}

// DurationDefault does the same as Duration(), but returns passed default value instead of error
func (t DurationString) DurationDefault(defaultDuration time.Duration) time.Duration {
	val, err := t.Duration()
	if err != nil {
		return defaultDuration
	}

	return val
}
