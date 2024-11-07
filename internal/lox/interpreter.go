package lox

import (
	"fmt"
	"io"
	"reflect"
)

const (
	OPERAND_MUST_BE_A_NUMBER                    = "Operand must be a number"
	OPERANDS_MUST_BE_TWO_NUMBERS_OR_TWO_STRINGS = "Operands must be two numbers or two strings"
)

type LoxCallable interface {
	Call(args []interface{}) (interface{}, error)
	Arity() int
}

type nativeFunction struct {
	Arity func() int
	Call  func(args []interface{}) (interface{}, error)
}

func BasicInterpret(expression Expr, stdout io.Writer, stderr io.Writer) {
	result, err := Eval(expression, NewGlobal(), stdout, stderr)
	if err != nil {
		LogRuntimeError(err, stderr)
		return
	}
	if result == nil {
		result = "nil"
	}
	fmt.Fprintln(stdout, result)
}

func Interpret(statements []Stmt, stdout io.Writer, stderr io.Writer) {
	env := NewGlobal()
	InitializeNativeFunctions(env)
	for _, stmt := range statements {
		_, err := Eval(stmt, env, stdout, stderr)
		if err != nil {
			LogRuntimeError(err, stderr)
			return
		}
	}
}

// Eval evaluates the given AST
func Eval(node Node, environment *Environment, stdout io.Writer, stderr io.Writer) (interface{}, error) {
	switch n := node.(type) {
	case *Literal:
		return n.Value, nil
	case *Grouping:
		return Eval(n.Expression, environment, stdout, stderr)
	case *Unary:
		right, err := Eval(n.Right, environment, stdout, stderr)
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
		left, err := Eval(n.Left, environment, stdout, stderr)
		if err != nil {
			return left, err
		}
		right, err := Eval(n.Right, environment, stdout, stderr)
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
			return nil, MakeRuntimeError(n.Operator, OPERANDS_MUST_BE_TWO_NUMBERS_OR_TWO_STRINGS)
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
	case *Print:
		value, err := Eval(n.Expression, environment, stdout, stderr)
		if err != nil {
			return value, err
		}
		fmt.Fprintln(stdout, value)
		return nil, nil
	case *Expression:
		r, err := Eval(n.Expression, environment, stdout, stderr)
		if err != nil {
			return r, err
		}
		return nil, nil
	case *Var:
		if n.Initializer != nil {
			value, err := Eval(n.Initializer, environment, stdout, stderr)
			if err != nil {
				return nil, err
			}
			environment.Define(n.Name.Lexeme, value)
		} else {
			// We initialize uninitialized variables with nil
			environment.Define(n.Name.Lexeme, nil)
		}
		return nil, nil
	case *Variable:
		return environment.Get(n.Name)
	case *Assign:
		value, err := Eval(n.Value, environment, stdout, stderr)
		if err != nil {
			return nil, err
		}
		if err = environment.Assign(n.Name, value); err == nil {
			return value, nil
		}
		return nil, err
	case *Block:
		newEnvironment := New(environment)
		for _, stmt := range n.Statements {
			_, err := Eval(stmt, newEnvironment, stdout, stderr)
			if err != nil {
				return nil, err
			}
		}
		return nil, nil
	case *If:
		condition, err := Eval(n.Condition, environment, stdout, stderr)
		if err != nil {
			return nil, err
		}
		if isTruthy(condition) {
			return Eval(n.ThenBranch, environment, stdout, stderr)
		} else if n.ElseBranch != nil {
			return Eval(n.ElseBranch, environment, stdout, stderr)
		}
		return nil, nil
	case *Logical:
		left, err := Eval(n.Left, environment, stdout, stderr)
		if err != nil {
			return nil, err
		}
		switch n.Operator.Type {
		case OR:
			if isTruthy(left) {
				return left, nil
			}
		case AND:
			if !isTruthy(left) {
				return left, nil
			}
		}
		return Eval(n.Right, environment, stdout, stderr)
	case *Call:
		callee, err := Eval(n.Callee, environment, stdout, stderr)
		if err != nil {
			return nil, err
		}

		args := make([]interface{}, 0)
		for _, arg := range n.Arguments {
			a, err := Eval(arg, environment, stdout, stderr)
			if err == nil {
				args = append(args, a)
			} else {
				return nil, err
			}
		}

		function, ok := callee.(LoxCallable)
		if !ok {
			return nil, MakeRuntimeError(n.Paren, "Can only call functions and classes.")
		}

		if function.Arity() != len(args) {
			return nil, MakeRuntimeError(n.Paren, fmt.Sprintf("Expected %d arguments but got %d.", function.Arity(), len(args)))
		}

		return function.Call(args)
	case *While:
		for {
			condition, err := Eval(n.Condition, environment, stdout, stderr)
			if err != nil {
				return nil, err
			}
			if !isTruthy(condition) {
				break
			}
			_, err = Eval(n.Statement, environment, stdout, stderr)
			if err != nil {
				return nil, err
			}
		}
		return nil, nil
	case nil:
		return nil, nil
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
	return MakeRuntimeError(operator, msg)
}

func ClearErrorFlags() {
	HadParseError = false
	HadRuntimeError = false
}
