package parser

import (
	"testing"

	"github.com/pspiagicw/fener/ast"
	"github.com/pspiagicw/fener/token"
)

func TestLambdaParser(t *testing.T) {
	input := `a = fn(x) x + 1 end`

	expectedTree := []ast.Statement{
		&ast.ExpressionStatement{
			Token: &token.Token{Type: token.IDENT, Value: "a"},
			Expression: &ast.AssignmentExpression{
				Token: &token.Token{Type: token.ASSIGN, Value: "="},
				Target: &ast.Identifier{
					Token: &token.Token{Type: token.IDENT, Value: "a"},
					Value: "a",
				},
				Value: &ast.Lambda{
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
		},
	}

	checkTree(t, input, expectedTree)
}

func TestParserFunction(t *testing.T) {

	input := `
    fn hello() 
        print("Hello, World!")
    end

    fn add(a, b)
        return a + b
    end
    `
	// Write expectedTree for above input

	expectedTree := []ast.Statement{
		&ast.FunctionStatement{
			Token: &token.Token{Type: token.FUNCTION, Value: "fn"},
			Target: &ast.Identifier{
				Token: &token.Token{Type: token.IDENT, Value: "hello"},
				Value: "hello",
			},
			Arguments: []*ast.Identifier{},
			Body: &ast.BlockStatement{
				Statements: []ast.Statement{
					&ast.ExpressionStatement{
						Token: &token.Token{Type: token.IDENT, Value: "print"},
						Expression: &ast.CallExpression{
							Token: &token.Token{Type: token.IDENT, Value: "print"},
							Function: &ast.Identifier{
								Token: &token.Token{Type: token.IDENT, Value: "print"},
								Value: "print",
							},
							Arguments: []ast.Expression{
								&ast.String{
									Token: &token.Token{Type: token.STRING, Value: "Hello, World!"},
									Value: "Hello, World!",
								},
							},
						},
					},
				},
			},
		},
		&ast.FunctionStatement{
			Token: &token.Token{Type: token.FUNCTION, Value: "fn"},
			Target: &ast.Identifier{
				Token: &token.Token{Type: token.IDENT, Value: "add"},
				Value: "add",
			},
			Arguments: []*ast.Identifier{
				&ast.Identifier{
					Token: &token.Token{Type: token.IDENT, Value: "a"},
					Value: "a",
				},

				&ast.Identifier{
					Token: &token.Token{Type: token.IDENT, Value: "b"},
					Value: "b",
				},
			},

			Body: &ast.BlockStatement{
				Statements: []ast.Statement{
					&ast.ReturnStatement{
						Token: &token.Token{Type: token.RETURN, Value: "return"},
						Value: &ast.InfixExpression{
							Left: &ast.Identifier{
								Token: &token.Token{Type: token.IDENT, Value: "a"},
								Value: "a",
							},
							Operator: token.PLUS,
							Right: &ast.Identifier{
								Token: &token.Token{Type: token.IDENT, Value: "b"},
								Value: "b",
							},
						},
					},
				},
			},
		},
	}

	checkTree(t, input, expectedTree)

}
