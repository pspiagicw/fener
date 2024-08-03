package object

import (
	"fmt"
	"strings"
)

func initBuiltins(env *Environment) {
	insertBuiltin := func(name string, fn func(...Object) (Object, error)) {
		env.Set(name, &Builtin{Fn: fn, Name: name})
	}
	insertBuiltin("print", printFunc)
	insertBuiltin("upper", upperFunc)
}
func printFunc(args ...Object) (Object, error) {
	for _, arg := range args {
		fmt.Println(arg.Pretty())
	}

	return &Null{}, nil
}
func upperFunc(args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("wrong number of arguments for UPPER. got=%d, want=1", len(args))
	}
	str, err := toString(args[0])
	if err != nil {
		return nil, err
	}
	return &String{Value: strings.ToUpper(str)}, nil
}
func toString(obj Object) (string, error) {
	if obj.Type() != STRING_OBJ {
		return "", fmt.Errorf("argument should be 'string', got %s", obj.Type())
	}
	str := obj.(*String)
	return str.Value, nil
}
