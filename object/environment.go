package object

type Environment struct {
	Outer    *Environment
	Bindings map[string]Object
}

func NewEnvironment() *Environment {
	return &Environment{
		Outer:    nil,
		Bindings: make(map[string]Object),
	}
}
func (e *Environment) Set(name string, value Object) {
	e.Bindings[name] = value
}
func (e *Environment) Get(name string) Object {
	obj, ok := e.Bindings[name]

	if !ok && e.Outer != nil {
		return e.Outer.Get(name)
	}

	return obj
}
