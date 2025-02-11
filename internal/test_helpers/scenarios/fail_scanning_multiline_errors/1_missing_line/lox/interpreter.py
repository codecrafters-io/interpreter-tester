from lox.types import *


def repr(value):
    if value is None:
        return 'nil'
    elif isinstance(value, bool):
        return str(value).lower()
    elif getattr(value, 'is_integer', False) and value.is_integer():
        return int(value)
    else:
        return value


def check_number_operand(operator: Token, operand):
    if isinstance(operand, float):
        return
    # print(operand, type(operand))
    raise RuntimeError(operator, 'Operand must be a number.')


def check_number_operands(operator: Token, left, right):
    if isinstance(left, float) and isinstance(right, float):
        return
    raise RuntimeError(operator, 'Operands must be numbers.')


def interpret(ast):
    def eval(expr: Expr):
        # fmt: off
        match expr: 
            case Binary(left, op, right):
                l, r = eval(left), eval(right)

                match op.type:
                    case TokenType.BANG_EQUAL : return l != r
                    case TokenType.EQUAL_EQUAL: return l == r
                    case TokenType.PLUS: 
                        if (isinstance(l, float) and isinstance(r, float) 
                         or isinstance(l, str)   and isinstance(r, str)):
                            return l + r
                        raise RuntimeError(op, 'Operands must be two numbers or two strings.')

                check_number_operands(op, l, r)
                match op.type:
                    case TokenType.MINUS        : return l - r
                    case TokenType.SLASH        : return l / r
                    case TokenType.STAR         : return l * r
                    case TokenType.GREATER      : return l > r
                    case TokenType.GREATER_EQUAL: return l >= r
                    case TokenType.LESS         : return l < r
                    case TokenType.LESS_EQUAL   : return l <= r

            case Grouping(expr)  : return eval(expr)
            case Literal(value)  : return value
            case Unary(op, right): 
                match op.type:
                    case TokenType.BANG : 
                        return not is_truthy(right)
                    case TokenType.MINUS: 
                        r = eval(right)
                        check_number_operand(op, r)
                        return -r
        # fmt: on

    def is_truthy(value: Expr):
        # fmt: off
        match value:
            case Literal('nil')  : return False
            case Literal('false'): return False
            case Literal('true') : return True
            case _               : return eval(value)
        # fmt: on

    print(repr(eval(ast)))
