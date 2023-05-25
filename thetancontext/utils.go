package thetancontext

import (
	"reflect"
	"runtime"
	"unsafe"
)

func getUnexportedField(field reflect.Value) interface{} {
	return reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).Elem().Interface()
}

// unsafe method - skip rule mustBeExported with unsafe address(UnsafeAddr)
func setUnexportedField(field reflect.Value, value interface{}) {
	reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).
		Elem().
		Set(reflect.ValueOf(value))
}

func Bind(data any, mapbind map[reflect.Type]any) {
	v := reflect.Indirect(reflect.ValueOf(data))
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if vlue, exist := mapbind[field.Type()]; exist {
			setUnexportedField(field, vlue)
		}
	}
}

// recover function
func recoverFn() {
	if err := recover(); err != nil {
		//log.Println("recovered from ", errors.Wrap(err, 2).ErrorStack())
	}
}

// trace log
func trace() (string, int, string) {
	pc, file, line, ok := runtime.Caller(2)
	if !ok {
		return "?", 0, "?"
	}

	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return file, line, "?"
	}

	return file, line, fn.Name()
}
