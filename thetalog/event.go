package thetalog

import "github.com/rs/zerolog"

type event struct {
	*zerolog.Event
}

type Event interface {
	Msg(msg string)
	Msgf(format string, v ...interface{})
	Send()
	Str(key, val string) Event
	Int(key string, i int) Event
	Op(val string) Event
	Bool(key string, b bool) Event
	Var(key string, i interface{}) Event
}

func (e *event) Str(key, val string) Event {
	e.Event.Str(key, val)
	return e
}

func (e *event) Int(key string, i int) Event {
	e.Event.Int(key, i)
	return e
}

func (e *event) Bool(key string, b bool) Event {
	e.Event.Bool(key, b)
	return e
}

func (e *event) Op(val string) Event {
	return e.Str("operation", val)
}

func (e *event) Var(key string, i interface{}) Event {
	e.Event.Interface(key, i)
	return e
}
