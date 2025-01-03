from lox.types import *


class Scanner:
    def __init__(self, lox, source: str):
        self._tokens: list[Token] = []
        self._start = 0
        self._current = 0
        self._line = 1
        self._lox = lox
        self._source = source

    def scan_tokens(self) -> list[Token]:
        while not self._is_at_end():
            self._scan_token()
        self._tokens += (Token(TokenType.EOF, '', None, self._line),)
        return self._tokens

    def _add_token(self, type: TokenType, literal=None):
        text = self._source[self._start : self._current]
        self._tokens += (Token(type, text, literal, self._line),)

    def _advance(self) -> str:
        c = self._source[self._current]
        self._current += 1
        return c

    def _error_char(self):
        char = self._source[self._start : self._current]
        self._lox.error_line(self._line, 'Unexpected character: ' + char)

    def _identifier(self):
        while (c := self._peek()).isalnum() or c == '_':
            self._advance()

        text = self._source[self._start : self._current]
        type = KEYWORDS.get(text) or TokenType.IDENTIFIER
        self._add_token(type)

    def _is_at_end(self) -> bool:
        return self._current >= len(self._source)

    def _match(self, expected: str) -> bool:
        if self._is_at_end() or (self._source[self._current] != expected):
            return False
        self._current += 1
        return True

    def _number(self):
        while self._peek().isdigit():
            self._advance()

        if self._peek() == '.' and self._peek_next().isdigit():
            self._advance()  # Consume the "."
            while self._peek().isdigit():
                self._advance()

        literal = self._source[self._start : self._current]
        self._add_token(TokenType.NUMBER, float(literal))

    def _peek(self) -> str:
        return '\0' if self._is_at_end() else self._source[self._current]

    def _peek_next(self) -> str:
        if self._current + 1 >= len(self._source):
            return '\0'
        return self._source[self._current + 1]

    def _scan_token(self):
        self._start = self._current
        # fmt: off
        match c := self._advance():
            case '(' : self._add_token(TokenType.LEFT_PAREN)
            case ')' : self._add_token(TokenType.RIGHT_PAREN)
            case '{' : self._add_token(TokenType.LEFT_BRACE)
            case '}' : self._add_token(TokenType.RIGHT_BRACE)
            case ',' : self._add_token(TokenType.COMMA)
            case '.' : self._add_token(TokenType.DOT)
            case '-' : self._add_token(TokenType.MINUS)
            case '+' : self._add_token(TokenType.PLUS)
            case ';' : self._add_token(TokenType.SEMICOLON)
            case '*' : self._add_token(TokenType.STAR)

            case '!' : self._add_token(TokenType.BANG_EQUAL    if self._match('=') else TokenType.BANG)
            case '=' : self._add_token(TokenType.EQUAL_EQUAL   if self._match('=') else TokenType.EQUAL)
            case '<' : self._add_token(TokenType.LESS_EQUAL    if self._match('=') else TokenType.LESS)
            case '>' : self._add_token(TokenType.GREATER_EQUAL if self._match('=') else TokenType.GREATER)

            case '/' : self._slash()

            case ' ' | '\t': pass
            case '\n': self._line += 1

            case '"' : self._string()

            case c if c.isdigit(): self._number()
            case c if c.isalpha() or c == '_': self._identifier()

            case _   : self._error_char()
        # fmt: on

    def _slash(self):
        if self._match('/'):
            while self._peek() != '\n' and not self._is_at_end():
                self._advance()
        else:
            self._add_token(TokenType.SLASH)

    def _string(self):
        while self._peek() != '"' and not self._is_at_end():
            if self._peek() == '\n':
                self._line += 1
            self._advance()

        if self._is_at_end():
            return self._lox.error_line(self._line, 'Unterminated string.')

        self._advance()  # The closing ".
        value = self._source[self._start + 1 : self._current - 1]
        self._add_token(TokenType.STRING, value)
