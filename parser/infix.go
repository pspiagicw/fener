package parser

import "github.com/pspiagicw/fener/ast"

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
