package lox

import (
	"fmt"
	"reflect"
	"strings"
)

// Node is the root class of AST nodes
type Node interface {
	String() string
}

// Expr is the root class of expression nodes
type Expr interface {
	Node
}

// Binary is used for binary operators
type Binary struct {
	Expr
	Left     Expr
	Operator Token
	Right    Expr
}

// String pretty prints the operator
func (b *Binary) String() string {
	var sb strings.Builder
	sb.WriteString("(")
	sb.WriteString(b.Operator.Lexeme)
	sb.WriteString(" ")
	sb.WriteString(b.Left.String())
	sb.WriteString(" ")
	sb.WriteString(b.Right.String())
	sb.WriteString(")")
	return sb.String()
}

// Grouping is used for parenthesized expressions
type Grouping struct {
	Expr
	Expression Expr
}

// String pretty prints the expression grouping
func (g *Grouping) String() string {
	var sb strings.Builder
	sb.WriteString("(group ")
	sb.WriteString(g.Expression.String())
	sb.WriteString(")")
	return sb.String()
}

// Literal values
type Literal struct {
	Expr
	Value interface{}
}

// String pretty prints the literal
func (l *Literal) String() string {
	var sb strings.Builder
	if l.Value == nil {
		sb.WriteString("nil")
	} else if reflect.TypeOf(l.Value).Kind() == reflect.Float64 {
		sb.WriteString(formatFloat(l.Value.(float64)))
	} else {
		sb.WriteString(fmt.Sprintf("%v", l.Value))
	}
	return sb.String()
}

// Unary is used for unary operators
type Unary struct {
	Expr
	Operator Token
	Right    Expr
}

// String pretty prints the unary operator
func (u *Unary) String() string {
	var sb strings.Builder
	sb.WriteString("(")
	sb.WriteString(u.Operator.Lexeme)
	sb.WriteString(" ")
	sb.WriteString(u.Right.String())
	sb.WriteString(")")
	return sb.String()
}

// Statements and state

// Assign is used for variable assignment
// name = value
type Assign struct {
	Expr
	Name     Token
	Value    Expr
	EnvIndex int
	EnvDepth int
}

// String pretty prints the assignment statement
func (a *Assign) String() string {
	var sb strings.Builder
	sb.WriteString("(")
	sb.WriteString("=")
	sb.WriteString(" ")
	sb.WriteString(a.Name.Lexeme)
	sb.WriteString(" ")
	sb.WriteString(a.Value.String())
	sb.WriteString(")")
	return sb.String()
}

// Variable access expression
// print x
type Variable struct {
	Expr
	Name     Token
	EnvIndex int
	EnvDepth int
}

// String pretty prints the assignment expression
func (v *Variable) String() string {
	var sb strings.Builder
	sb.WriteString(v.Name.Lexeme)
	return sb.String()
}

// Stmt form a second hierarchy of syntax nodes independent of expressions
type Stmt interface {
	Node
}

// Block is a curly-braced block statement that defines a local scope
//
//	{
//	  ...
//	}
type Block struct {
	Stmt
	Statements []Stmt
	EnvSize    int
}

// String pretty prints the block statement
func (b *Block) String() string {
	var sb strings.Builder
	sb.WriteString("(")
	for _, stmt := range b.Statements {
		sb.WriteString(stmt.String())
	}
	sb.WriteString(")")
	return sb.String()
}

// Expression statement
type Expression struct {
	Stmt
	Expression Expr
}

// String pretty prints the expression statement
func (e *Expression) String() string {
	var sb strings.Builder
	sb.WriteString("(")
	sb.WriteString(e.Expression.String())
	sb.WriteString(")")
	return sb.String()
}

// Print statement
// print 1 + 2
type Print struct {
	Stmt
	Expression Expr
}

// String pretty prints the print statement
func (p *Print) String() string {
	var sb strings.Builder
	sb.WriteString("(")
	sb.WriteString("print")
	sb.WriteString(" ")
	sb.WriteString(p.Expression.String())
	sb.WriteString(")")
	return sb.String()
}

// Var is the variable declaration statement
// var <name> = <initializer>
type Var struct {
	Stmt
	Name        Token
	Initializer Expr
	EnvIndex    int
}

// String pretty prints the var declaration
func (v *Var) String() string {
	var sb strings.Builder
	sb.WriteString("(")
	sb.WriteString("var")
	sb.WriteString(" ")
	sb.WriteString(v.Name.Lexeme)
	sb.WriteString(" ")
	if v.Initializer != nil {
		sb.WriteString(v.Initializer.String())
	} else {
		sb.WriteString("nil")
	}
	sb.WriteString(")")
	return sb.String()
}
