package lox

import (
	"time"
)

// GlobalEnv is the global environment
var GlobalEnv = NewGlobal()
var globals = GlobalEnv

func init() {
	GlobalEnv.Define("clock", &nativeFunction{
		Arity: func() int { return 0 },
		Call: func(args []interface{}) (interface{}, error) {
			return time.Now().Second(), nil
		},
	})
}

// ResetGlobalEnv resets the GlobalEnv to its original reference
func ResetGlobalEnv() {
	GlobalEnv = globals
}
