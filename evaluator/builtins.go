package evaluator

import (
	"fmt"
	"github.com/branislavlazic/bell/object"
)

var builtins = map[string]*object.Builtin{
	"head": {
		Fn: func(args ...object.Object) object.Object {
			lenArgs := len(args)
			if lenArgs != 1 {
				return &object.RuntimeError{
					Error: fmt.Sprintf("Insufficient number of arguments. Expected %d, got %d.", 1, lenArgs),
				}
			}
			switch arg := args[0].(type) {
			case *object.List:
				return arg.Objects[0]
			default:
				return &object.RuntimeError{Error: fmt.Sprintf("Function is not applicable for %s type.", arg.Type())}
			}
		},
	},
	"tail": {
		Fn: func(args ...object.Object) object.Object {
			lenArgs := len(args)
			if lenArgs != 1 {
				return &object.RuntimeError{
					Error: fmt.Sprintf("Insufficient number of arguments. Expected %d, got %d.", 1, lenArgs),
				}
			}
			switch arg := args[0].(type) {
			case *object.List:
				objects := arg.Objects[1:len(arg.Objects)]
				if len(objects) == 0 {
					return &object.Nil{}
				}
				return &object.List{Objects: objects}
			default:
				return &object.RuntimeError{Error: fmt.Sprintf("Function is not applicable for %s type.", arg.Type())}
			}
		},
	},
}
