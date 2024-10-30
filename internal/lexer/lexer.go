package lexer

import "unicode"

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

type Lexer struct {
	input        string
	position     int
	readPosition int
	currentChar  byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.currentChar = 0
	} else {
		l.currentChar = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

func (l *Lexer) NextToken() Token {
	var tok Token

	l.skipWhitespace()

	switch l.currentChar {
	case '=':
		if l.peekChar() == '=' {
			ch := l.currentChar
			l.readChar()
			tok = Token{Type: EQ, Literal: string(ch) + string(l.currentChar)}
		} else {
			tok = newToken(ASSIGN, l.currentChar)
		}
	case ';':
		if l.peekChar() == '=' {
			l.readChar()
			tok = Token{Type: ASSIGN, Literal: ";="}
		} else {
			tok = newToken(SEMICOLON, l.currentChar)
		}
	case '!':
		if l.peekChar() == '=' {
			ch := l.currentChar
			l.readChar()
			tok = Token{Type: NOT_EQ, Literal: string(ch) + string(l.currentChar)}
		} else {
			tok = newToken(BANG, l.currentChar)
		}
	case '+':
		if l.peekChar() == '+' {
			ch := l.currentChar
			l.readChar()
			tok = Token{Type: INC, Literal: string(ch) + string(l.currentChar)}
		} else {
			tok = newToken(PLUS, l.currentChar)
		}
	case '-':
		if l.peekChar() == '>' {
			ch := l.currentChar
			l.readChar()
			tok = Token{Type: ARROW, Literal: string(ch) + string(l.currentChar)}
		} else if l.peekChar() == '-' {
			ch := l.currentChar
			l.readChar()
			tok = Token{Type: DEC, Literal: string(ch) + string(l.currentChar)}
		} else {
			tok = newToken(MINUS, l.currentChar)
		}
	case '.':
		if l.peekChar() == '.' {
			ch := l.currentChar
			l.readChar()
			tok = Token{Type: CONCAT, Literal: string(ch) + string(l.currentChar)}
		} else {
			tok = newToken(ILLEGAL, l.currentChar)
		}
	case '*':
		tok = newToken(ASTERISK, l.currentChar)
	case '/':
		tok = newToken(SLASH, l.currentChar)
	case '%':
		tok = newToken(MOD, l.currentChar)
	case '<':
		tok = newToken(LT, l.currentChar)
	case '>':
		tok = newToken(GT, l.currentChar)
	case ',':
		tok = newToken(COMMA, l.currentChar)
	case ':':
		tok = newToken(COLON, l.currentChar)
	case '(':
		tok = newToken(LEFT_PAREN, l.currentChar)
	case ')':
		tok = newToken(RIGHT_PAREN, l.currentChar)
	case '{':
		tok = newToken(LEFT_BRACE, l.currentChar)
	case '}':
		tok = newToken(RIGHT_BRACE, l.currentChar)
	case '&':
		if l.peekChar() == '&' {
			ch := l.currentChar
			l.readChar()
			tok = Token{Type: AND, Literal: string(ch) + string(l.currentChar)}
		} else {
			tok = newToken(ILLEGAL, l.currentChar)
		}
	case '|':
		if l.peekChar() == '|' {
			ch := l.currentChar
			l.readChar()
			tok = Token{Type: OR, Literal: string(ch) + string(l.currentChar)}
		} else {
			tok = newToken(ILLEGAL, l.currentChar)
		}
	case '"':
		tok.Type = STRING
		tok.Literal = l.readString()
		return tok
	case '#':
		if l.peekChar() == '#' {
			l.readChar() // consume second #
			comment := l.readMultiLineComment()
			tok = Token{Type: COMMENT_MULTI, Literal: comment}
			return tok
		} else {
			comment := l.readSingleLineComment()
			tok = Token{Type: COMMENT_SINGLE, Literal: comment}
			return tok
		}
	case 0:
		tok.Literal = ""
		tok.Type = EOF
	default:
		if isLetter(l.currentChar) {
			tok.Literal = l.readIdentifier()
			tok.Type = LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.currentChar) {
			tok.Literal = l.readNumber()
			tok.Type = INT
			return tok
		} else {
			tok = newToken(ILLEGAL, l.currentChar)
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) skipWhitespace() {
	for unicode.IsSpace(rune(l.currentChar)) {
		l.readChar()
	}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.currentChar) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.currentChar) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readString() string {
	position := l.position + 1
	l.readChar()
	for l.currentChar != '"' && l.currentChar != 0 {
		l.readChar()
	}
	if l.currentChar == 0 {
		return l.input[position:l.position]
	}
	str := l.input[position:l.position]
	l.readChar()
	return str
}

func (l *Lexer) readSingleLineComment() string {
	position := l.position
	for l.currentChar != '\n' && l.currentChar != 0 {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readMultiLineComment() string {
	position := l.position - 1 // include the ##
	for {
		if l.currentChar == '#' && l.peekChar() == '#' {
			l.readChar()
			l.readChar()
			return l.input[position:l.position]
		}
		if l.currentChar == 0 {
			return l.input[position:l.position]
        }
		l.readChar()
	}
}

func newToken(tokenType TokenType, ch byte) Token {
	return Token{Type: tokenType, Literal: string(ch)}
}

func isLetter(ch byte) bool {
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z') || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
