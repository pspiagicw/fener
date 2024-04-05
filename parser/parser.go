package parser

import (
	"fmt"

	"github.com/pspiagicw/fener/ast"
	"github.com/pspiagicw/fener/lexer"
	"github.com/pspiagicw/fener/token"
)

const (
	LOWEST = iota
	ADDITION
)

var precedences = map[token.TokenType]int{}

type infixParseFn func(ast.Expression) ast.Expression
type prefixParseFn func() ast.Expression

type Parser struct {
	l      *lexer.Lexer
	errors []string

	curToken  *token.Token
	peekToken *token.Token

	infixParseFns  map[token.TokenType]infixParseFn
	prefixParseFns map[token.TokenType]prefixParseFn
}

func (p *Parser) Errors() []string {
	return p.errors
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	p.infixParseFns = map[token.TokenType]infixParseFn{}
	p.prefixParseFns = map[token.TokenType]prefixParseFn{}

	p.registerPrefix(token.INT, p.parseInteger)
	p.registerPrefix(token.STRING, p.parseString)

	p.advance()
	p.advance()

	return p

}
func (p *Parser) registerPrefix(t token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[t] = fn
}

func (p *Parser) Parse() *ast.Program {
	statements := []ast.Statement{}

	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			statements = append(statements, stmt)
		}
	}

	return &ast.Program{Statements: statements}
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		p.errors = append(p.errors, fmt.Sprintf("unknown token type %s", p.curToken.Type))
		p.advance()
	}

	return nil
}
func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken}

	p.advance()

	stmt.Value = p.parseExpression(LOWEST)

	return stmt
}

func (p *Parser) advance() {
	p.curToken = p.peekToken
	p.peekToken = p.l.Next()
}
func (p *Parser) parseExpression(precedence int) ast.Expression {

	prefix := p.prefixParseFns[p.curToken.Type]

	if prefix == nil {
		p.noPrefixParseFnError(p.curToken.Type)
		p.advance()
		return nil
	}

	leftExp := prefix()

	if !p.peektokenIs(token.EOF) && precedence < p.peekPrecedence() {
		infix := p.infixParseFns[p.peekToken.Type]

		if infix == nil {
			return leftExp
		}

		p.advance()

		leftExp = infix(leftExp)
	}

	return leftExp
}
func (p *Parser) peektokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}
func (p *Parser) peekPrecedence() int {
	if p, ok := precedences[p.peekToken.Type]; ok {
		return p
	}

	return LOWEST
}
func (p *Parser) noPrefixParseFnError(t token.TokenType) {
	msg := fmt.Sprintf("no prefix parse function for %s found", t)
	p.errors = append(p.errors, msg)
}
