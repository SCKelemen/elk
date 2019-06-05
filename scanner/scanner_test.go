package scanner

import (
	"testing"

	"github.com/SCKelemen/elk/token"
)

func TestNextToken(t *testing.T) {
	input := "_:(){}"

	tests := []struct {
		expectedType    token.TokenKind
		expectedLiteral string
	}{
		{token.UNDERSCORE, "_"},
		{token.COLON, ":"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		token := l.NextToken()

		if token.Kind != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, token.Kind)
		}

		if token.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, token.Literal)
		}
	}
}

func TestExampleProgram(t *testing.T) {
	input := `
	func itoa(index: int): string {
		index match {
			1:	"one"
			2:	"two"
			3:	"three"
			4:	"four"
			_:	"Error"
		}
	}
	`

	tests := []struct {
		expectedType    token.TokenKind
		expectedLiteral string
	}{
		{token.FUNC, "func"},
		{token.IDENTITY, "itoa"},
		{token.LPAREN, "("},
		{token.IDENTITY, "index"},
		{token.COLON, ":"},
		{token.IDENTITY, "int"},
		{token.RPAREN, ")"},
		{token.COLON, ":"},
		{token.IDENTITY, "string"},
		{token.LBRACE, "{"},
		{token.IDENTITY, "index"},
		{token.MATCH, "match"},
		{token.LBRACE, "{"},
		{token.INTEGER, "1"},
		{token.COLON, ":"},
		{token.STRING, "\"one\""},
		{token.INTEGER, "2"},
		{token.COLON, ":"},
		{token.STRING, "\"two\""},
		{token.INTEGER, "3"},
		{token.COLON, ":"},
		{token.STRING, "\"three\""},
		{token.INTEGER, "4"},
		{token.COLON, ":"},
		{token.STRING, "\"four\""},
		{token.UNDERSCORE, "_"},
		{token.COLON, ":"},
		{token.STRING, "\"Error\""},
		{token.RBRACE, "}"},
		{token.RBRACE, "}"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		token := l.NextToken()

		if token.Kind != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, token.Kind)
		}

		if token.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, token.Literal)
		}
	}
}
