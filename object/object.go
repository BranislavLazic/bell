package object

import "fmt"

type ObjectType string

const (
	IntegerObj      = "INTEGER"
	BooleanObj      = "BOOLEAN"
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
