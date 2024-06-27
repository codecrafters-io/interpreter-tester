package lox

import "fmt"

func ScanTokens(source string) ([]string, []string, int, error) {
	scanner := NewScanner(source)
	tokens, errors := scanner.ScanTokens()
	exitCode := 0

	var tokenLines []string
	errorLines := errors

	for _, token := range tokens {
		literal := token.Literal
		if literal == nil {
			literal = "null"
		}
		tokenLines = append(tokenLines, fmt.Sprintf("%s %s %s", getTokenName(token.Type), token.Lexeme, literal))
	}

	if len(errorLines) > 0 {
		exitCode = 65
	}

	return tokenLines, errorLines, exitCode, nil
}
