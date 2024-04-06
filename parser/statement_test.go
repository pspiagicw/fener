package parser

import (
	"testing"

	"github.com/pspiagicw/fener/ast"
	"github.com/pspiagicw/fener/token"
)

func TestParserReturnInt(t *testing.T) {
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

func TestParserReturnString(t *testing.T) {
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
