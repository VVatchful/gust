package lexer

import "unicode"

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

func (l *Lexer) NextToken() Token {
	var tok Token

	for unicode.IsSpace(rune(l.currentChar)) {
		l.readChar()
	}

	switch l.currentChar {
	case '=':
		tok = Token{Type: ASSIGN, Literal: string(l.currentChar)}
	case '+':
		tok = Token{Type: PLUS, Literal: string(l.currentChar)}
	case '-':
		tok = Token{Type: MINUS, Literal: string(l.currentChar)}
	case '!':
		tok = Token{Type: BANG, Literal: string(l.currentChar)}
	case '*':
		tok = Token{Type: ASTERISK, Literal: string(l.currentChar)}
	case '/':
		tok = Token{Type: SLASH, Literal: string(l.currentChar)}
	case 0:
		tok = Token{Type: EOF, Literal: ""}
	case ';':
		tok = Token{Type: SEMICOLON, Literal: string(l.currentChar)}
	case ',':
		tok = Token{Type: COMMA, Literal: string(l.currentChar)}
	case '(':
		tok = Token{Type: LEFT_PAREN, Literal: string(l.currentChar)}
	case ')':
		tok = Token{Type: RIGHT_PAREN, Literal: string(l.currentChar)}
	case '{':
		tok = Token{Type: LEFT_BRACE, Literal: string(l.currentChar)}
	case '}':
		tok = Token{Type: RIGHT_BRACE, Literal: string(l.currentChar)}
	default:
		if isLetter(l.currentChar) {
			tok.Literal = l.readIdentifier()
			tok.Type = IDENT
			return tok
    } else if isDigit(l.currentChar) {
			tok.Literal = l.readNumber()
			tok.Type = INT
			return tok
		} else {
			tok = Token{Type: ILLEGAL, Literal: string(l.currentChar)}
		}
	}

	l.readChar()
	return tok
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

func isLetter(ch byte) bool {
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z') || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
