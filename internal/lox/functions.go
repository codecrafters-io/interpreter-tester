package lox

import "io"

type loxCallable func([]interface{}) (interface{}, error)

// Callable is the generic interface for functions/classes in Lox
type Callable interface {
	Arity() int
	Call([]interface{}, *Environment, io.Writer, io.Writer) (interface{}, error)
}

// NativeFunction is a builtin Lox function
type NativeFunction struct {
	Callable
	nativeCall loxCallable
	arity      int
}

// Call is the operation that executes a builtin function
func (n *NativeFunction) Call(arguments []interface{}, globalEnv *Environment, stdout io.Writer, stderr io.Writer) (interface{}, error) {
	return n.nativeCall(arguments)
}

// Arity returns the number of allowed parameters for the native function
func (n *NativeFunction) Arity() int {
	return n.arity
}

// String returns the name of the native function
func (n *NativeFunction) String() string {
	return "<native fn>"
}

// UserFunction are functions defined in Lox code
type UserFunction struct {
	Callable
	Declaration *Function
}

// NewUserFunction creates a new UserFunction
func NewUserFunction(def *Function) *UserFunction {
	return &UserFunction{Declaration: def}
}

// Call executes a user-defined Lox function
func (u *UserFunction) Call(arguments []interface{}, globalEnv *Environment, stdout io.Writer, stderr io.Writer) (interface{}, error) {
	env := New(globalEnv)

	for i, param := range u.Declaration.Params {
		env.Define(param.Lexeme, arguments[i])
	}

	for _, stmt := range u.Declaration.Body {
		_, err := Eval(stmt, env, stdout, stderr)

		if err != nil {
			if r, ok := err.(ReturnError); ok {
				return r.value, nil
			}
			return nil, err
		}
	}

	return nil, nil
}

// Arity returns the number of arguments of the user-defined function
func (u *UserFunction) Arity() int {
	return len(u.Declaration.Params)
}

// String returns the name of the user-function
func (u *UserFunction) String() string {
	return u.Declaration.Name.Lexeme
}
