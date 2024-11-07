package lox

import (
	"time"
)

func InitializeNativeFunctions(env *Environment) {
	env.Define("clock", &nativeFunction{
		Arity: func() int { return 0 },
		Call: func(args []interface{}) (interface{}, error) {
			return time.Now().Second(), nil
		},
	})
}
