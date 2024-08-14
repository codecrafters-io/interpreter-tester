package lox

import (
	"fmt"
	"os"
)

func Evaluate(source string) (string, int, string) {
	scanner := NewScanner(source)
	tokens := scanner.ScanTokens(os.Stdout, os.Stderr)

	var existingErrors string
	exitCode := 0

	parser := NewParser(tokens)
	expression, err := parser.BasicParse()
	if err != nil {
		return "", 65, existingErrors + err.Error()
	}

	result, err := BasicInterpret(expression)
	if err != nil {
		return "", 70, existingErrors + err.Error()
	}

	if result == nil {
		return "nil", exitCode, existingErrors
	}
	return fmt.Sprintf("%v", result), exitCode, existingErrors
}
