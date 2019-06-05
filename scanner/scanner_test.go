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
