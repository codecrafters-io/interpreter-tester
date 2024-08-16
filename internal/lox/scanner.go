package lox

import (
	"fmt"
	"io"
	"strconv"
)

var keywords = map[string]Type{
	"and":    AND,
	"class":  CLASS,
	"else":   ELSE,
	"false":  FALSE,
	"for":    FOR,
	"fun":    FUN,
	"if":     IF,
	"nil":    NIL,
	"or":     OR,
	"print":  PRINT,
	"return": RETURN,
	"super":  SUPER,
	"this":   THIS,
	"true":   TRUE,
	"var":    VAR,
	"while":  WHILE,
}

type Scanner struct {
	source  string
	start   int
	current int
	line    int
	tokens  []Token
}

func NewScanner(source string) Scanner {
	scanner := Scanner{source: source, line: 1, tokens: make([]Token, 0)}
	return scanner
}

// ScanTokens transforms the source into an array of tokens.
// The last token is always an EOF
func (sc *Scanner) ScanTokens(stdout, stderr io.Writer) []Token {
	for !sc.isAtEnd() {
		sc.scanToken(stdout, stderr)

		// We're at the beginning of the next lexeme
		sc.start = sc.current
	}

	sc.addToken(EOF)

	return sc.tokens
}

func (sc *Scanner) addToken(tp Type) {
	sc.addTokenWithLiteral(tp, nil)
}

func (sc *Scanner) addTokenWithLiteral(tp Type, literal interface{}) {
	text := sc.source[sc.start:sc.current]
	sc.tokens = append(sc.tokens, Token{Type: tp, Lexeme: text, Literal: literal, Line: sc.line})
}

func (sc *Scanner) scanString(stdout, stderr io.Writer) {
	for sc.peek() != '"' && !sc.isAtEnd() {
		if sc.peek() == '\n' {
			sc.line++
		}
		sc.advance()
	}

	// unterminated string
	if sc.isAtEnd() {
		LogParseError(fmt.Errorf("[line %d] Error: Unterminated string.", sc.line), stderr)
		return
	}

	// the closing ".
	sc.advance()

	// trim the surrounding quotes
	value := sc.source[sc.start+1 : sc.current-1]
	sc.addTokenWithLiteral(STRING, value)
}

func (sc *Scanner) scanNumber() {
	for sc.isDigit(sc.peek()) {
		sc.advance()
	}

	// look for a fractional part
	if sc.peek() == '.' && sc.isDigit(sc.peekNext()) {
		sc.advance() // consume "."
		for sc.isDigit(sc.peek()) {
			sc.advance()
		}
	}

	number, err := strconv.ParseFloat(sc.source[sc.start:sc.current], 64)
	if err != nil {
		panic("Invalid number format")
	} else {
		sc.addTokenWithLiteral(NUMBER, number)
	}
}

func (sc *Scanner) scanIdentifier() {
	for sc.isAlphaNumeric(sc.peek()) {
		sc.advance()
	}

	// see if the identifier is a reserved word
	text := sc.source[sc.start:sc.current]
	tp, ok := keywords[text]
	if ok {
		sc.addToken(tp)
	} else {
		sc.addToken(IDENTIFIER)
	}
}

func (sc *Scanner) scanToken(stdout, stderr io.Writer) {
	c := sc.advance()

	switch c {
	case '(':
		sc.addToken(LEFTPAREN)
	case ')':
		sc.addToken(RIGHTPAREN)
	case '{':
		sc.addToken(LEFTBRACE)
	case '}':
		sc.addToken(RIGHTBRACE)
	case ',':
		sc.addToken(COMMA)
	case '.':
		sc.addToken(DOT)
	case '-':
		sc.addToken(MINUS)
	case '+':
		sc.addToken(PLUS)
	case '?':
		sc.addToken(QMARK)
	case ':':
		sc.addToken(COLON)
	case ';':
		sc.addToken(SEMICOLON)
	case '*':
		sc.addToken(STAR)
	case '!':
		if sc.match('=') {
			sc.addToken(BANGEQUAL)
		} else {
			sc.addToken(BANG)
		}
	case '=':
		if sc.match('=') {
			sc.addToken(EQUALEQUAL)
		} else {
			sc.addToken(EQUAL)
		}
	case '<':
		if sc.match('=') {
			sc.addToken(LESSEQUAL)
		} else {
			sc.addToken(LESS)
		}
	case '>':
		if sc.match('=') {
			sc.addToken(GREATEREQUAL)
		} else {
			sc.addToken(GREATER)
		}
	case '/':
		if sc.match('/') {
			// A comment goes until the end of the line
			for sc.peek() != '\n' && !sc.isAtEnd() {
				sc.advance()
			}
		} else {
			sc.addToken(SLASH)
		}
	case '\n':
		sc.line++
	case ' ', '\r', '\t':
		// do nothing
	case '"':
		sc.scanString(stdout, stderr)
	default:
		if sc.isDigit(c) {
			sc.scanNumber()
		} else if sc.isAlpha(c) {
			sc.scanIdentifier()
		} else {
			LogParseError(fmt.Errorf("[line %d] Error: Unexpected character: %c", sc.line, c), stderr)
		}
	}
}

func (sc *Scanner) isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func (sc *Scanner) isAlpha(c byte) bool {
	return (c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z') ||
		c == '_'
}

func (sc *Scanner) isAlphaNumeric(c byte) bool {
	return sc.isAlpha(c) || sc.isDigit(c)
}

func (sc *Scanner) isAtEnd() bool {
	return sc.current >= len(sc.source)
}

// advance returns the current character and advances to the next
func (sc *Scanner) advance() byte {
	sc.current++
	return sc.source[sc.current-1]
}

func (sc *Scanner) match(expected byte) bool {
	if sc.isAtEnd() {
		return false
	}
	if sc.source[sc.current] != expected {
		return false
	}
	sc.current++
	return true
}

func (sc *Scanner) peek() byte {
	if sc.isAtEnd() {
		return 0
	}
	return sc.source[sc.current]
}

func (sc *Scanner) peekNext() byte {
	if sc.current+1 >= len(sc.source) {
		return 0
	}
	return sc.source[sc.current+1]
}
