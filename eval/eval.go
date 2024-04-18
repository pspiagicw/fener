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
		return nil
	}
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
func (e *Evaluator) evalPrefixExpression(node *ast.PrefixExpression) object.Object {

	right := e.Eval(node.Right)

	switch node.Operator {
	case token.MINUS:
		number := toInteger(right)

		if number == nil {
			e.Error("Can't negate expression %T", right)
			return nil
		}

		return &object.Integer{Value: -number.Value}
	case token.BANG:
		return &object.Boolean{Value: !isTruthy(right)}
	default:
		e.Error("Unknown prefix operator: %s", node.Operator)
		return nil
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
