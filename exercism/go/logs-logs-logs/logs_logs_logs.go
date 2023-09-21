package logs

import (
	"strings"
	"unicode/utf8"
)

// Application identifies the application emitting the given log.
func Application(log string) string {

	for _, rune_ := range log {
		if rune_ == '❗' {
			return "recommendation"
		} else if rune_ == '🔍' {
			return "search"
		} else if rune_ == '☀' {
			return "weather"
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	return strings.ReplaceAll(log, string(oldRune), string(newRune))
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	charCount := utf8.RuneCountInString(log)
	if charCount <= limit {
		return true
	}
	return false
}
