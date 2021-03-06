package lexer

import "github.com/chousemath/monkey/token"

// this lexer only supports ASCII characters

// Lexer represent the lexer datatype
type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

// readChar gives us the next char and advances our position in the input
func (l *Lexer) readChar() {
	// check if at end of input
	if l.readPosition >= len(l.input) {
		// 0 is the ASCII code of NUL -> nothing read, or EOF
		l.ch = '0'
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

//=+-*/%(){},;

// NextToken will assign a token to the incoming input char
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	// assign the appropriate token value
	switch l.ch {
	case '=':
		tok = l.peekAndConcat(token.ASSIGN, token.EQ)
	case '+':
		tok = l.peekAndConcat(token.PLUS, token.PLUSEQ)
	case '-':
		tok = l.peekAndConcat(token.MINUS, token.MINUSEQ)
	case '*':
		tok = l.peekAndConcat(token.ASTERISK, token.ASTERISKEQ)
	case '/':
		tok = l.peekAndConcat(token.SLASH, token.SLASHEQ)
	case '|':
		tok = l.peekAndConcat(token.BAR, token.BAREQ)
	case '%':
		tok = newToken(token.MODULO, l.ch)
	case '^':
		tok = newToken(token.EXPONENT, l.ch)
	case '!':
		tok = l.peekAndConcat(token.BANG, token.NOTEQ)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case '[':
		tok = newToken(token.LBRACKET, l.ch)
	case ']':
		tok = newToken(token.RBRACKET, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '0':
		tok.Type = token.EOF
		tok.Literal = ""
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	// read in the next ch in the input string
	l.readChar()

	return tok
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readNumber() string {
	position := l.position

	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

func (l *Lexer) concatChar() string {
	ch := l.ch
	l.readChar()
	return string(ch) + string(l.ch)
}

func (l *Lexer) peekAndConcat(defType token.TokenType, altType token.TokenType) token.Token {
	if l.peekChar() == '=' {
		return token.Token{Type: altType, Literal: l.concatChar()}
	}
	return newToken(defType, l.ch)
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isLetter(ch byte) bool {
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z') || ch == '_'
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// New creates a reference to a lexer
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}
