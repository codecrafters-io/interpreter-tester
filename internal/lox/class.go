package lox

import (
	"fmt"
	"io"
)

// Class is a user defined class
type UserClass struct {
	Callable
	Name    string
	Methods map[string]*UserFunction
}

func (c *UserClass) String() string {
	return c.Name
}

// Call is the operation that executes a class constructor
func (c *UserClass) Call(arguments []any, globalEnv *Environment, stdout io.Writer, stderr io.Writer) (any, error) {
	instance := &UserClassInstance{Class: c, fields: make(map[string]any)}
	if initializer, ok := c.Methods["init"]; ok {
		_, err := initializer.Bind(instance).Call(arguments, globalEnv, stdout, stderr)
		if err != nil {
			return nil, err
		}
	}

	return instance, nil
}

// Arity returns the number of allowed parameters in the class constructor
// which is always 0
func (c *UserClass) Arity() int {
	if initializer, ok := c.Methods["init"]; ok {
		return initializer.Arity()
	}
	return 0
}

// UserClassInstance is a user defined class instance
type UserClassInstance struct {
	Class  *UserClass
	fields map[string]any
}

func (c *UserClassInstance) String() string {
	return fmt.Sprintf("%s instance", c.Class.Name)
}

// Get accesses the property
func (c *UserClassInstance) Get(name Token) (any, error) {
	if v, ok := c.fields[name.Lexeme]; ok {
		return v, nil
	}

	if m, ok := c.Class.Methods[name.Lexeme]; ok {
		return m.Bind(c), nil
	}
	return nil, MakeRuntimeError(name, fmt.Sprintf("Undefined property '%s'", name.Lexeme))
}

// Set accesses the property
func (c *UserClassInstance) Set(name Token, value any) (any, error) {
	c.fields[name.Lexeme] = value
	return nil, nil
}
