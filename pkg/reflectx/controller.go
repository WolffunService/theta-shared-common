package reflectx

import (
	"reflect"
	"runtime"
	"strings"
)

// FnName get func name, but only used for iris controller
func FnName(fn any) string {
	path := runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
	splitter := strings.Split(path, ".")
	result := splitter[len(splitter)-1]
	return result[:len(result)-3] // remove -fm
}
