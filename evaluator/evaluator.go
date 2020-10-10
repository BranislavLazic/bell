package evaluator

import (
	"github.com/branislavlazic/bell/ast"
	"github.com/branislavlazic/bell/object"
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		return evalStatements(node.Statements)
	case *ast.AddExpression:
		left := Eval(node.LeftExpr)
		right := Eval(node.RightExpr)
		return evalArithmeticExpression(node, left, right)
	case *ast.SubtractExpression:
		left := Eval(node.LeftExpr)
		right := Eval(node.RightExpr)
		return evalArithmeticExpression(node, left, right)
	case *ast.MultiplyExpression:
		left := Eval(node.LeftExpr)
		right := Eval(node.RightExpr)
		return evalArithmeticExpression(node, left, right)
	case *ast.DivideExpression:
		left := Eval(node.LeftExpr)
		right := Eval(node.RightExpr)
		return evalArithmeticExpression(node, left, right)
	case *ast.NegativeValueExpression:
		value := Eval(node.Expr)
		return evalNegateExpression(value)
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	default:
		return &object.Null{}
	}
}

func evalArithmeticExpression(exprType ast.Node, left object.Object, right object.Object) object.Object {
	switch {
	case left.Type() == object.IntegerObj && right.Type() == object.IntegerObj:
		leftVal := left.(*object.Integer)
		rightVal := right.(*object.Integer)
		return evalArithmeticOperation(exprType, leftVal, rightVal)
	default:
		return &object.Null{}
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
	default:
		return &object.Null{}
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

func evalStatements(stmts []ast.Statement) object.Object {
	var result object.Object
	for _, stmt := range stmts {
		result = Eval(stmt.Expr)
	}
	return result
}
