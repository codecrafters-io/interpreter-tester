package lox

import (
	"fmt"
)

// Environment associates variables to values
type Environment struct {
	values    map[string]interface{}
	enclosing *Environment
}

// New creates a new environment
func New(env *Environment) *Environment {
	return &Environment{values: make(map[string]interface{}), enclosing: env}
}

// NewGlobal creates a new global environment
func NewGlobal() *Environment {
	return New(nil)
}

// NewEnvironment creates a new local environment
func NewEnvironment() *Environment {
	return New(nil)
}

// Define binds a name to a new value
func (e *Environment) Define(name string, value interface{}) {
	e.values[name] = value
}

// Get lookups a variable given a Token
func (e *Environment) Get(name Token) (interface{}, error) {
	v, prs := e.values[name.Lexeme]
	if prs {
		if v == nil {
			// We want to allow uninitialized variables
			return "nil", nil // XXX: Hack: returning nil as string here
		}
		return v, nil
	}
	if e.enclosing != nil {
		return e.enclosing.Get(name)
	}
	return nil, MakeRuntimeError(name, fmt.Sprintf("Undefined variable '%s'.", name.Lexeme))
}

// Assign sets a new value to an old variable
func (e *Environment) Assign(name Token, value interface{}) error {
	if _, prs := e.values[name.Lexeme]; prs {
		e.values[name.Lexeme] = value
		return nil
	}
	if e.enclosing != nil {
		return e.enclosing.Assign(name, value)
	}
	return MakeRuntimeError(name, fmt.Sprintf("Undefined variable '%s'.", name.Lexeme))
}
