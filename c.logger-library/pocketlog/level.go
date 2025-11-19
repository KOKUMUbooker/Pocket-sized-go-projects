package pocketlog

// Level represents an available logging level
type Level byte;

const (
	// LevelDebug represents the lowest level of log, mostly used for debugging
	LevelDebug Level = iota
	// LevelInfo represents a logging level that contains information deemed meaningful
	LevelInfo
	// LevelError represents the highest logging level, only to be used to track errors/bugs
	LevelError
)