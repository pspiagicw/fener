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
func (i *Integer) String() string  { return fmt.Sprintf("%d", i.Value) }

type String struct {
	Token *token.Token
	Value string
}

func (s *String) Name() string    { return "String" }
func (s *String) expressionNode() {}
func (s *String) String() string  { return fmt.Sprintf("\"%s\"", s.Value) }

type ExpressionStatement struct {
	Expression Expression
	Token      *token.Token
}

func (es *ExpressionStatement) Name() string   { return "ExpressionStatement" }
func (es *ExpressionStatement) statementNode() {}
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

type Boolean struct {
	Token *token.Token
	Value bool
}

func (b *Boolean) Name() string    { return "Boolean" }
func (b *Boolean) expressionNode() {}
func (b *Boolean) String() string  { return fmt.Sprintf("%t", b.Value) }

type InfixExpression struct {
	Left     Expression
	Operator token.TokenType
	Right    Expression
}

func (ie *InfixExpression) Name() string    { return "InfixExpression" }
func (ie *InfixExpression) expressionNode() {}
func (ie *InfixExpression) String() string {
	return fmt.Sprintf("(%s %s %s)", ie.Left.String(), ie.Operator, ie.Right.String())
}

type Identifier struct {
	Token *token.Token
	Value string
}

func (i *Identifier) Name() string    { return "Identifier" }
func (i *Identifier) expressionNode() {}
func (i *Identifier) String() string  { return fmt.Sprintf("%s", i.Value) }

type AssignmentExpression struct {
	Token  *token.Token
	Value  Expression
	Target *Identifier
}

func (ae *AssignmentExpression) Name() string    { return "AssignmentExpression" }
func (ae *AssignmentExpression) expressionNode() {}
func (ae *AssignmentExpression) String() string {
	return fmt.Sprintf("%s = %s", ae.Target, ae.Value.String())
}

type IfExpression struct {
	Token       *token.Token
	Condition   Expression
	Consequence *BlockStatement
	Elif        map[Expression]*BlockStatement
	Alternative *BlockStatement
}

func (ie *IfExpression) Name() string    { return "IfExpression" }
func (ie *IfExpression) expressionNode() {}
func (ie *IfExpression) String() string {
	var out strings.Builder
	out.WriteString("if ")
	out.WriteString(ie.Condition.String())
	out.WriteString(" then\n")
	out.WriteString(ie.Consequence.String())
	for k, v := range ie.Elif {
		out.WriteString("elif ")
		out.WriteString(k.String())
		out.WriteString(" then\n")
		out.WriteString(v.String())
	}
	if ie.Alternative != nil {
		out.WriteString("else\n")
		out.WriteString(ie.Alternative.String())
	}
	out.WriteString("end")
	return out.String()
}

type BlockStatement struct {
	Token      *token.Token
	Statements []Statement
}

func (bs *BlockStatement) Name() string   { return "BlockStatement" }
func (bs *BlockStatement) statementNode() {}
func (bs *BlockStatement) String() string {
	var out strings.Builder
	for _, s := range bs.Statements {
		out.WriteString(s.String())
		out.WriteString("\n")
	}
	return out.String()
}

type CallExpression struct {
	Token     *token.Token
	Function  Expression
	Arguments []Expression
}

func (ce *CallExpression) Name() string    { return "CallExpression" }
func (ce *CallExpression) expressionNode() {}
func (ce *CallExpression) String() string {
	var out strings.Builder
	out.WriteString(ce.Function.String())
	out.WriteString("(")
	for i, arg := range ce.Arguments {
		out.WriteString(arg.String())
		if i != len(ce.Arguments)-1 {
			out.WriteString(", ")
		}
	}
	out.WriteString(")")
	return out.String()
}

type WhileStatement struct {
	Token       *token.Token
	Condition   Expression
	Consequence *BlockStatement
}

func (ws *WhileStatement) Name() string   { return "WhileStatement" }
func (ws *WhileStatement) statementNode() {}
func (ws *WhileStatement) String() string {
	var out strings.Builder
	out.WriteString("while ")
	out.WriteString(ws.Condition.String())
	out.WriteString(" then\n")
	out.WriteString(ws.Consequence.String())
	out.WriteString("end")
	return out.String()
}

type Lambda struct {
	Token     *token.Token
	Arguments []*Identifier
	Body      *BlockStatement
}

func (l *Lambda) Name() string    { return "Lambda" }
func (l *Lambda) expressionNode() {}
func (l *Lambda) String() string {
	var out strings.Builder
	out.WriteString("fn(")
	for i, arg := range l.Arguments {
		out.WriteString(arg.String())
		if i != len(l.Arguments)-1 {
			out.WriteString(", ")
		}
	}
	out.WriteString(")\n")
	out.WriteString(l.Body.String())
	out.WriteString("end")
	return out.String()
}

type FunctionStatement struct {
	Token     *token.Token
	Target    *Identifier
	Arguments []*Identifier
	Body      *BlockStatement
}

func (fs *FunctionStatement) Name() string   { return "FunctionStatement" }
func (fs *FunctionStatement) statementNode() {}
func (fs *FunctionStatement) String() string {
	var out strings.Builder
	out.WriteString("fn ")
	out.WriteString(fs.Target.String())
	out.WriteString("(")
	for i, arg := range fs.Arguments {
		out.WriteString(arg.String())
		if i != len(fs.Arguments)-1 {
			out.WriteString(", ")
		}
	}
	out.WriteString(")\n")
	out.WriteString(fs.Body.String())
	out.WriteString("end")
	return out.String()
}

type TestStatement struct {
	Token      *token.Token
	Target     *String
	Statements *BlockStatement
}

func (ts *TestStatement) Name() string   { return "TestStatement" }
func (ts *TestStatement) statementNode() {}
func (ts *TestStatement) String() string {
	var out strings.Builder
	out.WriteString("test ")
	out.WriteString(ts.Target.String())
	out.WriteString("\n")
	out.WriteString(ts.Statements.String())
	out.WriteString("end\n")
	return out.String()
}

type IndexExpression struct {
	Token *token.Token
	Left  Expression
	Index Expression
}

func (ie *IndexExpression) Name() string    { return "IndexExpression" }
func (ie *IndexExpression) expressionNode() {}
func (ie *IndexExpression) String() string {
	return fmt.Sprintf("(%s[%s])", ie.Left.String(), ie.Index.String())
}

type Array struct {
	Token    *token.Token
	Elements []Expression
}

func (a *Array) Name() string    { return "Array" }
func (a *Array) expressionNode() {}
func (a *Array) String() string {
	var out strings.Builder
	out.WriteString("[")
	for i, elem := range a.Elements {
		out.WriteString(elem.String())
		if i != len(a.Elements)-1 {
			out.WriteString(", ")
		}
	}
	out.WriteString("]")
	return out.String()
}

type PrefixExpression struct {
	Token    *token.Token
	Operator token.TokenType
	Right    Expression
}

func (pe *PrefixExpression) Name() string    { return "PrefixExpression" }
func (pe *PrefixExpression) expressionNode() {}
func (pe *PrefixExpression) String() string {
	return fmt.Sprintf("(%s%s)", pe.Operator, pe.Right.String())
}

type ClassStatement struct {
	Token   *token.Token
	Target  *Identifier
	Methods []*FunctionStatement
}

func (cs *ClassStatement) Name() string   { return "ClassStatement" }
func (cs *ClassStatement) statementNode() {}
func (cs *ClassStatement) String() string {
	var out strings.Builder
	out.WriteString("class ")
	out.WriteString(cs.Target.Value)
	out.WriteString("\n")
	for _, method := range cs.Methods {
		out.WriteString(method.String())
		out.WriteString("\n")
	}
	out.WriteString("end\n")
	return out.String()
}

type FieldExpression struct {
	Token  *token.Token
	Target Expression
	Field  *Identifier
	Method *CallExpression
}

func (fe *FieldExpression) Name() string    { return "FieldExpression" }
func (fe *FieldExpression) expressionNode() {}
func (fe *FieldExpression) String() string {
	return fmt.Sprintf("%s.%s", fe.Target.String(), fe.Field.String())
}
