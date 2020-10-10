package parser

import (
	"testing"

	"github.com/branislavlazic/bell/ast"
	"github.com/branislavlazic/bell/lexer"
	"github.com/branislavlazic/bell/token"
)

func TestParser_ParseLiteral(t *testing.T) {
	input := `(2)`
	l := lexer.New(input)
	p := New(l)
	prog := p.ParseProgram()

	if len(prog.Statements) != 1 {
		t.Fatalf("test - wrong number of statements. expected=%d, got=%d", 1, len(prog.Statements))
	}
	if len(p.Errors) != 0 {
		t.Fatalf("test - error list should be empty. expected=%d, got=%d", 0, len(p.Errors))
	}
	intLiteralExpr := prog.Statements[0].Expr.(*ast.IntegerLiteral)

	if intLiteralExpr.Value != 2 {
		t.Fatalf("test - wrong value of integer literal. expected=%d, got=%d", 2, intLiteralExpr.Value)
	}
}

func TestParser_ParseNegativeLiteral(t *testing.T) {
	input := `(- 2)`
	l := lexer.New(input)
	p := New(l)
	prog := p.ParseProgram()

	if len(prog.Statements) != 1 {
		t.Fatalf("test - wrong number of statements. expected=%d, got=%d", 1, len(prog.Statements))
	}
	if len(p.Errors) != 0 {
		t.Fatalf("test - error list should be empty. expected=%d, got=%d", 0, len(p.Errors))
	}
	negativeValueExpr := prog.Statements[0].Expr.(*ast.NegativeValueExpression)

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
	// fmt.Printf("%+v\n", prog)

	if len(prog.Statements) != 2 {
		t.Fatalf("test - wrong number of statements. expected=%d, got=%d", 2, len(prog.Statements))
	}

	if len(p.Errors) != 0 {
		t.Fatalf("test - error list should be empty. expected=%d, got=%d", 0, len(p.Errors))
	}

	plusExpr := prog.Statements[0].Expr.(*ast.AddExpression)
	if plusExpr.Token.Type != token.ADD {
		t.Fatalf("test - wrong token type for plus expression. expected=%s, got=%s", "+", plusExpr.Token.Type)
	}
	if plusExpr.LeftExpr.(*ast.IntegerLiteral).Value != 2 {
		t.Fatalf("test - wrong value for integer literal. expected=%d, got=%d", 2, plusExpr.LeftExpr.(*ast.IntegerLiteral).Value)
	}
	if plusExpr.RightExpr.(*ast.IntegerLiteral).Value != 3 {
		t.Fatalf("test - wrong value for integer literal. expected=%d, got=%d", 3, plusExpr.RightExpr.(*ast.IntegerLiteral).Value)
	}

	minusExpr := prog.Statements[1].Expr.(*ast.SubtractExpression)
	if minusExpr.Token.Type != token.SUBTRACT {
		t.Fatalf("test - wrong token type for minus expression. expected=%s, got=%s", "+", minusExpr.Token.Type)
	}
	if minusExpr.LeftExpr.(*ast.IntegerLiteral).Value != 7 {
		t.Fatalf("test - wrong value for integer literal. expected=%d, got=%d", 7, minusExpr.LeftExpr.(*ast.IntegerLiteral).Value)
	}
	if minusExpr.RightExpr.(*ast.IntegerLiteral).Value != 9 {
		t.Fatalf("test - wrong value for integer literal. expected=%d, got=%d", 9, minusExpr.RightExpr.(*ast.IntegerLiteral).Value)
	}
}
