import sys

from lox.interpreter import interpret
from lox.parser import Parser
from lox.scanner import Scanner, Token, TokenType


class Lox:
    def __init__(self, filename: str):
        self._tokens: list[Token] = []
        self._ast = None
        self._had_error = False
        self._had_runtime_error = False

        with open(filename) as f:
            self._source = f.read()

    def error_line(self, line: int, message: str):
        self._error_report(line, "", message)

    def error_token(self, token: Token, message: str):
        where = " at " + ("end" if token.type == TokenType.EOF else f"'{token.lexeme}'")
        self._error_report(token.line, where, message)

    def evaluate(self):
        self.parse(False)
        try:
            interpret(self._ast)
        except RuntimeError as e:
            self._runtime_error(e)
        self._check_errror_exit()

    def parse(self, should_print=True):
        self.scan(False)
        parser = Parser(self, self._tokens)
        self._ast = parser.parse()
        self._check_errror_exit()
        if should_print:
            Parser.print_ast(self._ast)

    def scan(self, should_print=True):
        scanner = Scanner(self, self._source)
        self._tokens = scanner.scan_tokens()
        if should_print:
            print(*self._tokens, sep="\n")
        self._check_errror_exit()

    def _check_errror_exit(self):
        if self._had_error:
            sys.exit(65)
        if self._had_runtime_error:
            sys.exit(70)

    def _error_report(self, line: int, where: str, message: str):
        self._had_error = True
        print(f"This line will be skipped in stderr", file=sys.stderr)

        output = f"[line {line}] Error{where}: {message}"
        print(output, file=sys.stderr)
        if message.endswith("#"):
            print("[line 888] extra line", file=sys.stderr)

    def _runtime_error(self, error: RuntimeError):
        self._had_runtime_error = True
        token, msg = error.args
        output = msg + f"\n[line {token.line}]"
        print(output, file=sys.stderr)
