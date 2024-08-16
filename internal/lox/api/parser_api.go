package loxapi

import (
	"bytes"

	"github.com/codecrafters-io/interpreter-tester/internal/lox"
)

func Parse(source string) (string, int, string) {
	lox.ClearErrorFlags()
	mockStdout := &bytes.Buffer{}
	mockStderr := &bytes.Buffer{}

	scanner := lox.NewScanner(source)
	tokens := scanner.ScanTokens(mockStdout, mockStderr)
	parser := lox.NewParser(tokens)
	expression := parser.BasicParse(mockStdout, mockStderr)
	capturedStderr := mockStderr.String()

	exitCode := 0
	if lox.HadParseError {
		exitCode = 65
	} else if lox.HadRuntimeError {
		exitCode = 70
	}

	if lox.HadParseError {
		return "", exitCode, capturedStderr
	}

	return expression.String(), exitCode, capturedStderr
}
