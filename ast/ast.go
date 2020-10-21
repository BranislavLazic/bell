package ast

import (
	"fmt"
	"strings"

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
	Token token.Token // Add operator
	Exprs []Expression
}

func (ae *AddExpression) TokenLiteral() string {
	return ae.Token.Literal
}
func (ae *AddExpression) String() string {
	return fmt.Sprintf("(+ %s)", concatExprsAsString(ae.Exprs))
}

type SubtractExpression struct {
	Token token.Token // Subtract operator
	Exprs []Expression
}

func (se *SubtractExpression) TokenLiteral() string {
	return se.Token.Literal
}
func (se *SubtractExpression) String() string {
	return fmt.Sprintf("(- %s)", concatExprsAsString(se.Exprs))
}

type MultiplyExpression struct {
	Token token.Token // Multiply operator
	Exprs []Expression
}

func (me *MultiplyExpression) TokenLiteral() string {
	return me.Token.Literal
}
func (me *MultiplyExpression) String() string {
	return fmt.Sprintf("(* %s)", concatExprsAsString(me.Exprs))
}

type DivideExpression struct {
	Token token.Token // Divide operator
	Exprs []Expression
}

func (de *DivideExpression) TokenLiteral() string {
	return de.Token.Literal
}
func (de *DivideExpression) String() string {
	return fmt.Sprintf("(/ %s)", concatExprsAsString(de.Exprs))
}

type EqualExpression struct {
	Token token.Token // Equal operator
	Exprs []Expression
}

func (ee *EqualExpression) TokenLiteral() string {
	return ee.Token.Literal
}
func (ee *EqualExpression) String() string {
	return fmt.Sprintf("(= %s)", concatExprsAsString(ee.Exprs))
}

type NotEqualExpression struct {
	Token token.Token // Not equal operator
	Exprs []Expression
}

func (nee *NotEqualExpression) TokenLiteral() string {
	return nee.Token.Literal
}
func (nee *NotEqualExpression) String() string {
	return fmt.Sprintf("(not= %s)", concatExprsAsString(nee.Exprs))
}

type AndExpression struct {
	Token token.Token // And operator
	Exprs []Expression
}

func (ae *AndExpression) TokenLiteral() string {
	return ae.Token.Literal
}
func (ae *AndExpression) String() string {
	return fmt.Sprintf("(and %s)", concatExprsAsString(ae.Exprs))
}

type OrExpression struct {
	Token token.Token // Or operator
	Exprs []Expression
}

func (oe *OrExpression) TokenLiteral() string {
	return oe.Token.Literal
}
func (oe *OrExpression) String() string {
	return fmt.Sprintf("(or %s)", concatExprsAsString(oe.Exprs))
}

type ModuloExpression struct {
	Token token.Token // Modulo operator
	Exprs []Expression
}

func (me *ModuloExpression) TokenLiteral() string {
	return me.Token.Literal
}
func (me *ModuloExpression) String() string {
	return fmt.Sprintf("(%% %s)", concatExprsAsString(me.Exprs))
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
	return fmt.Sprintf("(- %s)", ae.Expr.String())
}

type NotExpression struct {
	Token token.Token // Not operator
	Expr  Expression
}

func (ne *NotExpression) TokenLiteral() string {
	return ne.Token.Literal
}
func (ne *NotExpression) String() string {
	return fmt.Sprintf("(not %s)", ne.Expr.String())
}

type GreaterThanExpression struct {
	Token token.Token // Greater than operator
	Exprs []Expression
}

func (gte *GreaterThanExpression) TokenLiteral() string {
	return gte.Token.Literal
}
func (gte *GreaterThanExpression) String() string {
	return fmt.Sprintf("(> %s)", concatExprsAsString(gte.Exprs))
}

type LessThanExpression struct {
	Token token.Token // Less than operator
	Exprs []Expression
}

func (lte *LessThanExpression) TokenLiteral() string {
	return lte.Token.Literal
}
func (lte *LessThanExpression) String() string {
	return fmt.Sprintf("(< %s)", concatExprsAsString(lte.Exprs))
}

type GreaterThanEqualExpression struct {
	Token token.Token // Greater than equal operator
	Exprs []Expression
}

func (gtee *GreaterThanEqualExpression) TokenLiteral() string {
	return gtee.Token.Literal
}
func (gtee *GreaterThanEqualExpression) String() string {
	return fmt.Sprintf("(>= %s)", concatExprsAsString(gtee.Exprs))
}

type LessThanEqualExpression struct {
	Token token.Token // Less than equal operator
	Exprs []Expression
}

func (gtee *LessThanEqualExpression) TokenLiteral() string {
	return gtee.Token.Literal
}
func (gtee *LessThanEqualExpression) String() string {
	return fmt.Sprintf("(<= %s)", concatExprsAsString(gtee.Exprs))
}

func concatExprsAsString(exprs []Expression) string {
	var exprsAsStrArr []string
	for _, expr := range exprs {
		exprsAsStrArr = append(exprsAsStrArr, expr.String())
	}
	return strings.Join(exprsAsStrArr, " ")
}
