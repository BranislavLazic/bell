package parser

import (
	"fmt"
	"math"
	"strconv"

	"github.com/branislavlazic/bell/ast"
	"github.com/branislavlazic/bell/lexer"
	"github.com/branislavlazic/bell/token"
)

type Parser struct {
	lxr         *lexer.Lexer
	curToken    token.Token
	peekToken   token.Token
	parensCount int
	Errors      []string
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{lxr: l, Errors: []string{}}
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}
	for p.curToken.Type != token.EOF {
		var stmt *ast.Statement
		// Start by checking whether the current token is StartExpression.
		// If so, start parsing an expressions within the statement.
		if p.curToken.Type == token.StartExpression {
			p.parensCount++
			stmt = p.parseStatement()
		}
		if stmt != nil {
			program.Statements = append(program.Statements, *stmt)
		}
		p.nextToken()
	}
	if p.parensCount > 0 {
		p.Errors = append(p.Errors, fmt.Sprintf("Missing %d closing parenteses", p.parensCount))
	}
	if p.parensCount < 0 {
		p.Errors = append(p.Errors, fmt.Sprintf("Missing %d opening parenteses", int(math.Abs(float64(p.parensCount)))))
	}
	return program
}

func (p *Parser) parseStatement() *ast.Statement {
	stmt := &ast.Statement{}
	stmt.Expr = p.parseExpression()
	return stmt
}

func (p *Parser) parseExpression() ast.Expression {
	var expr ast.Expression
	switch p.peekToken.Type {
	case token.ADD:
		expr = p.parseArithmeticExpression()
	case token.SUBTRACT:
		expr = p.parseArithmeticExpression()
	case token.MULTIPLY:
		expr = p.parseArithmeticExpression()
	case token.DIVIDE:
		expr = p.parseArithmeticExpression()
	case token.INT:
		expr = p.parseIntLiteral()
	case token.StartExpression:
		p.parensCount++
		expr = p.parseNextExpression()
	case token.EndExpression:
		p.parensCount--
		expr = p.parseNextExpression()
	default:
		break
	}
	// Check for EndExpression token
	p.finalizeExpression()
	return expr
}

func (p *Parser) parseNextExpression() ast.Expression {
	p.nextToken()
	expr := p.parseExpression()
	return expr
}

func (p *Parser) parseArithmeticExpression() ast.Expression {
	p.nextToken()
	tok := p.curToken
	leftExpr := p.parseExpression()
	rightExpr := p.parseExpression()
	// If the prefix token is "-", the left expression is present,
	// and the right expression is nil, then it's a negated expression.
	if tok.Type == token.SUBTRACT && leftExpr != nil && rightExpr == nil {
		return &ast.NegativeValueExpression{Token: tok, Expr: leftExpr}
	}
	var expr ast.Expression
	switch tok.Type {
	case token.ADD:
		expr = &ast.AddExpression{Token: tok, LeftExpr: leftExpr, RightExpr: rightExpr}
	case token.SUBTRACT:
		expr = &ast.SubtractExpression{Token: tok, LeftExpr: leftExpr, RightExpr: rightExpr}
	case token.MULTIPLY:
		expr = &ast.MultiplyExpression{Token: tok, LeftExpr: leftExpr, RightExpr: rightExpr}
	case token.DIVIDE:
		expr = &ast.DivideExpression{Token: tok, LeftExpr: leftExpr, RightExpr: rightExpr}
	}
	return expr
}

func (p *Parser) parseIntLiteral() *ast.IntegerLiteral {
	p.nextToken()
	value, err := strconv.Atoi(p.curToken.Literal)
	if err != nil {
		p.Errors = append(p.Errors, "Failed to parse a value to integer.")
	}
	return &ast.IntegerLiteral{Token: p.curToken, Value: int64(value)}
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.lxr.NextToken()
}

func (p *Parser) finalizeExpression() {
	for p.peekToken.Type == token.EndExpression {
		p.parensCount--
		p.nextToken()
	}
}
