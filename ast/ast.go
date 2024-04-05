package ast

import (
	"fmt"
	"github.com/pspiagicw/fener/token"
	"strings"
)

type Node interface {
	Name() string
	String() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) Name() string { return "Program" }
func (p *Program) String() string {
	var out strings.Builder
	for _, s := range p.Statements {
		out.WriteString(s.String())
		out.WriteString("\n")
	}
	return out.String()
}

type ReturnStatement struct {
	Value Expression
	Token *token.Token
}

func (es *ReturnStatement) Name() string   { return "ReturnStatement" }
func (es *ReturnStatement) statementNode() {}
func (es *ReturnStatement) String() string {
	if es.Value != nil {
		return es.Value.String()
	}
	return ""
}

type Integer struct {
	Token *token.Token
	Value int64
}

func (i *Integer) Name() string    { return "Integer" }
func (i *Integer) expressionNode() {}
func (i *Integer) String() string  { return fmt.Sprintf("Integer(%d)", i.Value) }

type String struct {
	Token *token.Token
	Value string
}

func (s *String) Name() string    { return "String" }
func (s *String) expressionNode() {}
func (s *String) String() string  { return fmt.Sprintf("String(%s)", s.Value) }
