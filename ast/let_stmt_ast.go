package ast

import (
	"bytes"

	"github.com/mohamed-samir907/meow/token"
)

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

func (l *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(l.TokenLiteral() + " ")
	out.WriteString(l.Name.String())
	out.WriteString(" = ")

	if l.Value != nil {
		out.WriteString(l.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

type Idenetifier struct {
	Token token.Token
	Value string
}

func (i *Idenetifier) expressionNode() {}
func (i *Idenetifier) TokenLiteral() string {
	return i.Token.Literal
}

func (i *Idenetifier) String() string {
	return i.Value
}
