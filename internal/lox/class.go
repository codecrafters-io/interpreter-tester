package lox

import (
	"fmt"
)

// Class ...
type UserClass struct {
	Callable
	Name    string
	Methods map[string]*UserFunction
}

func (c *UserClass) String() string {
	return fmt.Sprintf("<class %s>", c.Name)
}

// Call is the operation that executes a class constructor
func (c *UserClass) Call(arguments []interface{}) (interface{}, error) {
	instance := &UserClassInstance{Class: c, fields: make(map[string]interface{})}
	if initializer, prs := c.Methods["init"]; prs {
		_, err := initializer.Bind(instance).Call(arguments, nil, nil, nil)
		if err != nil {
			return nil, err
		}
	}

	return instance, nil
}

// Arity returns the number of allowed parameters in the class constructor
// which is always 0
func (c *UserClass) Arity() int {
	if initializer, prs := c.Methods["init"]; prs {
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
	return fmt.Sprintf("<class-instance %s>", c.Class.Name)
}

// Get accesses the property
func (c *UserClassInstance) Get(name Token) (interface{}, error) {
	if v, prs := c.fields[name.Lexeme]; prs {
		return v, nil
	}

	if m, prs := c.Class.Methods[name.Lexeme]; prs {
		return m.Bind(c), nil
	}
	return nil, MakeRuntimeError(name, fmt.Sprintf("Undefined property '%s'", name.Lexeme))
}

// Set accesses the property
func (c *UserClassInstance) Set(name Token, value interface{}) (interface{}, error) {
	c.fields[name.Lexeme] = value
	return nil, nil
}
