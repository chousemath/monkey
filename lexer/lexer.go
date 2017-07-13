package lexer

// Lexer represent the lexer datatype
type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

// readChar gives us the next char and advances our position in the input
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

// New creates a reference to a lexer
func New(input string) *Lexer {
	l := &Lexer{input: input}
	return l
}
