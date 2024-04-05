package parser

import (
	"testing"

	"github.com/pspiagicw/fener/ast"
	"github.com/pspiagicw/fener/token"
)

func TestParserInfixSimple(t *testing.T) {
	input := `5 + 5`

	expectedTree := []ast.Statement{
		&ast.ExpressionStatement{
			Token: &token.Token{Type: token.INT, Value: "5", Line: 0},
			Expression: &ast.InfixExpression{
				Left:     &ast.Integer{Value: 5, Token: &token.Token{Type: token.INT, Value: "5", Line: 0}},
				Operator: token.PLUS,
				Right:    &ast.Integer{Value: 5, Token: &token.Token{Type: token.INT, Value: "5", Line: 0}},
			},
		},
	}

	checkTree(t, input, expectedTree)
}
func TestParserInfixMultiplication(t *testing.T) {
	input := `5 * 5`

	expectedTree := []ast.Statement{
		&ast.ExpressionStatement{
			Expression: &ast.InfixExpression{
				Left:     &ast.Integer{Value: 5, Token: &token.Token{Type: token.INT, Value: "5", Line: 0}},
				Operator: token.MULTIPLY,
				Right:    &ast.Integer{Value: 5, Token: &token.Token{Type: token.INT, Value: "5", Line: 0}},
			},
		},
	}

	checkTree(t, input, expectedTree)
}

func TestParserInfixSubtraction(t *testing.T) {
	input := `10 - 5`

	expectedTree := []ast.Statement{
		&ast.ExpressionStatement{
			Expression: &ast.InfixExpression{
				Left:     &ast.Integer{Value: 10, Token: &token.Token{Type: token.INT, Value: "10", Line: 0}},
				Operator: token.MINUS,
				Right:    &ast.Integer{Value: 5, Token: &token.Token{Type: token.INT, Value: "5", Line: 0}},
			},
		},
	}

	checkTree(t, input, expectedTree)
}

func TestParserInfixDivision(t *testing.T) {
	input := `10 / 2`

	expectedTree := []ast.Statement{
		&ast.ExpressionStatement{
			Expression: &ast.InfixExpression{
				Left:     &ast.Integer{Value: 10, Token: &token.Token{Type: token.INT, Value: "10", Line: 0}},
				Operator: token.DIVIDE,
				Right:    &ast.Integer{Value: 2, Token: &token.Token{Type: token.INT, Value: "2", Line: 0}},
			},
		},
	}

	checkTree(t, input, expectedTree)
}

func TestParserInfixModulus(t *testing.T) {
	input := `10 % 3`

	expectedTree := []ast.Statement{
		&ast.ExpressionStatement{
			Expression: &ast.InfixExpression{
				Left:     &ast.Integer{Value: 10, Token: &token.Token{Type: token.INT, Value: "10", Line: 0}},
				Operator: token.MOD,
				Right:    &ast.Integer{Value: 3, Token: &token.Token{Type: token.INT, Value: "3", Line: 0}},
			},
		},
	}

	checkTree(t, input, expectedTree)
}
func TestParserInfixComplexA(t *testing.T) {
	input := `5 + 2 * 10`

	expectedTree := []ast.Statement{
		&ast.ExpressionStatement{
			Expression: &ast.InfixExpression{
				Left:     &ast.Integer{Value: 5, Token: &token.Token{Type: token.INT, Value: "5", Line: 0}},
				Operator: token.PLUS,
				Right: &ast.InfixExpression{
					Left:     &ast.Integer{Value: 2, Token: &token.Token{Type: token.INT, Value: "2", Line: 0}},
					Operator: token.MULTIPLY,
					Right:    &ast.Integer{Value: 10, Token: &token.Token{Type: token.INT, Value: "10", Line: 0}},
				},
			},
		},
	}

	checkTree(t, input, expectedTree)
}
func TestParserInfixComplexB(t *testing.T) {
	input := `1 + 2 * 3 + 4 / 5 - 6`

	expectedTree := []ast.Statement{
		&ast.ExpressionStatement{
			Expression: &ast.InfixExpression{
				Left: &ast.InfixExpression{
					Left: &ast.InfixExpression{
						Left:     &ast.Integer{Value: 1, Token: &token.Token{Type: token.INT, Value: "1", Line: 0}},
						Operator: token.PLUS,
						Right: &ast.InfixExpression{
							Left:     &ast.Integer{Value: 2, Token: &token.Token{Type: token.INT, Value: "2", Line: 0}},
							Operator: token.MULTIPLY,
							Right:    &ast.Integer{Value: 3, Token: &token.Token{Type: token.INT, Value: "3", Line: 0}},
						},
					},
					Operator: token.PLUS,
					Right: &ast.InfixExpression{
						Left:     &ast.Integer{Value: 4, Token: &token.Token{Type: token.INT, Value: "4", Line: 0}},
						Operator: token.DIVIDE,
						Right:    &ast.Integer{Value: 5, Token: &token.Token{Type: token.INT, Value: "5", Line: 0}},
					},
				},
				Operator: token.MINUS,
				Right:    &ast.Integer{Value: 6, Token: &token.Token{Type: token.INT, Value: "6", Line: 0}},
			},
		},
	}

	checkTree(t, input, expectedTree)
}