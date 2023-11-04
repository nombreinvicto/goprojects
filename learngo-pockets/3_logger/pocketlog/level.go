package pocketlog

// Level represents available logging level
type Level byte

const (

	// LevelDebug represents the lowest level of log, used for debugging
	LevelDebug Level = iota

	// LevelInfo represents information level logging
	LevelInfo

	// LevelError represents highest logging level to be used to trace errors
	LevelError
)
