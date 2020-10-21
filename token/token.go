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
	// Operators
	ADD              = "ADD"
	SUBTRACT         = "SUBTRACT"
	MULTIPLY         = "MULTIPLY"
	DIVIDE           = "DIVIDE"
	MODULO           = "MODULO"
	EQUAL            = "EQUAL"
	NotEqual         = "NOT_EQUAL"
	NOT              = "NOT"
	AND              = "AND"
	OR               = "OR"
	GreaterThan      = "GREATER_THAN"
	LessThan         = "LESS_THAN"
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
