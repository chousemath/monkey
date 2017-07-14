package token

// TokenType will include various different token types
type TokenType string

// Token contains the token type and the literal value
type Token struct {
	Type    TokenType
	Literal string
}

// a list of possible TokenTypes
const (
	// ILLEGAL represents a token/character we do not know about
	ILLEGAL = "ILLEGAL"

	// EOF represents "End Of File", telling the Parser where to stop
	EOF = "EOF"

	// IDENT and the rest represents identifiers and literals
	IDENT = "IDENT"

	// INT represents the integer type
	INT = "INT"

	// List of Operators

	// ASSIGN is the assignment operator
	ASSIGN = "="

	// PLUS is the plus operator
	PLUS = "+"

	// MINUS is the minus operator
	MINUS = "-"

	// MULTIPLIEDBY is the multiplication operator
	ASTERISK = "*"

	// SLASH is the division operator
	SLASH = "/"

	// BAR is yet to be determined...
	BAR = "|"

	// MODULO is the modulo operator
	MODULO = "%"

	// EXPONENT is a dot product operator
	EXPONENT = "^"

	// BANG is the bang operator
	BANG = "!"

	// LT is the Less Than operator
	LT = "<"

	// GT is the Greater Than operator
	GT = ">"

	// EQ is the equality operator
	EQ = "=="

	// NOTEQ is the inequality operator
	NOTEQ = "!="

	// PLUSEQ is the equivalent of += in C
	PLUSEQ = "+="

	// MINUSEQ is the equivalent of -= in C
	MINUSEQ = "-="

	// ASTERISKEQ is the equivalent of *= in C
	ASTERISKEQ = "*="

	// SLASHEQ is the equivalent of /= in C
	SLASHEQ = "/="

	// Delimiters

	// COMMA separates input
	COMMA = ","

	// SEMICOLON indicates the end of a line
	SEMICOLON = ";"

	// LPAREN indicates the start of a function call
	LPAREN = "("

	// RPAREN indicates the end of a function call
	RPAREN = ")"

	// LBRACE indicates the start of a closure
	LBRACE = "{"

	// RBRACE indicates the end of a closure
	RBRACE = "}"

	// LBRACKET indicates the start of a list
	LBRACKET = "["

	// RBRACKET indicates the end of a list
	RBRACKET = "]"

	// Keywords

	// FUNCTION indicates a function definition
	FUNCTION = "FUNCTION"

	// LET indicates variable assignment
	LET = "LET"

	// TRUE indicates a boolean true
	TRUE = "TRUE"

	// FALSE indicates a boolean false
	FALSE = "FALSE"

	// IF indicates a conditional if
	IF = "IF"

	// ELSE indicates a conditional else
	ELSE = "ELSE"

	// RETURN indicates a return statement
	RETURN = "RETURN"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

// LookupIdent checks keyword hash map for keyword & returns TokenType or IDENT
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
