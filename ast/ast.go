package ast

import (
	textm "github.com/yuin/goldmark/text"
)

type NodeType int

const (
	TypeBlock NodeType = iota + 1

	TypeInline

	TypeDocument
)

type NodeKind int

func (k NodeKind) String() string { _ = "STUB: not implemented"; return "" }

var kindMax NodeKind
var kindNames = []string{""}

func NewNodeKind(name string) NodeKind { _ = "STUB: not implemented"; return *new(NodeKind) }

type Attribute struct {
	Name  []byte
	Value any
}

type Node interface {
	Type() NodeType

	Kind() NodeKind

	Pos() int

	SetPos(v int)

	NextSibling() Node

	PreviousSibling() Node

	Parent() Node

	SetParent(Node)

	SetPreviousSibling(Node)

	SetNextSibling(Node)

	HasChildren() bool

	ChildCount() int

	FirstChild() Node

	LastChild() Node

	AppendChild(self, child Node)

	RemoveChild(self, child Node)

	RemoveChildren(self Node)

	SortChildren(comparator func(n1, n2 Node) int)

	ReplaceChild(self, v1, insertee Node)

	InsertBefore(self, v1, insertee Node)

	InsertAfter(self, v1, insertee Node)

	OwnerDocument() *Document

	Dump(source []byte, level int)

	Text(source []byte) []byte

	HasBlankPreviousLines() bool

	SetBlankPreviousLines(v bool)

	Lines() *textm.Segments

	SetLines(*textm.Segments)

	IsRaw() bool

	SetAttribute(name []byte, value any)

	SetAttributeString(name string, value any)

	Attribute(name []byte) (any, bool)

	AttributeString(name string) (any, bool)

	Attributes() []Attribute

	RemoveAttributes()
}

type pos struct {
	has   bool
	value int
}

func (p *pos) Pos() int { _ = "STUB: not implemented"; return 0 }

func (p *pos) SetPos(v int) { _ = "STUB: not implemented"; return }

type BaseNode struct {
	firstChild Node
	lastChild  Node
	parent     Node
	next       Node
	prev       Node
	childCount int
	attributes []Attribute
	pos        pos
}

func ensureIsolated(v Node) { _ = "STUB: not implemented"; return }

func (n *BaseNode) Pos() int { _ = "STUB: not implemented"; return 0 }

func (n *BaseNode) SetPos(v int) { _ = "STUB: not implemented"; return }

func (n *BaseNode) HasChildren() bool { _ = "STUB: not implemented"; return false }

func (n *BaseNode) SetPreviousSibling(v Node) { _ = "STUB: not implemented"; return }

func (n *BaseNode) SetNextSibling(v Node) { _ = "STUB: not implemented"; return }

func (n *BaseNode) PreviousSibling() Node { _ = "STUB: not implemented"; return *new(Node) }

func (n *BaseNode) NextSibling() Node { _ = "STUB: not implemented"; return *new(Node) }

func (n *BaseNode) RemoveChild(self, v Node) { _ = "STUB: not implemented"; return }

func (n *BaseNode) RemoveChildren(self Node) { _ = "STUB: not implemented"; return }

func (n *BaseNode) SortChildren(comparator func(n1, n2 Node) int) {
	_ = "STUB: not implemented"
	return
}

func (n *BaseNode) FirstChild() Node { _ = "STUB: not implemented"; return *new(Node) }

func (n *BaseNode) LastChild() Node { _ = "STUB: not implemented"; return *new(Node) }

func (n *BaseNode) ChildCount() int { _ = "STUB: not implemented"; return 0 }

func (n *BaseNode) Parent() Node { _ = "STUB: not implemented"; return *new(Node) }

func (n *BaseNode) SetParent(v Node) { _ = "STUB: not implemented"; return }

func (n *BaseNode) AppendChild(self, v Node) { _ = "STUB: not implemented"; return }

func (n *BaseNode) ReplaceChild(self, v1, insertee Node) { _ = "STUB: not implemented"; return }

func (n *BaseNode) InsertAfter(self, v1, insertee Node) { _ = "STUB: not implemented"; return }

func (n *BaseNode) InsertBefore(self, v1, insertee Node) { _ = "STUB: not implemented"; return }

func (n *BaseNode) OwnerDocument() *Document { _ = "STUB: not implemented"; return nil }

func (n *BaseNode) Text(source []byte) []byte { _ = "STUB: not implemented"; return nil }

func (n *BaseNode) SetAttribute(name []byte, value any) { _ = "STUB: not implemented"; return }

func (n *BaseNode) SetAttributeString(name string, value any) { _ = "STUB: not implemented"; return }

func (n *BaseNode) Attribute(name []byte) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

func (n *BaseNode) AttributeString(s string) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

func (n *BaseNode) Attributes() []Attribute { _ = "STUB: not implemented"; return nil }

func (n *BaseNode) RemoveAttributes() { _ = "STUB: not implemented"; return }

func DumpHelper(v Node, source []byte, level int, kv map[string]string, cb func(int)) {
	_ = "STUB: not implemented"
	return
}

type WalkStatus int

const (
	WalkStop WalkStatus = iota + 1

	WalkSkipChildren

	WalkContinue
)

type Walker func(n Node, entering bool) (WalkStatus, error)

func Walk(n Node, walker Walker) error { _ = "STUB: not implemented"; return nil }

func walkHelper(n Node, walker Walker) (WalkStatus, error) {
	_ = "STUB: not implemented"
	return *new(WalkStatus), nil
}
