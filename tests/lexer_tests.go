package test

import (
	"testing"
  "github.com/voidwyrm-2/gust/internal/lexer"
)

func TestNextToken(t *testing.T) {
	input := `let five = 5;
              let ten = 10;
              let add = fn(x, y) {
                 x + y;
              };
              let result = add(five, ten);
              !-/*5;
              5 < 10 > 5;
              if (5 < 10) { return true; } else { return false; }
              10 == 10; 10 != 9;
              && ||`

	tests := []struct {
		expectedType    lexer.TokenType
		expectedLiteral string
	}{
		{lexer.LET, "let"},
		{lexer.IDENT, "five"},
		{lexer.ASSIGN, "="},
		{lexer.INT, "5"},
		{lexer.SEMICOLON, ";"},
		{lexer.LET, "let"},
		{lexer.IDENT, "ten"},
		{lexer.ASSIGN, "="},
		{lexer.INT, "10"},
		{lexer.SEMICOLON, ";"},
		{lexer.LET, "let"},
		{lexer.IDENT, "add"},
		{lexer.ASSIGN, "="},
		{lexer.FUNCTION, "fn"},
		{lexer.LEFT_PAREN, "("},
		{lexer.IDENT, "x"},
		{lexer.COMMA, ","},
		{lexer.IDENT, "y"},
		{lexer.RIGHT_PAREN, ")"},
		{lexer.LEFT_BRACE, "{"},
		{lexer.IDENT, "x"},
		{lexer.PLUS, "+"},
		{lexer.IDENT, "y"},
		{lexer.SEMICOLON, ";"},
		{lexer.RIGHT_BRACE, "}"},
		{lexer.SEMICOLON, ";"},
		{lexer.LET, "let"},
		{lexer.IDENT, "result"},
		{lexer.ASSIGN, "="},
		{lexer.IDENT, "add"},
		{lexer.LEFT_PAREN, "("},
		{lexer.IDENT, "five"},
		{lexer.COMMA, ","},
		{lexer.IDENT, "ten"},
		{lexer.RIGHT_PAREN, ")"},
		{lexer.SEMICOLON, ";"},
		{lexer.BANG, "!"},
		{lexer.MINUS, "-"},
		{lexer.SLASH, "/"},
		{lexer.ASTERISK, "*"},
		{lexer.INT, "5"},
		{lexer.SEMICOLON, ";"},
		{lexer.INT, "5"},
		{lexer.LT, "<"},
		{lexer.INT, "10"},
		{lexer.GT, ">"},
		{lexer.INT, "5"},
		{lexer.SEMICOLON, ";"},
		{lexer.IF, "if"},
		{lexer.LEFT_PAREN, "("},
		{lexer.INT, "5"},
		{lexer.LT, "<"},
		{lexer.INT, "10"},
		{lexer.RIGHT_PAREN, ")"},
		{lexer.LEFT_BRACE, "{"},
		{lexer.RETURN, "return"},
		{lexer.TRUE, "true"},
		{lexer.SEMICOLON, ";"},
		{lexer.RIGHT_BRACE, "}"},
		{lexer.ELSE, "else"},
		{lexer.LEFT_BRACE, "{"},
		{lexer.RETURN, "return"},
		{lexer.FALSE, "false"},
		{lexer.SEMICOLON, ";"},
		{lexer.RIGHT_BRACE, "}"},
		{lexer.INT, "10"},
		{lexer.EQ, "=="},
		{lexer.INT, "10"},
		{lexer.SEMICOLON, ";"},
		{lexer.INT, "10"},
		{lexer.NOT_EQ, "!="},
		{lexer.INT, "9"},
		{lexer.SEMICOLON, ";"},
		{lexer.AND, "&&"},
		{lexer.OR, "||"},
	}

	l := lexer.New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
