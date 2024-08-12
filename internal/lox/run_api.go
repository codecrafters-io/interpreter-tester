package lox

import (
	"fmt"
	"strings"
)

func Run(source string) (string, int, string) {
	scanner := NewScanner(source)
	tokens, errors := scanner.ScanTokens()

	var existingErrors string
	exitCode := 0
	if len(errors) > 0 {
		for _, err := range errors {
			existingErrors += err + "\n"
		}
		exitCode = 65
	}

	parser := NewParser(tokens)
	statements, err := parser.Parse()
	if err != nil {
		return "", 65, existingErrors + err.Error()
	}

	results, err := Interpret(statements, NewGlobal())
	if err != nil {
		return "", 70, existingErrors + err.Error()
	}

	if results == nil {
		return "nil", exitCode, existingErrors
	}

	resultsString := []string{}
	for _, res := range results {
		resultsString = append(resultsString, fmt.Sprintf("%v", res))
	}
	return strings.Join(resultsString, "\n"), exitCode, existingErrors
}
