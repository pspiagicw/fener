package parser

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/pspiagicw/fener/ast"
	"github.com/pspiagicw/fener/lexer"
)

func TestParserInt(t *testing.T) {
	input := `
    return 123
    return 456
    `

	expectedTree := []ast.Statement{
		&ast.ReturnStatement{
			Value: &ast.Integer{Value: 123},
		},
		&ast.ReturnStatement{
			Value: &ast.Integer{Value: 456},
		},
	}

	checkTree(t, input, expectedTree)
}

func TestParserString(t *testing.T) {
	input := `
    return "Hello"
    return "World"
    `

	expectedTree := []ast.Statement{
		&ast.ReturnStatement{
			Value: &ast.String{Value: "Hello"},
		},
		&ast.ReturnStatement{
			Value: &ast.String{Value: "World"},
		},
	}

	checkTree(t, input, expectedTree)
}

//
// func TestParserBoolean(t *testing.T) {
// 	input := `
//     true
//     false
//     `
//
// 	expectedTree := []ast.Statement{
// 		&ast.BoolExpression{
// 			Value: true,
// 		},
// 		&ast.BoolExpression{
// 			Value: false,
// 		},
// 	}
//
// 	checkTree(t, input, expectedTree)
// }
// func TestParserIdentifierExpressionOnly(t *testing.T) {
// 	input := `identifier`
//
// 	expectedTree := []ast.Statement{
// 		&ast.ExpressionStatement{
// 			Expression: &ast.IdentifierExpression{Token: token.Token{Type: token.IDENT, Value: "identifier"}},
// 		},
// 	}
//
// 	checkTree(t, input, expectedTree)
// }
// func TestParserAssignmentStatements(t *testing.T) {
// 	input := `
//     a = 1
//     name = "myName"
//     flag = true
//     result = 5 * 3 - 2
//     `
//
// 	expectedTree := []ast.Statement{
// 		&ast.AssignmentStatement{
// 			Name:  token.Token{Type: token.IDENT, Value: "a"},
// 			Value: &ast.IntExpression{Value: 1},
// 		},
// 		&ast.AssignmentStatement{
// 			Name:  token.Token{Type: token.IDENT, Value: "name"},
// 			Value: &ast.StringExpression{Value: "myName"},
// 		},
// 		&ast.AssignmentStatement{
// 			Name:  token.Token{Type: token.IDENT, Value: "flag"},
// 			Value: &ast.BoolExpression{Value: true},
// 		},
// 		&ast.AssignmentStatement{
// 			Name: token.Token{Type: token.IDENT, Value: "result"},
// 			Value: &ast.InfixExpression{
// 				Left: &ast.InfixExpression{
// 					Left:     &ast.IntExpression{Value: 5},
// 					Operator: token.MULTIPLY,
// 					Right:    &ast.IntExpression{Value: 3},
// 				},
// 				Operator: token.MINUS,
// 				Right:    &ast.IntExpression{Value: 2},
// 			},
// 		},
// 	}
//
// 	checkTree(t, input, expectedTree)
// }
//
// func TestParserInfixExpressions(t *testing.T) {
// 	input := `
//     1 + 2
//     3 - 4
//     5 * 6
//     7 / 8
//     9 % 10
//     true and false
//     false or true
//     `
//
// 	expectedTree := []ast.Statement{
// 		&ast.InfixExpression{
// 			Left:     &ast.IntExpression{Value: 1},
// 			Operator: token.PLUS,
// 			Right:    &ast.IntExpression{Value: 2},
// 		},
// 		&ast.InfixExpression{
// 			Left:     &ast.IntExpression{Value: 3},
// 			Operator: token.MINUS,
// 			Right:    &ast.IntExpression{Value: 4},
// 		},
// 		&ast.InfixExpression{
// 			Left:     &ast.IntExpression{Value: 5},
// 			Operator: token.MULTIPLY,
// 			Right:    &ast.IntExpression{Value: 6},
// 		},
// 		&ast.InfixExpression{
// 			Left:     &ast.IntExpression{Value: 7},
// 			Operator: token.DIVIDE,
// 			Right:    &ast.IntExpression{Value: 8},
// 		},
// 		&ast.InfixExpression{
// 			Left:     &ast.IntExpression{Value: 9},
// 			Operator: token.MODULO,
// 			Right:    &ast.IntExpression{Value: 10},
// 		},
// 		&ast.InfixExpression{
// 			Left:     &ast.BoolExpression{Value: true},
// 			Operator: token.AND,
// 			Right:    &ast.BoolExpression{Value: false},
// 		},
// 		&ast.InfixExpression{
// 			Left:     &ast.BoolExpression{Value: false},
// 			Operator: token.OR,
// 			Right:    &ast.BoolExpression{Value: true},
// 		},
// 	}
//
// 	checkTree(t, input, expectedTree)
// }

func checkTree(t *testing.T, input string, expectedTree []ast.Statement) {
	l := lexer.New(input)
	p := New(l)

	tree := p.Parse()

	if len(tree.Statements) != len(expectedTree) {
		t.Fatalf("Expected %d statements, got %d", len(expectedTree), len(tree.Statements))
	}

	if len(p.errors) != 0 {
		t.Fatalf("Parser has %d errors", len(p.errors))
	}

	spew.Config.DisablePointerAddresses = true
	expectedDump := spew.Sdump(expectedTree)
	actualDump := spew.Sdump(tree.Statements)

	if actualDump != expectedDump {
		t.Fatalf("Expected tree to be %s, got %s", expectedDump, actualDump)
	}

}
