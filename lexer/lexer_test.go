package lexer

import (
	"testing"

	"github.com/branislavlazic/bell/token"
)

func TestNextToken(t *testing.T) {
	input := `(+ 2 3)
	(- 4 5)
	(or false true)
	(and false true)
	(not false)`
	tests := []struct {
		expectedType    token.TokType
		expectedLiteral string
	}{
		// Add expr
		{token.StartExpression, "("},
		{token.ADD, "+"},
		{token.INT, "2"},
		{token.INT, "3"},
		{token.EndExpression, ")"},
		{token.EOL, ""},
		// Subtract expr
		{token.StartExpression, "("},
		{token.SUBTRACT, "-"},
		{token.INT, "4"},
		{token.INT, "5"},
		{token.EndExpression, ")"},
		{token.EOL, ""},
		// Or expr
		{token.StartExpression, "("},
		{token.OR, "or"},
		{token.BOOL, "false"},
		{token.BOOL, "true"},
		{token.EndExpression, ")"},
		{token.EOL, ""},
		// And expr
		{token.StartExpression, "("},
		{token.AND, "and"},
		{token.BOOL, "false"},
		{token.BOOL, "true"},
		{token.EndExpression, ")"},
		{token.EOL, ""},
		// Not expr
		{token.StartExpression, "("},
		{token.NOT, "not"},
		{token.BOOL, "false"},
		{token.EndExpression, ")"},
		{token.EOF, ""},
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
