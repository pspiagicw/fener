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
	LPAREN  = "("
	RPAREN  = ")"
	LBRACE  = "{"
	RBRACE  = "}"
	LSQUARE = "["
	RSQUARE = "]"

	// Keywords
	IF    = "IF"
	ELSE  = "ELSE"
	WHILE = "WHILE"
	ELIF  = "ELIF"

	FALSE = "FALSE"
	TRUE  = "TRUE"

	FUNCTION = "FUNCTION"
	RETURN   = "RETURN"
	END      = "END"
	THEN     = "THEN"
	TEST     = "TEST"
	CLASS    = "CLASS"

	AND = "AND"
	OR  = "OR"
	NOT = "NOT"

	BITAND = "BITAND"
	BITOR  = "BITOR"

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
