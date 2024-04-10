package parser

import (
	"github.com/pspiagicw/fener/ast"
	"github.com/pspiagicw/fener/token"
)

func (p *Parser) parseIndexExpression(left ast.Expression) ast.Expression {
	expression := &ast.IndexExpression{
		Token: p.curToken,
		Left:  left,
	}

	p.advance()

	expression.Index = p.parseExpression(LOWEST)

	if !p.expect(token.RSQUARE) {
		return nil
	}

	return expression
}
func (p *Parser) parseCallExpression(left ast.Expression) ast.Expression {
	expression := &ast.CallExpression{
		Token:    p.curToken,
		Function: left,
	}

	p.advance()

	expression.Arguments = p.parseExpressionList(token.RPAREN)

	return expression
}
func (p *Parser) parseExpressionList(end token.TokenType) []ast.Expression {
	expressions := []ast.Expression{}

	if p.curTokenIs(end) {
		p.advance()
		return expressions
	}

	for true {
		expression := p.parseExpression(LOWEST)
		if expression != nil {
			expressions = append(expressions, expression)
		}
		if p.curTokenIs(end) {
			p.advance()
			break
		}
		if !p.expect(token.COMMA) {
			return expressions
		}
	}

	return expressions
}

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	expression := &ast.InfixExpression{
		Left:     left,
		Operator: p.curToken.Type,
	}

	precedence := p.curPrecedence()
	p.advance()
	expression.Right = p.parseExpression(precedence)

	return expression
}
func (p *Parser) parseAssignmentExpression(left ast.Expression) ast.Expression {

	identifier, ok := left.(*ast.Identifier)

	if !ok {
		p.errors = append(p.errors, "Invalid assignment target")
		return nil
	}

	p.advance()

	assignment := &ast.AssignmentExpression{
		Token:  p.curToken,
		Target: identifier,
	}

	assignment.Value = p.parseExpression(LOWEST)

	return assignment
}
