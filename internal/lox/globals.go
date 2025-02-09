package lox

import (
	"time"
)

func InitializeNativeFunctions(env *Environment) {
	env.Define("clock", &NativeFunction{
		arity: 0,
		nativeCall: func(args []interface{}) (interface{}, error) {
			exponentNotation := float64(time.Now().Unix())

			return exponentNotation, nil
		},
	})
}
