package lox

import (
	"bytes"
)

func Parse(source string) (string, int, string) {
	ClearErrorFlags()
	mockStdout := &bytes.Buffer{}
	mockStderr := &bytes.Buffer{}

	scanner := NewScanner(source)
	tokens := scanner.ScanTokens(mockStdout, mockStderr)
	parser := NewParser(tokens)
	expression := parser.BasicParse(mockStdout, mockStderr)
	capturedStderr := mockStderr.String()

	exitCode := 0
	if HadParseError {
		exitCode = 65
	} else if HadRuntimeError {
		exitCode = 70
	}

	if HadParseError {
		return "", exitCode, capturedStderr
	}

	return expression.String(), exitCode, capturedStderr
}
