package lox

import "fmt"

// Type is the type of the token given as a string
type Type string

const (
	// single-character tokens

	LEFTPAREN  = "("
	RIGHTPAREN = ")"
	LEFTBRACE  = "{"
	RIGHTBRACE = "}"
	COMMA      = ","
	DOT        = "."
	MINUS      = "-"
	PLUS       = "+"
	SEMICOLON  = ";"
	SLASH      = "/"
	STAR       = "*"
	QMARK      = "?"
	COLON      = ":"

	// one or two character tokens

	BANG         = "!"
	BANGEQUAL    = "!="
	EQUAL        = "="
	EQUALEQUAL   = "=="
	GREATER      = ">"
	GREATEREQUAL = ">="
	LESS         = "<"
	LESSEQUAL    = "<="
	POWER        = "**"

	// literals

	IDENTIFIER = "IDENT"
	STRING     = "STRING"
	NUMBER     = "NUMBER"

	// keywords

	AND      = "and"
	CLASS    = "class"
	ELSE     = "else"
	FALSE    = "false"
	FUN      = "fun"
	FOR      = "for"
	IF       = "if"
	NIL      = "nil"
	OR       = "or"
	PRINT    = "print"
	RETURN   = "return"
	SUPER    = "super"
	THIS     = "this"
	TRUE     = "true"
	VAR      = "var"
	WHILE    = "while"
	BREAK    = "break"
	CONTINUE = "continue"
	EOF      = "eof"
	INVALID  = "__INVALID__"
)

// tokenNames maps the token values to their names
var tokenNames = map[Type]string{
	LEFTPAREN:  "LEFT_PAREN",
	RIGHTPAREN: "RIGHT_PAREN",
	LEFTBRACE:  "LEFT_BRACE",
	RIGHTBRACE: "RIGHT_BRACE",
	COMMA:      "COMMA",
	DOT:        "DOT",
	MINUS:      "MINUS",
	PLUS:       "PLUS",
	SEMICOLON:  "SEMICOLON",
	SLASH:      "SLASH",
	STAR:       "STAR",
	QMARK:      "QMARK",
	COLON:      "COLON",

	BANG:         "BANG",
	BANGEQUAL:    "BANG_EQUAL",
	EQUAL:        "EQUAL",
	EQUALEQUAL:   "EQUAL_EQUAL",
	GREATER:      "GREATER",
	GREATEREQUAL: "GREATER_EQUAL",
	LESS:         "LESS",
	LESSEQUAL:    "LESS_EQUAL",
	POWER:        "POWER",

	IDENTIFIER: "IDENTIFIER",
	STRING:     "STRING",
	NUMBER:     "NUMBER",

	AND:      "AND",
	CLASS:    "CLASS",
	ELSE:     "ELSE",
	FALSE:    "FALSE",
	FUN:      "FUN",
	FOR:      "FOR",
	IF:       "IF",
	NIL:      "NIL",
	OR:       "OR",
	PRINT:    "PRINT",
	RETURN:   "RETURN",
	SUPER:    "SUPER",
	THIS:     "THIS",
	TRUE:     "TRUE",
	VAR:      "VAR",
	WHILE:    "WHILE",
	BREAK:    "BREAK",
	CONTINUE: "CONTINUE",
	EOF:      "EOF",
	INVALID:  "INVALID",
}

// Token contains the lexeme read by the scanner
type Token struct {
	Type    Type
	Lexeme  string
	Literal interface{}
	Line    int
}

func (token *Token) String() string {
	return fmt.Sprintf("%s %s %v", token.Type, token.Lexeme, token.Literal)
}

// GetTokenName returns the name of the token given its value
func GetTokenName(value Type) string {
	if name, exists := tokenNames[value]; exists {
		return name
	}
	return "UNKNOWN"
}
