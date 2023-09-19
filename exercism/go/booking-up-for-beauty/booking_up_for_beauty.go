package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	layout := "1/02/2006 15:04:05"
	t, _ := time.Parse(layout, date)
	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05"
	t, _ := time.Parse(layout, date)

	if t.Before(time.Now()) {
		return true
	}
	return false

}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
	t, _ := time.Parse(layout, date)

	if t.Hour() >= 12 && t.Hour() < 18 {
		return true
	}
	return false
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	layout := "1/2/2006 15:04:05"
	t, _ := time.Parse(layout, date)

	message := fmt.Sprintf("You have an appointment on %v, %v %v, %v, at %v:%v.",
		t.Weekday().String(), t.Month().String(), t.Day(), t.Year(), t.Hour(), t.Minute())
	return message
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {

	currentYear := time.Now().Year()
	anniversary := fmt.Sprintf("%v-09-15 00:00:00 +0000 UTC", currentYear)
	t, _ := time.Parse("2006-01-02 15:04:05 -0700 MST", anniversary)
	return t
}
