package compile

import (
	"fmt"

	"github.com/pspiagicw/fener/ast"
	"github.com/pspiagicw/fener/code"
	"github.com/pspiagicw/fener/object"
	"github.com/pspiagicw/fener/token"
)

type Compiler struct {
	instructions []*code.Instruction
	constants    []object.Object

	symbols *SymbolTable

	constID int
}

func New() *Compiler {
	return &Compiler{
		instructions: []*code.Instruction{},
		constants:    []object.Object{},
		constID:      0,

		symbols: NewSymbolTable(),
	}
}

func (c *Compiler) Instructions() []*code.Instruction {
	return c.instructions
}
func (c *Compiler) Constants() []object.Object {
	return c.constants
}

func (c *Compiler) Compile(node ast.Node) error {
	switch node := node.(type) {
	case *ast.Program:
		return c.compileProgram(node)
	case *ast.ExpressionStatement:
		return c.compileExpressionStatement(node.Expression)
	case *ast.Integer:
		return c.compileInteger(node)
	case *ast.InfixExpression:
		return c.compileInfixExpression(node)
	case *ast.Boolean:
		return c.compileBoolean(node)
	case *ast.PrefixExpression:
		return c.compilePrefixExpression(node)
	case *ast.AssignmentExpression:
		return c.compileAssignmentExpression(node)
	case *ast.Identifier:
		return c.compileIdentifier(node)
	default:
		return fmt.Errorf("unknown node type %T", node)
	}
}
func (c *Compiler) compileIdentifier(node *ast.Identifier) error {
	sym, ok := c.symbols.Resolve(node.Value)
	if !ok {
		return fmt.Errorf("undefined variable %s", node.Value)
	}
	return c.emit(code.GET, sym.Index)
}
func (c *Compiler) addAssignment(target *ast.Identifier) error {
	sym, ok := c.symbols.Resolve(target.Value)

	if !ok {
		sym = c.symbols.Define(target.Value)

	}
	return c.emit(code.SET, sym.Index)
}
func (c *Compiler) compileAssignmentExpression(node *ast.AssignmentExpression) error {
	err := c.Compile(node.Value)
	if err != nil {
		return err
	}

	target := node.Target
	switch target := target.(type) {
	case *ast.Identifier:
		return c.addAssignment(target)
	default:
		return fmt.Errorf("unknown target type %T", target)
	}
}
func (c *Compiler) compilePrefixExpression(node *ast.PrefixExpression) error {
	err := c.Compile(node.Right)
	if err != nil {
		return err
	}

	switch node.Operator {
	case token.BANG:
		return c.emit(code.NOT)
	default:
		return fmt.Errorf("unknown prefix operator '%s'", node.Operator)
	}
}
func (c *Compiler) compileBoolean(node *ast.Boolean) error {

	if node.Value {
		return c.emit(code.TRUE)
	}

	return c.emit(code.FALSE)
}
func (c *Compiler) compileInfixExpression(node *ast.InfixExpression) error {
	err := c.Compile(node.Left)
	if err != nil {
		return err
	}

	err = c.Compile(node.Right)
	if err != nil {
		return err
	}

	switch node.Operator {
	case token.PLUS:
		return c.emit(code.ADD)
	case token.MINUS:
		return c.emit(code.SUB)
	case token.MULTIPLY:
		return c.emit(code.MUL)
	case token.DIVIDE:
		return c.emit(code.DIV)
	case token.LT:
		return c.emit(code.LT)
	case token.GT:
		return c.emit(code.GT)
	case token.EQ:
		return c.emit(code.EQ)
	case token.AND:
		return c.emit(code.AND)
	case token.OR:
		return c.emit(code.OR)
	case token.NOT_EQ:
		return c.emit(code.NEQ)
	default:
		return fmt.Errorf("unknown infix operator '%s'", node.Operator)
	}
}
func (c *Compiler) compileInteger(node *ast.Integer) error {
	integer := &object.Integer{Value: node.Value}

	cid := c.addConstant(integer)

	return c.emit(code.PUSH, cid)
}
func (c *Compiler) compileExpressionStatement(node ast.Expression) error {
	err := c.Compile(node)
	if err != nil {
		return err
	}
	return nil
}
func (c *Compiler) compileProgram(node *ast.Program) error {
	for _, statement := range node.Statements {
		err := c.Compile(statement)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Compiler) addConstant(obj object.Object) int {
	c.constants = append(c.constants, obj)
	c.constID++
	return c.constID - 1
}
func (c *Compiler) emit(op code.OpCode, operands ...int) error {
	i := code.Make(op, operands...)

	if i == nil {
		return fmt.Errorf("failed to make instruction with opcode %s", op)
	}

	c.instructions = append(c.instructions, i)

	return nil
}
