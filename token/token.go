package token

import "fmt"

type TokenType string

type Token struct {
	Type  TokenType
	Value string
	Line  int
}

func (t Token) String() string {
	return fmt.Sprintf("Token(%s, '%s')", t.Type, t.Value)
}

const (
	COMMENT = "COMMENT"

	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers and literals
	IDENT  = "IDENT" // add, foobar, x, y, ...
	INT    = "INT"   // 123456, 2.12
	STRING = "STRING"

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	MULTIPLY = "*"
	DIVIDE   = "/"
	MOD      = "%"
	BANG     = "!"

	// Delimiters
	LPAREN   = "("
	RPAREN   = ")"
	LBRACE   = "{"
	RBRACE   = "}"
	LBRACKET = "["
	RBRACKET = "]"

	// Keywords
	IF    = "IF"
	ELSE  = "ELSE"
	WHILE = "WHILE"

	FALSE = "FALSE"
	TRUE  = "TRUE"

	FUNCTION = "FUNCTION"
	RETURN   = "RETURN"
	END      = "END"
	THEN     = "THEN"

	AND = "AND"
	OR  = "OR"
	NOT = "NOT"

	// Comparison operators
	EQ     = "=="
	NOT_EQ = "!="
	LT     = "<"
	GT     = ">"
	LTE    = "<="
	GTE    = ">="

	// Misc

	DOT   = "."
	COMMA = ","
)
