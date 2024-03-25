package lexer

import "testing"
import "github.com/pspiagicw/fener/token"

func TestTokenIntegers(t *testing.T) {
	input := `123 -456 789`

	expectedTokens := []token.Token{
		{Type: token.INT, Value: "123"},
		{Type: token.MINUS, Value: "-"},
		{Type: token.INT, Value: "456"},
		{Type: token.INT, Value: "789"},
		{Type: token.EOF, Value: ""},
	}

	checkTokens(t, expectedTokens, input)
}

func TestTokenBooleans(t *testing.T) {
	input := `true false`

	expectedTokens := []token.Token{
		{Type: token.TRUE, Value: "true"},
		{Type: token.FALSE, Value: "false"},
		{Type: token.EOF, Value: ""},
	}

	checkTokens(t, expectedTokens, input)
}

func TestStringToken(t *testing.T) {
	// Test case for strings
	input := `"Hello, World!" 
    1`

	expectedTokens := []token.Token{
		{Type: token.STRING, Value: "Hello, World!"},
		{Type: token.INT, Value: "1"},
		{Type: token.EOF, Value: ""},
	}
	checkTokens(t, expectedTokens, input)
}

func TestCommentToken(t *testing.T) {
	// Test case for comments
	input := `;; This is a comment`

	expectedTokens := []token.Token{
		{Type: token.COMMENT, Value: " This is a comment"},
		{Type: token.EOF, Value: ""},
	}
	checkTokens(t, expectedTokens, input)
}

func TestSymbolTokens(t *testing.T) {
	// Test case for symbols
	input := "+-*/=!%"

	expectedTokens := []token.Token{
		{Type: token.PLUS, Value: "+"},
		{Type: token.MINUS, Value: "-"},
		{Type: token.MULTIPLY, Value: "*"},
		{Type: token.DIVIDE, Value: "/"},
		{Type: token.ASSIGN, Value: "="},
		{Type: token.BANG, Value: "!"},
		{Type: token.MOD, Value: "%"},
		{Type: token.EOF, Value: ""},
	}
	checkTokens(t, expectedTokens, input)
}

func TestIdentifierTokens(t *testing.T) {
	// Test case for identifiers
	input := "foo bar baz foobar"

	expectedTokens := []token.Token{
		{Type: token.IDENT, Value: "foo"},
		{Type: token.IDENT, Value: "bar"},
		{Type: token.IDENT, Value: "baz"},
		{Type: token.IDENT, Value: "foobar"},
		{Type: token.EOF, Value: ""},
	}
	checkTokens(t, expectedTokens, input)
}

func TestKeywordTokens(t *testing.T) {
	// Test case for keywords
	input := "if else while false true return fn end and or not then"

	expectedTokens := []token.Token{
		{Type: token.IF, Value: "if"},
		{Type: token.ELSE, Value: "else"},
		{Type: token.WHILE, Value: "while"},
		{Type: token.FALSE, Value: "false"},
		{Type: token.TRUE, Value: "true"},
		{Type: token.RETURN, Value: "return"},
		{Type: token.FUNCTION, Value: "fn"},
		{Type: token.END, Value: "end"},
		{Type: token.AND, Value: "and"},
		{Type: token.OR, Value: "or"},
		{Type: token.NOT, Value: "not"},
		{Type: token.THEN, Value: "then"},
		{Type: token.EOF, Value: ""},
	}
	checkTokens(t, expectedTokens, input)
}

func TestEqualityRelationalTokens(t *testing.T) {
	// Test case for equality and relational operators
	input := "== != < > <= >="

	expectedTokens := []token.Token{
		{Type: token.EQ, Value: "=="},
		{Type: token.NOT_EQ, Value: "!="},
		{Type: token.LT, Value: "<"},
		{Type: token.GT, Value: ">"},
		{Type: token.LTE, Value: "<="},
		{Type: token.GTE, Value: ">="},
		{Type: token.EOF, Value: ""},
	}
	checkTokens(t, expectedTokens, input)
}

func TestMiscellaneousTokens(t *testing.T) {
	input := "."

	expectedTokens := []token.Token{
		{Type: token.DOT, Value: "."},
		{Type: token.EOF, Value: ""},
	}

	checkTokens(t, expectedTokens, input)
}
func checkTokens(t *testing.T, expected []token.Token, input string) {
	t.Helper()

	lexer := NewLexer(input)

	for i, expectedToken := range expected {
		actualToken := lexer.Next()
		if actualToken == nil {
			t.Fatalf("Parsed token is nil")
		}
		matchToken(t, i, expectedToken, actualToken)
	}

}
func matchToken(t *testing.T, i int, expected token.Token, actual *token.Token) {
	t.Helper()
	if actual.Type != expected.Type {
		t.Errorf("Test [%d], Expected Type: '%v', Actual TokenType: '%v'", i, expected.Type, actual.Type)

	}
	if actual.Value != expected.Value {
		t.Errorf("Test [%d], Expected Value: '%v', Actual TokenValue: '%v'", i, expected.Value, actual.Value)
	}

}

func TestLexerTokenization(t *testing.T) {
	input := `
        fn factorial(n)
            if n <= 1 then
                return 1
            end
            
            return n * factorial(n-1)
        end

        result = factorial(5)
        print(result)
    `

	expectedTokens := []token.Token{
		{Type: token.FUNCTION, Value: "fn"},
		{Type: token.IDENT, Value: "factorial"},
		{Type: token.LPAREN, Value: "("},
		{Type: token.IDENT, Value: "n"},
		{Type: token.RPAREN, Value: ")"},
		{Type: token.IF, Value: "if"},
		{Type: token.IDENT, Value: "n"},
		{Type: token.LTE, Value: "<="},
		{Type: token.INT, Value: "1"},
		{Type: token.THEN, Value: "then"},
		{Type: token.RETURN, Value: "return"},
		{Type: token.INT, Value: "1"},
		{Type: token.END, Value: "end"},
		{Type: token.RETURN, Value: "return"},
		{Type: token.IDENT, Value: "n"},
		{Type: token.MULTIPLY, Value: "*"},
		{Type: token.IDENT, Value: "factorial"},
		{Type: token.LPAREN, Value: "("},
		{Type: token.IDENT, Value: "n"},
		{Type: token.MINUS, Value: "-"},
		{Type: token.INT, Value: "1"},
		{Type: token.RPAREN, Value: ")"},
		{Type: token.END, Value: "end"},
		{Type: token.IDENT, Value: "result"},
		{Type: token.ASSIGN, Value: "="},
		{Type: token.IDENT, Value: "factorial"},
		{Type: token.LPAREN, Value: "("},
		{Type: token.INT, Value: "5"},
		{Type: token.RPAREN, Value: ")"},
		{Type: token.IDENT, Value: "print"},
		{Type: token.LPAREN, Value: "("},
		{Type: token.IDENT, Value: "result"},
		{Type: token.RPAREN, Value: ")"},
	}

	checkTokens(t, expectedTokens, input)

}
