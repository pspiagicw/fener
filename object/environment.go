package object

type Environment struct {
	Outer    *Environment
	Bindings map[string]Object
}

func NewEnvironment() *Environment {
	env := &Environment{
		Outer:    nil,
		Bindings: make(map[string]Object),
	}
	initBuiltins(env)
	return env
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
