package lox

import (
	"fmt"
	"reflect"
)

const (
	OPERAND_MUST_BE_A_NUMBER                    = "Operand must be a number"
	OPERANDS_MUST_BE_TWO_NUMBERS_OR_TWO_STRINGS = "Operands must be two numbers or two strings"
)

func Interpret(expression Expr) (interface{}, error) {
	result, err := Eval(expression)
	if err != nil {
		return "", err
	}
	return result, nil
}

// Eval evaluates the given AST
func Eval(node Node) (interface{}, error) {
	switch n := node.(type) {
	case *Literal:
		return n.Value, nil
	case *Grouping:
		return Eval(n.Expression)
	case *Unary:
		right, err := Eval(n.Right)
		if err != nil {
			return right, err
		} else if n.Operator.Type == MINUS {
			err := checkNumberOperand(n.Operator, right, OPERAND_MUST_BE_A_NUMBER)
			if err != nil {
				return nil, err
			}
			return -right.(float64), nil
		} else if n.Operator.Type == BANG {
			return !isTruthy(right), nil
		}
	case *Binary:
		left, err := Eval(n.Left)
		if err != nil {
			return left, err
		}
		right, err := Eval(n.Right)
		if err != nil {
			return right, err
		}
		switch n.Operator.Type {
		case MINUS:
			err := checkNumberOperand(n.Operator, left, OPERAND_MUST_BE_A_NUMBER)
			if err != nil {
				return nil, err
			}
			err = checkNumberOperand(n.Operator, right, OPERAND_MUST_BE_A_NUMBER)
			if err != nil {
				return nil, err
			}
			return left.(float64) - right.(float64), nil
		case PLUS:
			switch lhs := left.(type) {
			case float64:
				switch rhs := right.(type) {
				case float64:
					return lhs + rhs, nil
				}
			case string:
				switch rhs := right.(type) {
				case string:
					return lhs + rhs, nil
				}
			}
			return nil, fmt.Errorf("%s", OPERANDS_MUST_BE_TWO_NUMBERS_OR_TWO_STRINGS)
		case SLASH:
			err := checkNumberOperand(n.Operator, left, OPERAND_MUST_BE_A_NUMBER)
			if err != nil {
				return nil, err
			}
			err = checkNumberOperand(n.Operator, right, OPERAND_MUST_BE_A_NUMBER)
			if err != nil {
				return nil, err
			}
			return left.(float64) / right.(float64), nil
		case STAR:
			err := checkNumberOperand(n.Operator, left, OPERAND_MUST_BE_A_NUMBER)
			if err != nil {
				return nil, err
			}
			err = checkNumberOperand(n.Operator, right, OPERAND_MUST_BE_A_NUMBER)
			if err != nil {
				return nil, err
			}
			return left.(float64) * right.(float64), nil
		case GREATER:
			err := checkNumberOperand(n.Operator, left, OPERAND_MUST_BE_A_NUMBER)
			if err != nil {
				return nil, err
			}
			err = checkNumberOperand(n.Operator, right, OPERAND_MUST_BE_A_NUMBER)
			if err != nil {
				return nil, err
			}
			return left.(float64) > right.(float64), nil
		case GREATEREQUAL:
			err := checkNumberOperand(n.Operator, left, OPERAND_MUST_BE_A_NUMBER)
			if err != nil {
				return nil, err
			}
			err = checkNumberOperand(n.Operator, right, OPERAND_MUST_BE_A_NUMBER)
			if err != nil {
				return nil, err
			}
			return left.(float64) >= right.(float64), nil
		case LESS:
			err := checkNumberOperand(n.Operator, left, OPERAND_MUST_BE_A_NUMBER)
			if err != nil {
				return nil, err
			}
			err = checkNumberOperand(n.Operator, right, OPERAND_MUST_BE_A_NUMBER)
			if err != nil {
				return nil, err
			}
			return left.(float64) < right.(float64), nil
		case LESSEQUAL:
			err := checkNumberOperand(n.Operator, left, OPERAND_MUST_BE_A_NUMBER)
			if err != nil {
				return nil, err
			}
			err = checkNumberOperand(n.Operator, right, OPERAND_MUST_BE_A_NUMBER)
			if err != nil {
				return nil, err
			}
			return left.(float64) <= right.(float64), nil
		case BANGEQUAL:
			return !isEqual(left, right), nil
		case EQUALEQUAL:
			return isEqual(left, right), nil
		}
	case *Expression:
		r, err := Eval(n.Expression)
		if err != nil {
			return r, err
		}
		return nil, nil
	case nil:
		return nil, MakeRuntimeError(Token{Lexeme: ""}, "Fatal interpreter error.")
	}
	panic(fmt.Sprintf("CodeCrafters Internal Error: Unexpected node type in interpreter: %s", reflect.TypeOf(node).String()))
}

func isTruthy(val interface{}) bool {
	if val == nil {
		return false
	} else if b, ok := val.(bool); ok {
		return b
	}
	return true
}

func isEqual(left interface{}, right interface{}) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil {
		return false
	}
	return left == right
}

func checkNumberOperand(operator Token, value interface{}, msg string) error {
	switch value.(type) {
	case int, float64:
		return nil
	}
	return fmt.Errorf("%v\n[line %v]", msg, operator.Line)
}
