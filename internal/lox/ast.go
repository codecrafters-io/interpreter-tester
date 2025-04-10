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
		sb.WriteString(FormatFloat(l.Value.(float64)))
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

// Control Flow

// If is the classic if statement
type If struct {
	Stmt
	Condition  Expr
	ThenBranch Stmt
	ElseBranch Stmt
}

// String pretty prints the if statement
func (i *If) String() string {
	var sb strings.Builder
	sb.WriteString("(")
	sb.WriteString("if")
	sb.WriteString(" ")
	sb.WriteString(i.Condition.String())
	sb.WriteString(" ")
	sb.WriteString(i.ThenBranch.String())
	sb.WriteString(" ")
	sb.WriteString(i.ElseBranch.String())
	sb.WriteString(")")
	return sb.String()
}

// While is the classic while statement
type While struct {
	Stmt
	Condition Expr
	Statement Stmt
}

// String pretty prints the while statement
func (w *While) String() string {
	var sb strings.Builder
	sb.WriteString("(")
	sb.WriteString("while")
	sb.WriteString(" ")
	sb.WriteString(w.Condition.String())
	sb.WriteString(" ")
	sb.WriteString(w.Statement.String())
	sb.WriteString(")")
	return sb.String()
}

// Logical is used for "or" and "and" operators.
type Logical struct {
	Expr
	Left     Expr
	Operator Token
	Right    Expr
}

// String pretty prints the unary operator
func (l *Logical) String() string {
	var sb strings.Builder
	sb.WriteString("(")
	sb.WriteString(l.Operator.Lexeme)
	sb.WriteString(" ")
	sb.WriteString(l.Left.String())
	sb.WriteString(" ")
	sb.WriteString(l.Right.String())
	sb.WriteString(")")
	return sb.String()
}

// Functions

// Call is the node of a function call
type Call struct {
	Callee    Expr
	Paren     Token
	Arguments []Expr
}

// String pretty prints the call operator
func (c *Call) String() string {
	var sb strings.Builder
	sb.WriteString("(")
	sb.WriteString("call")
	sb.WriteString(" ")
	sb.WriteString(c.Callee.String())
	sb.WriteString(" ")
	for _, e := range c.Arguments {
		sb.WriteString(e.String())
		sb.WriteString(" ")
	}
	sb.WriteString(")")
	return sb.String()
}

// Function is the function definition node
type Function struct {
	Name   Token
	Params []Token
	Body   []Stmt
}

// String pretty prints the function
func (f *Function) String() string {
	var sb strings.Builder
	sb.WriteString("(")
	sb.WriteString("fun")
	sb.WriteString(" ")
	sb.WriteString(f.Name.String())
	sb.WriteString(" ")
	sb.WriteString("(")
	for _, p := range f.Params {
		sb.WriteString(p.String())
		sb.WriteString(" ")
	}
	sb.WriteString(")")
	sb.WriteString(" ")
	sb.WriteString("(")
	for _, stmt := range f.Body {
		sb.WriteString(stmt.String())
		sb.WriteString(" ")
	}
	sb.WriteString(")")
	sb.WriteString(")")
	return sb.String()
}

// Return is used to return from a function
type Return struct {
	Stmt
	Keyword Token
	Value   Expr
}

// String pretty prints the function
func (r *Return) String() string {
	var sb strings.Builder
	sb.WriteString("(")
	sb.WriteString("return")
	sb.WriteString(" ")
	sb.WriteString(r.Value.String())
	sb.WriteString(" ")
	sb.WriteString(")")
	return sb.String()
}

// Class node
type Class struct {
	Stmt
	Name       Token
	Methods    []*Function
	EnvIndex   int
	SuperClass *Variable
}

// String pretty prints the class
func (c *Class) String() string {
	var sb strings.Builder
	sb.WriteString("(")
	sb.WriteString("class")
	sb.WriteString("")
	sb.WriteString(c.Name.Lexeme)
	sb.WriteString("")
	for _, f := range c.Methods {
		sb.WriteString(f.String())
		sb.WriteString(" ")
	}
	sb.WriteString(")")
	return sb.String()
}

// Get is used for property access
type Get struct {
	Expr
	Name       Token
	Expression Expr
}

// String pretty prints the class
func (g *Get) String() string {
	var sb strings.Builder
	sb.WriteString("(")
	sb.WriteString(".")
	sb.WriteString(" ")
	sb.WriteString(g.Expression.String())
	sb.WriteString(" ")
	sb.WriteString(g.Name.Lexeme)
	sb.WriteString(")")
	return sb.String()
}

// Set is used for writing to a property
type Set struct {
	Expr
	Object Expr
	Name   Token
	Value  Expr
}

// String pretty prints the setter
func (s *Set) String() string {
	var sb strings.Builder
	sb.WriteString("(")
	sb.WriteString("set")
	sb.WriteString(" ")
	sb.WriteString(s.Object.String())
	sb.WriteString(" ")
	sb.WriteString(s.Name.Lexeme)
	sb.WriteString(" ")
	sb.WriteString(s.Value.String())
	sb.WriteString(")")
	return sb.String()
}

// This is used to access the current instance
type This struct {
	Expr
	Keyword  Token
	EnvIndex int
	EnvDepth int
}

// String pretty prints the setter
func (t *This) String() string {
	var sb strings.Builder
	sb.WriteString("(")
	sb.WriteString("this")
	sb.WriteString(" ")
	sb.WriteString(t.Keyword.Lexeme)
	sb.WriteString(" ")
	sb.WriteString(")")
	return sb.String()
}

// Super is used to access the superclass
type Super struct {
	Expr
	Keyword  Token
	Method   Token
	EnvIndex int
	EnvDepth int
}

// String pretty prints the 'super' node
func (s *Super) String() string {
	var sb strings.Builder
	sb.WriteString("(")
	sb.WriteString("super")
	sb.WriteString(" ")
	sb.WriteString(s.Method.Lexeme)
	sb.WriteString(" ")
	sb.WriteString(")")
	return sb.String()
}

func FormatFloat(num float64) string {
	str := fmt.Sprintf("%f", num)
	parts := strings.Split(str, ".")
	if len(parts) != 2 {
		return str
	}
	integerPart := parts[0]
	decimalPart := parts[1]

	// Remove trailing zeros from the decimal part
	decimalPart = strings.TrimRight(decimalPart, "0")

	// Ensure at least one decimal place
	if decimalPart == "" {
		decimalPart = "0"
	}

	return integerPart + "." + decimalPart
}
