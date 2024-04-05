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
