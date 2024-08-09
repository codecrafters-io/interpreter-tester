package lox

/*
program    -> declaration* EOF ;
declaration-> varDecl
			| stmt ;
varDecl    -> "var" IDENTIFIER ( "=" expression )? ";" ;
stmt       -> exprStmt
			| printStmt ;
			| block ;
block      -> "{" declaration* "}" ;
exprStmt   -> expression ";" ;
printStmt  -> "print" expression ";" ;
expression -> assignment ;
assignment -> IDENTIFIER "=" assignment
			| equality ;
equality   -> comparison ( ( "!=" | "==") comparison )* ;
comparison -> term ( ( ">" | ">=" | "<" | "<=") term )* ;
term   -> factor ( ( "+" | "-" ) factor )* ;
factor -> unary ( ( "/" | "*" ) unary )* ;
unary      -> ( "!" | "-" ) unary | primary ;
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

func (p *Parser) Parse() (Expr, error) {
	return p.expression()
}

func (p *Parser) expression() (Expr, error) {
	return p.equality()
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

	return p.primary()
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
