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
