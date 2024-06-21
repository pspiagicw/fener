package object

import "fmt"

func initBuiltins(env *Environment) {
	insertBuiltin := func(name string, fn func(...Object) Object) {
		env.Set(name, &Builtin{Fn: fn, Name: name})
	}
	insertBuiltin("print", printFunc)
}
func printFunc(args ...Object) Object {
	for _, arg := range args {
		fmt.Println(arg.Pretty())
	}

	return &Null{}
}
