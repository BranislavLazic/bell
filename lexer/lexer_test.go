package lexer

import (
	"testing"

	"github.com/branislavlazic/bell/token"
)

func TestNextToken(t *testing.T) {
	input := `(+ 2 3)`
	tests := []struct {
		expectedType    token.TokType
		expectedLiteral string
	}{
		{token.StartExpression, "("},
		{token.ADD, "+"},
		{token.INT, "2"},
		{token.INT, "3"},
		{token.EndExpression, ")"},
	}
	l := New(input)

	for i, tokenType := range tests {
		tok := l.NextToken()
		if tok.Type != tokenType.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tokenType.expectedType, tok.Type)
		}
		if tok.Literal != tokenType.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tokenType.expectedLiteral, tok.Literal)
		}
	}
}
