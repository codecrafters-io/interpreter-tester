package loxapi

import (
	"bytes"

	"github.com/codecrafters-io/interpreter-tester/internal/lox"
)

func Run(source string) (string, int, string) {
	lox.ClearErrorFlags()
	mockStdout := &bytes.Buffer{}
	mockStderr := &bytes.Buffer{}

	scanner := lox.NewScanner(source)
	tokens := scanner.ScanTokens(mockStdout, mockStderr)
	parser := lox.NewParser(tokens)
	statements := parser.Parse(mockStdout, mockStderr)
	locals, err := lox.Resolve(statements)
	if err != nil && lox.HadSemanticError {
		return "", 65, err.Error()
	}

	env := lox.NewGlobal()
	lox.Interpret(statements, env, locals, mockStdout, mockStderr)

	exitCode := 0
	if lox.HadParseError || lox.HadSemanticError {
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
