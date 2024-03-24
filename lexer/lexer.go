package lexer

import (
	"github.com/pspiagicw/fener/token"
	"github.com/pspiagicw/goreland"
)

type Lexer struct {
	input        string
	position     int    // current position in input (points to current char)
	readPosition int    // current reading position in input (after current char)
	ch           string // current char under examination
	eof          bool   // signals end of file
	line         int    // current line number
}

func NewLexer(input string) *Lexer {
	l := &Lexer{
		input:        input,
		position:     -1,
		readPosition: 0,
	}
	return l
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
		goreland.LogFatal("Unterminated string at line %d", l.line)
	}

	l.advance()

	return l.input[position:l.position]
}

func (l *Lexer) keyword(ident string) *token.Token {
	switch ident {
	case "if":
		return &token.Token{Type: token.IF, Value: "if"}
	case "else":
		return &token.Token{Type: token.ELSE, Value: "else"}
	case "while":
		return &token.Token{Type: token.WHILE, Value: "while"}
	case "false":
		return &token.Token{Type: token.FALSE, Value: "false"}
	case "true":
		return &token.Token{Type: token.TRUE, Value: "true"}
	case "fn":
		return &token.Token{Type: token.FUNCTION, Value: "fn"}
	case "return":
		return &token.Token{Type: token.RETURN, Value: "return"}
	case "end":
		return &token.Token{Type: token.END, Value: "end"}
	case "and":
		return &token.Token{Type: token.AND, Value: "and"}
	case "or":
		return &token.Token{Type: token.OR, Value: "or"}
	case "not":
		return &token.Token{Type: token.NOT, Value: "not"}
	default:
		return &token.Token{Type: token.IDENT, Value: ident}
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

	return comment
}

func (l *Lexer) Next() *token.Token {
	l.advance()
	l.whitespace()

	if l.eof {
		return &token.Token{Type: token.EOF, Value: ""}
	}

	switch l.ch {
	case ".":
		return &token.Token{Type: token.DOT, Value: "."}
	case "(":
		return &token.Token{Type: token.LPAREN, Value: "("}
	case ")":
		return &token.Token{Type: token.RPAREN, Value: ")"}
	case "+":
		return &token.Token{Type: token.PLUS, Value: "+"}
	case "-":
		if isDigit(l.peek()) {
			value := l.number()
			return &token.Token{Type: token.INT, Value: value}
		}
		return &token.Token{Type: token.MINUS, Value: "-"}
	case "*":
		return &token.Token{Type: token.MULTIPLY, Value: "*"}
	case "/":
		return &token.Token{Type: token.DIVIDE, Value: "/"}
	case "%":
		return &token.Token{Type: token.MOD, Value: "%"}
	case ";":
		if l.peek() == ";" {
			l.advance()
			comment := l.comment()
			return &token.Token{Type: token.COMMENT, Value: comment}
		}
		return &token.Token{Type: token.ILLEGAL, Value: l.ch}
	case "<":
		if l.peek() == "=" {
			l.advance()
			return &token.Token{Type: token.LTE, Value: "<="}
		}
		return &token.Token{Type: token.LT, Value: "<"}
	case ">":
		if l.peek() == "=" {
			l.advance()
			return &token.Token{Type: token.GTE, Value: ">="}
		}
		return &token.Token{Type: token.GT, Value: ">"}
	case "=":
		if l.peek() == "=" {
			l.advance()
			return &token.Token{Type: token.EQ, Value: "=="}
		}
		return &token.Token{Type: token.ASSIGN, Value: "="}
	case "!":
		if l.peek() == "=" {
			l.advance()
			return &token.Token{Type: token.NOT_EQ, Value: "!="}
		}
		return &token.Token{Type: token.BANG, Value: "!"}
	case `"`:
		return &token.Token{Type: token.STRING, Value: l.string()}
	default:
		if isLetter(l.ch) {
			identifier := l.identifier()
			return l.keyword(identifier)
		} else if isDigit(l.ch) {
			return &token.Token{Type: token.INT, Value: l.number()}
		}
	}
	return &token.Token{Type: token.ILLEGAL, Value: l.ch}
}
