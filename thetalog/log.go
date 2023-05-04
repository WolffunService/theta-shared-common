package thetalog

import (
	ctx "context"
	"strconv"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	fullPath := func(file string, line int) string {
		return file + ":" + strconv.Itoa(line)
	}

	zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
		// Inspired from uber/zap TrimmedPath =))

		idx := strings.LastIndexByte(file, '/')
		if idx == -1 {
			return fullPath(file, line)
		}
		// Find the penultimate separator.
		idx = strings.LastIndexByte(file[:idx], '/')
		if idx == -1 {
			return fullPath(file, line)
		}

		buf := strings.Builder{}
		buf.WriteString(file[idx+1:])
		buf.WriteRune(':')
		buf.WriteString(strconv.Itoa(line))
		caller := buf.String()
		buf.Reset()

		return caller
	}
}

// Level defines log levels.
type Level int8

type logger struct {
	*zerolog.Logger
}

// Logger is a logger that supports log levels, context and structured logging.
type Logger interface {
	Err(err error) Event

	Trace() Event

	Debug() Event

	Info() Event

	Warn() Event

	Error() Event

	Fatal() Event

	Panic() Event

	WithLevel(level Level) Event

	Log() Event

	Print(v ...interface{})

	Printf(format string, v ...interface{})
}

func SetGlobalLevel(l Level) {
	zerolog.SetGlobalLevel(zerolog.Level(l))
}

func NewLogger() Logger {
	l := log.With().Caller().Logger()
	return &logger{&l}
}

func NewBizLogger(biz string) Logger {
	l := log.With().Caller().Str("business", biz).Logger()
	return &logger{&l}
}

func (l logger) Err(err error) Event {
	return &event{l.Logger.Err(err)}
}

func (l logger) Trace() Event {
	return &event{l.Logger.Trace()}
}

func (l logger) Debug() Event {
	return &event{l.Logger.Debug()}
}

func (l logger) Info() Event {
	return &event{l.Logger.Info()}
}

func (l logger) Warn() Event {
	return &event{l.Logger.Warn()}
}

func (l logger) Error() Event {
	return &event{l.Logger.Error()}
}

func (l logger) Fatal() Event {
	return &event{l.Logger.Fatal()}
}

func (l logger) Panic() Event {
	return &event{l.Logger.Panic()}
}

func (l logger) WithLevel(level Level) Event {
	return &event{l.Logger.WithLevel(zerolog.Level(level))}
}

func (l logger) Log() Event {
	return &event{l.Logger.Log()}
}

func (l logger) Ctx(ctx ctx.Context) Logger {
	l.Logger = zerolog.Ctx(ctx)
	return l
}

var (
	L = NewLogger()
)

// Err starts a new message with error level with err as a field if not nil or
// with info level if err is nil.
//
// You must call Msg on the returned event in order to send the event.
func Err(err error) Event {
	return L.Err(err)
}

// Trace starts a new message with trace level.
//
// You must call Msg on the returned event in order to send the event.
func Trace() Event {
	return L.Trace()
}

// Debug starts a new message with debug level.
//
// You must call Msg on the returned event in order to send the event.
func Debug() Event {
	return L.Debug()
}

// Info starts a new message with info level.
//
// You must call Msg on the returned event in order to send the event.
func Info() Event {
	return L.Info()
}

// Warn starts a new message with warn level.
//
// You must call Msg on the returned event in order to send the event.
func Warn() Event {
	return L.Warn()
}

// Error starts a new message with error level.
//
// You must call Msg on the returned event in order to send the event.
func Error() Event {
	return L.Error()
}

// Fatal starts a new message with fatal level. The os.Exit(1) function
// is called by the Msg method.
//
// You must call Msg on the returned event in order to send the event.
func Fatal() Event {
	return L.Fatal()
}

// Panic starts a new message with panic level. The message is also sent
// to the panic function.
//
// You must call Msg on the returned event in order to send the event.
func Panic() Event {
	return L.Panic()
}

// WithLevel starts a new message with level.
//
// You must call Msg on the returned event in order to send the event.
func WithLevel(level Level) Event {
	return L.WithLevel(level)
}

// Log starts a new message with no level. Setting zerolog.GlobalLevel to
// zerolog.Disabled will still disable events produced by this method.
//
// You must call Msg on the returned event in order to send the event.
func Log() Event {
	return L.Log()
}

var (
	// LevelTraceValue is the value used for the trace level field.
	LevelTraceValue = "trace"
	// LevelDebugValue is the value used for the debug level field.
	LevelDebugValue = "debug"
	// LevelInfoValue is the value used for the info level field.
	LevelInfoValue = "info"
	// LevelWarnValue is the value used for the warn level field.
	LevelWarnValue = "warn"
	// LevelErrorValue is the value used for the error level field.
	LevelErrorValue = "error"
	// LevelFatalValue is the value used for the fatal level field.
	LevelFatalValue = "fatal"
	// LevelPanicValue is the value used for the panic level field.
	LevelPanicValue = "panic"

	// LevelFieldMarshalFunc allows customization of global level field marshaling.
	LevelFieldMarshalFunc = func(l Level) string {
		return l.String()
	}
)
