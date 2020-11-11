package object

import (
	"fmt"
	"github.com/branislavlazic/bell/ast"
	"strings"
)

type ObjectType string

const (
	IntegerObj      = "INTEGER"
	BooleanObj      = "BOOLEAN"
	StringObj       = "STRING"
	ListObj         = "LIST"
	FunctionObj     = "FUNCTION"
	NilObj          = "NIL"
	NoopObj         = "NOOP"
	BuiltinObj      = "BUILTIN"
	RuntimeErrorObj = "RUNTIME_ERROR"
)

type Object interface {
	Type() ObjectType
	Inspect() string
}

type Integer struct {
	Value int64
}

func (i *Integer) Type() ObjectType {
	return IntegerObj
}
func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

type Boolean struct {
	Value bool
}

func (b *Boolean) Type() ObjectType {
	return BooleanObj
}
func (b *Boolean) Inspect() string {
	return fmt.Sprintf("%t", b.Value)
}

type String struct {
	Value string
}

func (s *String) Type() ObjectType {
	return StringObj
}
func (s *String) Inspect() string {
	return fmt.Sprintf("%s", s.Value)
}

type List struct {
	Objects []Object
}

func (l *List) Type() ObjectType {
	return ListObj
}
func (l *List) Inspect() string {
	var exprsAsStrArr []string
	for _, obj := range l.Objects {
		exprsAsStrArr = append(exprsAsStrArr, obj.Inspect())
	}
	return strings.Join(exprsAsStrArr, " ")
}

type Function struct {
	Identifier *ast.Identifier
	Params     []*ast.Identifier
	Body       []ast.Expression
}

func (f *Function) Type() ObjectType {
	return FunctionObj
}
func (f *Function) Inspect() string {
	var params []string
	for _, obj := range f.Params {
		params = append(params, obj.String())
	}
	joinedParams := strings.Join(params, " ")
	if params != nil {
		return fmt.Sprintf("(%s %s)", f.Identifier.String(), joinedParams)
	}
	return fmt.Sprintf("(%s)", f.Identifier.String())
}

type RuntimeError struct {
	Error string
}

func (re *RuntimeError) Type() ObjectType {
	return RuntimeErrorObj
}
func (re *RuntimeError) Inspect() string {
	return re.Error
}

type Nil struct{}

func (n *Nil) Type() ObjectType {
	return NilObj
}
func (n *Nil) Inspect() string {
	return "nil"
}

type Noop struct{}

func (n *Noop) Type() ObjectType {
	return NoopObj
}
func (n *Noop) Inspect() string {
	return ""
}

type BuiltinFunction func(args ...Object) Object

type Builtin struct {
	Fn BuiltinFunction
}

func (b *Builtin) Type() ObjectType {
	return BuiltinObj
}
func (b *Builtin) Inspect() string {
	return "builtin"
}
