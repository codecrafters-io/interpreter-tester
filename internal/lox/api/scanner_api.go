package loxapi

import (
	"fmt"
	"os"
	"reflect"

	"github.com/codecrafters-io/interpreter-tester/internal/lox"
)

func ScanTokens(source string) ([]string, []string, int, error) {
	lox.ClearErrorFlags()
	scanner := lox.NewScanner(source)
	tokens := scanner.ScanTokens(os.Stdout, os.Stderr)

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

	return tokenLines, nil, exitCode, nil
}
