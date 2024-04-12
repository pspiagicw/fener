package eval

import (
	"testing"

	"github.com/pspiagicw/fener/ast"
	"github.com/pspiagicw/fener/lexer"
	"github.com/pspiagicw/fener/object"
	"github.com/pspiagicw/fener/parser"
)

func TestEval(t *testing.T) {

	table := []struct {
		input string
		value interface{}
	}{
		{"5", 5},
		{`"some string"`, "some string"},
		// {"-10", -10},
	}

	for _, tt := range table {
		t.Run(tt.input, func(t *testing.T) {

			checkEval(t, tt.input, tt.value)

		})
	}
}

func checkEval(t *testing.T, input string, expected interface{}) {
	t.Helper()

	ast, errors := parse(input)

	if len(errors) > 0 {

		for _, err := range errors {
			t.Errorf(err)
		}

		t.Fatalf("Parsing failed")
	}

	e := New(func(err error) {
		t.Fatalf(err.Error())
	})

	value := e.Eval(ast)

	switch expected := expected.(type) {
	case int:
		checkIntegerObject(t, value, int64(expected))
	case string:
		checkStringObject(t, value, expected)
	default:
		t.Fatalf("Unknown type `%T` for testing", expected)
	}

}
func checkIntegerObject(t *testing.T, obj object.Object, expected int64) {
	t.Helper()

	result, ok := obj.(*object.Integer)

	if !ok {
		t.Fatalf("Object is not an Integer. Got: %T", obj)
	}

	if result.Value != expected {
		t.Fatalf("Expected %d, got %d", expected, result.Value)
	}
}
func checkStringObject(t *testing.T, obj object.Object, expected string) {
	t.Helper()

	result, ok := obj.(*object.String)

	if !ok {
		t.Fatalf("Object is not a String. Got: %T", obj)
	}

	if result.Value != expected {
		t.Fatalf("Expected %s, got %s", expected, result.Value)
	}
}

func parse(input string) (*ast.Program, []string) {

	l := lexer.New(input)
	p := parser.New(l)

	return p.Parse(), p.Errors()
}
