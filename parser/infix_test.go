package parser

import (
	"testing"

	"github.com/pspiagicw/fener/ast"
	"github.com/pspiagicw/fener/lexer"
	"github.com/pspiagicw/fener/token"
)

func TestBooleanParser(t *testing.T) {
	input := `
    true && true
    false || false

    1 & 2
    2 | 4
    `

	expectedTree := []ast.Statement{
		&ast.ExpressionStatement{
			Expression: &ast.InfixExpression{
				Left:     &ast.Boolean{Value: true, Token: &token.Token{Type: token.TRUE, Value: "true", Line: 0}},
				Operator: token.AND,
				Right:    &ast.Boolean{Value: true, Token: &token.Token{Type: token.TRUE, Value: "true", Line: 0}},
			},
		},
		&ast.ExpressionStatement{
			Expression: &ast.InfixExpression{
				Left:     &ast.Boolean{Value: false, Token: &token.Token{Type: token.FALSE, Value: "false", Line: 0}},
				Operator: token.OR,
				Right:    &ast.Boolean{Value: false, Token: &token.Token{Type: token.FALSE, Value: "false", Line: 0}},
			},
		},
		&ast.ExpressionStatement{
			Expression: &ast.InfixExpression{
				Left:     &ast.Integer{Value: 1, Token: &token.Token{Type: token.INT, Value: "1", Line: 0}},
				Operator: token.BITAND,
				Right:    &ast.Integer{Value: 2, Token: &token.Token{Type: token.INT, Value: "2", Line: 0}},
			},
		},
		&ast.ExpressionStatement{
			Expression: &ast.InfixExpression{
				Left:     &ast.Integer{Value: 2, Token: &token.Token{Type: token.INT, Value: "2", Line: 0}},
				Operator: token.BITOR,
				Right:    &ast.Integer{Value: 4, Token: &token.Token{Type: token.INT, Value: "4", Line: 0}},
			},
		},
	}

	checkTree(t, input, expectedTree)

}

func TestFunctionCall(t *testing.T) {
	input := `add() add(1, 2)`

	expectedTree := []ast.Statement{
		&ast.ExpressionStatement{
			Expression: &ast.CallExpression{
				Function:  &ast.Identifier{Value: "add", Token: &token.Token{Type: token.IDENT, Value: "add", Line: 0}},
				Arguments: []ast.Expression{},
			},
		},
		&ast.ExpressionStatement{
			Expression: &ast.CallExpression{
				Function: &ast.Identifier{Value: "add", Token: &token.Token{Type: token.IDENT, Value: "add", Line: 0}},
				Arguments: []ast.Expression{
					&ast.Integer{Value: 1, Token: &token.Token{Type: token.INT, Value: "1", Line: 0}},
					&ast.Integer{Value: 2, Token: &token.Token{Type: token.INT, Value: "2", Line: 0}},
				},
			},
		},
	}

	checkTree(t, input, expectedTree)
}

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
func TestParserInfixTable(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			"5 < 4 != 3 > 4",
			"((5 < 4) != (3 > 4))",
		},
		{
			"a + b + c",
			"((a + b) + c)",
		},
		{
			"a + b - c",
			"((a + b) - c)",
		},
		{
			"a * b * c",
			"((a * b) * c)",
		},
		{
			"a * b / c",
			"((a * b) / c)",
		},
		{
			"a + b / c",
			"(a + (b / c))",
		},
		{
			"a + b * c + d / e - f",
			"(((a + (b * c)) + (d / e)) - f)",
		},
		{
			"5 > 4 == 3 < 4",
			"((5 > 4) == (3 < 4))",
		},
		{
			"5 < 4 != 3 > 4",
			"((5 < 4) != (3 > 4))",
		},
		{
			"3 + 4 * 5 == 3 * 1 + 4 * 5",
			"((3 + (4 * 5)) == ((3 * 1) + (4 * 5)))",
		},
		{
			"3 + 4 * 5 == 3 * 1 + 4 * 5",
			"((3 + (4 * 5)) == ((3 * 1) + (4 * 5)))",
		},
		{
			"3 > 5 == true",
			"((3 > 5) == true)",
		},
		{
			"3 > 5 == false",
			"((3 > 5) == false)",
		},
		{
			"1 + (2 + 3) + 4",
			"((1 + (2 + 3)) + 4)",
		},
		{
			"(5 + 5) * 2",
			"((5 + 5) * 2)",
		},
		{
			"2 / (5 + 5)",
			"(2 / (5 + 5))",
		},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)

		tree := p.Parse()

		if len(p.errors) != 0 {
			for _, err := range p.errors {
				t.Logf("Parser error: %s", err)
			}
			t.Fatalf("Parser has %d errors", len(p.errors))
		}

		statement := tree.Statements[0]
		if tt.expected != statement.String() {
			t.Errorf("Expected %s, got %s", tt.expected, statement.String())
		}

	}
}
