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
		l.ch = 0
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

	// assign the appropriate token value
	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '*':
		tok = newToken(token.MULTIPLIEDBY, l.ch)
	case '/':
		tok = newToken(token.DIVIDEDBY, l.ch)
	case '%':
		tok = newToken(token.MODULO, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '0':
		tok.Type = token.EOF
		tok.Literal = ""
	}

	// read in the next ch in the input string
	l.readChar()

	return tok
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
