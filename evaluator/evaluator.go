package evaluator

import (
	"fmt"

	"github.com/branislavlazic/bell/ast"
	"github.com/branislavlazic/bell/object"
)

func Eval(node ast.Node, env *object.Environment) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		return evalExpressions(node.Expressions, env)
	case *ast.AddExpression:
		return evalExpression(node, node.Exprs, env)
	case *ast.SubtractExpression:
		return evalExpression(node, node.Exprs, env)
	case *ast.MultiplyExpression:
		return evalExpression(node, node.Exprs, env)
	case *ast.DivideExpression:
		return evalExpression(node, node.Exprs, env)
	case *ast.ModuloExpression:
		return evalExpression(node, node.Exprs, env)
	case *ast.EqualExpression:
		return evalEqualForAll(node, node.Exprs, env)
	case *ast.NotEqualExpression:
		return evalNotEqualForAll(node, node.Exprs, env)
	case *ast.AndExpression:
		return evalExpression(node, node.Exprs, env)
	case *ast.OrExpression:
		return evalExpression(node, node.Exprs, env)
	case *ast.GreaterThanExpression:
		return evalGreaterThan(node, node.Exprs, env)
	case *ast.GreaterThanEqualExpression:
		return evalGreaterThanEqual(node, node.Exprs, env)
	case *ast.LessThanExpression:
		return evalLessThan(node, node.Exprs, env)
	case *ast.LessThanEqualExpression:
		return evalLessThanEqual(node, node.Exprs, env)
	case *ast.NegativeValueExpression:
		return evalNegateExpression(Eval(node.Expr, env))
	case *ast.NotExpression:
		return evalNotExpression(Eval(node.Expr, env))
	case *ast.LetExpression:
		return evalLetExpression(node, env)
	case *ast.IfExpression:
		return evalIfExpression(node, env)
	case *ast.Identifier:
		return evalIdentifier(node, env)
	case *ast.ListExpression:
		return evalListExpression(node, env)
	case *ast.Function:
		return evalFunctionExpression(node, env)
	case *ast.CallFunction:
		return evalCallFunctionExpression(node, env)
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	case *ast.BooleanLiteral:
		return &object.Boolean{Value: node.Value}
	default:
		return &object.Nil{}
	}
}

func evalExpression(exprType ast.Node, exprs []ast.Expression, env *object.Environment) object.Object {
	var accumResult object.Object
	for _, expr := range exprs {
		evalExpr := Eval(expr, env)
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

func evalEqualForAll(exprType ast.Node, exprs []ast.Expression, env *object.Environment) object.Object {
	var accumResult object.Object
	for _, expr := range exprs {
		evalExpr := Eval(expr, env)
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

func evalNotEqualForAll(exprType ast.Node, exprs []ast.Expression, env *object.Environment) object.Object {
	res := evalEqualForAll(exprType, exprs, env)
	switch res.(type) {
	case *object.Boolean:
		return &object.Boolean{Value: !res.(*object.Boolean).Value}
	default:
		return res
	}
}

func evalComparison(exprType ast.Node, exprs []ast.Expression, env *object.Environment, comparator func(left int64, right int64) bool) object.Object {
	var accumResult object.Object
	for _, expr := range exprs {
		evalExpr := Eval(expr, env)
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

func evalGreaterThan(exprType ast.Node, exprs []ast.Expression, env *object.Environment) object.Object {
	return evalComparison(exprType, exprs, env, func(left int64, right int64) bool {
		return left <= right
	})
}

func evalGreaterThanEqual(exprType ast.Node, exprs []ast.Expression, env *object.Environment) object.Object {
	return evalComparison(exprType, exprs, env, func(left int64, right int64) bool {
		return left < right
	})
}

func evalLessThan(exprType ast.Node, exprs []ast.Expression, env *object.Environment) object.Object {
	return evalComparison(exprType, exprs, env, func(left int64, right int64) bool {
		return left >= right
	})
}

func evalLessThanEqual(exprType ast.Node, exprs []ast.Expression, env *object.Environment) object.Object {
	return evalComparison(exprType, exprs, env, func(left int64, right int64) bool {
		return left > right
	})
}

func evalNegateExpression(value object.Object) object.Object {
	if value.Type() == object.IntegerObj {
		v := value.(*object.Integer).Value
		return &object.Integer{Value: -1 * v}
	}
	return &object.RuntimeError{
		Error: fmt.Sprintf("Negation of arithmetic expressions is not applicable for %s type.", value.Type()),
	}
}

func evalNotExpression(value object.Object) object.Object {
	if value.Type() == object.BooleanObj {
		v := value.(*object.Boolean).Value
		return &object.Boolean{Value: !v}
	}
	return &object.RuntimeError{
		Error: fmt.Sprintf("Negation of logical expressions is not applicable for %s types.", value.Type()),
	}
}

func evalLetExpression(letExpr *ast.LetExpression, env *object.Environment) object.Object {
	ident := letExpr.Identifier.String()
	val := Eval(letExpr.Expr, env)
	env.Set(ident, val)
	return val
}

func evalFunctionExpression(f *ast.Function, env *object.Environment) object.Object {
	ident := f.Identifier.String()
	fn := &object.Function{Identifier: f.Identifier, Params: f.Params, Body: f.Body}
	env.Set(ident, fn)
	return fn
}

func evalCallFunctionExpression(cf *ast.CallFunction, env *object.Environment) object.Object {
	fnName := cf.Identifier.String()
	val, ok := env.Get(fnName)
	if !ok {
		return &object.RuntimeError{Error: fmt.Sprintf("Function %s is undefined", fnName)}
	}
	switch fn := val.(type) {
	case *object.Function:
		paramsCount := len(fn.Params)
		argsCount := len(cf.Args)
		if argsCount > paramsCount {
			return &object.RuntimeError{
				Error: fmt.Sprintf("Too many arguments. Expected %d, got %d.", paramsCount, argsCount),
			}
		}
		if argsCount < paramsCount {
			return &object.RuntimeError{
				Error: fmt.Sprintf("Insufficient number of arguments. Expected %d, got %d.", paramsCount, argsCount),
			}
		}
		for idx, param := range fn.Params {
			env.Set(param.Value, Eval(cf.Args[idx], env))
		}
		return Eval(fn.Body, env)
	}
	return val
}

func evalIfExpression(ifExpr *ast.IfExpression, env *object.Environment) object.Object {
	cond := Eval(ifExpr.Condition, env)
	switch cnd := cond.(type) {
	case *object.Boolean:
		if cnd.Value {
			return Eval(ifExpr.ThenExpr, env)
		}
		return Eval(ifExpr.ElseExpr, env)
	}
	return &object.RuntimeError{
		Error: fmt.Sprintf("Condition for if expression should evaluate to BOOLEAN type. Found %s type.", cond.Type()),
	}
}

func evalListExpression(listExpression *ast.ListExpression, env *object.Environment) object.Object {
	list := &object.List{Objects: []object.Object{}}
	for _, expr := range listExpression.Exprs {
		res := Eval(expr, env)
		list.Objects = append(list.Objects, res)
	}
	return list
}

func evalIdentifier(ident *ast.Identifier, env *object.Environment) object.Object {
	val, ok := env.Get(ident.Value)
	if !ok {
		return &object.Nil{}
	}
	return val
}

func evalExpressions(exprs []ast.Expression, env *object.Environment) object.Object {
	var result object.Object
	for _, expr := range exprs {
		result = Eval(expr, env)
	}
	return result
}
