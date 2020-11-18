package token

type TokType string

type Token struct {
	Type    TokType
	Literal string
}

const (
	StartExpression = "START_EXPRESSION"
	EndExpression   = "END_EXPRESSION"
	INT             = "INT"
	BOOL            = "BOOL"
	LET             = "LET"
	IF              = "IF"
	LIST            = "LIST"
	STRING          = "STRING"
	OPEN            = "OPEN"
	IDENT           = "IDENT"
	NIL             = "NIL"
	StartParamList  = "START_PARAM_LIST"
	EndParamList    = "END_PARAM_LIST"
	EOF             = "EOF"
	EOL             = "EOL"
	ILLEGAL         = "ILLEGAL"
	// Operators
	ADD              = "ADD"
	SUBTRACT         = "SUBTRACT"
	MULTIPLY         = "MULTIPLY"
	DIVIDE           = "DIVIDE"
	MODULO           = "MODULO"
	POW              = "POW"
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
	"let":   LET,
	"if":    IF,
	"list":  LIST,
	"open":  OPEN,
	"nil":   NIL,
}

func LookupKeyword(instruction string) TokType {
	if tok, ok := keywords[instruction]; ok {
		return tok
	}
	return ILLEGAL
}

var OperatorLiterals = []string{
	"+", "-", "%",
	"*", "-", "/",
	"=", "not=", "and",
	"or", ">", ">=",
	"<", "<=", "not",
	"list", "if",
	"^", "open"}
