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

    l.skipWhitespace()

	switch l.currentChar {
	case '=':
		if l.peekChar() == '=' {
			l.readChar()
			tok = Token{Type: EQ, Literal: "=="}
		} else {
			tok = Token{Type: ASSIGN, Literal: string(l.currentChar)}
		}
	case '+':
		tok = Token{Type: PLUS, Literal: string(l.currentChar)}
	case '-':
		tok = Token{Type: MINUS, Literal: string(l.currentChar)}
	case '!':
		if l.peekChar() == '=' {
			l.readChar()
			tok = Token{Type: NOT_EQ, Literal: "!="}
		} else {
			tok = Token{Type: BANG, Literal: string(l.currentChar)}
		}
	case '*':
		tok = Token{Type: ASTERISK, Literal: string(l.currentChar)}
	case '/':
		tok = Token{Type: SLASH, Literal: string(l.currentChar)}
	case '&':
		if l.peekChar() == '&' {
			l.readChar()
			tok = Token{Type: AND, Literal: "&&"}
		} else {
			tok = Token{Type: ILLEGAL, Literal: string(l.currentChar)}
		}
	case '|':
		if l.peekChar() == '|' {
			l.readChar()
			tok = Token{Type: OR, Literal: "||"}
		} else {
			tok = Token{Type: ILLEGAL, Literal: string(l.currentChar)}
		}
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
  case '"':
    tok.Literal = l.readString()
    tok.Type = STRING
        return tok
	case 0:
		tok = Token{Type: EOF, Literal: ""}
	default:
		if isLetter(l.currentChar) {
			tok.Literal = l.readIdentifier()
			tok.Type = LookUpIdent(tok.Literal)
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

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
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
    l.readChar()
    return l.input[position:l.position-1]
}

func (l *Lexer) skipWhitespace() {
    for unicode.IsSpace(rune(l.currentChar)) {
        l.readChar()
    }
}

func isLetter(ch byte) bool {
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z') || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
