package test

import (
	"fmt"
	"os"
	"testing"
	"github.com/voidwyrm-2/gust/internal/lexer"
)

func TestLexer(t *testing.T) {
	data, err := os.ReadFile("../examples/hello_world.gt")
	if err != nil {
		t.Fatalf("Could not read example file: %v", err)
	}

	input := string(data)
	l := lexer.New(input)

	expectedTokens := []lexer.Token{
		// println("Hello World")
		{Type: lexer.IDENT, Literal: "println"},
		{Type: lexer.LEFT_PAREN, Literal: "("},
		{Type: lexer.STRING, Literal: "Hello World"},
		{Type: lexer.RIGHT_PAREN, Literal: ")"},
		{Type: lexer.SEMICOLON, Literal: ";"},
		{Type: lexer.COMMENT_SINGLE, Literal: "# this is a single line comment"},
		{Type: lexer.COMMENT_MULTI, Literal: "## this is a multi-line comment ##"},

		// fn greet(name: str) -> str {
		{Type: lexer.FUNCTION, Literal: "fn"},
		{Type: lexer.IDENT, Literal: "greet"},
		{Type: lexer.LEFT_PAREN, Literal: "("},
		{Type: lexer.IDENT, Literal: "name"},
		{Type: lexer.COLON, Literal: ":"},
		{Type: lexer.IDENT, Literal: "str"},
		{Type: lexer.RIGHT_PAREN, Literal: ")"},
		{Type: lexer.ARROW, Literal: "->"},
		{Type: lexer.IDENT, Literal: "str"},
		{Type: lexer.LEFT_BRACE, Literal: "{"},

		// return "hello " .. name .. "! Nice to meet you!"
		{Type: lexer.RETURN, Literal: "return"},
		{Type: lexer.STRING, Literal: "hello "},
		{Type: lexer.CONCAT, Literal: ".."},
		{Type: lexer.IDENT, Literal: "name"},
		{Type: lexer.CONCAT, Literal: ".."},
		{Type: lexer.STRING, Literal: "! Nice to meet you!"},
		{Type: lexer.SEMICOLON, Literal: ";"},
		{Type: lexer.RIGHT_BRACE, Literal: "}"},

		// let greeting = greet("nick")
		{Type: lexer.LET, Literal: "let"},
		{Type: lexer.IDENT, Literal: "greeting"},
		{Type: lexer.ASSIGN, Literal: "="},
		{Type: lexer.IDENT, Literal: "greet"},
		{Type: lexer.LEFT_PAREN, Literal: "("},
		{Type: lexer.STRING, Literal: "nick"},
		{Type: lexer.RIGHT_PAREN, Literal: ")"},
		{Type: lexer.SEMICOLON, Literal: ";"},

		// println(greeting)
		{Type: lexer.IDENT, Literal: "println"},
		{Type: lexer.LEFT_PAREN, Literal: "("},
		{Type: lexer.IDENT, Literal: "greeting"},
		{Type: lexer.RIGHT_PAREN, Literal: ")"},
		{Type: lexer.SEMICOLON, Literal: ";"},

		// fn fizzbuzz() { ... }
		{Type: lexer.FUNCTION, Literal: "fn"},
		{Type: lexer.IDENT, Literal: "fizzbuzz"},
		{Type: lexer.LEFT_PAREN, Literal: "("},
		{Type: lexer.RIGHT_PAREN, Literal: ")"},
		{Type: lexer.LEFT_BRACE, Literal: "{"},

		// for i ;= 1, i != 100, i++ {
		{Type: lexer.FOR, Literal: "for"},
		{Type: lexer.IDENT, Literal: "i"},
		{Type: lexer.ASSIGN, Literal: ";="},
		{Type: lexer.INT, Literal: "1"},
		{Type: lexer.COMMA, Literal: ","},
		{Type: lexer.IDENT, Literal: "i"},
		{Type: lexer.NOT_EQ, Literal: "!="},
		{Type: lexer.INT, Literal: "100"},
		{Type: lexer.COMMA, Literal: ","},
		{Type: lexer.IDENT, Literal: "i"},
		{Type: lexer.INC, Literal: "++"},
		{Type: lexer.LEFT_BRACE, Literal: "{"},

		// if i % 3 == 0 && i % 5 == 0 {
		{Type: lexer.IF, Literal: "if"},
		{Type: lexer.IDENT, Literal: "i"},
		{Type: lexer.MOD, Literal: "%"},
		{Type: lexer.INT, Literal: "3"},
		{Type: lexer.EQ, Literal: "=="},
		{Type: lexer.INT, Literal: "0"},
		{Type: lexer.AND, Literal: "&&"},
		{Type: lexer.IDENT, Literal: "i"},
		{Type: lexer.MOD, Literal: "%"},
		{Type: lexer.INT, Literal: "5"},
		{Type: lexer.EQ, Literal: "=="},
		{Type: lexer.INT, Literal: "0"},
		{Type: lexer.LEFT_BRACE, Literal: "{"},



		{Type: lexer.RIGHT_BRACE, Literal: "}"},
		{Type: lexer.EOF, Literal: ""},
	}

	for i, expected := range expectedTokens {
		tok := l.NextToken()
		if tok.Type != expected.Type {
			t.Fatalf("Token type mismatch at index %d: expected %v, got %v", i, expected.Type, tok.Type)
		}
		if tok.Literal != expected.Literal {
			t.Fatalf("Token literal mismatch at index %d: expected %q, got %q", i, expected.Literal, tok.Literal)
		}

		fmt.Printf("Token %d: type=%v, literal=%v\n", i, tok.Type, tok.Literal)
	}
}

