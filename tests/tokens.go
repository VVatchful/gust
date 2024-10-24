package tests

import (
	"testing"
  "github.com/voidwyrm-2/gust/internal/lexer"
)

func TestToken(t *testing.T) {
	tests := []struct {
		input    string
		expected []lexer.Token
	}{
		{
			"let x = 10;",
			[]lexer.Token{
				{Type: lexer.LET, Literal: "let"},
				{Type: lexer.IDENT, Literal: "x"},
				{Type: lexer.ASSIGN, Literal: "="},
				{Type: lexer.INT, Literal: "10"},
				{Type: lexer.SEMICOLON, Literal: ";"},
				{Type: lexer.EOF, Literal: ""},
			},
		},
		{
			"if (x > 10) { return true; } else { return false; }",
			[]lexer.Token{
				{Type: lexer.IF, Literal: "if"},
				{Type: lexer.LEFT_PAREN, Literal: "("},
				{Type: lexer.IDENT, Literal: "x"},
				{Type: lexer.GT, Literal: ">"},
				{Type: lexer.INT, Literal: "10"},
				{Type: lexer.RIGHT_PAREN, Literal: ")"},
				{Type: lexer.LEFT_BRACE, Literal: "{"},
				{Type: lexer.RETURN, Literal: "return"},
				{Type: lexer.TRUE, Literal: "true"},
				{Type: lexer.SEMICOLON, Literal: ";"},
				{Type: lexer.RIGHT_BRACE, Literal: "}"},
				{Type: lexer.ELSE, Literal: "else"},
				{Type: lexer.LEFT_BRACE, Literal: "{"},
				{Type: lexer.RETURN, Literal: "return"},
				{Type: lexer.FALSE, Literal: "false"},
				{Type: lexer.SEMICOLON, Literal: ";"},
				{Type: lexer.RIGHT_BRACE, Literal: "}"},
				{Type: lexer.EOF, Literal: ""},
			},
		},
		{
			"true && false",
			[]lexer.Token{
				{Type: lexer.TRUE, Literal: "true"},
				{Type: lexer.AND, Literal: "&&"},
				{Type: lexer.FALSE, Literal: "false"},
				{Type: lexer.EOF, Literal: ""},
			},
		},
		{
			"if (x < 10) { return false; }",
			[]lexer.Token{
				{Type: lexer.IF, Literal: "if"},
				{Type: lexer.LEFT_PAREN, Literal: "("},
				{Type: lexer.IDENT, Literal: "x"},
				{Type: lexer.LT, Literal: "<"},
				{Type: lexer.INT, Literal: "10"},
				{Type: lexer.RIGHT_PAREN, Literal: ")"},
				{Type: lexer.LEFT_BRACE, Literal: "{"},
				{Type: lexer.RETURN, Literal: "return"},
				{Type: lexer.FALSE, Literal: "false"},
				{Type: lexer.SEMICOLON, Literal: ";"},
				{Type: lexer.RIGHT_BRACE, Literal: "}"},
				{Type: lexer.EOF, Literal: ""},
			},
		},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)

		for i, expectedToken := range tt.expected {
			tok := l.NextToken()
			if tok.Type != expectedToken.Type {
				t.Fatalf("tests[%d] - tok.type wrong. expected=%q, got=%q", i, expectedToken.Type, tok.Type)
			}
			if tok.Literal != expectedToken.Literal {
				t.Fatalf("tests[%d] - tok.literal wrong. expected=%q, got=%q", i, expectedToken.Literal, tok.Literal)
			}
		}
	}
}

