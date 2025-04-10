package lox

import (
	"fmt"
)

// Locals is the output of the resolve phase
type Locals = map[Expr]int

// Scope represents a Lox scope
type Scope = map[string]bool

// Resolve performs name resolution to the given statements
func Resolve(statements []Stmt) (Locals, error) {
	locals := make(Locals)
	resolver := &Resolver{scopes: make([]Scope, 0), currentFunctionType: ftNone, currentClassType: ctNone}
	err := resolver.resolveStatements(statements, locals)
	return locals, err
}

// FunctionType represents the type of a function
const (
	ftNone        = iota
	ftFunction    = iota
	ftMethod      = iota
	ftInitializer = iota
)

// ClassType represents the type of a class
const (
	ctNone     = iota
	ctClass    = iota
	ctSubClass = iota
)

// Resolver performs variable resolution on an AST
type Resolver struct {
	scopes              []Scope
	currentFunctionType int
	currentClassType    int
}

func (r *Resolver) resolve(node Node, locals Locals) error {
	switch n := node.(type) {
	case *Block:
		r.pushScope()
		defer r.popScope()
		for _, stmt := range n.Statements {
			if err := r.resolve(stmt, locals); err != nil {
				return err
			}
		}
	case *Var:
		if err := r.declare(n.Name); err != nil {
			return MakeSemanticError("Already a variable with this name in this scope.")
		}
		if n.Initializer != nil {
			if err := r.resolve(n.Initializer, locals); err != nil {
				return MakeSemanticError("Can't read local variable in its own initializer.")
			}
		}
		r.define(n.Name)
	case *Variable:
		if len(r.scopes) != 0 {
			// Local scope
			if b, ok := r.scopes[len(r.scopes)-1][n.Name.Lexeme]; ok && !b {
				return MakeSemanticError("Cannot read local variable in its own initializer.")
			}
		}
		r.resolveLocal(n, n.Name, locals)
	case *Assign:
		if err := r.resolve(n.Value, locals); err != nil {
			return err
		}
		r.resolveLocal(n, n.Name, locals)
	case *Function:
		if err := r.declare(n.Name); err != nil {
			return err
		}
		r.define(n.Name)
		if err := r.resolveFunction(n, locals, ftFunction); err != nil {
			return err
		}
	case *Expression:
		if err := r.resolve(n.Expression, locals); err != nil {
			return err
		}
	case *If:
		if err := r.resolve(n.Condition, locals); err != nil {
			return err
		}
		if err := r.resolve(n.ThenBranch, locals); err != nil {
			return err
		}
		if n.ElseBranch != nil {
			if err := r.resolve(n.ElseBranch, locals); err != nil {
				return err
			}
		}
	case *Print:
		if err := r.resolve(n.Expression, locals); err != nil {
			return err
		}
	case *Return:
		if r.currentFunctionType == ftNone {
			return MakeSemanticError("Cannot return from top-level code.")
		}
		if n.Value != nil {
			if r.currentFunctionType == ftInitializer {
				return MakeSemanticError("Cannot return a value from an initializer.")
			}
			if err := r.resolve(n.Value, locals); err != nil {
				return err
			}
		}
	case *While:
		if err := r.resolve(n.Condition, locals); err != nil {
			return err
		}
		if err := r.resolve(n.Statement, locals); err != nil {
			return err
		}
	case *Binary:
		if err := r.resolve(n.Left, locals); err != nil {
			return err
		}
		if err := r.resolve(n.Right, locals); err != nil {
			return err
		}
	case *Call:
		if err := r.resolve(n.Callee, locals); err != nil {
			return err
		}

		for _, e := range n.Arguments {
			if err := r.resolve(e, locals); err != nil {
				return err
			}
		}
	case *Grouping:
		if err := r.resolve(n.Expression, locals); err != nil {
			return err
		}
	case *Logical:
		if err := r.resolve(n.Left, locals); err != nil {
			return err
		}
		if err := r.resolve(n.Right, locals); err != nil {
			return err
		}
	case *Unary:
		if err := r.resolve(n.Right, locals); err != nil {
			return err
		}
	case *Class:
		enclosingClass := r.currentClassType
		r.currentClassType = ctClass

		resetCurrentClass := func() {
			r.currentClassType = enclosingClass
		}

		defer resetCurrentClass()

		if err := r.declare(n.Name); err != nil {
			return err
		}
		r.define(n.Name)

		if n.SuperClass != nil {
			if n.Name.Lexeme == n.SuperClass.Name.Lexeme {
				return MakeSemanticError("A class cannot inherit from itself.")
			}
			r.currentClassType = ctSubClass
			err := r.resolve(n.SuperClass, locals)
			if err != nil {
				return err
			}

			r.pushScope()
			defer r.popScope()

			top := r.scopes[len(r.scopes)-1]
			top["super"] = true
			r.scopes[len(r.scopes)-1] = top
		}

		r.pushScope()
		defer r.popScope()

		top := r.scopes[len(r.scopes)-1]
		top["this"] = true
		r.scopes[len(r.scopes)-1] = top

		for _, method := range n.Methods {
			declaration := ftMethod
			if method.Name.Lexeme == "init" {
				declaration = ftInitializer
			}
			if err := r.resolveFunction(method, locals, declaration); err != nil {
				return err
			}
		}
	case *Super:
		if r.currentClassType == ctNone {
			return MakeSemanticError("Cannot use 'super' outside of a class.")
		} else if r.currentClassType != ctSubClass {
			return MakeSemanticError("Cannot use 'super' in a class with no superclass.")
		}
		r.resolveLocal(n, n.Keyword, locals)
	case *Get:
		if err := r.resolve(n.Expression, locals); err != nil {
			return err
		}
	case *Set:
		if err := r.resolve(n.Value, locals); err != nil {
			return err
		}
		if err := r.resolve(n.Object, locals); err != nil {
			return err
		}
	case *This:
		if r.currentClassType == ctNone {
			return MakeSemanticError("Cannot use 'this' outside of a class.")
		}
		r.resolveLocal(n, n.Keyword, locals)
	}
	return nil
}

func (r *Resolver) resolveStatements(statements []Stmt, locals Locals) error {
	for _, stmt := range statements {
		if err := r.resolve(stmt, locals); err != nil {
			return err
		}
	}
	return nil
}

func (r *Resolver) resolveFunction(function *Function, locals Locals, functionType int) error {
	enclosingFunctionType := r.currentFunctionType
	r.currentFunctionType = functionType

	resetCurrentFunction := func() {
		r.currentFunctionType = enclosingFunctionType
	}

	defer resetCurrentFunction()

	r.pushScope()
	defer r.popScope()

	for _, param := range function.Params {
		if err := r.declare(param); err != nil {
			return err
		}
		r.define(param)
	}
	return r.resolveStatements(function.Body, locals)
}

func (r *Resolver) resolveLocal(expr Expr, name Token, locals Locals) {
	for i := len(r.scopes) - 1; i >= 0; i-- {
		if _, ok := r.scopes[i][name.Lexeme]; ok {
			locals[expr] = len(r.scopes) - i - 1
			return
		}
	}
}

func (r *Resolver) pushScope() {
	r.scopes = append(r.scopes, make(Scope))
}

func (r *Resolver) popScope() {
	r.scopes = r.scopes[:len(r.scopes)-1]
}

func (r *Resolver) declare(name Token) error {
	if len(r.scopes) != 0 {
		scope := r.scopes[len(r.scopes)-1]
		if _, ok := scope[name.Lexeme]; ok {
			return MakeSemanticError(
				fmt.Sprintf("Variable '%s' already declared in this scope.", name.Lexeme))
		}
		scope[name.Lexeme] = false
	}
	return nil
}

func (r *Resolver) define(name Token) {
	if len(r.scopes) != 0 {
		scope := r.scopes[len(r.scopes)-1]
		scope[name.Lexeme] = true
	}
}
