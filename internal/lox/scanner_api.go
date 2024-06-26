package lox

import "fmt"

func ScanTokens(source string) ([]string, []string, error) {
	scanner := NewScanner(source)
	tokens, errors := scanner.ScanTokens()

	tokenLines := []string{}
	errorLines := errors

	for _, token := range tokens {
		literal := token.Literal
		if literal == nil {
			literal = "null"
		}
		tokenLines = append(tokenLines, fmt.Sprintf("%s %s %s", getTokenName(token.Type), token.Lexeme, literal))
	}

	return tokenLines, errorLines, nil
}
