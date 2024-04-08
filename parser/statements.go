package parser

import (
	"github.com/pspiagicw/fener/ast"
	"github.com/pspiagicw/fener/token"
)

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken}

	p.advance()

	stmt.Value = p.parseExpression(LOWEST)

	return stmt
}
func (p *Parser) parseWhileStatement() *ast.WhileStatement {
	stmt := &ast.WhileStatement{Token: p.curToken}

	p.advance()

	stmt.Condition = p.parseExpression(LOWEST)

	if !p.expect(token.THEN) {
		return nil
	}

	stmt.Consequence = p.parseBlockStatement()

	if !p.expect(token.END) {
		return nil
	}

	return stmt
}