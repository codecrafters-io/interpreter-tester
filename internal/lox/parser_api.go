package lox

import "os"

func Parse(source string) (string, int, string) {
	scanner := NewScanner(source)
	tokens := scanner.ScanTokens(os.Stdout, os.Stderr)

	var existingErrors string
	exitCode := 0
	// if len(errors) > 0 {
	// 	for _, err := range errors {
	// 		existingErrors += err + "\n"
	// 	}
	// 	exitCode = 65
	// }

	parser := NewParser(tokens)
	expression, err := parser.BasicParse()
	if err != nil {
		return "", 65, existingErrors + err.Error()
	}

	return expression.String(), exitCode, ""
}
