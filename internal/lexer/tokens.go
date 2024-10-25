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
	MOD
	CONCAT
	EQ
	NOT_EQ
	LT
	GT
	INC
	DEC
	AND
	OR
	SEMICOLON
	COMMA
	COLON
	LEFT_PAREN
	RIGHT_PAREN
	LEFT_BRACE
	RIGHT_BRACE
	ARROW
	FUNCTION
	LET
	RETURN
	FOR
	IF
	ELSE
	TRUE
	FALSE
	COMMENT_SINGLE
  COMMENT_MULTI
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"return": RETURN,
	"for":    FOR,
	"if":     IF,
	"else":   ELSE,
	"true":   TRUE,
	"false":  FALSE,
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
	case "++":
		return Token{Type: INC, Literal: "++"}
	case "--":
		return Token{Type: DEC, Literal: "--"}
	case "%":
		return Token{Type: MOD, Literal: "%"}
	case ":":
		return Token{Type: COLON, Literal: ":"}
	case "->":
		return Token{Type: ARROW, Literal: "->"}
	case "#":
		return Token{Type: COMMENT_SINGLE, Literal: "#"}
	case "##":
		return Token{Type: COMMENT_MULTI, Literal: "##"}
	default:
		return Token{Type: ILLEGAL, Literal: input}
	}
}

