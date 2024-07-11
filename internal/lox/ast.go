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
	// fmt.Println("Binary")
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
	// fmt.Println("Grouping")
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
	// fmt.Println("Literal")
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
	// fmt.Println("Unary")
	var sb strings.Builder
	sb.WriteString("(")
	sb.WriteString(u.Operator.Lexeme)
	sb.WriteString(" ")
	sb.WriteString(u.Right.String())
	sb.WriteString(")")
	return sb.String()
}

// Expression statement
type Expression struct {
	Expression Expr
}

// String pretty prints the expression statement
func (e *Expression) String() string {
	// fmt.Println("Expression")
	var sb strings.Builder
	sb.WriteString("(")
	sb.WriteString(e.Expression.String())
	sb.WriteString(")")
	return sb.String()
}
