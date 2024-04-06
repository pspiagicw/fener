package parser

import (
	"fmt"
	"strconv"

	"github.com/pspiagicw/fener/ast"
	"github.com/pspiagicw/fener/token"
)

func (p *Parser) parseInteger() ast.Expression {
	value := p.curToken

	num, err := strconv.ParseInt(value.Value, 10, 64)

	if err != nil {
		message := fmt.Sprintf("could not parse %q as integer", value.Value)
		p.errors = append(p.errors, message)
		return nil
	}

	p.advance()

	return &ast.Integer{Token: value, Value: num}
}

func (p *Parser) parseString() ast.Expression {
	s := &ast.String{Token: p.curToken, Value: p.curToken.Value}

	p.advance()

	return s
}
func (p *Parser) parseBoolean() ast.Expression {
	b := &ast.Boolean{Token: p.curToken}

	b.Value = p.curTokenIs(token.TRUE)

	p.advance()

	return b
}
func (p *Parser) parseIdent() ast.Expression {
	i := &ast.Identifier{Token: p.curToken, Value: p.curToken.Value}

	p.advance()

	return i
}
func (p *Parser) parseGroupedExpression() ast.Expression {
	p.advance()

	exp := p.parseExpression(LOWEST)

	if p.curTokenIs(token.RPAREN) {
		p.advance()
	} else {
		p.errors = append(p.errors, "expected )")
	}

	return exp
}

func (p *Parser) parseIfExpression() ast.Expression {
	ifExp := &ast.IfExpression{Token: p.curToken}
	p.advance()

	ifExp.Elif = map[ast.Expression]*ast.BlockStatement{}

	ifExp.Condition = p.parseExpression(LOWEST)

	if !p.expect(token.THEN) {
		return nil
	}

	ifExp.Consequence = p.parseBlockStatement()

	for p.curTokenIs(token.ELIF) {
		p.advance()
		condition := p.parseExpression(LOWEST)
		if !p.expect(token.THEN) {
			return nil
		}
		body := p.parseBlockStatement()
		ifExp.Elif[condition] = body
	}

	if p.curTokenIs(token.ELSE) {
		p.advance()
		ifExp.Alternative = p.parseBlockStatement()
	}

	if !p.expect(token.END) {
		return nil
	}

	return ifExp

}
func (p *Parser) parseBlockStatement() *ast.BlockStatement {

	b := &ast.BlockStatement{Token: p.curToken}

	b.Statements = []ast.Statement{}

	for !p.curTokenIs(token.END) && !p.curTokenIs(token.EOF) && !p.curTokenIs(token.ELSE) && !p.curTokenIs(token.ELIF) {
		b.Statements = append(b.Statements, p.parseStatement())
	}

	return b
}
