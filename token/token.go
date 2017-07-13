package token

// TokenType will include various different token types
type TokenType string

// Token contains the token type and the literal value
type Token struct {
	Type    TokenType
	Literal string
}

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
	MULTIPLIEDBY = "*"

	// DIVIDEDBY is the division operator
	DIVIDEDBY = "/"

	// MODULO is the modulo operator
	MODULO = "%"

	// Delimiters

	// COMMA separates input
	COMMA = ","

	// SEMICOLON indicates the end of a line
	SEMICOLON = ";"

	// LPAREN indicates the opening of a function call
	LPAREN = "("

	// RPAREN indicates the closing of a function call
	RPAREN = ")"

	// LBRACE indicates the opening of a closure
	LBRACE = "{"

	// RBRACE indicates the closing of a closure
	RBRACE = "}"

	// Keywords

	// FUNCTION indicates a function definition
	FUNCTION = "FUNCTION"

	// LET indicates variable assignment
	LET = "LET"
)
