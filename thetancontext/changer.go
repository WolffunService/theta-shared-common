package thetancontext

import (
	"github.com/func25/mongofunc/moper"
)

type (
	ChangedDataKey struct{}

	ChangedData interface {
		GetChanged() moper.D

		//set change data
		Set(key string, value any)
		//unset change data
		Unset(key string)
		//increase data
		Inc(key string, value any)
	}

	Changer struct {
		changed moper.D
	}
)

func NewChanger() ChangedData {
	return &Changer{}
}

func (c *Changer) GetChanged() moper.D {
	return c.changed
}

func (c *Changer) Set(key string, value any) {
	c.changed.Set(moper.P{key, value})
}

func (c *Changer) Unset(key string) {
	c.changed.Unset(key)
}

func (c *Changer) Inc(key string, value any) {
	c.changed.Inc(moper.P{key, value})
}
