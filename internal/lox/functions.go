package lox

import (
	"io"
)

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
	Declaration   *Function
	Closure       *Environment
	Locals        Locals // TODO: Pass pointer to Locals
	IsInitializer bool
}

// NewUserFunction creates a new UserFunction
func NewUserFunction(declaration *Function, closure *Environment, locals Locals) *UserFunction {
	return &UserFunction{
		Declaration:   declaration,
		Closure:       closure,
		Locals:        locals,
		IsInitializer: false,
	}
}

// Call executes a user-defined Lox function
func (u *UserFunction) Call(arguments []interface{}, globalEnv *Environment, stdout io.Writer, stderr io.Writer) (interface{}, error) {
	// ToDo: Will the closure always encapsulate all global vars ?
	env := New(u.Closure)

	for i, param := range u.Declaration.Params {
		env.Define(param.Lexeme, arguments[i])
	}

	for _, stmt := range u.Declaration.Body {
		_, err := Eval(stmt, env, u.Locals, stdout, stderr)

		if err != nil {
			if r, ok := err.(ReturnError); ok {
				if u.IsInitializer {
					return u.Closure.GetAt(0, Token{Lexeme: "this"})
				}
				return r.value, nil
			}
			return nil, err
		}
	}

	if u.IsInitializer {
		return u.Closure.GetAt(0, Token{Lexeme: "this"})
	}
	return nil, nil
}

// Arity returns the number of arguments of the user-defined function
func (u *UserFunction) Arity() int {
	return len(u.Declaration.Params)
}

// String returns the name of the user-function
func (u *UserFunction) String() string {
	return "<fn " + u.Declaration.Name.Lexeme + ">"
}

// Bind creates a new instance method
func (u *UserFunction) Bind(instance *UserClassInstance) *UserFunction {
	thisEnv := New(u.Closure)
	thisEnv.Define("this", instance)
	return &UserFunction{
		Declaration:   u.Declaration,
		Closure:       thisEnv,
		Locals:        u.Locals,
		IsInitializer: u.IsInitializer,
	}
}
