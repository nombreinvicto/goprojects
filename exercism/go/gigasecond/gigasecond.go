// Package gigasecond contains function that adds a Gigasecond to Time.
package gigasecond

import "time"

// AddGigasecond takes in a time and adds 1 Gigaseconds to it and returns the time.
func AddGigasecond(t time.Time) time.Time {
	t = t.Add(time.Second * time.Duration(1_000_000_000))
	return t
}
