package parser

import (
	"testing"

	"github.com/pspiagicw/fener/ast"
	"github.com/pspiagicw/fener/token"
)

func TestNegatePrefix(t *testing.T) {
	input := `-5`

	expectedTree := []ast.Statement{
		&ast.ExpressionStatement{
			Expression: &ast.PrefixExpression{
				Operator: token.MINUS,
				Right:    &ast.Integer{Value: 5},
			},
		},
	}

	checkTree(t, input, expectedTree)
}
func TestBangPrefix(t *testing.T) {
	input := `!true !false`

	expectedTree := []ast.Statement{
		&ast.ExpressionStatement{
			Expression: &ast.PrefixExpression{
				Operator: token.BANG,
				Right:    &ast.Boolean{Value: true},
			},
		},
		&ast.ExpressionStatement{
			Expression: &ast.PrefixExpression{
				Operator: token.BANG,
				Right:    &ast.Boolean{Value: false},
			},
		},
	}

	checkTree(t, input, expectedTree)
}
