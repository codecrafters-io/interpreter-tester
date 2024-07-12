package internal

var PARENS = []string{"(", ")"}

var BRACES = []string{"{", "}"}

var SINGLECHAROPERATORS = []string{"+", "-", "*", ".", ",", ";"}

var LEXICALERRORS = []string{"@", "$", "#", "%"}

var EQUALS = []string{"=", "=="}

var NEGATIONS = []string{"!", "!="}

var RELATIONALS = []string{"<", ">", "<=", ">="}

var DIVISION = []string{"/"}

var WHITESPACES = []string{SP, TAB, LF}
var SP = " "
var TAB = "	"
var LF = "\n"

var IDENTIFIERS = []string{"_hello", "world_", "f00", "6ar", "6az", "foo", "bar", "baz"}

var KEYWORDS = []string{"and", "class", "else", "false", "for", "fun", "if", "nil", "or", "print", "return", "super", "this", "true", "var", "while"}
var CAPITALIZEDKEYWORDS = []string{"AND", "CLASS", "ELSE", "FALSE", "FOR", "FUN", "IF", "NIL", "OR", "PRINT", "RETURN", "SUPER", "THIS", "TRUE", "VAR", "WHILE"}

var BOOLEANS = []string{"true", "false"}

var QUOTEDSTRINGS = []string{"\"hello\"", "\"world\"", "\"foo\"", "\"bar\"", "\"baz\""}

var STRINGS = []string{"foo", "bar", "baz", "quz", "hello", "world"}
