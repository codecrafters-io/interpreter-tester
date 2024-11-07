package lox

import (
	"io"
)

/*
program    -> declaration* EOF ;
declaration-> funDecl
			| varDecl
			| stmt ;
funDecl    -> "fun" function ;
function   -> IDENTIFIER "(" parameters? ")" block ;
parameters -> IDENTIFIER ( "," IDENTIFIER )* ;
varDecl    -> "var" IDENTIFIER ( "=" expression )? ";" ;
stmt       -> exprStmt
			| forStmt
			| ifStmt
			| printStmt
			| whileStmt
			| block ;
block      -> "{" declaration* "}" ;
exprStmt   -> expression ";" ;
forStmt    -> "for" "(" ( varDecl | exprStmt | ";" )
              expression? ";"
              expression? ")" statement ;
ifStmt     -> "if" "(" expression ")" statement ( "else" statement )? ;
printStmt  -> "print" expression ";" ;
whileStmt  -> "while" "(" expression ")" statement ;
expression -> assignment ;
assignment -> IDENTIFIER "=" assignment
			| logic_or ;
logic_or   -> logic_and ( "or" logic_and )* ;
logic_and  -> equality ( "and" equality )* ;
equality   -> comparison ( ( "!=" | "==") comparison )* ;
comparison -> term ( ( ">" | ">=" | "<" | "<=") term )* ;
term   -> factor ( ( "+" | "-" ) factor )* ;
factor -> unary ( ( "/" | "*" ) unary )* ;
unary      -> ( "!" | "-" ) unary | call ;
call       -> primary ( "(" arguments? ")" )* ;
arguments  -> expression ( "," expression )* ;
primary    -> "true" | "false" | "nil"
			| NUMBER | STRING
			| "(" expression ")"
			IDENTIFIER ;
*/

// Parser will transform an array of tokens to an AST.
// Use NewParser to create a new Parser.
type Parser struct {
	tokens  []Token
	current int
}

// NewParser creates a new parser
func NewParser(tokens []Token) Parser {
	return Parser{tokens, 0}
}

func (p *Parser) BasicParse(stdout, stderr io.Writer) Expr {
	expr, err := p.expression()
	if err != nil {
		LogParseError(err, stderr)
	}
	return expr
}

func (p *Parser) Parse(stdout, stderr io.Writer) []Stmt {
	statements := make([]Stmt, 0)
	for !p.isAtEnd() {
		stmt, err := p.declaration()
		if err != nil {
			LogParseError(err, stderr)
		}
		statements = append(statements, stmt)
	}
	return statements
}

func (p *Parser) declaration() (Stmt, error) {
	var stmt Stmt
	var err error

	checkError := func(stmt Stmt, err error) (Stmt, error) {
		if err != nil {
			p.synchronize()
			return nil, err
		}
		return stmt, nil
	}

	if p.match(VAR) {
		stmt, err = p.varDeclaration()
	} else if p.match(FUN) {
		stmt, err = p.funDeclaration("function")
	} else {
		stmt, err = p.statement()
	}

	return checkError(stmt, err)
}

func (p *Parser) varDeclaration() (Stmt, error) {
	name, err := p.consume(IDENTIFIER, "Expected variable name.")
	if err != nil {
		return nil, err
	}

	var initializer Expr
	if p.match(EQUAL) {
		initializer, err = p.expression()
		if err != nil {
			return nil, err
		}
	}
	_, err = p.consume(SEMICOLON, "Expected ';' after variable declaration.")
	if err != nil {
		return nil, err
	}
	return &Var{Name: name, Initializer: initializer}, nil
}

func (p *Parser) funDeclaration(kind string) (*Function, error) {
	name, err := p.consume(IDENTIFIER, "Expected variable name.")
	if err != nil {
		return nil, err
	}

	parameters, err := p.functionArguments(kind)
	if err != nil {
		return nil, err
	}

	_, err = p.consume(LEFTBRACE, "Expect '{' before "+kind+" body.")
	if err != nil {
		return nil, err
	}

	body, err := p.block()
	if err != nil {
		return nil, err
	}

	return &Function{Name: name, Params: parameters, Body: body}, nil
}

func (p *Parser) functionArguments(kind string) ([]Token, error) {
	_, err := p.consume(LEFTPAREN, "Expect '(' after "+kind+" name.")
	if err != nil {
		return nil, err
	}

	parameters := make([]Token, 0)
	if !p.check(RIGHTPAREN) {
		for {
			if len(parameters) >= 255 {
				return nil, MakeParseError(p.peek(), "Cannot have more than 255 parameters.")
			}

			param, err2 := p.consume(IDENTIFIER, "Expect parameter name.")
			if err2 != nil {
				return nil, err2
			}

			parameters = append(parameters, param)

			if !p.match(COMMA) {
				break
			}
		}
	}

	_, err = p.consume(RIGHTPAREN, "Expect ')' after parameters.")
	return parameters, err
}

func (p *Parser) statement() (Stmt, error) {
	if p.match(FOR) {
		return p.forStatement()
	} else if p.match(IF) {
		return p.ifStatement()
	} else if p.match(PRINT) {
		return p.printStatement()
	} else if p.match(WHILE) {
		return p.whileStatement()
	} else if p.match(LEFTBRACE) {
		statements, err := p.block()
		if err == nil {
			return &Block{Statements: statements}, nil
		}
		return nil, err
	}
	return p.expressionStatement()
}

func (p *Parser) block() ([]Stmt, error) {
	statements := make([]Stmt, 0)
	for !p.check(RIGHTBRACE) && !p.isAtEnd() {
		stmt, err := p.declaration()
		if err != nil {
			return nil, err
		}
		statements = append(statements, stmt)
	}
	_, err := p.consume(RIGHTBRACE, "Expected '}' after block.")
	if err != nil {
		return nil, err
	}
	return statements, nil
}

func (p *Parser) printStatement() (Stmt, error) {
	expr, err := p.expression()
	if err != nil {
		return nil, err
	}
	_, err = p.consume(SEMICOLON, "Expected ';' after value.")
	if err != nil {
		return nil, err
	}
	return &Print{Expression: expr}, nil
}

func (p *Parser) expressionStatement() (Stmt, error) {
	expr, err := p.expression()
	if err != nil {
		return nil, err
	}
	_, err = p.consume(SEMICOLON, "Expected ';' after value.")
	if err != nil {
		return nil, err
	}
	return &Expression{Expression: expr}, nil
}

func (p *Parser) forStatement() (Stmt, error) {
	if _, err := p.consume(LEFTPAREN, "Expected '(' after 'for'."); err != nil {
		return nil, err
	}

	var err error
	var initializer Stmt

	if p.match(SEMICOLON) {
		initializer = nil
	} else if p.match(VAR) {
		initializer, err = p.varDeclaration()
	} else {
		initializer, err = p.expressionStatement()
	}

	if err != nil {
		return nil, err
	}

	var condition Expr
	if !p.check(SEMICOLON) {
		condition, err = p.expression()
		if err != nil {
			return nil, err
		}
	}
	_, err = p.consume(SEMICOLON, "Expected ';' after loop condition.")
	if err != nil {
		return nil, err
	}

	var increment Expr
	if !p.check(RIGHTPAREN) {
		increment, err = p.expression()
		if err != nil {
			return nil, err
		}
	}
	_, err = p.consume(RIGHTPAREN, "Expected ')' after for clauses.")
	if err != nil {
		return nil, err
	}

	body, err := p.statement()
	if err != nil {
		return nil, err
	}

	if increment != nil {
		body = &Block{Statements: []Stmt{body, &Expression{Expression: increment}}}
	}

	if condition == nil {
		condition = &Literal{Value: true}
	}
	body = &While{Condition: condition, Statement: body}

	if initializer != nil {
		body = &Block{Statements: []Stmt{initializer, body}}
	}

	return body, nil
}

func (p *Parser) ifStatement() (Stmt, error) {
	if _, err := p.consume(LEFTPAREN, "Expected '(' after 'if'."); err != nil {
		return nil, err
	}
	condition, err := p.expression()
	if err != nil {
		return nil, err
	}
	if _, err := p.consume(RIGHTPAREN, "Expected ')' after 'if' condition."); err != nil {
		return nil, err
	}
	thenBranch, err := p.statement()
	if err != nil {
		return nil, err
	}
	if p.match(ELSE) {
		elseBranch, err := p.statement()
		if err != nil {
			return nil, err
		}
		return &If{Condition: condition, ThenBranch: thenBranch, ElseBranch: elseBranch}, nil
	}
	return &If{Condition: condition, ThenBranch: thenBranch}, nil
}

func (p *Parser) whileStatement() (Stmt, error) {
	_, err := p.consume(LEFTPAREN, "Expected '(' after 'while'.")
	if err != nil {
		return nil, err
	}
	condition, err := p.expression()
	if err != nil {
		return nil, err
	}
	_, err = p.consume(RIGHTPAREN, "Expected ')' after condition.")
	if err != nil {
		return nil, err
	}
	body, err := p.statement()
	if err != nil {
		return nil, err
	}
	return &While{Condition: condition, Statement: body}, nil
}

func (p *Parser) expression() (Expr, error) {
	return p.assignment()
}

func (p *Parser) assignment() (Expr, error) {
	expr, err := p.or()
	if err != nil {
		return nil, err
	}

	if p.match(EQUAL) {
		equals := p.previous()
		value, err := p.assignment()
		if err != nil {
			return nil, err
		}

		if variable, ok := expr.(*Variable); ok {
			return &Assign{Name: variable.Name, Value: value}, nil
		}
		return nil, MakeParseError(equals, "Invalid assignment target.")
	}
	return expr, nil
}

func (p *Parser) or() (Expr, error) {
	expr, err := p.and()
	if err != nil {
		return nil, err
	}

	for p.match(OR) {
		operator := p.previous()
		right, err := p.and()
		if err != nil {
			return nil, err
		}
		expr = &Logical{Left: expr, Operator: operator, Right: right}
	}
	return expr, nil
}

func (p *Parser) and() (Expr, error) {
	expr, err := p.equality()
	if err != nil {
		return nil, err
	}
	for p.match(AND) {
		operator := p.previous()
		right, err := p.equality()
		if err != nil {
			return nil, err
		}
		expr = &Logical{Left: expr, Operator: operator, Right: right}
	}
	return expr, nil
}

func (p *Parser) equality() (Expr, error) {
	expr, err := p.comparison()
	if err != nil {
		return nil, err
	}

	for p.match(BANGEQUAL, EQUALEQUAL) {
		operator := p.previous()
		right, err := p.comparison()
		if err != nil {
			return nil, err
		}
		expr = &Binary{Left: expr, Operator: operator, Right: right}
	}

	return expr, nil
}

func (p *Parser) comparison() (Expr, error) {
	expr, err := p.term()
	if err != nil {
		return nil, err
	}

	for p.match(GREATER, GREATEREQUAL, LESS, LESSEQUAL) {
		operator := p.previous()
		right, err := p.term()
		if err != nil {
			return nil, err
		}
		expr = &Binary{Left: expr, Operator: operator, Right: right}
	}

	return expr, nil
}

func (p *Parser) term() (Expr, error) {
	expr, err := p.factor()
	if err != nil {
		return nil, err
	}

	for p.match(PLUS, MINUS) {
		operator := p.previous()
		right, err := p.factor()
		if err != nil {
			return nil, err
		}
		expr = &Binary{Left: expr, Operator: operator, Right: right}
	}

	return expr, nil
}

func (p *Parser) factor() (Expr, error) {
	expr, err := p.unary()
	if err != nil {
		return nil, err
	}

	for p.match(STAR, SLASH) {
		operator := p.previous()
		right, err := p.unary()
		if err != nil {
			return nil, err
		}
		expr = &Binary{Left: expr, Operator: operator, Right: right}
	}

	return expr, nil
}

func (p *Parser) unary() (Expr, error) {
	if p.match(BANG, MINUS) {
		operator := p.previous()
		right, err := p.unary()
		if err != nil {
			return nil, err
		}
		return &Unary{Operator: operator, Right: right}, nil
	}

	return p.call()
}

func (p *Parser) call() (Expr, error) {
	expr, err := p.primary()
	if err != nil {
		return nil, err
	}

	for {
		if p.match(LEFTPAREN) {
			expr, err = p.finishCall(expr)
			if err != nil {
				return nil, err
			}
		} else {
			break
		}
	}

	return expr, nil
}

func (p *Parser) finishCall(callee Expr) (Expr, error) {
	// Parse argument list
	args := make([]Expr, 0)
	if !p.check(RIGHTPAREN) {
		for {
			arg, err := p.assignment() // we don't want the comma operator here
			if err != nil {
				return nil, err
			}
			if len(args) >= 255 {
				return nil, MakeParseError(p.peek(), "Cannot have more than 255 arguments.")
			}
			args = append(args, arg)
			if !p.match(COMMA) {
				break
			}
		}
	}

	paren, err := p.consume(RIGHTPAREN, "Expected ')' after arguments.")
	if err != nil {
		return nil, err
	}
	return &Call{Callee: callee, Paren: paren, Arguments: args}, nil
}

func (p *Parser) primary() (Expr, error) {
	if p.match(FALSE) {
		return &Literal{Value: false}, nil
	} else if p.match(TRUE) {
		return &Literal{Value: true}, nil
	} else if p.match(NIL) {
		return &Literal{Value: nil}, nil
	} else if p.match(NUMBER, STRING) {
		return &Literal{Value: p.previous().Literal}, nil
	} else if p.match(LEFTPAREN) {
		expr, err := p.expression()
		if err != nil {
			return nil, err
		}
		_, err = p.consume(RIGHTPAREN, "Expect ')' after expression.")
		if err != nil {
			return nil, err
		}
		return &Grouping{Expression: expr}, nil
	} else if p.match(IDENTIFIER) {
		return &Variable{Name: p.previous()}, nil
	}
	return nil, MakeParseError(p.peek(), "Expect expression.")
}

func (p *Parser) consume(tp Type, message string) (Token, error) {
	if p.check(tp) {
		return p.advance(), nil
	}
	return p.previous(), MakeParseError(p.peek(), message)
}

func (p *Parser) advance() Token {
	if !p.isAtEnd() {
		p.current++
	}
	return p.previous()
}

func (p *Parser) match(types ...Type) bool {
	for _, tp := range types {
		if p.check(tp) {
			p.advance()
			return true
		}
	}
	return false
}

// check checks if the next token is of the given type
func (p *Parser) check(tp Type) bool {
	if p.isAtEnd() {
		return false
	}
	return p.peek().Type == tp
}

func (p *Parser) isAtEnd() bool {
	return p.peek().Type == EOF
}

// peek returns the next token
func (p *Parser) peek() Token {
	return p.tokens[p.current]
}

// peek returns the current token
func (p *Parser) previous() Token {
	return p.tokens[p.current-1]
}

func (p *Parser) synchronize() {
	p.advance()
	for !p.isAtEnd() {
		if p.previous().Type == SEMICOLON {
			return
		}
		switch p.peek().Type {
		case VAR, PRINT:
			return
		}
		p.advance()
	}
}
