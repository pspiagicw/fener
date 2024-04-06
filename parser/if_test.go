package parser

import (
	"testing"

	"github.com/pspiagicw/fener/ast"
	"github.com/pspiagicw/fener/token"
)

func TestParserIf(t *testing.T) {
	input := `
    if true then 10 else 20 end
    if true then 10 end

    `

	expectedTree := []ast.Statement{
		&ast.ExpressionStatement{
			Expression: &ast.IfExpression{
				Token: &token.Token{Type: token.IF, Value: "if", Line: 0},
				Condition: &ast.Boolean{
					Token: &token.Token{Type: "TRUE", Value: "true", Line: 0},
					Value: true,
				},
				Consequence: &ast.BlockStatement{
					Token: &token.Token{Type: "INT", Value: "10", Line: 0},
					Statements: []ast.Statement{
						&ast.ExpressionStatement{
							Expression: &ast.Integer{
								Token: &token.Token{Type: "INT", Value: "10", Line: 0},
								Value: 10,
							},
							Token: &token.Token{Type: "INT", Value: "10", Line: 0},
						},
					},
				},
				Elif: map[ast.Expression]*ast.BlockStatement{},
				Alternative: &ast.BlockStatement{
					Token: &token.Token{Type: "INT", Value: "20", Line: 0},
					Statements: []ast.Statement{
						&ast.ExpressionStatement{
							Expression: &ast.Integer{
								Token: &token.Token{Type: "INT", Value: "20", Line: 0},
								Value: 20,
							},
							Token: &token.Token{Type: "INT", Value: "20", Line: 0},
						},
					},
				},
			},
			Token: &token.Token{Type: "IF", Value: "if", Line: 0},
		},
		&ast.ExpressionStatement{
			Expression: &ast.IfExpression{
				Token: &token.Token{Type: "IF", Value: "if", Line: 0},
				Condition: &ast.Boolean{
					Token: &token.Token{Type: "TRUE", Value: "true", Line: 0},
					Value: true,
				},
				Consequence: &ast.BlockStatement{
					Token: &token.Token{Type: "INT", Value: "10", Line: 0},
					Statements: []ast.Statement{
						&ast.ExpressionStatement{
							Expression: &ast.Integer{
								Token: &token.Token{Type: "INT", Value: "10", Line: 0},
								Value: 10,
							},
							Token: &token.Token{Type: "INT", Value: "10", Line: 0},
						},
					},
				},
				Elif:        map[ast.Expression]*ast.BlockStatement{},
				Alternative: nil,
			},
			Token: &token.Token{Type: "IF", Value: "if", Line: 0},
		},

		// &ast.ExpressionStatement{
		// 	Expression: &ast.IfExpression{
		// 		Token: &token.Token{Type: "IF", Value: "if", Line: 4},
		// 		Condition: &ast.InfixExpression{
		// 			Left: &ast.Integer{
		// 				Token: &token.Token{Type: "INT", Value: "1", Line: 4},
		// 				Value: 1,
		// 			},
		// 			Operator: "==",
		// 			Right: &ast.Integer{
		// 				Token: &token.Token{Type: "INT", Value: "2", Line: 4},
		// 				Value: 2,
		// 			},
		// 		},
		// 		Consequence: &ast.BlockStatement{
		// 			Token: &token.Token{Type: "INT", Value: "10", Line: 5},
		// 			Statements: []ast.Statement{
		// 				&ast.ExpressionStatement{
		// 					Expression: &ast.Integer{
		// 						Token: &token.Token{Type: "INT", Value: "10", Line: 5},
		// 						Value: 10,
		// 					},
		// 					Token: &token.Token{Type: "INT", Value: "10", Line: 5},
		// 				},
		// 			},
		// 		},
		// 		Elif: map[ast.Expression]*ast.BlockStatement{
		// 			&ast.InfixExpression{
		// 				Left: &ast.Integer{
		// 					Token: &token.Token{Type: "INT", Value: "2", Line: 6},
		// 					Value: 2,
		// 				},
		// 				Operator: "==",
		// 				Right: &ast.Integer{
		// 					Token: &token.Token{Type: "INT", Value: "3", Line: 6},
		// 					Value: 3,
		// 				},
		// 			}: &ast.BlockStatement{
		// 				Token: &token.Token{Type: "INT", Value: "20", Line: 7},
		// 				Statements: []ast.Statement{
		// 					&ast.ExpressionStatement{
		// 						Expression: &ast.Integer{
		// 							Token: &token.Token{Type: "INT", Value: "20", Line: 7},
		// 							Value: 20,
		// 						},
		// 						Token: &token.Token{Type: "INT", Value: "20", Line: 7},
		// 					},
		// 				},
		// 			},
		// 			&ast.InfixExpression{
		// 				Left: &ast.Integer{
		// 					Token: &token.Token{Type: "INT", Value: "3", Line: 8},
		// 					Value: 3,
		// 				},
		// 				Operator: "==",
		// 				Right: &ast.Integer{
		// 					Token: &token.Token{Type: "INT", Value: "4", Line: 8},
		// 					Value: 4,
		// 				},
		// 			}: &ast.BlockStatement{
		// 				Token: &token.Token{Type: "INT", Value: "30", Line: 9},
		// 				Statements: []ast.Statement{
		// 					&ast.ExpressionStatement{
		// 						Expression: &ast.Integer{
		// 							Token: &token.Token{Type: "INT", Value: "30", Line: 9},
		// 							Value: 30,
		// 						},
		// 						Token: &token.Token{Type: "INT", Value: "30", Line: 9},
		// 					},
		// 				},
		// 			},
		// 		},
		// 		Alternative: &ast.BlockStatement{
		// 			Token: &token.Token{Type: "INT", Value: "40", Line: 11},
		// 			Statements: []ast.Statement{
		// 				&ast.ExpressionStatement{
		// 					Expression: &ast.Integer{
		// 						Token: &token.Token{Type: "INT", Value: "40", Line: 11},
		// 						Value: 40,
		// 					},
		// 					Token: &token.Token{Type: "INT", Value: "40", Line: 11},
		// 				},
		// 			},
		// 		},
		// 	},
		// 	Token: &token.Token{Type: "IF", Value: "if", Line: 4},
		// },
	}

	checkTree(t, input, expectedTree)
}
