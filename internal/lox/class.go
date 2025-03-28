package lox

import (
	"fmt"
	"io"
)

// Class is a user defined class
type UserClass struct {
	Callable
	Name       string
	Methods    map[string]*UserFunction
	SuperClass *UserClass
}

func (c *UserClass) String() string {
	return c.Name
}

func (c *UserClass) FindMethod(name Token) (*UserFunction, error) {
	if m, prs := c.Methods[name.Lexeme]; prs {
		return m, nil
	} else if c.SuperClass != nil {
		return c.SuperClass.FindMethod(name)
	}
	return nil, MakeRuntimeError(name, fmt.Sprintf("Undefined property '%s'", name.Lexeme))
}

// Call is the operation that executes a class constructor
func (c *UserClass) Call(arguments []interface{}, globalEnv *Environment, stdout io.Writer, stderr io.Writer) (interface{}, error) {
	instance := &UserClassInstance{Class: c, fields: make(map[string]interface{})}

	// Look for initializer in this class or any superclass
	initializer, _ := c.FindMethod(Token{Lexeme: "init"})

	if initializer != nil {
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
	fields map[string]interface{}
}

func (c *UserClassInstance) String() string {
	return fmt.Sprintf("%s instance", c.Class.Name)
}

// Get accesses the property
func (c *UserClassInstance) Get(name Token) (interface{}, error) {
	// First check if the property exists in this instance
	if v, ok := c.fields[name.Lexeme]; ok {
		return v, nil
	}

	// If not found, look for a method
	method, err := c.Class.FindMethod(name)
	if err != nil {
		return nil, err
	}

	if method != nil {
		return method.Bind(c), nil
	}

	// If we get here, the property doesn't exist
	return nil, MakeRuntimeError(name, fmt.Sprintf("Undefined property '%s'", name.Lexeme))
}

// Set accesses the property
func (c *UserClassInstance) Set(name Token, value interface{}) (interface{}, error) {
	c.fields[name.Lexeme] = value
	return value, nil
}
