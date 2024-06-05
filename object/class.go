package object

import "fmt"

type Class struct {
	Name string
}

func (c Class) String() string {
	return fmt.Sprintf("Class: %s", c.Name)
}
func (c Class) Type() ObjectType {
	return CLASS_OBJ
}

type Instance struct {
	Class *Class
}

func (i Instance) String() string {
	return fmt.Sprintf("Instance of %s", i.Class.Name)
}
func (i Instance) Type() ObjectType {
	return INSTANCE_OBJ
}
