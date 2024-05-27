package eval

import (
	"fmt"

	"github.com/pspiagicw/fener/ast"
	"github.com/pspiagicw/fener/object"
	"github.com/pspiagicw/fener/token"
)

type Evaluator struct {
	ErrorHandler func(error)
}

func New(handler func(error)) *Evaluator {
	return &Evaluator{
		ErrorHandler: handler,
	}
}

func (e *Evaluator) Error(message string, args ...interface{}) {

	err := fmt.Errorf(message, args...)

	e.ErrorHandler(err)
}

func (e *Evaluator) Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.BlockStatement:
		return e.EvalBlockStatement(node)
	case *ast.IfExpression:
		return e.evalIfExpression(node)
	case *ast.InfixExpression:
		return e.evalInfixExpression(node)
	case *ast.PrefixExpression:
		return e.evalPrefixExpression(node)
	case *ast.Boolean:
		return evalBoolean(node)
	case *ast.String:
		return evalString(node)
	case *ast.Integer:
		return evalInteger(node)
	case *ast.ExpressionStatement:
		return e.Eval(node.Expression)
	case *ast.Program:
		return e.evalProgram(node)
	default:
		e.Error("Unknown node type: %T", node)
		return &object.Null{}
	}
}
func (e *Evaluator) EvalBlockStatement(node *ast.BlockStatement) object.Object {
	var result object.Object

	result = &object.Null{}

	for _, statement := range node.Statements {
		result = e.Eval(statement)
	}

	return result
}
func (e *Evaluator) evalIfExpression(node *ast.IfExpression) object.Object {
	condition := e.Eval(node.Condition)

	if isTruthy(condition) {
		return e.Eval(node.Consequence)
	} else if node.Alternative != nil {
		return e.Eval(node.Alternative)
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

func (e *Evaluator) evalProgram(node *ast.Program) object.Object {
	var result object.Object

	for _, statement := range node.Statements {
		result = e.Eval(statement)
	}

	return result
}

func (e *Evaluator) evalInfixExpression(node *ast.InfixExpression) object.Object {
	switch node.Operator {
	case token.MINUS, token.MULTIPLY, token.DIVIDE, token.PLUS:
		return e.evalInfixArithmetic(node)
	case token.EQ, token.NOT_EQ, token.GT, token.LT:
		return e.evalInfixComparison(node)
	}

	return &object.Null{}
}
func isEqual(left object.Object, right object.Object) bool {
	if left.Type() == right.Type() && left.String() == right.String() {
		return true
	}
	return false
}
func (e *Evaluator) evalInfixComparison(node *ast.InfixExpression) object.Object {
	left := e.Eval(node.Left)
	right := e.Eval(node.Right)

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
func (e *Evaluator) evalInfixArithmetic(node *ast.InfixExpression) object.Object {
	left := toInteger(e.Eval(node.Left))
	if left == nil {
		e.Error("Can't perform infix operation on left expression %T", node.Left)
	}
	right := toInteger(e.Eval(node.Right))
	if right == nil {
		e.Error("Can't perform infx operation on right expression %T", node.Right)
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
func (e *Evaluator) evalPrefixExpression(node *ast.PrefixExpression) object.Object {

	right := e.Eval(node.Right)

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
