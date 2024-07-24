package lox

import "fmt"

func Evaluate(source string) (string, int, string) {
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
	expression, err := parser.Parse()
	if err != nil {
		return "", 65, existingErrors + err.Error()
	}

	result, err := Interpret(expression)
	if err != nil {
		return "", 70, existingErrors + err.Error()
	}

	if result == nil {
		return "nil", exitCode, existingErrors
	}
	return fmt.Sprintf("%v", result), exitCode, existingErrors
}
