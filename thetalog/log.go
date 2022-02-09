package thetalog

import (
	ctx "context"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Level defines log levels.
type Level int8

type logger struct {
	*zerolog.Logger
}

// Logger is a logger that supports log levels, context and structured logging.
type Logger interface {
	//With() Context

	//Output(w io.Writer) Logger

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

	//Ctx(ctx ctx.Context) Logger
}

func SetGlobalLevel(l Level) {
	zerolog.SetGlobalLevel(zerolog.Level(l))
}

func NewLogger() Logger {
	l := &log.Logger
	return &logger{l}
}
func NewBizLogger(biz string) Logger {
	zeroLogger := log.Logger.With().Str("business", biz).Logger()
	l := &zeroLogger
	return &logger{l}
}

//func (l logger) With() Context {
//	if l.Logger == nil {
//		l.Logger = &log.Logger
//	}
//	return context{l.Logger.With()}
//}

//func (l logger) Output(w io.Writer) Logger {
//	return l.Output(w)
//}

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

//// Output duplicates the global logger and sets w as its output.
//func Output(w io.Writer) zerolog.Logger {
//	return L.Output(w)
//}
//
// With creates a child logger with the field added to its context.
//func With() Context {
//	return &logger{L.With()}
//}

//
//// Err starts a new message with error level with err as a field if not nil or
//// with info level if err is nil.
////
//// You must call Msg on the returned event in order to send the event.
//func Err(err error) *zerolog.Event {
//	return L.Err(err)
//}
//
//// Trace starts a new message with trace level.
////
//// You must call Msg on the returned event in order to send the event.
//func Trace() *zerolog.Event {
//	return L.Trace()
//}
//
//// Debug starts a new message with debug level.
////
//// You must call Msg on the returned event in order to send the event.
//func Debug() *zerolog.Event {
//	return L.Debug()
//}
//
//// Info starts a new message with info level.
////
//// You must call Msg on the returned event in order to send the event.
//func Info() *zerolog.Event {
//	return L.Info()
//}
//
//// Warn starts a new message with warn level.
////
//// You must call Msg on the returned event in order to send the event.
//func (l logger) Warn() Event {
//	return &l{L.Warn()}
//}
//
//// Error starts a new message with error level.
////
//// You must call Msg on the returned event in order to send the event.
//func Error() *zerolog.Event {
//	return L.Error()
//}
//
//// Fatal starts a new message with fatal level. The os.Exit(1) function
//// is called by the Msg method.
////
//// You must call Msg on the returned event in order to send the event.
//func Fatal() *zerolog.Event {
//	return L.Fatal()
//}
//
// Panic starts a new message with panic level. The message is also sent
// to the panic function.
//
// You must call Msg on the returned event in order to send the event.
//func Panic() Event {
//	return L.Panic()
//}

//
//// WithLevel starts a new message with level.
////
//// You must call Msg on the returned event in order to send the event.
//func WithLevel(level zerolog.Level) *zerolog.Event {
//	return L.WithLevel(level)
//}
//
//// Log starts a new message with no level. Setting zerolog.GlobalLevel to
//// zerolog.Disabled will still disable events produced by this method.
////
//// You must call Msg on the returned event in order to send the event.
//func Log() *zerolog.Event {
//	return L.Log()
//}
//
//// Print sends a log event using debug level and no extra field.
//// Arguments are handled in the manner of fmt.Print.
//func Print(v ...interface{}) {
//	L.Debug().CallerSkipFrame(1).Msg(fmt.Sprint(v...))
//}
//
//// Printf sends a log event using debug level and no extra field.
//// Arguments are handled in the manner of fmt.Printf.
//func (l logger) Printf(format string, v ...interface{}) {
//	return NewLogger().Debug().CallerSkipFrame(1).Msgf(format, v...)
//}

// Ctx returns the Logger associated with the ctx. If no logger
// is associated, a disabled logger is returned.
//func (l logger) Ctx(ctx context.Context) Logger {
//	return &logger{zerolog.Ctx(ctx)}
//}
