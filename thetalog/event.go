package thetalog

import "github.com/rs/zerolog"

type event struct {
	*zerolog.Event
}

type Event interface {
	Msg(msg string)
	Send()
	Msgf(format string, v ...interface{})
	Str(key, val string) Event
	Int(key string, i int) Event
	Op(val string) Event
}

func (e *event) Str(key, val string) Event {
	e.Event.Str(key, val)
	return e
}

func (e *event) Int(key string, i int) Event {
	e.Event.Int(key, i)
	return e
}

func (e *event) Op(val string) Event {
	return e.Str("operation", val)
}
