package ast

import (
	"fmt"

	"github.com/branislavlazic/bell/token"
)

type Node interface {
	TokenLiteral() string
	String() string
}

type Expression interface {
	Node
}

type Program struct {
	Expressions []Expression
}

func (p *Program) TokenLiteral() string {
	return ""
}
func (p *Program) String() string {
	var exprs string
	for _, expr := range p.Expressions {
		exprs = fmt.Sprintf("%s\n", expr.String())
	}
	return exprs
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

type EqualExpression struct {
	Token     token.Token // Equal operator
	LeftExpr  Expression
	RightExpr Expression
}

func (ne *EqualExpression) TokenLiteral() string {
	return ne.Token.Literal
}
func (ne *EqualExpression) String() string {
	return fmt.Sprintf("Equal(%s %s)", ne.LeftExpr.String(), ne.RightExpr.String())
}

type NotEqualExpression struct {
	Token     token.Token // Not equal operator
	LeftExpr  Expression
	RightExpr Expression
}

func (nee *NotEqualExpression) TokenLiteral() string {
	return nee.Token.Literal
}
func (nee *NotEqualExpression) String() string {
	return fmt.Sprintf("NotEqual(%s %s)", nee.LeftExpr.String(), nee.RightExpr.String())
}

type AndExpression struct {
	Token     token.Token // And operator
	LeftExpr  Expression
	RightExpr Expression
}

func (ae *AndExpression) TokenLiteral() string {
	return ae.Token.Literal
}
func (ae *AndExpression) String() string {
	return fmt.Sprintf("And(%s %s)", ae.LeftExpr.String(), ae.RightExpr.String())
}

type OrExpression struct {
	Token     token.Token // Pr operator
	LeftExpr  Expression
	RightExpr Expression
}

func (oe *OrExpression) TokenLiteral() string {
	return oe.Token.Literal
}
func (oe *OrExpression) String() string {
	return fmt.Sprintf("Or(%s %s)", oe.LeftExpr.String(), oe.RightExpr.String())
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

type BooleanLiteral struct {
	Token token.Token
	Value bool
}

func (bl *BooleanLiteral) TokenLiteral() string {
	return bl.Token.Literal
}
func (bl *BooleanLiteral) String() string {
	return bl.Token.Literal
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

type NotExpression struct {
	Token token.Token // Not operator
	Expr  Expression
}

func (ne *NotExpression) TokenLiteral() string {
	return ne.Token.Literal
}
func (ne *NotExpression) String() string {
	return fmt.Sprintf("Not(%s)", ne.Expr.String())
}

type NoopExpression struct{}

func (ne *NoopExpression) TokenLiteral() string {
	return "NOOP"
}
func (ne *NoopExpression) String() string {
	return "NOOP"
}
