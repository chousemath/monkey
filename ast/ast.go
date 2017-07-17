package ast

import "github.com/chousemath/monkey/token"

// an interface is defined behavior

// Node represents a node on the AST tree
type Node interface {
	// TokenLiteral will only be used for debugging and testing
	TokenLiteral() string
}

// Statement does not return a value
type Statement interface {
	Node
	// helps guid the Go compiler
	statementNode()
}

// Expression returns a value
type Expression interface {
	Node
	// helps guid the Go compiler
	expressionNode()
}

// Program represents a generic Monkey program, root node of every AST
type Program struct {
	Statements []Statement
}

// TokenLiteral as part of the Program struct makes it of type Node as well
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

// LetStatement has 3 fields, name of var, expression of binding, node token
type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

// TokenLiteral satisfies the book-keeping requirements of the Node interface
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}
func (ls *LetStatement) statementNode() {}

// Identifier represents the identifier of variable binding
type Identifier struct {
	Token token.Token
	Value string
}

// TokenLiteral satisfies the book-keeping requirements of the Node interface
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
func (i *Identifier) expressionNode() {}
