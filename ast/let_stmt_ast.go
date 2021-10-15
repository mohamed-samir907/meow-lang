package ast

import "github.com/mohamed-samir907/meow/token"

// let statement
type LetStatement struct {
	Token token.Token
	Name  *Idenetifier
	Value Expression
}

func (l *LetStatement) statementNode() {}
func (l *LetStatement) TokenLiteral() string {
	return l.Token.Literal
}

type Idenetifier struct {
	Token token.Token
	Value string
}

func (i *Idenetifier) expressionNode() {}
func (i *Idenetifier) TokenLiteral() string {
	return i.Token.Literal
}
