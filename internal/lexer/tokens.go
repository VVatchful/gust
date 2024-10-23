package lexer

type TokenType int

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL TokenType = iota
	EOF
	IDENT
	INT
	STRING
	ASSIGN
	PLUS
	MINUS
	BANG
	ASTERISK
	SLASH
	CONCAT
	EQ
	NOT_EQ
	LT
	GT
	INC
	SEMICOLON
	COMMA
	LEFT_PAREN
	RIGHT_PAREN
	LEFT_BRACE
	RIGHT_BRACE
	FUNCTION
	LET
	RETURN
	FOR
	AND
	OR
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"return": RETURN,
	"for":    FOR,
}

func LookUpIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

func NextToken(input string) Token {
	switch input {
	case "&&":
		return Token{Type: AND, Literal: "&&"}
	case "||":
		return Token{Type: OR, Literal: "||"}
	}
	return Token{Type: ILLEGAL, Literal: input}
}
