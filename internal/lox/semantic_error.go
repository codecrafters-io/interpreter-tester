package lox

import (
	"fmt"
	"os"
)

// PrintSemanticError reports a semantic error
func PrintSemanticError(message string) {
	fmt.Fprintf(os.Stderr, "%v\n", message)
	HadSemanticError = true
}

// MakeSemanticError creates a new semantic error
func MakeSemanticError(message string) error {
	HadSemanticError = true
	return fmt.Errorf("%s", message)
}

// HadSemanticError is true if an evaluation error was encountered
var HadSemanticError = false
