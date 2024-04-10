package parser

import (
	"testing"

	"github.com/pspiagicw/fener/ast"
	"github.com/pspiagicw/fener/token"
)

func TestLambdaParser(t *testing.T) {
	input := `fn(x) x + 1 end`

	expectedTree := []ast.Statement{
		&ast.ExpressionStatement{
			Token: &token.Token{Type: token.FUNCTION, Value: "fn"},
			Expression: &ast.Lambda{
				Token: &token.Token{Type: token.FUNCTION, Value: "fn"},
				Arguments: []*ast.Identifier{
					&ast.Identifier{
						Token: &token.Token{Type: token.IDENT, Value: "x"},
						Value: "x",
					},
				},
				Body: &ast.BlockStatement{
					Statements: []ast.Statement{
						&ast.ExpressionStatement{
							Token: &token.Token{Type: token.IDENT, Value: "x"},
							Expression: &ast.InfixExpression{
								Left: &ast.Identifier{
									Token: &token.Token{Type: token.IDENT, Value: "x"},
									Value: "x",
								},
								Operator: token.PLUS,
								Right: &ast.Integer{
									Token: &token.Token{Type: token.INT, Value: "1"},
									Value: 1,
								},
							},
						},
					},
				},
			},
		},
	}

	checkTree(t, input, expectedTree)
}

func TestParserFunction(t *testing.T) {

	// Write expectedTree for above input

}
