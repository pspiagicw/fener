package parser

import "testing"

func TestParserIdentifierExpressionOnly(t *testing.T) {
	input := `identifier`

	expectedTree := []ast.Statement{
		&ast.ExpressionStatement{
			Expression: &ast.IdentifierExpression{Token: token.Token{TokenType: token.IDENT, TokenValue: "identifier"}},
		},
	}

	checkTree(t, input, expectedTree)
}
func TestParserAssignmentStatements(t *testing.T) {
	input := `
    a = 1
    name = "myName"
    flag = true
    result = 5 * 3 - 2
    `

	expectedTree := []ast.Statement{
		&ast.AssignmentStatement{
			Name:  token.Token{TokenType: token.IDENT, TokenValue: "a"},
			Value: &ast.IntExpression{Value: 1},
		},
		&ast.AssignmentStatement{
			Name:  token.Token{TokenType: token.IDENT, TokenValue: "name"},
			Value: &ast.StringExpression{Value: "myName"},
		},
		&ast.AssignmentStatement{
			Name:  token.Token{TokenType: token.IDENT, TokenValue: "flag"},
			Value: &ast.BoolExpression{Value: true},
		},
		&ast.AssignmentStatement{
			Name: token.Token{TokenType: token.IDENT, TokenValue: "result"},
			Value: &ast.InfixExpression{
				Left: &ast.InfixExpression{
					Left:     &ast.IntExpression{Value: 5},
					Operator: token.MULTIPLY,
					Right:    &ast.IntExpression{Value: 3},
				},
				Operator: token.MINUS,
				Right:    &ast.IntExpression{Value: 2},
			},
		},
	}

	checkTree(t, input, expectedTree)
}

func TestParserInfixExpressions(t *testing.T) {
	input := `
    1 + 2
    3 - 4
    5 * 6
    7 / 8
    9 % 10
    true and false
    false or true
    `

	expectedTree := []ast.Statement{
		&ast.InfixExpression{
			Left:     &ast.IntExpression{Value: 1},
			Operator: token.PLUS,
			Right:    &ast.IntExpression{Value: 2},
		},
		&ast.InfixExpression{
			Left:     &ast.IntExpression{Value: 3},
			Operator: token.MINUS,
			Right:    &ast.IntExpression{Value: 4},
		},
		&ast.InfixExpression{
			Left:     &ast.IntExpression{Value: 5},
			Operator: token.MULTIPLY,
			Right:    &ast.IntExpression{Value: 6},
		},
		&ast.InfixExpression{
			Left:     &ast.IntExpression{Value: 7},
			Operator: token.DIVIDE,
			Right:    &ast.IntExpression{Value: 8},
		},
		&ast.InfixExpression{
			Left:     &ast.IntExpression{Value: 9},
			Operator: token.MODULO,
			Right:    &ast.IntExpression{Value: 10},
		},
		&ast.InfixExpression{
			Left:     &ast.BoolExpression{Value: true},
			Operator: token.AND,
			Right:    &ast.BoolExpression{Value: false},
		},
		&ast.InfixExpression{
			Left:     &ast.BoolExpression{Value: false},
			Operator: token.OR,
			Right:    &ast.BoolExpression{Value: true},
		},
	}

	checkTree(t, input, expectedTree)
}

func TestParserString(t *testing.T) {
	input := `
    "Hello"
    "World"
    `

	expectedTree := []ast.Statement{
		&ast.StringExpression{
			Value: "Hello",
		},
		&ast.StringExpression{
			Value: "World",
		},
	}

	checkTree(t, input, expectedTree)
}

func TestParserInt(t *testing.T) {
	input := `
    123
    456
    `

	expectedTree := []ast.Statement{
		&ast.IntExpression{
			Value: 123,
		},
		&ast.IntExpression{
			Value: 456,
		},
	}

	checkTree(t, input, expectedTree)
}

func TestParserBoolean(t *testing.T) {
	input := `
    true
    false
    `

	expectedTree := []ast.Statement{
		&ast.BoolExpression{
			Value: true,
		},
		&ast.BoolExpression{
			Value: false,
		},
	}

	checkTree(t, input, expectedTree)
}
