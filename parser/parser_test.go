package parser

import (
	"testing"

	"github.com/branislavlazic/bell/ast"
	"github.com/branislavlazic/bell/lexer"
	"github.com/branislavlazic/bell/token"
)

func TestParser_ParseLiteral(t *testing.T) {
	input := `(
2)
	(true)`
	l := lexer.New(input)
	p := New(l)
	prog := p.ParseProgram()

	if len(prog.Expressions) != 2 {
		t.Fatalf("test - wrong number of expressions. expected=%d, got=%d", 2, len(prog.Expressions))
	}
	if len(p.Errors) != 0 {
		t.Fatalf("test - error list should be empty. expected=%d, got=%d", 0, len(p.Errors))
	}
	intLiteralExpr := prog.Expressions[0].(*ast.IntegerLiteral)
	if intLiteralExpr.Value != 2 {
		t.Fatalf("test - wrong value of integer literal. expected=%d, got=%d", 2, intLiteralExpr.Value)
	}

	boolLiteralExpr := prog.Expressions[1].(*ast.BooleanLiteral)
	if boolLiteralExpr.Value != true {
		t.Fatalf("test - wrong value of boolean literal. expected=%t, got=%t", true, boolLiteralExpr.Value)
	}
}

func TestParser_ParseNegativeLiteral(t *testing.T) {
	input := `(- 2)`
	l := lexer.New(input)
	p := New(l)
	prog := p.ParseProgram()

	if len(prog.Expressions) != 1 {
		t.Fatalf("test - wrong number of expressions. expected=%d, got=%d", 1, len(prog.Expressions))
	}
	if len(p.Errors) != 0 {
		t.Fatalf("test - error list should be empty. expected=%d, got=%d", 0, len(p.Errors))
	}
	negativeValueExpr := prog.Expressions[0].(*ast.NegativeValueExpression)

	if negativeValueExpr.Token.Type != token.SUBTRACT {
		t.Fatalf("test - wrong value for the token. expected=%s, got=%s", token.SUBTRACT, negativeValueExpr.Token.Type)
	}
	intLit := negativeValueExpr.Expr.(*ast.IntegerLiteral)
	if intLit.Value != 2 {
		t.Fatalf("test - wrong value of integer literal. expected=%d, got=%d", 2, intLit.Value)
	}

}

func TestParser_ParseTwoPlusExpression(t *testing.T) {
	input := `(+ 2 3)
	(- 7 9)`
	l := lexer.New(input)
	p := New(l)
	prog := p.ParseProgram()

	if len(prog.Expressions) != 2 {
		t.Fatalf("test - wrong number of expressions. expected=%d, got=%d", 2, len(prog.Expressions))
	}

	if len(p.Errors) != 0 {
		t.Fatalf("test - error list should be empty. expected=%d, got=%d", 0, len(p.Errors))
	}

	plusExpr := prog.Expressions[0].(*ast.AddExpression)
	if plusExpr.Token.Type != token.ADD {
		t.Fatalf("test - wrong token type for plus expression. expected=%s, got=%s", "+", plusExpr.Token.Type)
	}
	if plusExpr.Exprs[0].(*ast.IntegerLiteral).Value != 2 {
		t.Fatalf("test - wrong value for integer literal. expected=%d, got=%d", 2, plusExpr.Exprs[0].(*ast.IntegerLiteral).Value)
	}
	if plusExpr.Exprs[1].(*ast.IntegerLiteral).Value != 3 {
		t.Fatalf("test - wrong value for integer literal. expected=%d, got=%d", 3, plusExpr.Exprs[1].(*ast.IntegerLiteral).Value)
	}

	minusExpr := prog.Expressions[1].(*ast.SubtractExpression)
	if minusExpr.Token.Type != token.SUBTRACT {
		t.Fatalf("test - wrong token type for minus expression. expected=%s, got=%s", "+", minusExpr.Token.Type)
	}
	if minusExpr.Exprs[0].(*ast.IntegerLiteral).Value != 7 {
		t.Fatalf("test - wrong value for integer literal. expected=%d, got=%d", 7, minusExpr.Exprs[0].(*ast.IntegerLiteral).Value)
	}
	if minusExpr.Exprs[1].(*ast.IntegerLiteral).Value != 9 {
		t.Fatalf("test - wrong value for integer literal. expected=%d, got=%d", 9, minusExpr.Exprs[1].(*ast.IntegerLiteral).Value)
	}
}

func TestParser_LogicalOperations(t *testing.T) {
	input := `(not false)
	(not true)
	(and true false)
	(or true false)
	(not= true false)`
	l := lexer.New(input)
	p := New(l)
	prog := p.ParseProgram()

	if len(prog.Expressions) != 5 {
		t.Fatalf("test - wrong number of expressions. expected=%d, got=%d", 5, len(prog.Expressions))
	}

	if len(p.Errors) != 0 {
		t.Fatalf("test - error list should be empty. expected=%d, got=%d", 0, len(p.Errors))
	}
	notExpr := prog.Expressions[0].(*ast.NotExpression)
	if notExpr.Token.Type != token.NOT {
		t.Fatalf("test - wrong token type for not expression. expected=%s, got=%s", token.NOT, notExpr.Token.Type)
	}
	if notExpr.Expr.(*ast.BooleanLiteral).Value != false {
		t.Fatalf("test - wrong value for boolean literal. expected=%t, got=%t", false, notExpr.Expr.(*ast.BooleanLiteral).Value)
	}
	notExprTrue := prog.Expressions[1].(*ast.NotExpression)
	if notExprTrue.Token.Type != token.NOT {
		t.Fatalf("test - wrong token type for not expression. expected=%s, got=%s", token.NOT, notExprTrue.Token.Type)
	}
	if notExprTrue.Expr.(*ast.BooleanLiteral).Value != true {
		t.Fatalf("test - wrong value for boolean literal. expected=%t, got=%t", true, notExprTrue.Expr.(*ast.BooleanLiteral).Value)
	}
	andExpr := prog.Expressions[2].(*ast.AndExpression)
	if andExpr.Token.Type != token.AND {
		t.Fatalf("test - wrong token type for and expression. expected=%s, got=%s", token.AND, andExpr.Token.Type)
	}
	if andExpr.Exprs[0].(*ast.BooleanLiteral).Value != true {
		t.Fatalf("test - wrong value for boolean literal. expected=%t, got=%t", true, andExpr.Exprs[0].(*ast.BooleanLiteral).Value)
	}
	if andExpr.Exprs[1].(*ast.BooleanLiteral).Value != false {
		t.Fatalf("test - wrong value for boolean literal. expected=%t, got=%t", false, andExpr.Exprs[1].(*ast.BooleanLiteral).Value)
	}
	orExpr := prog.Expressions[3].(*ast.OrExpression)
	if orExpr.Token.Type != token.OR {
		t.Fatalf("test - wrong token type for and expression. expected=%s, got=%s", token.OR, orExpr.Token.Type)
	}
	if orExpr.Exprs[0].(*ast.BooleanLiteral).Value != true {
		t.Fatalf("test - wrong value for boolean literal. expected=%t, got=%t", true, orExpr.Exprs[0].(*ast.BooleanLiteral).Value)
	}
	if orExpr.Exprs[1].(*ast.BooleanLiteral).Value != false {
		t.Fatalf("test - wrong value for boolean literal. expected=%t, got=%t", false, orExpr.Exprs[1].(*ast.BooleanLiteral).Value)
	}
	notEq := prog.Expressions[4].(*ast.NotEqualExpression)
	if notEq.Token.Type != token.NotEqual {
		t.Fatalf("test - wrong token type for not expression. expected=%s, got=%s", token.NotEqual, orExpr.Token.Type)
	}
	if notEq.Exprs[0].(*ast.BooleanLiteral).Value != true {
		t.Fatalf("test - wrong value for boolean literal. expected=%t, got=%t", true, notEq.Exprs[0].(*ast.BooleanLiteral).Value)
	}
	if orExpr.Exprs[1].(*ast.BooleanLiteral).Value != false {
		t.Fatalf("test - wrong value for boolean literal. expected=%t, got=%t", false, notEq.Exprs[1].(*ast.BooleanLiteral).Value)
	}
}
