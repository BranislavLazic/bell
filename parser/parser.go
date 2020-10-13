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
	program.Expressions = []ast.Expression{}
	for p.curToken.Type != token.EOF {
		var expr ast.Expression
		// Start by checking whether the current token is StartExpression.
		// If so, start parsing an expression.
		if p.curToken.Type == token.StartExpression {
			p.parensCount++
			expr = p.parseExpression()
		}
		if expr != nil {
			program.Expressions = append(program.Expressions, expr)
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
	case token.EQUAL:
		expr = p.parseArithmeticExpression()
	case token.NotEqual:
		expr = p.parseArithmeticExpression()
	case token.AND:
		expr = p.parseLogicalExpression()
	case token.OR:
		expr = p.parseLogicalExpression()
	case token.NOT:
		expr = p.parseLogicalExpression()
	case token.BOOL:
		expr = p.parseBoolLiteral()
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
	if rightExpr == nil {
		p.Errors = append(p.Errors, "Missing right expression for an arithmetic operation.")
		return nil
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
	case token.EQUAL:
		expr = &ast.EqualExpression{Token: tok, LeftExpr: leftExpr, RightExpr: rightExpr}
	case token.NotEqual:
		expr = &ast.NotEqualExpression{Token: tok, LeftExpr: leftExpr, RightExpr: rightExpr}
	}
	return expr
}
func (p *Parser) parseLogicalExpression() ast.Expression {
	p.nextToken()
	tok := p.curToken
	leftExpr := p.parseExpression()
	rightExpr := p.parseExpression()
	if tok.Type == token.NOT && leftExpr != nil && rightExpr == nil {
		return &ast.NotExpression{Token: tok, Expr: leftExpr}
	}
	if rightExpr == nil {
		p.Errors = append(p.Errors, "Missing right expression for a boolean operation.")
		return nil
	}
	var expr ast.Expression
	switch tok.Type {
	case token.AND:
		expr = &ast.AndExpression{Token: tok, LeftExpr: leftExpr, RightExpr: rightExpr}
	case token.OR:
		expr = &ast.OrExpression{Token: tok, LeftExpr: leftExpr, RightExpr: rightExpr}
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

func (p *Parser) parseBoolLiteral() *ast.BooleanLiteral {
	p.nextToken()
	value, err := strconv.ParseBool(p.curToken.Literal)
	if err != nil {
		p.Errors = append(p.Errors, "Failed to parse a value to bool.")
	}
	return &ast.BooleanLiteral{Token: p.curToken, Value: value}
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
