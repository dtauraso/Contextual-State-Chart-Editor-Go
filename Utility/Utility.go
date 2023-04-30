package Utility

import (
	"reflect"
	"runtime"
	"strings"
	// csc "github.com/dtauraso/Contextual-State-Chart-Editor-Go/ContextualStateChart"
)

// func ReturnTrue(test csc.Graph) bool {return true}

func GetFunctionName(temp interface{}) string {
	strs := strings.Split((runtime.FuncForPC(reflect.ValueOf(temp).Pointer()).Name()), ".")
	return strs[len(strs)-1]
}
func GetType(myvar interface{}) string {
	if t := reflect.TypeOf(myvar); t.Kind() == reflect.Ptr {
		return "*" + t.Elem().Name()
	} else {
		return t.Name()
	}
}
