package ast

import (
	"fmt"

	"github.com/branislavlazic/bell/token"
)

type Node interface {
	TokenLiteral() string
	String() string
}

type Statement struct {
	Expr Expression
}

type Expression interface {
	Node
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	return ""
}
func (p *Program) String() string {
	var stmts string
	for _, stmt := range p.Statements {
		stmts = fmt.Sprintf("%s\n", stmt.Expr.String())
	}
	return stmts
}

type AddExpression struct {
	Token     token.Token // Add operator
	LeftExpr  Expression
	RightExpr Expression
}

func (ae *AddExpression) TokenLiteral() string {
	return ae.Token.Literal
}
func (ae *AddExpression) String() string {
	return fmt.Sprintf("Add(%s, %s)", ae.LeftExpr.String(), ae.RightExpr.String())
}

type SubtractExpression struct {
	Token     token.Token // Subtract operator
	LeftExpr  Expression
	RightExpr Expression
}

func (se *SubtractExpression) TokenLiteral() string {
	return se.Token.Literal
}
func (se *SubtractExpression) String() string {
	return fmt.Sprintf("Subtract(%s, %s)", se.LeftExpr.String(), se.RightExpr.String())
}

type MultiplyExpression struct {
	Token     token.Token // Multiply operator
	LeftExpr  Expression
	RightExpr Expression
}

func (me *MultiplyExpression) TokenLiteral() string {
	return me.Token.Literal
}
func (me *MultiplyExpression) String() string {
	return fmt.Sprintf("Multiply(%s, %s)", me.LeftExpr.String(), me.RightExpr.String())
}

type DivideExpression struct {
	Token     token.Token // Divide operator
	LeftExpr  Expression
	RightExpr Expression
}

func (de *DivideExpression) TokenLiteral() string {
	return de.Token.Literal
}
func (de *DivideExpression) String() string {
	return fmt.Sprintf("Divide(%s %s)", de.LeftExpr.String(), de.RightExpr.String())
}

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) TokenLiteral() string {
	return il.Token.Literal
}
func (il *IntegerLiteral) String() string {
	return il.Token.Literal
}

type NegativeValueExpression struct {
	Token token.Token // Subtract operator
	Expr  Expression
}

func (ae *NegativeValueExpression) TokenLiteral() string {
	return ae.Token.Literal
}
func (ae *NegativeValueExpression) String() string {
	return fmt.Sprintf("NegativeValue(%s)", ae.Expr.String())
}

type NoopExpression struct{}

func (ne *NoopExpression) TokenLiteral() string {
	return "NOOP"
}
func (ne *NoopExpression) String() string {
	return "NOOP"
}
