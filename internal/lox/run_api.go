package lox

import (
	"bytes"
)

func Run(source string) (string, int, string) {
	ClearErrorFlags()
	mockStdout := &bytes.Buffer{}
	mockStderr := &bytes.Buffer{}

	scanner := NewScanner(source)
	tokens := scanner.ScanTokens(mockStdout, mockStderr)
	parser := NewParser(tokens)
	statements := parser.Parse(mockStdout, mockStderr)
	Interpret(statements, mockStdout, mockStderr)

	exitCode := 0
	if HadParseError {
		exitCode = 65
	} else if HadRuntimeError {
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
