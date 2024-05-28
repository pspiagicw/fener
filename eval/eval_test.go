package eval

import (
	"testing"

	"github.com/pspiagicw/fener/ast"
	"github.com/pspiagicw/fener/lexer"
	"github.com/pspiagicw/fener/object"
	"github.com/pspiagicw/fener/parser"
)

type testCase struct {
	input string
	value interface{}
}

func TestTest(t *testing.T) {
	table := []testCase{
		{
			`test "simple test"
                assert(1,2)
             end`,
			nil,
		},
	}
	runTableTests(t, table)
}
func TestBuiltin(t *testing.T) {
	table := []testCase{
		{
			`print("Hello!")`,
			nil,
		},
	}
	runTableTests(t, table)
}

func TestClosure(t *testing.T) {
	table := []testCase{
		{
			`fn counter()
                count = 0
                return fn() count = count + 1 end
             end
             c = counter()
             c()
             `,
			1,
		},
	}
	runTableTests(t, table)
}

func TestLambda(t *testing.T) {

	table := []testCase{
		{"a = fn() 5 end a()", 5},
		{"add = fn(x,y) x - y end add(3,2)", 1},
	}

	runTableTests(t, table)
}

func TestFunction(t *testing.T) {
	table := []testCase{
		{"fn example() 5 end example()", 5},
		{"fn add(x,y) x + y end add(2,3)", 5},
	}

	runTableTests(t, table)
}

func TestAssignment(t *testing.T) {
	table := []testCase{
		{"a = 5 a", 5},
		{"a = 5 b = 3 a+b", 8},
		{"a = 5 b = 3  c = 6 a+b+6", 14},
		{"a = 5 b = a  c = 6 a+b+6", 16},
	}

	runTableTests(t, table)
}

func TestIf(t *testing.T) {
	table := []testCase{
		{"if true then 5 end", 5},
		{"if 2 == 2 then 5 end", 5},
		{"if 2 == 3 then 5 end", nil},
		{"if 2 != 2 then 5 else 2 end", 2},
		{"if 2 != 2 then elif 2 == 3 then 7 else 9 end", 9},
	}

	runTableTests(t, table)

}

func TestInfix(t *testing.T) {

	table := []testCase{
		{"5", 5},
		{`"some string"`, "some string"},
		{"-10", -10},
		{"true", true},
		{"false", false},
		{"!true", false},
		{"!false", true},
		{"!!!false", true},
		{"!10", false},
		{"!true", false},
		{"!false", true},
		{"!5", false},
		{"!!true", true},
		{"!!false", false},
		{"!!5", true},
		{"true", true},
		{"false", false},
		{"1 < 2", true},
		{"1 > 2", false},
		{"1 < 1", false},
		{"1 > 1", false},
		{"1 == 1", true},
		{"1 != 1", false},
		{"1 == 2", false},
		{"1 != 2", true},
		{"true == true", true},
		{"false == false", true},
		{"true == false", false},
		{"true != false", true},
		{"false != true", true},
		{"(1 < 2) == true", true},
		{"(1 < 2) == false", false},
		{"(1 > 2) == true", false},
		{"(1 > 2) == false", true},
		{"5", 5},
		{"15", 15},
		{"-5", -5},
		{"-10", -10},
		{"5 + 5 + 5 + 5 - 10", 10},
		{"2 * 2 * 2 * 2 * 2", 32},
		{"-50 + 100 + -50", 0},
		{"5 * 2 + 10", 20},
		{"5 + 2 * 10", 25},
		{"20 + 2 * -10", 0},
		{"50 / 2 * 2 + 10", 60},
		{"2 * (5 + 10)", 30},
		{"3 * 3 * 3 + 10", 37},
		{"3 * (3 * 3) + 10", 37},
		{"(5 + 10 * 2 + 15 / 3) * 2 + -10", 50},
	}

	runTableTests(t, table)
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
	env := object.NewEnvironment()

	e := New(func(err error) {
		t.Fatalf(err.Error())
	})

	value := e.Eval(ast, env)

	switch expected := expected.(type) {
	case int:
		checkIntegerObject(t, value, int64(expected))
	case string:
		checkStringObject(t, value, expected)
	case bool:
		checkBooleanObject(t, value, expected)
	case nil:
		checkNullObject(t, value)
	default:
		t.Fatalf("Unknown type `%T` for testing", expected)
	}

}
func checkNullObject(t *testing.T, obj object.Object) {
	t.Helper()

	if obj.Type() != object.NULL_OBJ {
		t.Fatalf("Expected NULL, got %T", obj)
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
func checkBooleanObject(t *testing.T, obj object.Object, expected bool) {
	t.Helper()

	result, ok := obj.(*object.Boolean)

	if !ok {
		t.Fatalf("Object is not a Boolean. Got: %T", obj)
	}

	if result.Value != expected {
		t.Fatalf("Expected %t, got %t", expected, result.Value)
	}
}
func runTableTests(t *testing.T, table []testCase) {
	for _, tt := range table {
		t.Run(tt.input, func(t *testing.T) {

			checkEval(t, tt.input, tt.value)

		})
	}
}
