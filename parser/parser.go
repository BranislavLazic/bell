package parser

import (
	"fmt"
	"strconv"

	"github.com/branislavlazic/bell/ast"
	"github.com/branislavlazic/bell/lexer"
	"github.com/branislavlazic/bell/token"
)

type Parser struct {
	lxr       *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
	Errors    []string
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
			expr = p.parseExpression()
		}
		// If there are any errors after parsing an expression,
		// break any further parsing.
		if len(p.Errors) > 0 {
			break
		}
		if expr != nil {
			program.Expressions = append(program.Expressions, expr)
		}
		p.nextToken()
	}
	if len(program.Expressions) == 0 && len(p.Errors) == 0 {
		p.Errors = append(p.Errors, "No expression given.")
	}
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
	case token.MODULO:
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
	case token.GreaterThan:
		expr = p.parseOperationExpression()
	case token.LessThan:
		expr = p.parseOperationExpression()
	case token.GreaterThanEqual:
		expr = p.parseOperationExpression()
	case token.LessThanEqual:
		expr = p.parseOperationExpression()
	case token.LET:
		expr = p.parseLetExpression()
	case token.IF:
		expr = p.parseIfExpression()
	case token.LIST:
		expr = p.parseOperationExpression()
	case token.IDENT:
		expr = p.parseIdentifier()
	case token.BOOL:
		expr = p.parseBoolLiteral()
	case token.INT:
		expr = p.parseIntLiteral()
	case token.StartExpression:
		p.nextToken()
		expr = p.parseExpression()
	case token.EndExpression:
		p.nextToken()
		expr = p.parseExpression()
	case token.StartParamList:
		p.nextToken()
		expr = p.parseExpression()
	case token.EndParamList:
		p.nextToken()
		break
	case token.EOL:
		p.nextToken()
		expr = p.parseExpression()
	case token.ILLEGAL:
		p.Errors = append(
			p.Errors,
			fmt.Sprintf("Illegal character '%s' found at index %d.", p.peekToken.Literal, p.lxr.Position-1),
		)
		break
	default:
		break
	}
	return expr
}

// Parse arithmetic, relational, logic operations and list
func (p *Parser) parseOperationExpression() ast.Expression {
	p.nextToken()
	tok := p.curToken
	var exprs []ast.Expression
	leadingExpr := p.parseExpression()
	// If the prefix is "not", then only a single expression
	// must be present.
	isNotOperation := tok.Type == token.NOT && leadingExpr != nil
	if isNotOperation && p.peekToken.Type != token.EndExpression {
		p.nextToken()
		p.Errors = append(
			p.Errors,
			"'not' operation either contains more than one expression or lacks a closing parentheses.",
		)
		return nil
	}
	// If the prefix token is "not", one expression is present,
	// and there are no other expressions, then it's a logical not expression.
	if isNotOperation {
		p.nextToken()
		return &ast.NotExpression{Token: tok, Expr: leadingExpr}
	}
	if leadingExpr == nil {
		p.Errors = append(p.Errors, fmt.Sprintf("Missing at least one expression for operation '%s'.", tok.Literal))
		return nil
	}
	exprs = append(exprs, leadingExpr)
	// Parse expressions until '(' is the next token.
	for p.peekToken.Type != token.EndExpression {
		if p.isPeekEOF() || p.isPeekIllegal() || p.isPeekOperator() {
			return nil
		}
		exprs = append(exprs, p.parseExpression())
	}
	p.nextToken()
	// If the prefix token is "-", one expression is present,
	// then it's a negated expression.
	if tok.Type == token.SUBTRACT && len(exprs) == 1 {
		return &ast.NegativeValueExpression{Token: tok, Expr: leadingExpr}
	}
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
	case token.MODULO:
		expr = &ast.ModuloExpression{Token: tok, Exprs: exprs}
	case token.EQUAL:
		expr = &ast.EqualExpression{Token: tok, Exprs: exprs}
	case token.NotEqual:
		expr = &ast.NotEqualExpression{Token: tok, Exprs: exprs}
	case token.AND:
		expr = &ast.AndExpression{Token: tok, Exprs: exprs}
	case token.OR:
		expr = &ast.OrExpression{Token: tok, Exprs: exprs}
	case token.GreaterThan:
		expr = &ast.GreaterThanExpression{Token: tok, Exprs: exprs}
	case token.LessThan:
		expr = &ast.LessThanExpression{Token: tok, Exprs: exprs}
	case token.GreaterThanEqual:
		expr = &ast.GreaterThanEqualExpression{Token: tok, Exprs: exprs}
	case token.LessThanEqual:
		expr = &ast.LessThanEqualExpression{Token: tok, Exprs: exprs}
	case token.LIST:
		expr = &ast.ListExpression{Token: tok, Exprs: exprs}
	}
	return expr
}

func (p *Parser) parseLetExpression() ast.Expression {
	p.nextToken()
	letTok := p.curToken
	if p.peekToken.Type != token.IDENT {
		p.Errors = append(p.Errors, "'let' should be followed by an identifier.")
		return nil
	}
	ident := p.parseIdentifier()
	var params []*ast.Identifier
	isFunction := false
	if p.peekToken.Type == token.StartParamList {
		isFunction = true
		for p.curToken.Type != token.EndParamList {
			if p.isPeekEOF() || p.isPeekIllegal() || p.isPeekOperator() {
				return nil
			}
			switch ex := p.parseExpression().(type) {
			// Collect only identifiers
			case *ast.Identifier:
				params = append(params, ex)
			default:
				break
			}
		}
	}
	expr := p.parseExpression()
	if expr == nil {
		p.Errors = append(p.Errors, fmt.Sprintf("Missing assigned expression."))
		return nil
	}
	// Check whether the expression is closed.
	// If not, then anything following the expression is
	// an illegal token except of EOF (which gives Unexpected EOF error).
	if p.peekToken.Type != token.EndExpression {
		if !p.isPeekEOF() {
			p.Errors = append(
				p.Errors,
				fmt.Sprintf("Illegal character '%s' found at index %d.", p.peekToken.Literal, p.lxr.Position-1),
			)
			return nil
		}
	}
	p.nextToken()
	if isFunction {
		return &ast.Function{Token: letTok, Identifier: ident.(*ast.Identifier), Params: params, Body: expr}
	}
	return &ast.LetExpression{Token: letTok, Identifier: ident.(*ast.Identifier), Expr: expr}
}

func (p *Parser) parseIfExpression() ast.Expression {
	p.nextToken()
	ifTok := p.curToken
	cond := p.parseExpression()
	if cond == nil {
		p.Errors = append(p.Errors, fmt.Sprintf("If expression is missing condition."))
		return nil
	}
	expr := p.parseExpression()
	if expr == nil {
		p.Errors = append(p.Errors, fmt.Sprintf("If expression is missing then expression."))
		return nil
	}
	elseExpr := p.parseExpression()
	if elseExpr == nil {
		p.Errors = append(p.Errors, fmt.Sprintf("If expression is missing else expression."))
		return nil
	}
	if p.peekToken.Type != token.EndExpression {
		if !p.isPeekEOF() {
			p.Errors = append(
				p.Errors,
				fmt.Sprintf("Illegal character '%s' found at index %d.", p.peekToken.Literal, p.lxr.Position-1),
			)
			return nil
		}
	}
	p.nextToken()
	return &ast.IfExpression{Token: ifTok, Condition: cond, ThenExpr: expr, ElseExpr: elseExpr}
}

func (p *Parser) parseIdentifier() ast.Expression {
	// If an identifier is at the beginning of an
	// expression, then the expression is treated
	// as a function call.
	if p.curToken.Type == token.StartExpression {
		p.nextToken()
		ident := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
		var args []ast.Expression
		for p.peekToken.Type != token.EndExpression {
			if p.isPeekEOF() || p.isPeekIllegal() || p.isPeekOperator() {
				break
			}
			args = append(args, p.parseExpression())
		}
		p.nextToken()
		return &ast.CallFunction{Identifier: ident, Args: args}
	}
	p.nextToken()
	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
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

func (p *Parser) isPeekEOF() bool {
	if p.peekToken.Type == token.EOF {
		p.Errors = append(p.Errors, fmt.Sprintf("Unexpected EOF at index %d.", p.lxr.Position-1))
		return true
	}
	return false
}

func (p *Parser) isPeekIllegal() bool {
	if p.peekToken.Type == token.ILLEGAL {
		p.Errors = append(p.Errors, fmt.Sprintf("Illegal character '%s' found at index %d.", p.peekToken.Literal, p.lxr.Position-1))
		return true
	}
	return false
}

func (p *Parser) isPeekOperator() bool {
	for _, op := range token.OperatorLiterals {
		if p.peekToken.Literal == op {
			p.Errors = append(p.Errors, fmt.Sprintf("Illegal use of operator '%s' at index %d.", p.peekToken.Literal, p.lxr.Position-1))
			return true
		}
	}
	return false
}
