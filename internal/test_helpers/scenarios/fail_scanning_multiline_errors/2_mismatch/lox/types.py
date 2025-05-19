from enum import Enum
from typing import NamedTuple

# fmt: off
TokenType = Enum('TokenType', [
    # Single-character tokens.
    'LEFT_PAREN', 'RIGHT_PAREN', 'LEFT_BRACE', 'RIGHT_BRACE',
    'COMMA', 'DOT', 'MINUS', 'PLUS', 'SEMICOLON', 'SLASH', 'STAR',

    # One or two character tokens.
    'BANG', 'BANG_EQUAL',
    'EQUAL', 'EQUAL_EQUAL',
    'GREATER', 'GREATER_EQUAL',
    'LESS', 'LESS_EQUAL',

    # Literals.
    'IDENTIFIER', 'NUMBER', 'STRING', 

    # Keywords.
    'AND', 'CLASS', 'ELSE', 'FALSE', 'FUN', 'FOR', 'IF', 'NIL', 'OR',
    'PRINT', 'RETURN', 'SUPER', 'THIS', 'TRUE', 'VAR', 'WHILE',

    'EOF'
])
# fmt: on

KEYWORDS = {
    'and': TokenType.AND,
    'class': TokenType.CLASS,
    'else': TokenType.ELSE,
    'false': TokenType.FALSE,
    'for': TokenType.FOR,
    'fun': TokenType.FUN,
    'if': TokenType.IF,
    'nil': TokenType.NIL,
    'or': TokenType.OR,
    'print': TokenType.PRINT,
    'return': TokenType.RETURN,
    'super': TokenType.SUPER,
    'this': TokenType.THIS,
    'true': TokenType.TRUE,
    'var': TokenType.VAR,
    'while': TokenType.WHILE,
}


class Token(NamedTuple):
    type: TokenType
    lexeme: str
    literal: object
    line: int

    def __repr__(self):
        return f'{self.type.name} {self.lexeme} {"null" if self.literal is None else self.literal}'


class ParseError(Exception):
    pass


class Binary(NamedTuple):
    left: 'Expr'
    op: Token
    right: 'Expr'


class Grouping(NamedTuple):
    expr: 'Expr'


class Literal(NamedTuple):
    value: object


class Unary(NamedTuple):
    op: Token
    right: 'Expr'


Expr = Binary | Grouping | Literal | Unary
