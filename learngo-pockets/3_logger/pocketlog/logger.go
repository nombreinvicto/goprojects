package pocketlog

type Logger struct {
	threshold Level
}

// Debugf formats and prints a message if log level is debug or higher
func (l *Logger) Debugf(format string, args ...any) {
	//implement me
}

// Infof formats and prints a message if the log level is info or higher
func (l *Logger) Infof(format string, args ...any) {
	// implement me
}

// Errorf formats and prints a message if the log level is error or higher
func (l *Logger) Errorf(format string, args ...any) {
	// implement me
}

// /////////////////////////////////////////////////
// New returns you a logger, ready tolog at required threshold
func New(threshold Level) *Logger {
	return &Logger{
		threshold: threshold,
	}
}
