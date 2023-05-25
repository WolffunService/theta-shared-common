package thetancontext

import (
	"fmt"
	"reflect"
)

type (
	//provide and bind data in a flexible way, allowing for a wide variety of use cases.
	BindData interface {
		Provider(v any)
		BindingUserData(v any)
	}

	Binder struct {
		bindData map[reflect.Type]any
	}
)

func NewBinder() Binder {
	return Binder{
		map[reflect.Type]any{},
	}
}

func (b *Binder) Provider(v any) {
	vType := reflect.TypeOf(v)
	if vType.Kind() != reflect.Ptr {
		fmt.Printf("%+v must is pointer \n", v)
		return
	}
	if _, exist := b.bindData[vType]; exist {
		fmt.Printf("%+v already exist type \n", v)
		return
	}
	b.bindData[vType] = v
}

func (b *Binder) BindingUserData(data any) {
	//utils
	Bind(data, b.bindData)
}
