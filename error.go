package loop

import (
	"reflect"
)

var errorType = reflect.TypeOf(make([]error, 1)).Elem()
