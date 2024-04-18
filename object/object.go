package object

import "fmt"

type ObjectType string

const (
	INTEGER_OBJ = "INTEGER"
	STRING_OBJ  = "STRING"
	BOOLEAN_OBJ = "BOOLEAN"
)

type Object interface {
	Type() ObjectType
	String() string
}

type Integer struct {
	ObjType ObjectType
	Value   int64
}

func (i *Integer) Type() ObjectType { return i.ObjType }
func (i *Integer) String() string   { return fmt.Sprintf("%d", i.Value) }

type String struct {
	Value   string
	ObjType ObjectType
}

func (s *String) Type() ObjectType { return s.ObjType }
func (s *String) String() string   { return s.Value }

type Boolean struct {
	Value   bool
	ObjType ObjectType
}

func (b *Boolean) Type() ObjectType { return b.ObjType }
func (b *Boolean) String() string   { return fmt.Sprintf("%t", b.Value) }
