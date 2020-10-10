package token

type TokType string

const (
	StartExpression = "START_EXPRESSION"
	EndExpression   = "END_EXPRESSION"
	INT             = "INT"
	// Operators
	ADD      = "ADD"
	SUBTRACT = "SUBTRACT"
	MULTIPLY = "MULTIPLY"
	DIVIDE   = "DIVIDE"
	EOF      = "EOF"
	EOL      = "EOL"
	ILLEGAL  = "ILLEGAL"
)

var instructions = map[string]TokType{
	"(": StartExpression,
	")": EndExpression,
}

func LookupInstruction(instruction string) TokType {
	if tok, ok := instructions[instruction]; ok {
		return tok
	}
	return ILLEGAL
}

type Token struct {
	Type    TokType
	Literal string
}
