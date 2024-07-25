package lox

import (
	"fmt"
	"os"
)

// PrintRuntimeError reports a runtime error
func PrintRuntimeError(message string) {
	fmt.Fprintf(os.Stderr, "%v\n", message)
	HadRuntimeError = true
}

// MakeRuntimeError creates a new runtime error
func MakeRuntimeError(token Token, message string) error {
	return fmt.Errorf("[line %d]: %s\n", token.Line, message)
}

// HadRuntimeError is true if an evaluation error was encountered
var HadRuntimeError = false
