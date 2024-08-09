package vm

import (
	"testing"

	"github.com/pspiagicw/fener/code"
	"github.com/pspiagicw/fener/compile"
	"github.com/pspiagicw/fener/lexer"
	"github.com/pspiagicw/fener/object"
	"github.com/pspiagicw/fener/parser"
)

type vmTest struct {
	input string
	value interface{}
}

func TestVM(t *testing.T) {

	tt := []vmTest{
		{"1", 1},
	}

	testVM(t, tt)
}

func testVM(t *testing.T, tt []vmTest) {
	for _, tc := range tt {
		t.Run(tc.input, func(t *testing.T) {
			bytes, constants := compileInput(t, tc.input)

			vm := New(bytes, constants)

			err := vm.Run()

			if err != nil {
				t.Fatalf("error running vm: %v", err)
			}

			val := vm.StackTop()

			testValue(t, val, tc.value)
		})
	}
}
func compileInput(t *testing.T, input string) ([]byte, []object.Object) {
	l := lexer.New(input)
	p := parser.New(l)

	program := p.Parse()

	if len(p.Errors()) != 0 {
		for _, msg := range p.Errors() {
			t.Errorf("parser error: %v", msg)
		}
		t.FailNow()
	}

	c := compile.New()

	err := c.Compile(program)

	if err != nil {
		t.Fatalf("compiler error: %v", err)
	}

	bytecode := c.Instructions()
	constants := c.Constants()

	bytes := code.ToBytes(bytecode)

	return bytes, constants
}
func testValue(t *testing.T, val object.Object, expected interface{}) {

	switch expected := expected.(type) {
	case int:
		testIntegerObject(t, val, int64(expected))
	}
}
func testIntegerObject(t *testing.T, obj object.Object, expected int64) {
	result, ok := obj.(*object.Integer)

	if !ok {
		t.Fatalf("object is not Integer, while expected one. got=%T (%+v)", obj, obj)
	}

	if result.Value != expected {
		t.Errorf("object has wrong value. got=%d, want=%d", result.Value, expected)
	}
}
