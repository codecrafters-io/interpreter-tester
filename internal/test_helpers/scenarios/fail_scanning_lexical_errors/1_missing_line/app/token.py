from enum import StrEnum
from typing import Optional


class TokenType(StrEnum):
    COMMA = "COMMA"
    DOT = "DOT"
    EOF = "EOF"
    LEFT_BRACE = "LEFT_BRACE"
    LEFT_PAREN = "LEFT_PAREN"
    MINUS = "MINUS"
    PLUS = "PLUS"
    RIGHT_BRACE = "RIGHT_BRACE"
    RIGHT_PAREN = "RIGHT_PAREN"
    SEMICOLON = "SEMICOLON"
    STAR = "STAR"


class Token:
    type: TokenType
    lexeme: Optional[str]
    literal: str | int | None
    line_number: int

    def __init__(
        self,
        type: TokenType,
        lexeme: Optional[str],
        literal: str | int | None,
        line_number: int,
    ):
        self.type = type
        self.lexeme = lexeme
        self.literal = literal
        self.line_number = line_number

    def __str__(self) -> str:
        return f"{self.type} {self.lexeme} {self.literal or 'null'}"
