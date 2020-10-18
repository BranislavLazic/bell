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
	p.checkParenthesesCount()
	return program
}

func (p *Parser) parseExpression() ast.Expression {
	var expr ast.Expression
	switch p.peekToken.Type {
	case token.ADD:
		expr = p.parseOperationExpression()
	case token.SUBTRACT:
		expr = p.parseOperationExpression()
	case token.MULTIPLY:
		expr = p.parseOperationExpression()
	case token.DIVIDE:
		expr = p.parseOperationExpression()
	case token.EQUAL:
		expr = p.parseOperationExpression()
	case token.NotEqual:
		expr = p.parseOperationExpression()
	case token.AND:
		expr = p.parseOperationExpression()
	case token.OR:
		expr = p.parseOperationExpression()
	case token.NOT:
		expr = p.parseOperationExpression()
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

// Parse all mathematical operations
func (p *Parser) parseOperationExpression() ast.Expression {
	p.nextToken()
	tok := p.curToken
	var exprs []ast.Expression
	leadingExpr := p.parseExpression()
	// If the prefix is "not", then only a single expression
	// must be present.
	isNotOperation := tok.Type == token.NOT && leadingExpr != nil
	if isNotOperation && p.curToken.Type != token.EndExpression {
		p.Errors = append(p.Errors, "'not' operation contains more than one expression or lacks a closing parentheses.")
		return nil
	}
	// If the prefix token is "not", one expression is present,
	// and there are no other expressions, then it's a logical not expression.
	if isNotOperation {
		return &ast.NotExpression{Token: tok, Expr: leadingExpr}
	}
	if leadingExpr == nil {
		p.Errors = append(p.Errors, fmt.Sprintf("Missing at least one expression for operation '%s'.", tok.Literal))
		return nil
	}
	exprs = append(exprs, leadingExpr)
	for p.curToken.Type == token.EndExpression && (p.peekToken.Type == token.INT || p.peekToken.Type == token.BOOL || p.peekToken.Type == token.StartExpression) {
		exprs = append(exprs, p.parseExpression())
	}
	for p.curToken.Type != token.EndExpression && p.peekToken.Type != token.ILLEGAL && p.peekToken.Type != token.EOF {
		exprs = append(exprs, p.parseExpression())
	}
	// If the prefix token is "-", one expression is present,
	// then it's a negated expression.
	if tok.Type == token.SUBTRACT && len(exprs) == 1 {
		return &ast.NegativeValueExpression{Token: tok, Expr: leadingExpr}
	}
	return mapToExpression(tok, exprs)
}

func mapToExpression(tok token.Token, exprs []ast.Expression) ast.Expression {
	var expr ast.Expression
	switch tok.Type {
	case token.ADD:
		expr = &ast.AddExpression{Token: tok, Exprs: exprs}
	case token.SUBTRACT:
		expr = &ast.SubtractExpression{Token: tok, Exprs: exprs}
	case token.MULTIPLY:
		expr = &ast.MultiplyExpression{Token: tok, Exprs: exprs}
	case token.DIVIDE:
		expr = &ast.DivideExpression{Token: tok, Exprs: exprs}
	case token.EQUAL:
		expr = &ast.EqualExpression{Token: tok, Exprs: exprs}
	case token.NotEqual:
		expr = &ast.NotEqualExpression{Token: tok, Exprs: exprs}
	case token.AND:
		expr = &ast.AndExpression{Token: tok, Exprs: exprs}
	case token.OR:
		expr = &ast.OrExpression{Token: tok, Exprs: exprs}
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

func (p *Parser) checkParenthesesCount() {
	if len(p.Errors) == 0 {
		if p.parensCount > 0 {
			p.Errors = append(p.Errors, fmt.Sprintf("Missing %d closing parentheses", p.parensCount))
		}
		if p.parensCount < 0 {
			p.Errors = append(p.Errors, fmt.Sprintf("Missing %d opening parentheses", int(math.Abs(float64(p.parensCount)))))
		}
	}
}
