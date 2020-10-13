package token

type TokType string

const (
	StartExpression = "START_EXPRESSION"
	EndExpression   = "END_EXPRESSION"
	INT             = "INT"
	BOOL            = "BOOL"
	EOF             = "EOF"
	EOL             = "EOL"
	ILLEGAL         = "ILLEGAL"
	IdentTooLong    = "IDENT_TOO_LONG"
	// Operators
	ADD              = "ADD"
	SUBTRACT         = "SUBTRACT"
	MULTIPLY         = "MULTIPLY"
	DIVIDE           = "DIVIDE"
	EQUAL            = "EQUAL"
	NotEqual         = "NOT_EQUAL"
	NOT              = "NOT"
	AND              = "AND"
	OR               = "OR"
	GreaterThen      = "GREATER_THEN"
	LessThen         = "LESS_THEN"
	GreaterThanEqual = "GREATER_THAN_EQUAL"
	LessThanEqual    = "LESS_THAN_EQUAL"
)

var keywords = map[string]TokType{
	"true":  BOOL,
	"false": BOOL,
	"and":   AND,
	"or":    OR,
	"not":   NOT,
	"not=":  NotEqual,
}

func LookupKeyword(instruction string) TokType {
	if tok, ok := keywords[instruction]; ok {
		return tok
	}
	return ILLEGAL
}

type Token struct {
	Type    TokType
	Literal string
}
