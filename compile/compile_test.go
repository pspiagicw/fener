package compile

import (
	"testing"

	"github.com/pspiagicw/fener/code"
	"github.com/pspiagicw/fener/lexer"
	"github.com/pspiagicw/fener/object"
	"github.com/pspiagicw/fener/parser"
)

// func TestLambda(t *testing.T) {
// 	input := `fn something() end`
//
// 	constants := []interface{}{
// 		[]*code.Instruction{},
// 	}
//
// 	bytecode := []*code.Instruction{
// 		code.Make(code.PUSH, 0), // Push 0
// 		code.Make(code.SET, 0),  // Set something
// 	}
//
// 	testBytecode(t, input, bytecode, constants)
// }

func TestWhile(t *testing.T) {
	input := `while 2 < 3 then 10 end`

	constants := []interface{}{2, 3, 10}

	bytecode := []*code.Instruction{
		code.Make(code.PUSH, 0),  // Push 2 0000
		code.Make(code.PUSH, 1),  // Push 3 0003
		code.Make(code.LT),       // Less than 0006
		code.Make(code.JCMP, 16), // Jump if not true 0007
		code.Make(code.PUSH, 2),  // Push 10 0010
		code.Make(code.JMP, 0),   // Jump 0013
	}

	testBytecode(t, input, bytecode, constants)
}

func TestElif(t *testing.T) {
	t.Skip()
	input := `if 2 < 3 then 10 elif 3 < 4 then 20 else 30 end`

	constants := []interface{}{2, 3, 10, 4, 20, 30}

	bytecode := []*code.Instruction{
		code.Make(code.PUSH, 0),  // Push 2 0000
		code.Make(code.PUSH, 1),  // Push 3 0003
		code.Make(code.LT),       // Less than 0006
		code.Make(code.JCMP, 16), // Jump if not true 0007
		code.Make(code.PUSH, 2),  // Push 10 0010
		code.Make(code.JMP, 35),  // Jump 0013
		code.Make(code.PUSH, 3),  // Push 3 0016
		code.Make(code.PUSH, 4),  // Push 4 0019
		code.Make(code.LT),       // Less than 0022
		code.Make(code.JCMP, 32), // Jump if not true 0023
		code.Make(code.PUSH, 5),  // Push 20 0026
		code.Make(code.JMP, 35),  // Jump 0029
		code.Make(code.PUSH, 6),  // Push 30 0032
	}

	testBytecode(t, input, bytecode, constants)
}

func TestIf(t *testing.T) {
	input := `if 2 < 3 then 10 end`

	constants := []interface{}{2, 3, 10}

	bytecode := []*code.Instruction{
		code.Make(code.PUSH, 0),  // Push 2 0000
		code.Make(code.PUSH, 1),  // Push 3 0003
		code.Make(code.LT),       // Less than 0006
		code.Make(code.JCMP, 13), // Jump if not true 0007
		code.Make(code.PUSH, 2),  // Push 10 0010
	}
	testBytecode(t, input, bytecode, constants)
}

func TestIfExpressions(t *testing.T) {
	input := `if 2 < 3 then 10 else 20 end`

	constants := []interface{}{2, 3, 10, 20}

	bytecode := []*code.Instruction{
		code.Make(code.PUSH, 0),  // Push 2 0000
		code.Make(code.PUSH, 1),  // Push 3 0003
		code.Make(code.LT),       // Less than 0006
		code.Make(code.JCMP, 16), // Jump if not true 0007
		code.Make(code.PUSH, 2),  // Push 10 0010
		code.Make(code.JMP, 19),  // Jump 0013
		code.Make(code.PUSH, 3),  // Push 20 0016
	}

	testBytecode(t, input, bytecode, constants)
}

func TestVariableReassignment(t *testing.T) {
	input := `a = 5 a = 10`

	constants := []interface{}{5, 10}
	bytecode := []*code.Instruction{
		code.Make(code.PUSH, 0), // Push 5
		code.Make(code.SET, 0),  // Set variable 'a'
		code.Make(code.PUSH, 1), // Push 5
		code.Make(code.SET, 0),  // Set variable 'a'
	}
	testBytecode(t, input, bytecode, constants)
}

func TestVariableAssignment(t *testing.T) {
	input := `a = 5`

	constants := []interface{}{5}

	bytecode := []*code.Instruction{
		code.Make(code.PUSH, 0), // Push 5
		code.Make(code.SET, 0),  // Set variable 'a'
	}

	testBytecode(t, input, bytecode, constants)
}

func TestVariableRetrieval(t *testing.T) {
	input := `a = 5  a`

	constants := []interface{}{5}

	bytecode := []*code.Instruction{
		code.Make(code.PUSH, 0), // Push 5
		code.Make(code.SET, 0),  // Set variable 'a'
		code.Make(code.GET, 0),  // Get variable 'a'
	}

	testBytecode(t, input, bytecode, constants)
}

func TestVariableOperations(t *testing.T) {
	input := `a = 5  b = a + 2`

	constants := []interface{}{5, 2}

	bytecode := []*code.Instruction{
		code.Make(code.PUSH, 0), // Push 5
		code.Make(code.SET, 0),  // Set variable 'a'
		code.Make(code.GET, 0),  // Get variable 'a'
		code.Make(code.PUSH, 1), // Push 2
		code.Make(code.ADD),     // Add -> a + 2
		code.Make(code.SET, 1),  // Set variable 'b'
	}

	testBytecode(t, input, bytecode, constants)
}

func TestVariableReuse(t *testing.T) {
	input := `a = 5  b = 10  c = a + b`

	constants := []interface{}{5, 10}

	bytecode := []*code.Instruction{
		code.Make(code.PUSH, 0), // Push 5
		code.Make(code.SET, 0),  // Set variable 'a'
		code.Make(code.PUSH, 1), // Push 10
		code.Make(code.SET, 1),  // Set variable 'b'
		code.Make(code.GET, 0),  // Get variable 'a'
		code.Make(code.GET, 1),  // Get variable 'b'
		code.Make(code.ADD),     // Add -> a + b
		code.Make(code.SET, 2),  // Set variable 'c'
	}

	testBytecode(t, input, bytecode, constants)
}

func TestVariableInLogicalOperation(t *testing.T) {
	input := `a = true  b = false  c = a && b`

	constants := []interface{}{}

	bytecode := []*code.Instruction{
		code.Make(code.TRUE),   // Push true
		code.Make(code.SET, 0), // Set variable 'a'
		code.Make(code.FALSE),  // Push false
		code.Make(code.SET, 1), // Set variable 'b'
		code.Make(code.GET, 0), // Get variable 'a'
		code.Make(code.GET, 1), // Get variable 'b'
		code.Make(code.AND),    // Logical AND -> a && b
		code.Make(code.SET, 2), // Set variable 'c'
	}

	testBytecode(t, input, bytecode, constants)
}

func TestBooleanTrue(t *testing.T) {
	input := `true`

	constants := []interface{}{}

	bytecode := []*code.Instruction{
		code.Make(code.TRUE), // Push true
	}

	testBytecode(t, input, bytecode, constants)
}

func TestBooleanFalse(t *testing.T) {
	input := `false`

	constants := []interface{}{}

	bytecode := []*code.Instruction{
		code.Make(code.FALSE), // Push false
	}

	testBytecode(t, input, bytecode, constants)
}

func TestBooleanAnd(t *testing.T) {
	input := `true && false`

	constants := []interface{}{}

	bytecode := []*code.Instruction{
		code.Make(code.TRUE),  // Push true
		code.Make(code.FALSE), // Push false
		code.Make(code.AND),   // Logical AND
	}

	testBytecode(t, input, bytecode, constants)
}

func TestBooleanOr(t *testing.T) {
	input := `true || false`

	constants := []interface{}{}

	bytecode := []*code.Instruction{
		code.Make(code.TRUE),  // Push true
		code.Make(code.FALSE), // Push false
		code.Make(code.OR),    // Logical OR
	}

	testBytecode(t, input, bytecode, constants)
}

func TestBooleanNot(t *testing.T) {
	input := `!true`

	constants := []interface{}{}

	bytecode := []*code.Instruction{
		code.Make(code.TRUE), // Push true
		code.Make(code.NOT),  // Logical NOT
	}

	testBytecode(t, input, bytecode, constants)
}

func TestBooleanComplex(t *testing.T) {
	input := `(true && false) || !false`

	constants := []interface{}{}

	bytecode := []*code.Instruction{
		code.Make(code.TRUE),  // Push true
		code.Make(code.FALSE), // Push false
		code.Make(code.AND),   // Logical AND -> true && false
		code.Make(code.FALSE), // Push false
		code.Make(code.NOT),   // Logical NOT -> !false
		code.Make(code.OR),    // Logical OR -> (true && false) || !false
	}

	testBytecode(t, input, bytecode, constants)
}

func TestComplexLogicalOperations1(t *testing.T) {
	input := `(1 + 2) < (3 - 1)`

	constants := []interface{}{1, 2, 3, 1}

	bytecode := []*code.Instruction{
		code.Make(code.PUSH, 0), // Push 1
		code.Make(code.PUSH, 1), // Push 2
		code.Make(code.ADD),     // Add -> (1 + 2)
		code.Make(code.PUSH, 2), // Push 3
		code.Make(code.PUSH, 3), // Push 1
		code.Make(code.SUB),     // Subtract -> (3 - 1)
		code.Make(code.LT),      // Less than -> (1 + 2) < (3 - 1)
	}

	testBytecode(t, input, bytecode, constants)
}
func TestNotEQ(t *testing.T) {
	input := `1 != 2`

	constants := []interface{}{1, 2}

	bytecode := []*code.Instruction{
		code.Make(code.PUSH, 0), // Push 1
		code.Make(code.PUSH, 1), // Push 2
		code.Make(code.NEQ),     // Not Equal
	}
	testBytecode(t, input, bytecode, constants)
}

func TestComplexLogicalOperations2(t *testing.T) {
	input := `((1 + 2) > (3 - 1)) == (4 == 4)`

	constants := []interface{}{1, 2, 3, 1, 4, 4}

	bytecode := []*code.Instruction{
		code.Make(code.PUSH, 0), // Push 1
		code.Make(code.PUSH, 1), // Push 2
		code.Make(code.ADD),     // Add -> (1 + 2)
		code.Make(code.PUSH, 2), // Push 3
		code.Make(code.PUSH, 3), // Push 1
		code.Make(code.SUB),     // Subtract -> (3 - 1)
		code.Make(code.GT),      // Greater than -> (1 + 2) > (3 - 1)
		code.Make(code.PUSH, 4), // Push 4
		code.Make(code.PUSH, 5), // Push 4
		code.Make(code.EQ),      // Equal -> 4 == 4
		code.Make(code.EQ),      // Equal -> ((1 + 2) > (3 - 1)) == (4 == 4)
	}

	testBytecode(t, input, bytecode, constants)
}

func TestComplexLogicalOperations3(t *testing.T) {
	input := `((1 * 2) < (3 + 4)) && ((5 / 1) == 5)`

	constants := []interface{}{1, 2, 3, 4, 5, 1, 5}

	bytecode := []*code.Instruction{
		code.Make(code.PUSH, 0), // Push 1
		code.Make(code.PUSH, 1), // Push 2
		code.Make(code.MUL),     // Multiply -> (1 * 2)
		code.Make(code.PUSH, 2), // Push 3
		code.Make(code.PUSH, 3), // Push 4
		code.Make(code.ADD),     // Add -> (3 + 4)
		code.Make(code.LT),      // Less than -> (1 * 2) < (3 + 4)
		code.Make(code.PUSH, 4), // Push 5
		code.Make(code.PUSH, 5), // Push 1
		code.Make(code.DIV),     // Divide -> (5 / 1)
		code.Make(code.PUSH, 6), // Push 5
		code.Make(code.EQ),      // Equal -> (5 / 1) == 5
		code.Make(code.AND),     // Logical AND -> ((1 * 2) < (3 + 4)) && ((5 / 1) == 5)
	}

	testBytecode(t, input, bytecode, constants)
}

func TestComplexLogicalOperations4(t *testing.T) {
	input := `!(3 < 2) || (4 > 1)`

	constants := []interface{}{3, 2, 4, 1}

	bytecode := []*code.Instruction{
		code.Make(code.PUSH, 0), // Push 3
		code.Make(code.PUSH, 1), // Push 2
		code.Make(code.LT),      // Less than -> 3 < 2
		code.Make(code.NOT),     // Logical NOT -> !(3 < 2)
		code.Make(code.PUSH, 2), // Push 4
		code.Make(code.PUSH, 3), // Push 1
		code.Make(code.GT),      // Greater than -> 4 > 1
		code.Make(code.OR),      // Logical OR -> !(3 < 2) || (4 > 1)
	}

	testBytecode(t, input, bytecode, constants)
}

func TestLessThan(t *testing.T) {
	input := `1 < 2`

	constants := []interface{}{1, 2}

	bytecode := []*code.Instruction{
		code.Make(code.PUSH, 0), // Push 1
		code.Make(code.PUSH, 1), // Push 2
		code.Make(code.LT),      // Less than
	}

	testBytecode(t, input, bytecode, constants)
}
func TestGreaterThan(t *testing.T) {
	input := `3 > 2`

	constants := []interface{}{3, 2}

	bytecode := []*code.Instruction{
		code.Make(code.PUSH, 0), // Push 3
		code.Make(code.PUSH, 1), // Push 2
		code.Make(code.GT),      // Greater than
	}

	testBytecode(t, input, bytecode, constants)
}

func TestEqual(t *testing.T) {
	input := `4 == 4`

	constants := []interface{}{4, 4}

	bytecode := []*code.Instruction{
		code.Make(code.PUSH, 0), // Push 4
		code.Make(code.PUSH, 1), // Push 4
		code.Make(code.EQ),      // Equal
	}

	testBytecode(t, input, bytecode, constants)
}

func TestAllArithmeticOperations(t *testing.T) {
	input := `(1 + 2) * (3 - 4) / 5`

	constants := []interface{}{1, 2, 3, 4, 5}

	bytecode := []*code.Instruction{
		code.Make(code.PUSH, 0), // Push 1
		code.Make(code.PUSH, 1), // Push 2
		code.Make(code.ADD),     // Add -> (1 + 2)
		code.Make(code.PUSH, 2), // Push 3
		code.Make(code.PUSH, 3), // Push 4
		code.Make(code.SUB),     // Subtract -> (3 - 4)
		code.Make(code.MUL),     // Multiply -> (1 + 2) * (3 - 4)
		code.Make(code.PUSH, 4), // Push 5
		code.Make(code.DIV),     // Divide -> ((1 + 2) * (3 - 4)) / 5
	}

	testBytecode(t, input, bytecode, constants)
}

func TestSubtraction(t *testing.T) {
	input := `3 - 1`

	constants := []interface{}{3, 1}

	bytecode := []*code.Instruction{
		code.Make(code.PUSH, 0),
		code.Make(code.PUSH, 1),
		code.Make(code.SUB),
	}

	testBytecode(t, input, bytecode, constants)
}

func TestMultiplication(t *testing.T) {
	input := `2 * 3`

	constants := []interface{}{2, 3}

	bytecode := []*code.Instruction{
		code.Make(code.PUSH, 0),
		code.Make(code.PUSH, 1),
		code.Make(code.MUL),
	}

	testBytecode(t, input, bytecode, constants)
}

func TestDivision(t *testing.T) {
	input := `6 / 2`

	constants := []interface{}{6, 2}

	bytecode := []*code.Instruction{
		code.Make(code.PUSH, 0),
		code.Make(code.PUSH, 1),
		code.Make(code.DIV),
	}

	testBytecode(t, input, bytecode, constants)
}

func TestAddition(t *testing.T) {
	input := `1 + 2`

	constants := []interface{}{1, 2}

	bytecode := []*code.Instruction{
		code.Make(code.PUSH, 0),
		code.Make(code.PUSH, 1),
		code.Make(code.ADD),
	}

	testBytecode(t, input, bytecode, constants)
}

func TestPush(t *testing.T) {

	input := `1`

	constants := []interface{}{1}

	bytecode := []*code.Instruction{
		code.Make(code.PUSH, 0),
	}

	testBytecode(t, input, bytecode, constants)
}

func testBytecode(t *testing.T, input string, expected []*code.Instruction, expectedConstants []interface{}) {
	l := lexer.New(input)
	p := parser.New(l)

	program := p.Parse()

	if len(p.Errors()) != 0 {
		for _, err := range p.Errors() {
			t.Errorf("parser error: %q", err)
		}
		t.Fatalf("parser has %d errors", len(p.Errors()))
	}

	c := New()

	err := c.Compile(program)

	if err != nil {
		t.Fatalf("compiler error: %v", err)
	}

	bytecode := c.Instructions()
	constants := c.Constants()

	checkInstructions(t, bytecode, expected)
	checkConstants(t, constants, expectedConstants)
}
func checkInstructions(t *testing.T, actual []*code.Instruction, expected []*code.Instruction) {
	t.Helper()

	if len(actual) != len(expected) {
		t.Fatalf("wrong instructions length.\nwant=%d\ngot=%d\n", len(expected), len(actual))
	}

	for i, ins := range expected {
		testInstruction(t, actual[i], ins)
	}
}
func testInstruction(t *testing.T, actual *code.Instruction, expected *code.Instruction) {
	t.Helper()

	if actual == nil {
		t.Fatalf("actual instruction received is nil")
	}

	if actual.Opcode != expected.Opcode {
		t.Fatalf("wrong opcode.\nwant=%q\ngot=%q\n", expected.Opcode, actual.Opcode)
	}

	if len(actual.Operands) != len(expected.Operands) {
		t.Fatalf("wrong operands length.\nwant=%q\ngot=%q\n", expected.Operands, actual.Operands)
	}

	for i, operand := range expected.Operands {
		if actual.Operands[i] != operand {
			t.Fatalf("wrong operand at %d.\nwant=%d\ngot=%d\n", i, operand, actual.Operands[i])
		}
	}
}
func checkConstants(t *testing.T, actual []object.Object, expected []interface{}) {
	t.Helper()

	if len(actual) != len(expected) {
		t.Fatalf("wrong constants length.\nwant=%q\ngot=%q\n", expected, actual)
	}

	for i, constant := range expected {
		switch constant := constant.(type) {
		case int:
			testIntegerObject(t, actual[i], constant)
		default:
			t.Fatalf("Can't compare constant of type %T", constant)
		}
	}
}
func testIntegerObject(t *testing.T, obj object.Object, expected int) {
	t.Helper()

	result, ok := obj.(*object.Integer)

	if !ok {
		t.Fatalf("object is not Integer. got=%T (%+v)", obj, obj)
	}

	if result.Value != int64(expected) {
		t.Fatalf("object has wrong value. got=%d, want=%d", result.Value, expected)
	}
}
