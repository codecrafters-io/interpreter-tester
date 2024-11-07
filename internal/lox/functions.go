package lox

type loxCallable func([]interface{}) (interface{}, error)

// Callable is the generic interface for functions/classes in Lox
type Callable interface {
	Arity() int
	Call([]interface{}) (interface{}, error)
}

// NativeFunction is a builtin Lox function
type NativeFunction struct {
	Callable
	nativeCall loxCallable
	arity      int
}

// Call is the operation that executes a builtin function
func (n *NativeFunction) Call(arguments []interface{}) (interface{}, error) {
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
