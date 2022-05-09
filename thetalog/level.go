package thetalog

import (
	"fmt"
	"strconv"
)

const (
	// DebugLevel defines debug log level.
	DebugLevel Level = iota
	// InfoLevel defines info log level.
	InfoLevel
	// WarnLevel defines warn log level.
	WarnLevel
	// ErrorLevel defines error log level.
	ErrorLevel
	// FatalLevel defines fatal log level.
	FatalLevel
	// PanicLevel defines panic log level.
	PanicLevel
	// NoLevel defines an absent log level.
	NoLevel
	// Disabled disables the logger.
	Disabled

	// TraceLevel defines trace log level.
	TraceLevel Level = -1
	// Values less than TraceLevel are handled as numbers.
)

func (l Level) String() string {
	switch l {
	case TraceLevel:
		return LevelTraceValue
	case DebugLevel:
		return LevelDebugValue
	case InfoLevel:
		return LevelInfoValue
	case WarnLevel:
		return LevelWarnValue
	case ErrorLevel:
		return LevelErrorValue
	case FatalLevel:
		return LevelFatalValue
	case PanicLevel:
		return LevelPanicValue
	case Disabled:
		return "disabled"
	case NoLevel:
		return ""
	}
	return strconv.Itoa(int(l))
}

// ParseLevel converts a level string into a zerolog Level value.
// returns an error if the input string does not match known values.
func ParseLevel(levelStr string) (Level, error) {
	switch levelStr {
	case LevelFieldMarshalFunc(TraceLevel):
		return TraceLevel, nil
	case LevelFieldMarshalFunc(DebugLevel):
		return DebugLevel, nil
	case LevelFieldMarshalFunc(InfoLevel):
		return InfoLevel, nil
	case LevelFieldMarshalFunc(WarnLevel):
		return WarnLevel, nil
	case LevelFieldMarshalFunc(ErrorLevel):
		return ErrorLevel, nil
	case LevelFieldMarshalFunc(FatalLevel):
		return FatalLevel, nil
	case LevelFieldMarshalFunc(PanicLevel):
		return PanicLevel, nil
	case LevelFieldMarshalFunc(Disabled):
		return Disabled, nil
	case LevelFieldMarshalFunc(NoLevel):
		return NoLevel, nil
	}
	i, err := strconv.Atoi(levelStr)
	if err != nil {
		return NoLevel, fmt.Errorf("Unknown Level String: '%s', defaulting to NoLevel", levelStr)
	}
	if i > 127 || i < -128 {
		return NoLevel, fmt.Errorf("Out-Of-Bounds Level: '%d', defaulting to NoLevel", i)
	}
	return Level(i), nil
}
