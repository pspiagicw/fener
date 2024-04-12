package parser

import (
	"fmt"
	"strconv"

	"github.com/pspiagicw/fener/ast"
	"github.com/pspiagicw/fener/token"
)

func (p *Parser) parsePrefixExpression() ast.Expression {
	expression := &ast.PrefixExpression{Token: p.curToken, Operator: p.curToken.Type}

	p.advance()

	expression.Right = p.parseExpression(PREFIX)

	return expression
}

func (p *Parser) parseArray() ast.Expression {
	array := &ast.Array{Token: p.curToken}

	if !p.expect(token.LSQUARE) {
		return nil
	}

	array.Elements = p.parseExpressionList(token.RSQUARE)

	return array
}

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
		statement := p.parseStatement()

		if statement != nil {
			b.Statements = append(b.Statements, statement)
		}
	}

	return b
}

func (p *Parser) parseLambda() ast.Expression {
	lambda := &ast.Lambda{Token: p.curToken}

	p.advance()

	if !p.expect(token.LPAREN) {
		return nil
	}

	lambda.Arguments = p.parseFunctionParameters()

	lambda.Body = p.parseBlockStatement()

	if !p.expect(token.END) {
		return nil
	}

	return lambda
}
func (p *Parser) parseFunctionParameters() []*ast.Identifier {
	expressions := p.parseExpressionList(token.RPAREN)

	identifiers := []*ast.Identifier{}

	for _, exp := range expressions {
		ident, ok := exp.(*ast.Identifier)
		if !ok {
			message := fmt.Sprintf("Expected identifier, got %v", exp.Name())
			p.errors = append(p.errors, message)
			return nil
		}
		identifiers = append(identifiers, ident)
	}

	return identifiers
}
