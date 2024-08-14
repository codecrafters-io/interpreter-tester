package lox

import (
	"fmt"
	"io"
)

// HadParseError is true if a scanner/parser error was encountered
var HadParseError = false

// LogParseError reports in stderr an error encountered during parsing
func LogParseError(err error, stderr io.Writer) {
	fmt.Fprintf(stderr, "%v\n", err.Error())
	HadParseError = true
}

// MakeParseError renders a parsing error as a string
func MakeParseError(tok Token, message string) error {
	if tok.Type == EOF {
		return fmt.Errorf("[line %v] Error at end: %s", tok.Line, message)
	}
	return fmt.Errorf("[line %v] Error at '%s': %s", tok.Line, tok.Lexeme, message)
}

func ReportParseError(line int, where string, message string, stderr io.Writer) {
	LogParseError(MakeParseError(Token{Line: line, Lexeme: where}, message), stderr)
}
