package salt

type Level int8

const (
	TraceLevel Level = iota
	DebugLevel
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
	PanicLevel
)

func (l Level) String() string {
	all := []string{
		"TRACE",
		"DEBUG",
		"INFO",
		"WARN",
		"ERROR",
		"FATAL",
		"PANIC",
	}
	return all[l]
}

// Options specifies output
type Options struct {
	// WithoutLevels if false then all logs will be printed as a row message. Example:
	//
	// WithoutLevels = false
	//
	// INFO [2022-09-04T21:54:27.162-MSK]: user creation
	//
	// WithoutLevels = true
	//
	// user creation
	WithoutLevels bool

	// TimeUnix specify the way time would be printed.
	//
	// If false, then will print in standard format
	TimeUnix bool
}
