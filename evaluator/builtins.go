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
			case *object.String:
				arr := []rune(arg.Value)
				if len(arr) == 0 {
					return &object.String{Value: ""}
				}
				return &object.String{Value: string(arr[0])}
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
			case *object.String:
				arr := []rune(arg.Value)
				if len(arr) == 0 {
					return &object.String{Value: ""}
				}
				return &object.String{Value: string(arr[1:])}
			default:
				return &object.RuntimeError{Error: fmt.Sprintf("Function is not applicable for %s type.", arg.Type())}
			}
		},
	},
	"size": {
		Fn: func(args ...object.Object) object.Object {
			lenArgs := len(args)
			if lenArgs != 1 {
				return &object.RuntimeError{
					Error: fmt.Sprintf("Insufficient number of arguments. Expected %d, got %d.", 1, lenArgs),
				}
			}
			switch arg := args[0].(type) {
			case *object.List:
				return &object.Integer{Value: int64(len(arg.Objects))}
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			default:
				return &object.RuntimeError{Error: fmt.Sprintf("Function is not applicable for %s type.", arg.Type())}
			}
		},
	},
	"write": {
		Fn: func(args ...object.Object) object.Object {
			for _, arg := range args {
				fmt.Print(arg.Inspect())
			}
			return &object.Nil{}
		},
	},
	"writeln": {
		Fn: func(args ...object.Object) object.Object {
			for _, arg := range args {
				fmt.Println(arg.Inspect())
			}
			return &object.Nil{}
		},
	},
}
