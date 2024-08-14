package lox

import (
	"fmt"
	"os"
	"strings"
)

// PrintRuntimeError reports a runtime error
func LogRuntimeError(err error) {
	fmt.Fprintf(os.Stderr, "%v\n", strings.TrimSpace(err.Error()))
	HadRuntimeError = true
}

// MakeRuntimeError creates a new runtime error
func MakeRuntimeError(token Token, message string) error {
	return fmt.Errorf("[line %d]: %s\n", token.Line, message)
}

// HadRuntimeError is true if an evaluation error was encountered
var HadRuntimeError = false

func ReportRuntimeError(token Token, message string) {
	LogRuntimeError(MakeRuntimeError(token, message))
}
