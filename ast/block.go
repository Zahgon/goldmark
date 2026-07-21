package ast

import (
	textm "github.com/yuin/goldmark/text"
)

type BaseBlock struct {
	BaseNode
	lines              textm.Segments
	blankPreviousLines bool
}

func (b *BaseBlock) Type() NodeType { _ = "STUB: not implemented"; return *new(NodeType) }

func (b *BaseBlock) IsRaw() bool { _ = "STUB: not implemented"; return false }

func (b *BaseBlock) HasBlankPreviousLines() bool { _ = "STUB: not implemented"; return false }

func (b *BaseBlock) SetBlankPreviousLines(v bool) { _ = "STUB: not implemented"; return }

func (b *BaseBlock) Lines() *textm.Segments { _ = "STUB: not implemented"; return nil }

func (b *BaseBlock) SetLines(v *textm.Segments) { _ = "STUB: not implemented"; return }

type Document struct {
	BaseBlock

	meta map[string]any
}

var KindDocument = NewNodeKind("Document")

func (n *Document) Dump(source []byte, level int) { _ = "STUB: not implemented"; return }

func (n *Document) Type() NodeType { _ = "STUB: not implemented"; return *new(NodeType) }

func (n *Document) Pos() int { _ = "STUB: not implemented"; return 0 }

func (n *Document) Kind() NodeKind { _ = "STUB: not implemented"; return *new(NodeKind) }

func (n *Document) OwnerDocument() *Document { _ = "STUB: not implemented"; return nil }

func (n *Document) Meta() map[string]any { _ = "STUB: not implemented"; return nil }

func (n *Document) SetMeta(meta map[string]any) { _ = "STUB: not implemented"; return }

func (n *Document) AddMeta(key string, value any) { _ = "STUB: not implemented"; return }

func NewDocument() *Document { _ = "STUB: not implemented"; return nil }

type TextBlock struct {
	BaseBlock
}

func (n *TextBlock) Dump(source []byte, level int) { _ = "STUB: not implemented"; return }

func (n *TextBlock) Pos() int { _ = "STUB: not implemented"; return 0 }

var KindTextBlock = NewNodeKind("TextBlock")

func (n *TextBlock) Kind() NodeKind { _ = "STUB: not implemented"; return *new(NodeKind) }

func (n *TextBlock) Text(source []byte) []byte { _ = "STUB: not implemented"; return nil }

func NewTextBlock() *TextBlock { _ = "STUB: not implemented"; return nil }

type Paragraph struct {
	BaseBlock
}

func (n *Paragraph) Dump(source []byte, level int) { _ = "STUB: not implemented"; return }

func (n *Paragraph) Pos() int { _ = "STUB: not implemented"; return 0 }

var KindParagraph = NewNodeKind("Paragraph")

func (n *Paragraph) Kind() NodeKind { _ = "STUB: not implemented"; return *new(NodeKind) }

func (n *Paragraph) Text(source []byte) []byte { _ = "STUB: not implemented"; return nil }

func NewParagraph() *Paragraph { _ = "STUB: not implemented"; return nil }

func IsParagraph(node Node) bool { _ = "STUB: not implemented"; return false }

type Heading struct {
	BaseBlock

	Level int
}

func (n *Heading) Dump(source []byte, level int) { _ = "STUB: not implemented"; return }

var KindHeading = NewNodeKind("Heading")

func (n *Heading) Kind() NodeKind { _ = "STUB: not implemented"; return *new(NodeKind) }

func NewHeading(level int) *Heading { _ = "STUB: not implemented"; return nil }

type ThematicBreak struct {
	BaseBlock
}

func (n *ThematicBreak) Dump(source []byte, level int) { _ = "STUB: not implemented"; return }

var KindThematicBreak = NewNodeKind("ThematicBreak")

func (n *ThematicBreak) Kind() NodeKind { _ = "STUB: not implemented"; return *new(NodeKind) }

func NewThematicBreak() *ThematicBreak { _ = "STUB: not implemented"; return nil }

type CodeBlock struct {
	BaseBlock
}

func (n *CodeBlock) IsRaw() bool { _ = "STUB: not implemented"; return false }

func (n *CodeBlock) Dump(source []byte, level int) { _ = "STUB: not implemented"; return }

var KindCodeBlock = NewNodeKind("CodeBlock")

func (n *CodeBlock) Kind() NodeKind { _ = "STUB: not implemented"; return *new(NodeKind) }

func (n *CodeBlock) Text(source []byte) []byte { _ = "STUB: not implemented"; return nil }

func NewCodeBlock() *CodeBlock { _ = "STUB: not implemented"; return nil }

type FencedCodeBlock struct {
	BaseBlock

	Info *Text

	language []byte
}

func (n *FencedCodeBlock) Language(source []byte) []byte { _ = "STUB: not implemented"; return nil }

func (n *FencedCodeBlock) IsRaw() bool { _ = "STUB: not implemented"; return false }

func (n *FencedCodeBlock) Dump(source []byte, level int) { _ = "STUB: not implemented"; return }

var KindFencedCodeBlock = NewNodeKind("FencedCodeBlock")

func (n *FencedCodeBlock) Kind() NodeKind { _ = "STUB: not implemented"; return *new(NodeKind) }

func (n *FencedCodeBlock) Text(source []byte) []byte { _ = "STUB: not implemented"; return nil }

func NewFencedCodeBlock(info *Text) *FencedCodeBlock { _ = "STUB: not implemented"; return nil }

type Blockquote struct {
	BaseBlock
}

func (n *Blockquote) Dump(source []byte, level int) { _ = "STUB: not implemented"; return }

var KindBlockquote = NewNodeKind("Blockquote")

func (n *Blockquote) Kind() NodeKind { _ = "STUB: not implemented"; return *new(NodeKind) }

func NewBlockquote() *Blockquote { _ = "STUB: not implemented"; return nil }

type List struct {
	BaseBlock

	Marker byte

	IsTight bool

	Start int
}

func (l *List) IsOrdered() bool { _ = "STUB: not implemented"; return false }

func (l *List) CanContinue(marker byte, isOrdered bool) bool {
	_ = "STUB: not implemented"
	return false
}

func (l *List) Dump(source []byte, level int) { _ = "STUB: not implemented"; return }

var KindList = NewNodeKind("List")

func (l *List) Kind() NodeKind { _ = "STUB: not implemented"; return *new(NodeKind) }

func NewList(marker byte) *List { _ = "STUB: not implemented"; return nil }

type ListItem struct {
	BaseBlock

	Offset int
}

func (n *ListItem) Dump(source []byte, level int) { _ = "STUB: not implemented"; return }

var KindListItem = NewNodeKind("ListItem")

func (n *ListItem) Kind() NodeKind { _ = "STUB: not implemented"; return *new(NodeKind) }

func NewListItem(offset int) *ListItem { _ = "STUB: not implemented"; return nil }

type HTMLBlockType int

const (
	HTMLBlockType1 HTMLBlockType = iota + 1

	HTMLBlockType2

	HTMLBlockType3

	HTMLBlockType4

	HTMLBlockType5

	HTMLBlockType6

	HTMLBlockType7
)

type HTMLBlock struct {
	BaseBlock

	HTMLBlockType HTMLBlockType

	ClosureLine textm.Segment
}

func (n *HTMLBlock) IsRaw() bool { _ = "STUB: not implemented"; return false }

func (n *HTMLBlock) HasClosure() bool { _ = "STUB: not implemented"; return false }

func (n *HTMLBlock) Dump(source []byte, level int) { _ = "STUB: not implemented"; return }

var KindHTMLBlock = NewNodeKind("HTMLBlock")

func (n *HTMLBlock) Kind() NodeKind { _ = "STUB: not implemented"; return *new(NodeKind) }

func (n *HTMLBlock) Text(source []byte) []byte { _ = "STUB: not implemented"; return nil }

func NewHTMLBlock(typ HTMLBlockType) *HTMLBlock { _ = "STUB: not implemented"; return nil }

type LinkReferenceDefinition struct {
	BaseBlock

	Label []byte

	Destination []byte

	Title []byte
}

func (l *LinkReferenceDefinition) IsRaw() bool { _ = "STUB: not implemented"; return false }

func (l *LinkReferenceDefinition) Pos() int { _ = "STUB: not implemented"; return 0 }

func (l *LinkReferenceDefinition) Dump(source []byte, level int) { _ = "STUB: not implemented"; return }

var KindLinkReferenceDefinition = NewNodeKind("LinkReferenceDefinition")

func (l *LinkReferenceDefinition) Kind() NodeKind { _ = "STUB: not implemented"; return *new(NodeKind) }

func NewLinkReferenceDefinition(label, destination, title []byte) *LinkReferenceDefinition {
	_ = "STUB: not implemented"
	return nil
}
