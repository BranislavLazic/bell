package evaluator

import (
	"fmt"
	"github.com/branislavlazic/bell/ast"
	"github.com/branislavlazic/bell/object"
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		return evalExpressions(node.Expressions)
	case *ast.AddExpression:
		return evalExpression(node, node.Exprs)
	case *ast.SubtractExpression:
		return evalExpression(node, node.Exprs)
	case *ast.MultiplyExpression:
		return evalExpression(node, node.Exprs)
	case *ast.DivideExpression:
		return evalExpression(node, node.Exprs)
	case *ast.ModuloExpression:
		return evalExpression(node, node.Exprs)
	case *ast.EqualExpression:
		return evalEqualForAll(node, node.Exprs)
	case *ast.NotEqualExpression:
		return evalNotEqualForAll(node, node.Exprs)
	case *ast.AndExpression:
		return evalExpression(node, node.Exprs)
	case *ast.OrExpression:
		return evalExpression(node, node.Exprs)
	case *ast.GreaterThanExpression:
		return evalGreaterThan(node, node.Exprs)
	case *ast.GreaterThanEqualExpression:
		return evalGreaterThanEqual(node, node.Exprs)
	case *ast.LessThanExpression:
		return evalLessThan(node, node.Exprs)
	case *ast.LessThanEqualExpression:
		return evalLessThanEqual(node, node.Exprs)
	case *ast.NegativeValueExpression:
		value := Eval(node.Expr)
		return evalNegateExpression(value)
	case *ast.NotExpression:
		value := Eval(node.Expr)
		return evalNotExpression(value)
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	case *ast.BooleanLiteral:
		return &object.Boolean{Value: node.Value}
	default:
		return &object.Null{}
	}
}

func evalExpression(exprType ast.Node, exprs []ast.Expression) object.Object {
	var accumResult object.Object
	for _, expr := range exprs {
		evalExpr := Eval(expr)
		// Declare first value as an accumulator
		if accumResult == nil {
			accumResult = evalExpr
		} else {
			switch {
			case evalExpr.Type() == object.IntegerObj && accumResult.Type() == object.IntegerObj:
				nextValue := evalExpr.(*object.Integer)
				// Evaluate operation by passing current accumulated value and next value
				accumResult = evalArithmeticOperation(exprType, accumResult.(*object.Integer), nextValue)
			case evalExpr.Type() == object.BooleanObj && accumResult.Type() == object.BooleanObj:
				nextValue := evalExpr.(*object.Boolean)
				accumResult = evalLogicalOperation(exprType, accumResult.(*object.Boolean), nextValue)
			default:
				return &object.RuntimeError{
					Error: fmt.Sprintf("Operation %s cannot be performed for types: %s and %s",
						exprType.String(), accumResult.Type(), evalExpr.Type()),
				}
			}
		}
	}
	return accumResult
}

func evalArithmeticOperation(exprType ast.Node, left *object.Integer, right *object.Integer) object.Object {
	switch exprType.(type) {
	case *ast.AddExpression:
		return &object.Integer{Value: left.Value + right.Value}
	case *ast.SubtractExpression:
		if right.Value < 0 {
			return &object.Integer{Value: left.Value + right.Value}
		}
		return &object.Integer{Value: left.Value - right.Value}
	case *ast.MultiplyExpression:
		return &object.Integer{Value: left.Value * right.Value}
	case *ast.DivideExpression:
		return &object.Integer{Value: left.Value / right.Value}
	case *ast.ModuloExpression:
		return &object.Integer{Value: left.Value % right.Value}
	case *ast.EqualExpression:
		return &object.Boolean{Value: left.Value == right.Value}
	case *ast.NotEqualExpression:
		return &object.Boolean{Value: left.Value != right.Value}
	default:
		return &object.RuntimeError{Error: fmt.Sprintf("Non-existing operation %s for INTEGER types.", exprType.String())}
	}
}

func evalLogicalOperation(exprType ast.Node, left *object.Boolean, right *object.Boolean) object.Object {
	switch exprType.(type) {
	case *ast.AndExpression:
		return &object.Boolean{Value: left.Value && right.Value}
	case *ast.OrExpression:
		return &object.Boolean{Value: left.Value || right.Value}
	case *ast.EqualExpression:
		return &object.Boolean{Value: left.Value == right.Value}
	case *ast.NotEqualExpression:
		return &object.Boolean{Value: left.Value != right.Value}
	default:
		return &object.RuntimeError{Error: fmt.Sprintf("Non-existing operation %s for BOOLEAN types.", exprType.String())}
	}
}

func evalEqualForAll(exprType ast.Node, exprs []ast.Expression) object.Object {
	var accumResult object.Object
	for _, expr := range exprs {
		evalExpr := Eval(expr)
		if accumResult == nil {
			accumResult = evalExpr
		} else {
			switch {
			case evalExpr.Type() == object.IntegerObj && accumResult.Type() == object.IntegerObj:
				if accumResult.(*object.Integer).Value != evalExpr.(*object.Integer).Value {
					return &object.Boolean{Value: false}
				}
			case evalExpr.Type() == object.BooleanObj && accumResult.Type() == object.BooleanObj:
				if accumResult.(*object.Boolean).Value != evalExpr.(*object.Boolean).Value {
					return &object.Boolean{Value: false}
				}
			default:
				return &object.RuntimeError{
					Error: fmt.Sprintf("Operation %s cannot be performed for types: %s and %s",
						exprType.String(), accumResult.Type(), evalExpr.Type()),
				}
			}
		}
	}
	return &object.Boolean{Value: true}
}

func evalNotEqualForAll(exprType ast.Node, exprs []ast.Expression) object.Object {
	res := evalEqualForAll(exprType, exprs)
	switch res.(type) {
	case *object.Boolean:
		return &object.Boolean{Value: !res.(*object.Boolean).Value}
	default:
		return res
	}
}

func evalComparison(exprType ast.Node, exprs []ast.Expression, comparator func(left int64, right int64) bool) object.Object {
	var accumResult object.Object
	for _, expr := range exprs {
		evalExpr := Eval(expr)
		if accumResult == nil {
			accumResult = evalExpr
		} else {
			switch {
			case evalExpr.Type() == object.IntegerObj && accumResult.Type() == object.IntegerObj:
				if comparator(accumResult.(*object.Integer).Value, evalExpr.(*object.Integer).Value) {
					return &object.Boolean{Value: false}
				}
			default:
				return &object.RuntimeError{
					Error: fmt.Sprintf("Operation %s cannot be performed for types: %s and %s",
						exprType.String(), accumResult.Type(), evalExpr.Type()),
				}
			}
		}
	}
	return &object.Boolean{Value: true}
}

func evalGreaterThan(exprType ast.Node, exprs []ast.Expression) object.Object {
	return evalComparison(exprType, exprs, func(left int64, right int64) bool {
		return left <= right
	})
}

func evalGreaterThanEqual(exprType ast.Node, exprs []ast.Expression) object.Object {
	return evalComparison(exprType, exprs, func(left int64, right int64) bool {
		return left < right
	})
}

func evalLessThan(exprType ast.Node, exprs []ast.Expression) object.Object {
	return evalComparison(exprType, exprs, func(left int64, right int64) bool {
		return left >= right
	})
}

func evalLessThanEqual(exprType ast.Node, exprs []ast.Expression) object.Object {
	return evalComparison(exprType, exprs, func(left int64, right int64) bool {
		return left > right
	})
}

func evalNegateExpression(value object.Object) object.Object {
	if value.Type() == object.IntegerObj {
		v := value.(*object.Integer).Value
		return &object.Integer{Value: -1 * v}
	} else {
		return &object.RuntimeError{
			Error: fmt.Sprintf("Negation of arithmetic expressions is not applicable for %s type.", value.Type()),
		}
	}
}

func evalNotExpression(value object.Object) object.Object {
	if value.Type() == object.BooleanObj {
		v := value.(*object.Boolean).Value
		return &object.Boolean{Value: !v}
	} else {
		return &object.RuntimeError{
			Error: fmt.Sprintf("Negation of logical expressions is not applicable for %s types.", value.Type()),
		}
	}
}

func evalExpressions(exprs []ast.Expression) object.Object {
	var result object.Object
	for _, expr := range exprs {
		result = Eval(expr)
	}
	return result
}
