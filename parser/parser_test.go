package parser

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/pspiagicw/fener/ast"
	"github.com/pspiagicw/fener/lexer"
	"github.com/pspiagicw/fener/token"
)

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

func TestParserInt(t *testing.T) {
	input := `
    return 123
    return 456
    `

	expectedTree := []ast.Statement{
		&ast.ReturnStatement{
			Token: &token.Token{Type: token.RETURN, Value: "return", Line: 1},
			Value: &ast.Integer{Value: 123, Token: &token.Token{Type: token.INT, Value: "123", Line: 1}},
		},
		&ast.ReturnStatement{
			Token: &token.Token{Type: token.RETURN, Value: "return", Line: 2},
			Value: &ast.Integer{Value: 456, Token: &token.Token{Type: token.INT, Value: "456", Line: 2}},
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
			Token: &token.Token{Type: token.RETURN, Value: "return", Line: 1},
			Value: &ast.String{Value: "Hello", Token: &token.Token{Type: token.STRING, Value: "Hello", Line: 1}},
		},
		&ast.ReturnStatement{
			Token: &token.Token{Type: token.RETURN, Value: "return", Line: 2},
			Value: &ast.String{Value: "World", Token: &token.Token{Type: token.STRING, Value: "World", Line: 2}},
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
