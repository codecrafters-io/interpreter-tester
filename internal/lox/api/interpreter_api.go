package loxapi

import (
	"bytes"

	"github.com/codecrafters-io/interpreter-tester/internal/lox"
)

func Evaluate(source string) (string, int, string) {
	lox.ClearErrorFlags()
	mockStdout := &bytes.Buffer{}
	mockStderr := &bytes.Buffer{}

	scanner := lox.NewScanner(source)
	tokens := scanner.ScanTokens(mockStdout, mockStderr)
	parser := lox.NewParser(tokens)
	expression := parser.BasicParse(mockStdout, mockStderr)
	lox.BasicInterpret(expression, mockStdout, mockStderr)

	exitCode := 0
	if lox.HadParseError {
		exitCode = 65
	} else if lox.HadRuntimeError {
		exitCode = 70
	}

	capturedStdout := mockStdout.String()
	if len(capturedStdout) > 0 {
		capturedStdout = capturedStdout[:len(capturedStdout)-1]
	}
	capturedStderr := mockStderr.String()
	if len(capturedStderr) > 0 {
		capturedStderr = capturedStderr[:len(capturedStderr)-1]
	}

	return capturedStdout, exitCode, capturedStderr
}
