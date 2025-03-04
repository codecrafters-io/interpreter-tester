package lox

import (
	"time"
)

// GlobalEnv is the global environment
var GlobalEnv = NewGlobal()
var globals = GlobalEnv

func InitializeNativeFunctions() {
	GlobalEnv.Define("clock", &NativeFunction{
		arity: 0,
		nativeCall: func(args []interface{}) (interface{}, error) {
			exponentNotation := float64(time.Now().Unix())

			return exponentNotation, nil
		},
	})
}

// ResetGlobalEnv resets the GlobalEnv to its original reference
func ResetGlobalEnv() {
	GlobalEnv = globals
}
