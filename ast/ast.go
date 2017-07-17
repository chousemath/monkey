package ast

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
