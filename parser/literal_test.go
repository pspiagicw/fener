package parser

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/pspiagicw/fener/ast"
	"github.com/pspiagicw/fener/lexer"
	"github.com/pspiagicw/fener/token"
)

func TestArrayParser(t *testing.T) {
	input := `
    [1, 2, 3]
    ;; this comment is required
    ;; or else the next line is parsed as a index expressioo
    [1 + 2, someArr[2] , someFunc(arg, arg, arg)]
    `

	expectedTree := []ast.Statement{
		&ast.ExpressionStatement{
			Expression: &ast.Array{
				Elements: []ast.Expression{
					&ast.Integer{Value: 1, Token: &token.Token{Type: token.INT, Value: "1", Line: 1}},
					&ast.Integer{Value: 2, Token: &token.Token{Type: token.INT, Value: "2", Line: 1}},
					&ast.Integer{Value: 3, Token: &token.Token{Type: token.INT, Value: "3", Line: 1}},
				},
				Token: &token.Token{Type: token.LSQUARE, Value: "[", Line: 1},
			},
		},
		&ast.ExpressionStatement{
			Expression: &ast.Array{
				Elements: []ast.Expression{
					&ast.InfixExpression{
						Left:     &ast.Integer{Value: 1, Token: &token.Token{Type: token.INT, Value: "1", Line: 2}},
						Operator: token.PLUS,
						Right:    &ast.Integer{Value: 2, Token: &token.Token{Type: token.INT, Value: "2", Line: 2}},
					},
					&ast.IndexExpression{
						Left:  &ast.Identifier{Value: "someArr", Token: &token.Token{Type: token.IDENT, Value: "someArr", Line: 2}},
						Index: &ast.Integer{Value: 2, Token: &token.Token{Type: token.INT, Value: "2", Line: 2}},
					},
					&ast.CallExpression{
						Function: &ast.Identifier{Value: "someFunc", Token: &token.Token{Type: token.IDENT, Value: "someFunc", Line: 2}},
						Arguments: []ast.Expression{
							&ast.Identifier{Value: "arg", Token: &token.Token{Type: token.IDENT, Value: "arg", Line: 2}},
							&ast.Identifier{Value: "arg", Token: &token.Token{Type: token.IDENT, Value: "arg", Line: 2}},
							&ast.Identifier{Value: "arg", Token: &token.Token{Type: token.IDENT, Value: "arg", Line: 2}},
						},
					},
				},
				Token: &token.Token{Type: token.LSQUARE, Value: "[", Line: 2},
			},
		},
	}

	checkTree(t, input, expectedTree)
}

func TestParserIdentifiers(t *testing.T) {
	input := `
    identA identB
    `

	expectedTree := []ast.Statement{
		&ast.ExpressionStatement{
			Expression: &ast.Identifier{Value: "identA", Token: &token.Token{Type: token.IDENT, Value: "identA", Line: 1}},
		},
		&ast.ExpressionStatement{
			Expression: &ast.Identifier{Value: "identB", Token: &token.Token{Type: token.IDENT, Value: "identB", Line: 1}},
		},
	}

	checkTree(t, input, expectedTree)
}

func TestParserIntegerExpression(t *testing.T) {
	input := `
    123
    456
    `

	expectedTree := []ast.Statement{
		&ast.ExpressionStatement{
			Token:      &token.Token{Type: token.INT, Value: "123", Line: 1},
			Expression: &ast.Integer{Value: 123, Token: &token.Token{Type: token.INT, Value: "123", Line: 1}},
		},
		&ast.ExpressionStatement{
			Token:      &token.Token{Type: token.INT, Value: "456", Line: 2},
			Expression: &ast.Integer{Value: 456, Token: &token.Token{Type: token.INT, Value: "456", Line: 2}},
		},
	}

	checkTree(t, input, expectedTree)
}

func TestParserStringExpression(t *testing.T) {
	input := `
    "Hello"
    "World"
    `

	expectedTree := []ast.Statement{
		&ast.ExpressionStatement{
			Token:      &token.Token{Type: token.STRING, Value: "Hello", Line: 1},
			Expression: &ast.String{Value: "Hello", Token: &token.Token{Type: token.STRING, Value: "Hello", Line: 1}},
		},
		&ast.ExpressionStatement{
			Token:      &token.Token{Type: token.STRING, Value: "World", Line: 2},
			Expression: &ast.String{Value: "World", Token: &token.Token{Type: token.STRING, Value: "World", Line: 2}},
		},
	}

	checkTree(t, input, expectedTree)
}

func TestParserBooleanExpression(t *testing.T) {
	input := `
    true
    false
    `

	expectedTree := []ast.Statement{
		&ast.ExpressionStatement{
			Token:      &token.Token{Type: token.TRUE, Value: "true", Line: 1},
			Expression: &ast.Boolean{Value: true, Token: &token.Token{Type: token.TRUE, Value: "true", Line: 1}},
		},
		&ast.ExpressionStatement{
			Token:      &token.Token{Type: token.FALSE, Value: "false", Line: 2},
			Expression: &ast.Boolean{Value: false, Token: &token.Token{Type: token.FALSE, Value: "false", Line: 2}},
		},
	}

	checkTree(t, input, expectedTree)
}

func TestParserBoolean(t *testing.T) {
	input := `
    return true
    return false
    `

	expectedTree := []ast.Statement{
		&ast.ReturnStatement{
			Token: &token.Token{Type: token.RETURN, Value: "return", Line: 1},
			Value: &ast.Boolean{Value: true, Token: &token.Token{Type: token.TRUE, Value: "true", Line: 1}},
		},
		&ast.ReturnStatement{
			Token: &token.Token{Type: token.RETURN, Value: "return", Line: 2},
			Value: &ast.Boolean{Value: false, Token: &token.Token{Type: token.FALSE, Value: "false", Line: 2}},
		},
	}

	checkTree(t, input, expectedTree)
}

func checkTree(t *testing.T, input string, expectedTree []ast.Statement) {
	t.Helper()
	l := lexer.New(input)
	p := New(l)

	tree := p.Parse()

	if len(p.errors) != 0 {
		for _, err := range p.errors {
			t.Logf("Parser error: %s", err)
		}
		t.Fatalf("Parser has %d errors", len(p.errors))
	}

	if len(tree.Statements) != len(expectedTree) {
		t.Fatalf("Expected %d statements, got %d", len(expectedTree), len(tree.Statements))
	}

	spew.Config.DisablePointerAddresses = true
	expectedDump := spew.Sdump(expectedTree)
	actualDump := spew.Sdump(tree.Statements)

	if actualDump != expectedDump {
		t.Fatalf("Expected tree to be %s, got %s", expectedDump, actualDump)
	}

}
