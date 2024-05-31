package parser

import (
	"github.com/pspiagicw/fener/ast"
	"github.com/pspiagicw/fener/token"
)

func (p *Parser) parseClassStatement() ast.Statement {
	stmt := &ast.ClassStatement{Token: p.curToken}

	p.advance()

	value := p.parseExpression(LOWEST)

	ident, ok := value.(*ast.Identifier)

	if !ok {
		p.errors = append(p.errors, "Expected identifier target for class statement")
	}

	stmt.Target = ident

	stmt.Methods = []*ast.FunctionStatement{}

	for !p.curTokenIs(token.END) && !p.curTokenIs(token.EOF) {
		stmt.Methods = append(stmt.Methods, p.parseFunctionStatement())
	}

	if !p.expect(token.END) {
		return nil
	}

	return stmt
}

func (p *Parser) parseTestStatement() *ast.TestStatement {
	stmt := &ast.TestStatement{Token: p.curToken}

	p.advance()

	target := p.parseExpression(LOWEST)

	if target == nil {
		p.errors = append(p.errors, "Expected target for test statement")
		return nil
	}

	t, ok := target.(*ast.String)

	if !ok {
		p.errors = append(p.errors, "Expected string target for test statement")
		return nil
	}

	stmt.Target = t

	stmt.Statements = p.parseBlockStatement()

	if !p.expect(token.END) {
		return nil
	}

	return stmt
}

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
func (p *Parser) parseFunctionStatement() *ast.FunctionStatement {
	stmt := &ast.FunctionStatement{Token: p.curToken}

	p.advance()

	stmt.Target = &ast.Identifier{Token: p.curToken, Value: p.curToken.Value}

	if !p.expect(token.IDENT) {
		return nil
	}

	if !p.expect(token.LPAREN) {
		return nil
	}

	stmt.Arguments = p.parseFunctionParameters()

	stmt.Body = p.parseBlockStatement()

	if !p.expect(token.END) {
		return nil
	}

	return stmt
}
