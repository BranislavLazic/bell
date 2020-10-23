package object

import (
	"fmt"
	"strings"
)

type ObjectType string

const (
	IntegerObj      = "INTEGER"
	BooleanObj      = "BOOLEAN"
	ListObj         = "LIST"
	NilObj          = "NIL"
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
