package pocketlog

import (
	"fmt"
	"io"
	"os"
)

type Logger struct {
	threshold Level
	output    io.Writer
}

// Debugf formats and prints a message if log level is debug or higher
func (l *Logger) Debugf(format string, args ...any) {

	// in case users dont use New but use declaration
	// to create a logger object with nil values
	if l.output == nil {
		l.output = os.Stdout
	}

	if l.threshold <= LevelDebug {
		l.logf(format, args...)
	}
}

// Infof formats and prints a message if the log level is info or higher
func (l *Logger) Infof(format string, args ...any) {
	if l.output == nil {
		l.output = os.Stdout
	}

	if l.threshold <= LevelInfo {
		l.logf(format, args...)
	}
}

// Errorf formats and prints a message if the log level is error or higher
func (l *Logger) Errorf(format string, args ...any) {
	if l.output == nil {
		l.output = os.Stdout
	}

	if l.threshold <= LevelError {
		l.logf(format, args...)
	}
}

// /////////////////////////////////////////////////
// New returns you a logger, ready to log at required threshold
// the default output is stdout
func New(threshold Level, output io.Writer) *Logger {
	return &Logger{
		threshold: threshold,
		output:    output,
	}
}

func (l *Logger) logf(format string, args ...any) {
	fmt.Fprintf(l.output, format+"\n", args...)
}
