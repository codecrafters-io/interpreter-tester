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
	resultsString := []string{}
	for _, res := range results {
		if res != nil {
			resultsString = append(resultsString, fmt.Sprintf("%v", res))
		}
	}
	output := strings.Join(resultsString, "\n")

	if err != nil {
		return output, 70, existingErrors + err.Error()
	}

	return output, exitCode, existingErrors
}
