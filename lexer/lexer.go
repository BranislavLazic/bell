package lexer

import (
	"github.com/branislavlazic/bell/token"
)

type Lexer struct {
	input        string
	Position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	lexer := &Lexer{input: input}
	lexer.readChar()
	return lexer
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.skipWhitespace()
	switch l.ch {
	case '+':
		tok = newToken(token.ADD, l.ch)
	case '-':
		tok = newToken(token.SUBTRACT, l.ch)
	case '*':
		tok = newToken(token.MULTIPLY, l.ch)
	case '/':
		tok = newToken(token.DIVIDE, l.ch)
	case '%':
		tok = newToken(token.MODULO, l.ch)
	case '^':
		tok = newToken(token.POW, l.ch)
	case '=':
		tok = newToken(token.EQUAL, l.ch)
	case '>':
		if l.peekChar() == '=' {
			l.readChar()
			tok.Literal = ">="
			tok.Type = token.GreaterThanEqual
		} else {
			tok = newToken(token.GreaterThan, l.ch)
		}
	case '<':
		if l.peekChar() == '=' {
			l.readChar()
			tok.Literal = "<="
			tok.Type = token.LessThanEqual
		} else {
			tok = newToken(token.LessThan, l.ch)
		}
	case '(':
		tok = newToken(token.StartExpression, l.ch)
	case ')':
		tok = newToken(token.EndExpression, l.ch)
	case '[':
		tok = newToken(token.StartParamList, l.ch)
	case ']':
		tok = newToken(token.EndParamList, l.ch)
	case '"':
		tok.Type = token.STRING
		tok.Literal = l.readString()
	case '\n':
		tok.Literal = ""
		tok.Type = token.EOL
	case '\r':
		tok.Literal = ""
		tok.Type = token.EOL
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tokType := token.LookupKeyword(tok.Literal)
			if tokType == token.ILLEGAL {
				tok.Type = token.IDENT
			} else {
				tok.Type = token.LookupKeyword(tok.Literal)
			}
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.Position = l.readPosition
	l.readPosition++
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

func (l *Lexer) readNumber() string {
	position := l.Position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.Position]
}

// Identifier is a sequence of characters and
// it must begins with a letter
func (l *Lexer) readIdentifier() string {
	position := l.Position
	for isLetter(l.ch) || isAllowedFollowingIdentChar(l.ch) {
		l.readChar()
	}
	return l.input[position:l.Position]
}

func (l *Lexer) readString() string {
	position := l.Position + 1
	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
	}
	return l.input[position:l.Position]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' {
		l.readChar()
	}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isAllowedFollowingIdentChar(ch byte) bool {
	return isDigit(ch) || ch == '-' || ch == '?' || ch == '='
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func newToken(tokenType token.TokType, ch byte) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(ch),
	}
}
