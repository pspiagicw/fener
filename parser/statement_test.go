package parser

import (
	"testing"

	"github.com/pspiagicw/fener/ast"
	"github.com/pspiagicw/fener/token"
)

func TestEmptyClassStatement(t *testing.T) {
	input := `
    class className
    end
    `

	expectedTree := []ast.Statement{
		&ast.ClassStatement{
			Target:  &ast.Identifier{Value: "className", Token: &token.Token{Type: token.IDENT, Value: "className", Line: 1}},
			Methods: []*ast.FunctionStatement{},
		},
	}
	checkTree(t, input, expectedTree)
}
func TestMethodClassStatement(t *testing.T) {

	input := `
    class Dog
        fn bark()
            print("Woof")
        end
        fn eat()
            print("Yum")
        end
    end
    `

	expectedTree := []ast.Statement{
		&ast.ClassStatement{
			Target: &ast.Identifier{Value: "Dog", Token: &token.Token{Type: token.IDENT, Value: "Dog", Line: 1}},
			Methods: []*ast.FunctionStatement{
				{
					Target:    &ast.Identifier{Value: "bark", Token: &token.Token{Type: token.IDENT, Value: "bark", Line: 2}},
					Arguments: []*ast.Identifier{},
					Body: &ast.BlockStatement{
						Statements: []ast.Statement{
							&ast.ExpressionStatement{
								Expression: &ast.CallExpression{
									Function: &ast.Identifier{Value: "print", Token: &token.Token{Type: token.IDENT, Value: "print", Line: 3}},
									Arguments: []ast.Expression{
										&ast.String{Value: "Woof", Token: &token.Token{Type: token.STRING, Value: "Woof", Line: 3}},
									},
								},
							},
						},
					},
				},
				{
					Target:    &ast.Identifier{Value: "eat", Token: &token.Token{Type: token.IDENT, Value: "eat", Line: 5}},
					Arguments: []*ast.Identifier{},
					Body: &ast.BlockStatement{
						Statements: []ast.Statement{
							&ast.ExpressionStatement{
								Expression: &ast.CallExpression{
									Function: &ast.Identifier{Value: "print", Token: &token.Token{Type: token.IDENT, Value: "print", Line: 6}},
									Arguments: []ast.Expression{
										&ast.String{Value: "Yum", Token: &token.Token{Type: token.STRING, Value: "Yum", Line: 6}},
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

func TestParserTest(t *testing.T) {
	input := `
    test "Test 1"
        print("Hello, World")
    end
    `

	expectedTree := []ast.Statement{
		&ast.TestStatement{
			Target: &ast.String{Value: "Test 1", Token: &token.Token{Type: token.STRING, Value: "Test 1", Line: 1}},
			Statements: &ast.BlockStatement{
				Statements: []ast.Statement{
					&ast.ExpressionStatement{
						Expression: &ast.CallExpression{
							Function: &ast.Identifier{Value: "print", Token: &token.Token{Type: token.IDENT, Value: "print", Line: 2}},
							Arguments: []ast.Expression{
								&ast.String{Value: "Hello, World", Token: &token.Token{Type: token.STRING, Value: "Hello, World", Line: 2}},
							},
							Token: &token.Token{Type: token.LPAREN, Value: "(", Line: 2},
						},
						Token: &token.Token{Type: token.IDENT, Value: "print", Line: 2},
					},
				},
				Token: &token.Token{Type: token.THEN, Value: "then", Line: 1},
			},
			Token: &token.Token{Type: token.TEST, Value: "test", Line: 1},
		},
	}

	checkTree(t, input, expectedTree)

}

func TestWhileStatement(t *testing.T) {
	input := `
    while true then
        print("Hello, World")
    end
    `

	expectedTree := []ast.Statement{
		&ast.WhileStatement{
			Condition: &ast.Boolean{Value: true, Token: &token.Token{Type: token.TRUE, Value: "true", Line: 1}},
			Consequence: &ast.BlockStatement{
				Statements: []ast.Statement{
					&ast.ExpressionStatement{
						Expression: &ast.CallExpression{
							Function: &ast.Identifier{Value: "print", Token: &token.Token{Type: token.IDENT, Value: "print", Line: 2}},
							Arguments: []ast.Expression{
								&ast.String{Value: "Hello, World", Token: &token.Token{Type: token.STRING, Value: "Hello, World", Line: 2}},
							},
							Token: &token.Token{Type: token.LPAREN, Value: "(", Line: 2},
						},
						Token: &token.Token{Type: token.IDENT, Value: "print", Line: 2},
					},
				},
				Token: &token.Token{Type: token.THEN, Value: "then", Line: 1},
			},
			Token: &token.Token{Type: token.WHILE, Value: "while", Line: 1},
		},
	}

	checkTree(t, input, expectedTree)

}

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
