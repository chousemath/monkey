package parser

import (
	"fmt"

	"github.com/chousemath/monkey/ast"
	"github.com/chousemath/monkey/lexer"
	"github.com/chousemath/monkey/token"
)

// Parser bla bla
type Parser struct {
	// l is a pointer to an instance of a Lexer
	l *lexer.Lexer
	// curToken acts like `position` of Lexer
	// the current token tells us what to do next
	curToken token.Token
	// peekToken acts like `readPosition` of Lexer
	// the peek token gives us more information if current token is not enough
	peekToken token.Token
	errors    []string
}

// New creates a new parser instance
func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}

	// read two tokens so that curToken and peekToken are set
	p.nextToken()
	p.nextToken()

	return p
}

// Errors just returns the string slice containing all accumulated errors
func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// ParseProgram parses the program...
func (p *Parser) ParseProgram() *ast.Program {
	// first, construct the root node of the AST
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for !p.curTokenIs(token.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}
	// when there is nothing left to parse, return the root node
	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	default:
		return nil
	}
}

// creates a node, advances the token while making assertions about the token
func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}

	// first it expects to find a variable identifier
	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	// then it expects to see an equal sign for assignment
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	// skip over the value of the expression for now

	// basically skip over the semicolon as well
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

// an assertion function, enforces correctness of order of tokens
func (p *Parser) expectPeek(t token.TokenType) bool {
	// advances the token only if the type of the next token is correct
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	}
	p.peekError(t)
	return false
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, but got %s", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}
