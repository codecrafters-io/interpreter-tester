package lox

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

func ScanTokens(source string) ([]string, []string, int, error) {
	scanner := NewScanner(source)
	tokens := scanner.ScanTokens(os.Stdout, os.Stderr)
	exitCode := 0

	var tokenLines []string
	// errorLines := errors

	for _, token := range tokens {
		literal := token.Literal
		if literal == nil {
			literal = "null"
		} else if reflect.TypeOf(literal).Kind() == reflect.Float64 {
			literal = formatFloat(literal.(float64))
		}
		tokenLines = append(tokenLines, fmt.Sprintf("%s %s %s", getTokenName(token.Type), token.Lexeme, literal))
	}

	// if len(errorLines) > 0 {
	// 	exitCode = 65
	// }

	return tokenLines, nil, exitCode, nil
}

func formatFloat(num float64) string {
	str := fmt.Sprintf("%f", num)
	parts := strings.Split(str, ".")
	if len(parts) != 2 {
		return str
	}
	integerPart := parts[0]
	decimalPart := parts[1]

	// Remove trailing zeros from the decimal part
	decimalPart = strings.TrimRight(decimalPart, "0")

	// Ensure at least one decimal place
	if decimalPart == "" {
		decimalPart = "0"
	}

	return integerPart + "." + decimalPart
}
