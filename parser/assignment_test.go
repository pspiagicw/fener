package parser

import (
	"testing"

	"github.com/pspiagicw/fener/ast"
	"github.com/pspiagicw/fener/token"
)

func TestParserAssignment(t *testing.T) {
	input := `
    a = 1
    b = "something"
    `

	expectedTree := []ast.Statement{
		&ast.ExpressionStatement{
			Expression: &ast.AssignmentExpression{
				Token:  &token.Token{Type: token.ASSIGN, Value: "=", Line: 1},
				Target: &ast.Identifier{Value: "a", Token: &token.Token{Type: token.IDENT, Value: "a", Line: 1}},
				Value:  &ast.Integer{Value: 1, Token: &token.Token{Type: token.INT, Value: "1", Line: 1}},
			},
		},
		&ast.ExpressionStatement{
			Expression: &ast.AssignmentExpression{
				Token:  &token.Token{Type: token.ASSIGN, Value: "=", Line: 2},
				Target: &ast.Identifier{Value: "b", Token: &token.Token{Type: token.IDENT, Value: "b", Line: 2}},
				Value:  &ast.String{Value: "something", Token: &token.Token{Type: token.STRING, Value: "something", Line: 2}},
			},
		},
	}

	checkTree(t, input, expectedTree)
}
