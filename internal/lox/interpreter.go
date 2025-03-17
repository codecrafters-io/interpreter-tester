package lox

import (
	"fmt"
	"io"
	"reflect"
	"strconv"
)

const (
	OPERAND_MUST_BE_A_NUMBER                    = "Operand must be a number"
	OPERANDS_MUST_BE_TWO_NUMBERS_OR_TWO_STRINGS = "Operands must be two numbers or two strings"
)

type ReturnError struct {
	error
	value interface{}
}

func BasicInterpret(expression Expr, stdout io.Writer, stderr io.Writer) {
	result, err := Eval(expression, NewGlobal(), Locals{}, stdout, stderr)
	if err != nil {
		LogRuntimeError(err, stderr)
		return
	}
	if result == nil {
		result = "nil"
	}
	fmt.Fprintln(stdout, result)
}

func Interpret(statements []Stmt, env *Environment, locals Locals, stdout io.Writer, stderr io.Writer) {
	OldGlobalEnv := GlobalEnv
	GlobalEnv = env
	InitializeNativeFunctions()

	for _, stmt := range statements {
		_, err := Eval(stmt, env, locals, stdout, stderr)
		if err != nil {
			LogRuntimeError(err, stderr)
			return
		}
	}
	GlobalEnv = OldGlobalEnv
}

// Eval evaluates the given AST
func Eval(node Node, environment *Environment, locals Locals, stdout io.Writer, stderr io.Writer) (interface{}, error) {
	switch n := node.(type) {
	case *Literal:
		return n.Value, nil
	case *Grouping:
		return Eval(n.Expression, environment, locals, stdout, stderr)
	case *Unary:
		right, err := Eval(n.Right, environment, locals, stdout, stderr)
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
		left, err := Eval(n.Left, environment, locals, stdout, stderr)
		if err != nil {
			return left, err
		}
		right, err := Eval(n.Right, environment, locals, stdout, stderr)
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
		value, err := Eval(n.Expression, environment, locals, stdout, stderr)
		if err != nil {
			return value, err
		}
		switch value := value.(type) {
		case float64:
			// Print without exponent notation
			fmt.Fprintln(stdout, strconv.FormatFloat(value, 'f', -1, 64))
		case nil:
			fmt.Fprintln(stdout, "nil")
		default:
			fmt.Fprintln(stdout, value) // Uses the Format()/String() method of the interface value
		}
		return nil, nil
	case *Expression:
		r, err := Eval(n.Expression, environment, locals, stdout, stderr)
		if err != nil {
			return r, err
		}
		return nil, nil
	case *Var:
		if n.Initializer != nil {
			value, err := Eval(n.Initializer, environment, locals, stdout, stderr)
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
		if distance, ok := locals[n]; ok {
			return environment.GetAt(distance, n.Name)
		}
		return GlobalEnv.Get(n.Name)
	case *Assign:
		value, err := Eval(n.Value, environment, locals, stdout, stderr)
		if err != nil {
			return nil, err
		}
		if distance, ok := locals[n]; ok {
			if err := environment.AssignAt(distance, n.Name, value); err == nil {
				return value, nil
			}
			return nil, err
		} else if err := GlobalEnv.Assign(n.Name, value); err == nil {
			return value, nil
		}
		return nil, err
	case *Block:
		newEnvironment := New(environment)
		for _, stmt := range n.Statements {
			_, err := Eval(stmt, newEnvironment, locals, stdout, stderr)
			if err != nil {
				return nil, err
			}
		}
		return nil, nil
	case *If:
		condition, err := Eval(n.Condition, environment, locals, stdout, stderr)
		if err != nil {
			return nil, err
		}
		if isTruthy(condition) {
			return Eval(n.ThenBranch, environment, locals, stdout, stderr)
		} else if n.ElseBranch != nil {
			return Eval(n.ElseBranch, environment, locals, stdout, stderr)
		}
		return nil, nil
	case *Logical:
		left, err := Eval(n.Left, environment, locals, stdout, stderr)
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
		return Eval(n.Right, environment, locals, stdout, stderr)
	case *Call:
		callee, err := Eval(n.Callee, environment, locals, stdout, stderr)
		if err != nil {
			return nil, err
		}

		args := make([]interface{}, 0)
		for _, arg := range n.Arguments {
			a, err := Eval(arg, environment, locals, stdout, stderr)
			if err == nil {
				args = append(args, a)
			} else {
				return nil, err
			}
		}

		function, ok := callee.(Callable)
		if !ok {
			return nil, MakeRuntimeError(n.Paren, "Can only call functions and classes.")
		}

		if function.Arity() != len(args) {
			return nil, MakeRuntimeError(n.Paren, fmt.Sprintf("Expected %d arguments but got %d.", function.Arity(), len(args)))
		}

		return function.Call(args, environment, stdout, stderr)
	case *While:
		for {
			condition, err := Eval(n.Condition, environment, locals, stdout, stderr)
			if err != nil {
				return nil, err
			}
			if !isTruthy(condition) {
				break
			}
			_, err = Eval(n.Statement, environment, locals, stdout, stderr)
			if err != nil {
				return nil, err
			}
		}
		return nil, nil
	case *Function:
		function := NewUserFunction(n, environment, locals)
		environment.Define(n.Name.Lexeme, function)
		return nil, nil
	case *Return:
		var value interface{}
		var err error
		if n.Value != nil {
			value, err = Eval(n.Value, environment, locals, stdout, stderr)
			if err != nil {
				return nil, err
			}
		}
		return nil, ReturnError{value: value}
	case *Class:
		var superclass *UserClass
		if n.SuperClass != nil {
			sc, err := Eval(n.SuperClass, environment, locals, stdout, stderr)
			if err != nil {
				return nil, err
			} else if sup, ok := sc.(*UserClass); ok {
				superclass = sup
			} else {
				return nil, MakeRuntimeError(n.SuperClass.Name, "Superclass must be a class.")
			}
		}

		environment.Define(n.Name.Lexeme, nil)

		if superclass != nil {
			environment = New(environment)
			environment.Define("super", superclass)
		}

		methods := make(map[string]*UserFunction)
		for _, method := range n.Methods {
			function := NewUserFunction(method, environment, locals)
			methods[method.Name.Lexeme] = function
			if method.Name.Lexeme == "init" {
				function.IsInitializer = true
			}
		}

		if superclass != nil {
			environment = environment.Ancestor(1)
		}

		userClass := &UserClass{SuperClass: superclass, Name: n.Name.Lexeme, Methods: methods}
		environment.Assign(n.Name, userClass)

		return nil, nil
	case *Get:
		value, err := Eval(n.Expression, environment, locals, stdout, stderr)
		if err != nil {
			return nil, err
		}
		if obj, ok := value.(*UserClassInstance); ok {
			return obj.Get(n.Name)
		}
		return nil, MakeRuntimeError(n.Name, "Only instances have properties.")
	case *Set:
		obj, err := Eval(n.Object, environment, locals, stdout, stderr)
		if err != nil {
			return nil, err
		}
		if instance, ok := obj.(*UserClassInstance); ok {
			value, err := Eval(n.Value, environment, locals, stdout, stderr)
			if err != nil {
				return nil, err
			}
			return instance.Set(n.Name, value)
		}
		return nil, MakeRuntimeError(n.Name, "Only instances have properties.")
	case *This:
		if distance, ok := locals[n]; ok {
			return environment.GetAt(distance, n.Keyword)
		}
		return GlobalEnv.Get(n.Keyword)
	case *Super:
		sc, err := environment.GetAt(n.EnvDepth, Token{Lexeme: "super"})
		if err != nil {
			return nil, err
		}
		if superclass, ok := sc.(*UserClass); ok {
			tc, err2 := environment.GetAt(n.EnvDepth-1, Token{Lexeme: "this"})
			if err2 != nil {
				return nil, err2
			}

			if thisclass, ok2 := tc.(*UserClassInstance); ok2 {
				method, err3 := superclass.FindMethod(n.Method)
				if err3 != nil {
					return nil, err3
				}
				return method.Bind(thisclass), nil
			}
		}
		return nil, MakeRuntimeError(n.Keyword, "Fatal error: 'super' not a class instance ?")
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
	HadSemanticError = false
}
