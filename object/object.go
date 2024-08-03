package object

import (
	"fmt"

	"github.com/pspiagicw/fener/ast"
)

type ObjectType string

const (
	INTEGER_OBJ = "INTEGER"
	STRING_OBJ  = "STRING"
	BOOLEAN_OBJ = "BOOLEAN"

	NULL_OBJ = "NULL"

	FUNCTION_OBJ = "FUNCTION"
	BULITIN_OBJ  = "BUILTIN"
	RETURN_OBJ   = "RETURN"
)

type Object interface {
	Type() ObjectType
	String() string
	Pretty() string
}

type Integer struct {
	Value int64
}

func (i *Integer) Type() ObjectType { return INTEGER_OBJ }
func (i *Integer) String() string {
	return fmt.Sprintf("int(%d)", i.Value)
}
func (i *Integer) Pretty() string {
	return fmt.Sprintf("%d", i.Value)
}

type String struct {
	Value string
}

func (s *String) Type() ObjectType { return STRING_OBJ }
func (s *String) String() string {
	return fmt.Sprintf("str(%s)", s.Value)
}
func (s *String) Pretty() string { return s.Value }

type Boolean struct {
	Value bool
}

func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }
func (b *Boolean) String() string {
	return fmt.Sprintf("bool(%t)", b.Value)
}
func (b *Boolean) Pretty() string { return fmt.Sprintf("%t", b.Value) }

type Null struct {
}

func (n *Null) Type() ObjectType { return NULL_OBJ }
func (n *Null) String() string   { return "null" }
func (n *Null) Pretty() string   { return "null" }

type Function struct {
	Env       *Environment
	Arguments []string
	Body      *ast.BlockStatement
}

func (f *Function) Type() ObjectType { return FUNCTION_OBJ }
func (f *Function) String() string {
	return fmt.Sprintf("function(%d args) at [%p]", len(f.Arguments), f)
}
func (f *Function) Pretty() string { return f.String() }

type Builtin struct {
	Name string
	Fn   func(args ...Object) (Object, error)
}

func (b *Builtin) Type() ObjectType { return BULITIN_OBJ }
func (b *Builtin) String() string {
	return fmt.Sprintf("builtin %s() at [%p]", b.Name, b)
}
func (b *Builtin) Pretty() string { return b.String() }

type Return struct {
	Value Object
}

func (r *Return) Type() ObjectType { return RETURN_OBJ }
func (r *Return) String() string   { return r.Value.String() }
func (r *Return) Pretty() string   { return r.Value.Pretty() }
