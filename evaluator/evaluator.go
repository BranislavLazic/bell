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
		return evalBinaryExpression(node, node.LeftExpr, node.RightExpr)
	case *ast.SubtractExpression:
		return evalBinaryExpression(node, node.LeftExpr, node.RightExpr)
	case *ast.MultiplyExpression:
		return evalBinaryExpression(node, node.LeftExpr, node.RightExpr)
	case *ast.DivideExpression:
		return evalBinaryExpression(node, node.LeftExpr, node.RightExpr)
	case *ast.EqualExpression:
		return evalBinaryExpression(node, node.LeftExpr, node.RightExpr)
	case *ast.AndExpression:
		return evalBinaryExpression(node, node.LeftExpr, node.RightExpr)
	case *ast.OrExpression:
		return evalBinaryExpression(node, node.LeftExpr, node.RightExpr)
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

func evalBinaryExpression(node ast.Node, leftExpr ast.Expression, rightExpr ast.Expression) object.Object {
	left := Eval(leftExpr)
	right := Eval(rightExpr)
	return evalExpression(node, left, right)
}

func evalExpression(exprType ast.Node, left object.Object, right object.Object) object.Object {
	switch {
	case left.Type() == object.IntegerObj && right.Type() == object.IntegerObj:
		leftVal := left.(*object.Integer)
		rightVal := right.(*object.Integer)
		return evalArithmeticOperation(exprType, leftVal, rightVal)
	case left.Type() == object.BooleanObj && right.Type() == object.BooleanObj:
		leftVal := left.(*object.Boolean)
		rightVal := right.(*object.Boolean)
		return evalLogicalOperation(exprType, leftVal, rightVal)
	default:
		return &object.RuntimeError{
			Error: fmt.Sprintf("Operation %s cannot be performed for types: %s and %s", exprType.String(), left.Type(), right.Type()),
		}
	}
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
	case *ast.EqualExpression:
		return &object.Boolean{Value: left.Value == right.Value}
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
	default:
		return &object.RuntimeError{Error: fmt.Sprintf("Non-existing operation %s for BOOLEAN types.", exprType.String())}
	}
}

func evalNegateExpression(value object.Object) object.Object {
	if value.Type() == object.IntegerObj {
		v := value.(*object.Integer).Value
		return &object.Integer{Value: -1 * v}
	} else {
		return &object.Null{}
	}
}

func evalNotExpression(value object.Object) object.Object {
	if value.Type() == object.BooleanObj {
		v := value.(*object.Boolean).Value
		return &object.Boolean{Value: !v}
	} else {
		return &object.Null{}
	}
}

func evalExpressions(exprs []ast.Expression) object.Object {
	var result object.Object
	for _, expr := range exprs {
		result = Eval(expr)
	}
	return result
}
