package evaluator

import (
	"fmt"
	"github.com/branislavlazic/bell/ast"
)

func Eval(node ast.Node) {
	switch node := node.(type) {
	case *ast.IntegerLiteral:
		fmt.Println(node.Value)
	}
}

func evalStatements(stmts []ast.Statement) {
	for _, stmt := range stmts {
		Eval(stmt.Expr)
	}
}
