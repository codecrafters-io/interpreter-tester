package loxapi

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"

	"github.com/codecrafters-io/interpreter-tester/internal/lox"
)

func ScanTokens(source string) ([]string, []string, int, error) {
	lox.ClearErrorFlags()
	mockStdout := &bytes.Buffer{}
	mockStderr := &bytes.Buffer{}

	scanner := lox.NewScanner(source)
	tokens := scanner.ScanTokens(mockStdout, mockStderr)

	var tokenLines []string

	for _, token := range tokens {
		literal := token.Literal
		if literal == nil {
			literal = "null"
		} else if reflect.TypeOf(literal).Kind() == reflect.Float64 {
			literal = lox.FormatFloat(literal.(float64))
		}
		tokenLines = append(tokenLines, fmt.Sprintf("%s %s %s", lox.GetTokenName(token.Type), token.Lexeme, literal))
	}

	exitCode := 0
	if lox.HadParseError {
		exitCode = 65
	} else if lox.HadRuntimeError {
		exitCode = 70
	}

	stderr := mockStderr.String()

	var stderrLines []string

	if len(stderr) > 0 {
		stderr = stderr[:len(stderr)-1]
		stderrLines = strings.Split(stderr, "\n")
	} else {
		stderrLines = []string{}
	}
	return tokenLines, stderrLines, exitCode, nil
}
