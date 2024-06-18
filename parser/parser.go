package parser

import (
	"fmt"

	"github.com/pspiagicw/fener/ast"
	"github.com/pspiagicw/fener/lexer"
	"github.com/pspiagicw/fener/token"
)

const (
	_ = iota
	LOWEST
	EQUALS
	COMPARE
	ADDITION
	MULTIPLY
	MOD
	BOOLEAN
	PREFIX
	FIELD
	CALL
	INDEX
	ASSIGNMENT
)

var precedences = map[token.TokenType]int{
	token.ASSIGN:   ASSIGNMENT,
	token.PLUS:     ADDITION,
	token.MINUS:    ADDITION,
	token.MULTIPLY: MULTIPLY,
	token.DIVIDE:   MULTIPLY,
	token.EQ:       EQUALS,
	token.NOT_EQ:   EQUALS,
	token.LT:       COMPARE,
	token.GT:       COMPARE,
	token.LTE:      COMPARE,
	token.GTE:      COMPARE,
	token.MOD:      MOD,
	token.LPAREN:   CALL,
	token.AND:      BOOLEAN,
	token.OR:       BOOLEAN,
	token.BITAND:   BOOLEAN,
	token.BITOR:    BOOLEAN,
	token.LSQUARE:  INDEX,
	token.DOT:      FIELD,
}

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

	p.prefixParseFns = map[token.TokenType]prefixParseFn{}

	p.registerPrefix(token.INT, p.parseInteger)
	p.registerPrefix(token.STRING, p.parseString)
	p.registerPrefix(token.TRUE, p.parseBoolean)
	p.registerPrefix(token.FALSE, p.parseBoolean)
	p.registerPrefix(token.IDENT, p.parseIdent)
	p.registerPrefix(token.LPAREN, p.parseGroupedExpression)
	p.registerPrefix(token.IF, p.parseIfExpression)
	p.registerPrefix(token.FUNCTION, p.parseLambda)
	p.registerPrefix(token.LSQUARE, p.parseArray)
	p.registerPrefix(token.MINUS, p.parsePrefixExpression)
	p.registerPrefix(token.BANG, p.parsePrefixExpression)

	p.infixParseFns = map[token.TokenType]infixParseFn{}

	p.registerInfix(token.PLUS, p.parseInfixExpression)
	p.registerInfix(token.MINUS, p.parseInfixExpression)
	p.registerInfix(token.MULTIPLY, p.parseInfixExpression)
	p.registerInfix(token.DIVIDE, p.parseInfixExpression)
	p.registerInfix(token.MOD, p.parseInfixExpression)
	p.registerInfix(token.EQ, p.parseInfixExpression)
	p.registerInfix(token.NOT_EQ, p.parseInfixExpression)
	p.registerInfix(token.LT, p.parseInfixExpression)
	p.registerInfix(token.GT, p.parseInfixExpression)
	p.registerInfix(token.LTE, p.parseInfixExpression)
	p.registerInfix(token.GTE, p.parseInfixExpression)
	p.registerInfix(token.ASSIGN, p.parseAssignmentExpression)
	p.registerInfix(token.LPAREN, p.parseCallExpression)
	p.registerInfix(token.AND, p.parseInfixExpression)
	p.registerInfix(token.OR, p.parseInfixExpression)
	p.registerInfix(token.BITAND, p.parseInfixExpression)
	p.registerInfix(token.BITOR, p.parseInfixExpression)
	p.registerInfix(token.LSQUARE, p.parseIndexExpression)

	p.advance()
	p.advance()

	return p

}

func (p *Parser) registerPrefix(t token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[t] = fn
}
func (p *Parser) registerInfix(t token.TokenType, fn infixParseFn) {
	p.infixParseFns[t] = fn
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
	case token.WHILE:
		return p.parseWhileStatement()
	case token.COMMENT:
		p.advance()
		return nil
	case token.FUNCTION:
		return p.parseFunctionStatement()
	case token.TEST:
		return p.parseTestStatement()
	case token.CLASS:
		return p.parseClassStatement()
	default:
		return p.parseExpressionStatement()
	}
}
func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.curToken}

	stmt.Expression = p.parseExpression(LOWEST)

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

	// This consumes the current token and advances to the next one
	leftExp := prefix()
	// So we compare the precedence of the current token.

	for !p.curTokenIs(token.EOF) && precedence < p.curPrecedence() {
		infix := p.infixParseFns[p.curToken.Type]

		if infix == nil {
			return leftExp
		}

		leftExp = infix(leftExp)
	}

	return leftExp
}
func (p *Parser) peektokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}
func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
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
func (p *Parser) curPrecedence() int {
	if p, ok := precedences[p.curToken.Type]; ok {
		return p
	}

	return LOWEST
}
func (p *Parser) expect(t token.TokenType) bool {
	if p.curTokenIs(t) {
		p.advance()
		return true
	}

	p.errors = append(p.errors, fmt.Sprintf("expected %s, got %s", t, p.curToken.Type))
	return false
}
