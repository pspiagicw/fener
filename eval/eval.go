package eval

import (
	"fmt"

	"github.com/pspiagicw/fener/ast"
	"github.com/pspiagicw/fener/object"
	"github.com/pspiagicw/fener/token"
)

type Evaluator struct {
	ErrorHandler func(error)
	Test         bool
}

func New(handler func(error)) *Evaluator {
	return &Evaluator{
		Test:         false,
		ErrorHandler: handler,
	}
}

func (e *Evaluator) Error(message string, args ...interface{}) {

	err := fmt.Errorf(message, args...)

	e.ErrorHandler(err)
}

func (e *Evaluator) Eval(node ast.Node, env *object.Environment) object.Object {
	switch node := node.(type) {
	case *ast.WhileStatement:
		return e.evalWhileStatement(node, env)
	case *ast.TestStatement:
		return e.evalTestStatement(node, env)
	case *ast.ReturnStatement:
		return e.evalReturnStatement(node, env)
	case *ast.Lambda:
		return e.evalLambdaExpression(node, env)
	case *ast.CallExpression:
		return e.evalCallExpression(node, env)
	case *ast.FunctionStatement:
		return e.evalFunctionStatement(node, env)
	case *ast.Identifier:
		return e.evalIdentifier(node, env)
	case *ast.AssignmentExpression:
		return e.evalAssignmentExpression(node, env)
	case *ast.BlockStatement:
		return e.evalBlockStatement(node, env)
	case *ast.IfExpression:
		return e.evalIfExpression(node, env)
	case *ast.InfixExpression:
		return e.evalInfixExpression(node, env)
	case *ast.PrefixExpression:
		return e.evalPrefixExpression(node, env)
	case *ast.Boolean:
		return evalBoolean(node)
	case *ast.String:
		return evalString(node)
	case *ast.Integer:
		return evalInteger(node)
	case *ast.ExpressionStatement:
		return e.Eval(node.Expression, env)
	case *ast.Program:
		return e.evalProgram(node, env)
	default:
		e.Error("Unknown node type: %T", node)
		return &object.Null{}
	}
}
func (e *Evaluator) evalWhileStatement(node *ast.WhileStatement, env *object.Environment) object.Object {
	for true {
		value := e.Eval(node.Condition, env)

		if isTruthy(value) {
			e.Eval(node.Consequence, env)
		} else {
			break
		}
	}
	return &object.Null{}
}
func (e *Evaluator) evalIdentifier(node *ast.Identifier, env *object.Environment) object.Object {
	value := env.Get(node.Value)

	if value == nil {
		e.Error("Identifier not found: %s", node.Value)
		return &object.Null{}
	}

	return value
}
func (e *Evaluator) evalAssignmentExpression(node *ast.AssignmentExpression, env *object.Environment) object.Object {
	value := e.Eval(node.Value, env)

	env.Set(node.Target.Value, value)

	return value
}
func (e *Evaluator) evalBlockStatement(node *ast.BlockStatement, env *object.Environment) object.Object {
	var result object.Object

	result = &object.Null{}

	for _, statement := range node.Statements {
		result = e.Eval(statement, env)
		if result.Type() == object.RETURN_OBJ {
			return result
		}
	}

	return result
}
func (e *Evaluator) evalIfExpression(node *ast.IfExpression, env *object.Environment) object.Object {
	condition := e.Eval(node.Condition, env)

	if isTruthy(condition) {
		return e.Eval(node.Consequence, env)
	} else if len(node.Elif) > 0 {
		for condition, block := range node.Elif {
			if isTruthy(e.Eval(condition, env)) {
				return e.Eval(block, env)
			}
		}
	}
	if node.Alternative != nil {
		return e.Eval(node.Alternative, env)
	}
	return &object.Null{}
}
func evalString(node *ast.String) object.Object {
	return &object.String{Value: node.Value}
}

func evalInteger(node *ast.Integer) object.Object {
	return &object.Integer{Value: node.Value}
}
func evalBoolean(node *ast.Boolean) object.Object {
	return &object.Boolean{Value: node.Value}
}

func (e *Evaluator) evalProgram(node *ast.Program, env *object.Environment) object.Object {
	var result object.Object

	for _, statement := range node.Statements {
		result = e.Eval(statement, env)
	}

	return result
}

func (e *Evaluator) evalInfixExpression(node *ast.InfixExpression, env *object.Environment) object.Object {
	switch node.Operator {
	case token.MINUS, token.MULTIPLY, token.DIVIDE, token.PLUS:
		return e.evalInfixArithmetic(node, env)
	case token.EQ, token.NOT_EQ, token.GT, token.LT:
		return e.evalInfixComparison(node, env)
	}

	return &object.Null{}
}
func isEqual(left object.Object, right object.Object) bool {
	if left.Type() == right.Type() && left.String() == right.String() {
		return true
	}
	return false
}
func (e *Evaluator) evalInfixComparison(node *ast.InfixExpression, env *object.Environment) object.Object {
	left := e.Eval(node.Left, env)
	right := e.Eval(node.Right, env)

	switch node.Operator {
	case token.EQ:
		return &object.Boolean{Value: isEqual(left, right)}
	case token.NOT_EQ:
		return &object.Boolean{Value: !isEqual(left, right)}
	case token.GT:
		return e.evalGreaterThan(left, right)
	case token.LT:
		return e.evalLessThan(left, right)
	default:
		e.Error("Unknown infix comparison operator: %s", node.Operator)
		return &object.Null{}
	}
}
func (e *Evaluator) evalGreaterThan(left, right object.Object) object.Object {
	leftInt := toInteger(left)
	rightInt := toInteger(right)

	if leftInt == nil || rightInt == nil {
		e.Error("Can't compare expressions %T and %T", left, right)
		return &object.Null{}
	}
	return &object.Boolean{Value: leftInt.Value > rightInt.Value}

}
func (e *Evaluator) evalLessThan(left, right object.Object) object.Object {
	leftInt := toInteger(left)
	rightInt := toInteger(right)

	if leftInt == nil || rightInt == nil {
		e.Error("Can't compare expressions %T and %T", left, right)
		return &object.Null{}
	}
	return &object.Boolean{Value: leftInt.Value < rightInt.Value}
}
func (e *Evaluator) evalInfixArithmetic(node *ast.InfixExpression, env *object.Environment) object.Object {
	left := toInteger(e.Eval(node.Left, env))
	if left == nil {
		e.Error("Can't perform infix operation on left expression %T", node.Left)
		return &object.Null{}
	}
	right := toInteger(e.Eval(node.Right, env))
	if right == nil {
		e.Error("Can't perform infx operation on right expression %T", node.Right)
		return &object.Null{}
	}

	switch node.Operator {
	case token.PLUS:
		return &object.Integer{Value: left.Value + right.Value}
	case token.MINUS:
		return &object.Integer{Value: left.Value - right.Value}
	case token.MULTIPLY:
		return &object.Integer{Value: left.Value * right.Value}
	case token.DIVIDE:
		return &object.Integer{Value: left.Value / right.Value}
	default:
		e.Error("Unknown arithmetic infix operator: %s", node.Operator)
		return &object.Null{}
	}
}
func (e *Evaluator) negateValue(value object.Object) object.Object {
	number := toInteger(value)

	if number == nil {
		e.Error("Can't negate expression %T", value)
		return &object.Null{}
	}

	return &object.Integer{Value: -number.Value}

}
func (e *Evaluator) evalPrefixExpression(node *ast.PrefixExpression, env *object.Environment) object.Object {

	right := e.Eval(node.Right, env)

	switch node.Operator {
	case token.MINUS:
		return e.negateValue(right)
	case token.BANG:
		return &object.Boolean{Value: !isTruthy(right)}
	default:
		e.Error("Unknown prefix operator: %s", node.Operator)
		return &object.Null{}
	}
}
func toInteger(obj object.Object) *object.Integer {

	i, ok := obj.(*object.Integer)

	if !ok {
		return nil
	}

	return i
}

func isTruthy(obj object.Object) bool {
	switch obj := obj.(type) {
	case *object.Integer:
		return obj.Value != 0
	case *object.Boolean:
		return obj.Value
	default:
		return false
	}
}
func getArgumentNames(args []*ast.Identifier) []string {
	var names []string

	for _, arg := range args {
		names = append(names, arg.Value)
	}

	return names
}
func (e *Evaluator) evalFunctionStatement(node *ast.FunctionStatement, env *object.Environment) object.Object {

	args := getArgumentNames(node.Arguments)

	fn := &object.Function{
		Arguments: args,
		Body:      node.Body,
		Env:       env,
	}
	name := node.Target.Value

	env.Set(name, fn)

	return &object.Null{}
}
func (e *Evaluator) evalLambdaExpression(node *ast.Lambda, env *object.Environment) object.Object {
	args := getArgumentNames(node.Arguments)

	fn := &object.Function{
		Arguments: args,
		Body:      node.Body,
		Env:       env,
	}

	return fn
}
func newEnclosedEnvironment(outer *object.Environment) *object.Environment {
	env := object.NewEnvironment()

	env.Outer = outer

	return env
}
func (e *Evaluator) evalArgs(args []ast.Expression, env *object.Environment) []object.Object {
	var evaluated []object.Object

	for _, arg := range args {
		evaluated = append(evaluated, e.Eval(arg, env))
	}

	return evaluated
}
func (e *Evaluator) evalCallExpression(node *ast.CallExpression, env *object.Environment) object.Object {
	ex := e.Eval(node.Function, env)
	args := e.evalArgs(node.Arguments, env)

	switch ex := ex.(type) {
	case *object.Function:
		return e.evalFunctionCall(ex, args)
	case *object.Builtin:
		return e.evalBuiltinCall(ex, args)
	default:
		e.Error("Can't call expression %T", ex)
		return &object.Null{}
	}
}
func (e *Evaluator) evalBuiltinCall(fn *object.Builtin, args []object.Object) object.Object {
	return fn.Fn(args...)
}
func (e *Evaluator) evalFunctionCall(fn *object.Function, args []object.Object) object.Object {
	newEnv := newEnclosedEnvironment(fn.Env)

	e.applyArguments(fn.Arguments, args, newEnv)

	evaluated := e.Eval(fn.Body, newEnv)

	return unwrapReturnValue(evaluated)
}
func unwrapReturnValue(obj object.Object) object.Object {
	if returnValue, ok := obj.(*object.Return); ok {
		return returnValue.Value
	}
	return obj
}
func (e *Evaluator) applyArguments(params []string, args []object.Object, env *object.Environment) {

	if len(params) != len(args) {
		e.Error("Expected %d arguments, got %d", len(params), len(args))
		return
	}

	for i, arg := range args {
		env.Set(params[i], arg)
	}
}
func toFunction(node object.Object) *object.Function {
	fn, ok := node.(*object.Function)
	if !ok {
		return nil
	}
	return fn
}

func (e *Evaluator) evalReturnStatement(node *ast.ReturnStatement, env *object.Environment) object.Object {
	return &object.Return{Value: e.Eval(node.Value, env)}
}
func (e *Evaluator) evalTestStatement(node *ast.TestStatement, env *object.Environment) object.Object {
	if e.Test {
		return e.Eval(node.Statements, env)
	}
	return &object.Null{}
}
