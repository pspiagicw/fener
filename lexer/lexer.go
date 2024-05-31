package lexer

import (
	"fmt"

	"github.com/pspiagicw/fener/token"
)

type lexerState int

const (
	lexerStateNormal lexerState = iota
	lexerStateExec
)

type Lexer struct {
	input        string
	position     int    // current position in input (points to current char)
	readPosition int    // current reading position in input (after current char)
	ch           string // current char under examination
	eof          bool   // signals end of file
	line         int    // current line number
	err          error  // errors encountered during lexing

	currentState lexerState
	currentDepth int
}

func New(input string) *Lexer {
	l := &Lexer{
		input:        input,
		position:     -1,
		readPosition: 0,
		err:          nil,

		currentState: lexerStateNormal,
		currentDepth: 0,
	}
	return l
}
func (l *Lexer) token(ttype token.TokenType, value string) *token.Token {
	return &token.Token{Type: ttype, Value: value, Line: l.line}
}

func (l *Lexer) advance() {
	l.position = l.readPosition

	if l.readPosition < len(l.input) {
		l.ch = string(l.input[l.readPosition])
	} else {
		l.eof = true
		l.ch = ""
	}
	l.readPosition++
}
func isLetter(ch string) bool {
	return (ch >= "a" && ch <= "z") || (ch >= "A" && ch <= "Z") || ch == "_"
}

func isDigit(ch string) bool {
	return ch >= "0" && ch <= "9"
}

func (l *Lexer) identifier() string {
	position := l.position
	for isLetter(l.peek()) {
		l.advance()
	}
	return l.input[position : l.position+1]
}
func (l *Lexer) number() string {
	position := l.position
	for isDigit(l.peek()) {
		l.advance()
	}
	return l.input[position : l.position+1]
}
func (l *Lexer) string() string {
	l.advance()
	position := l.position
	for !l.eof && l.peek() != `"` {
		l.advance()
	}

	if l.eof {
		l.error("Unterminated string at line %d", l.line)
		return ""
	}

	l.advance()

	return l.input[position:l.position]
}
func (l *Lexer) error(format string, args ...interface{}) {
	l.err = fmt.Errorf(format, args...)
}
func (l *Lexer) Error() error {
	return l.err
}

func (l *Lexer) keyword(ident string) *token.Token {
	switch ident {
	case "if":
		return l.token(token.IF, "if")
	case "else":
		return l.token(token.ELSE, "else")
	case "while":
		return l.token(token.WHILE, "while")
	case "false":
		return l.token(token.FALSE, "false")
	case "true":
		return l.token(token.TRUE, "true")
	case "fn":
		return l.token(token.FUNCTION, "fn")
	case "return":
		return l.token(token.RETURN, "return")
	case "end":
		return l.token(token.END, "end")
	case "not":
		return l.token(token.NOT, "not")
	case "then":
		return l.token(token.THEN, "then")
	case "elif":
		return l.token(token.ELIF, "elif")
	case "test":
		return l.token(token.TEST, "test")
	case "class":
		return l.token(token.CLASS, "class")
	default:
		return l.token(token.IDENT, ident)
	}
}

func (l *Lexer) peek() string {
	if l.readPosition < len(l.input) {
		return string(l.input[l.readPosition])
	}
	return ""
}

func (l *Lexer) whitespace() {
	for l.ch == " " || l.ch == "\t" || l.ch == "\n" || l.ch == "\r" {
		if l.ch == "\n" {
			l.line++
		}
		l.advance()
	}
}
func (l *Lexer) comment() string {
	// Move over the second semicolon
	l.advance()
	var comment string
	for l.peek() != "\n" && !l.eof {
		comment += l.ch
		l.advance()
	}
	comment += l.ch

	return comment
}

func (l *Lexer) Next() *token.Token {

	l.advance()
	l.whitespace()

	if l.eof {
		return l.token(token.EOF, "")
	}

	switch l.ch {
	case "&":
		if l.peek() == "&" {
			l.advance()
			return l.token(token.AND, "&&")
		}
		return l.token(token.BITAND, l.ch)
	case "|":
		if l.peek() == "|" {
			l.advance()
			return l.token(token.OR, "||")
		}
		return l.token(token.BITOR, l.ch)
	case ".":
		return l.token(token.DOT, ".")
	case ",":
		return l.token(token.COMMA, ",")
	case "(":
		return l.token(token.LPAREN, "(")
	case ")":
		return l.token(token.RPAREN, ")")
	case "[":
		return l.token(token.LSQUARE, "[")
	case "]":
		return l.token(token.RSQUARE, "]")
	case "+":
		return l.token(token.PLUS, "+")
	case "-":
		return l.token(token.MINUS, "-")
	case "*":
		return l.token(token.MULTIPLY, "*")
	case "/":
		return l.token(token.DIVIDE, "/")
	case "%":
		return l.token(token.MOD, "%")
	case ";":
		if l.peek() == ";" {
			l.advance()
			comment := l.comment()
			return l.token(token.COMMENT, comment)
		}
		return l.token(token.ILLEGAL, l.ch)
	case "<":
		if l.peek() == "=" {
			l.advance()
			return l.token(token.LTE, "<=")
		}
		return l.token(token.LT, "<")
	case ">":
		if l.peek() == "=" {
			l.advance()
			return l.token(token.GTE, ">=")
		}
		return l.token(token.GT, ">")
	case "=":
		if l.peek() == "=" {
			l.advance()
			return l.token(token.EQ, "==")
		}
		return l.token(token.ASSIGN, "=")
	case "!":
		if l.peek() == "=" {
			l.advance()
			return l.token(token.NOT_EQ, "!=")
		}
		return l.token(token.BANG, "!")
	case `"`:
		return l.token(token.STRING, l.string())
	default:
		if isLetter(l.ch) {
			identifier := l.identifier()
			return l.keyword(identifier)
		} else if isDigit(l.ch) {
			return l.token(token.INT, l.number())
		}
	}
	return l.token(token.ILLEGAL, l.ch)
}
