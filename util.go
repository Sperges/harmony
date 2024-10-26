package harmony

import (
	"reflect"
	"strings"
)

func structNameAsJsonString[T any](t T) string {
	name := reflect.TypeOf(t).Name()
	return strings.ToLower(string(name[0])) + name[1:]
}
