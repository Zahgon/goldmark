package ast

import (
	textm "github.com/yuin/goldmark/text"
)

type BaseInline struct {
	BaseNode
}

func (b *BaseInline) Type() NodeType { _ = "STUB: not implemented"; return *new(NodeType) }

func (b *BaseInline) IsRaw() bool { _ = "STUB: not implemented"; return false }

func (b *BaseInline) HasBlankPreviousLines() bool { _ = "STUB: not implemented"; return false }

func (b *BaseInline) SetBlankPreviousLines(v bool) { _ = "STUB: not implemented"; return }

func (b *BaseInline) Lines() *textm.Segments { _ = "STUB: not implemented"; return nil }

func (b *BaseInline) SetLines(v *textm.Segments) { _ = "STUB: not implemented"; return }

type Text struct {
	BaseInline

	Segment textm.Segment

	flags uint8
}

const (
	textSoftLineBreak = 1 << iota
	textHardLineBreak
	textRaw
	textCode
)

func textFlagsString(flags uint8) string { _ = "STUB: not implemented"; return "" }

func (n *Text) Inline() { _ = "STUB: not implemented"; return }

func (n *Text) Pos() int { _ = "STUB: not implemented"; return 0 }

func (n *Text) SoftLineBreak() bool { _ = "STUB: not implemented"; return false }

func (n *Text) SetSoftLineBreak(v bool) { _ = "STUB: not implemented"; return }

func (n *Text) IsRaw() bool { _ = "STUB: not implemented"; return false }

func (n *Text) SetRaw(v bool) { _ = "STUB: not implemented"; return }

func (n *Text) HardLineBreak() bool { _ = "STUB: not implemented"; return false }

func (n *Text) SetHardLineBreak(v bool) { _ = "STUB: not implemented"; return }

func (n *Text) Merge(node Node, source []byte) bool { _ = "STUB: not implemented"; return false }

func (n *Text) Text(source []byte) []byte { _ = "STUB: not implemented"; return nil }

func (n *Text) Value(source []byte) []byte { _ = "STUB: not implemented"; return nil }

func (n *Text) Dump(source []byte, level int) { _ = "STUB: not implemented"; return }

var KindText = NewNodeKind("Text")

func (n *Text) Kind() NodeKind { _ = "STUB: not implemented"; return *new(NodeKind) }

func NewText() *Text { _ = "STUB: not implemented"; return nil }

func NewTextSegment(v textm.Segment) *Text { _ = "STUB: not implemented"; return nil }

func NewRawTextSegment(v textm.Segment) *Text { _ = "STUB: not implemented"; return nil }

func MergeOrAppendTextSegment(parent Node, s textm.Segment) { _ = "STUB: not implemented"; return }

func MergeOrReplaceTextSegment(parent Node, n Node, s textm.Segment) {
	_ = "STUB: not implemented"
	return
}

type String struct {
	BaseInline

	Value []byte
	flags uint8
}

func (n *String) Inline() { _ = "STUB: not implemented"; return }

func (n *String) Pos() int { _ = "STUB: not implemented"; return 0 }

func (n *String) IsRaw() bool { _ = "STUB: not implemented"; return false }

func (n *String) SetRaw(v bool) { _ = "STUB: not implemented"; return }

func (n *String) IsCode() bool { _ = "STUB: not implemented"; return false }

func (n *String) SetCode(v bool) { _ = "STUB: not implemented"; return }

func (n *String) Text(source []byte) []byte { _ = "STUB: not implemented"; return nil }

func (n *String) Dump(source []byte, level int) { _ = "STUB: not implemented"; return }

var KindString = NewNodeKind("String")

func (n *String) Kind() NodeKind { _ = "STUB: not implemented"; return *new(NodeKind) }

func NewString(v []byte) *String { _ = "STUB: not implemented"; return nil }

type CodeSpan struct {
	BaseInline
}

func (n *CodeSpan) Inline() { _ = "STUB: not implemented"; return }

func (n *CodeSpan) IsBlank(source []byte) bool { _ = "STUB: not implemented"; return false }

func (n *CodeSpan) Dump(source []byte, level int) { _ = "STUB: not implemented"; return }

var KindCodeSpan = NewNodeKind("CodeSpan")

func (n *CodeSpan) Kind() NodeKind { _ = "STUB: not implemented"; return *new(NodeKind) }

func NewCodeSpan() *CodeSpan { _ = "STUB: not implemented"; return nil }

type Emphasis struct {
	BaseInline

	Level int
}

func (n *Emphasis) Dump(source []byte, level int) { _ = "STUB: not implemented"; return }

var KindEmphasis = NewNodeKind("Emphasis")

func (n *Emphasis) Kind() NodeKind { _ = "STUB: not implemented"; return *new(NodeKind) }

func NewEmphasis(level int) *Emphasis { _ = "STUB: not implemented"; return nil }

type baseLink struct {
	BaseInline

	Destination []byte

	Title []byte

	Reference *ReferenceLink
}

func (n *baseLink) Inline() { _ = "STUB: not implemented"; return }

type ReferenceLinkType int

const (
	ReferenceLinkFull ReferenceLinkType = iota + 1

	ReferenceLinkCollapsed

	ReferenceLinkShortcut
)

func (t ReferenceLinkType) String() string { _ = "STUB: not implemented"; return "" }

type ReferenceLink struct {
	Type ReferenceLinkType

	Value []byte
}

func NewReferenceLink(typ ReferenceLinkType, value []byte) *ReferenceLink {
	_ = "STUB: not implemented"
	return nil
}

type Link struct {
	baseLink
}

func (n *Link) Dump(source []byte, level int) { _ = "STUB: not implemented"; return }

var KindLink = NewNodeKind("Link")

func (n *Link) Kind() NodeKind { _ = "STUB: not implemented"; return *new(NodeKind) }

func NewLink() *Link { _ = "STUB: not implemented"; return nil }

type Image struct {
	baseLink
}

func (n *Image) Dump(source []byte, level int) { _ = "STUB: not implemented"; return }

var KindImage = NewNodeKind("Image")

func (n *Image) Kind() NodeKind { _ = "STUB: not implemented"; return *new(NodeKind) }

func NewImage(link *Link) *Image { _ = "STUB: not implemented"; return nil }

type AutoLinkType int

const (
	AutoLinkEmail AutoLinkType = iota + 1

	AutoLinkURL
)

type AutoLink struct {
	BaseInline

	AutoLinkType AutoLinkType

	Protocol []byte

	value *Text
}

func (n *AutoLink) Inline() { _ = "STUB: not implemented"; return }

func (n *AutoLink) Dump(source []byte, level int) { _ = "STUB: not implemented"; return }

var KindAutoLink = NewNodeKind("AutoLink")

func (n *AutoLink) Kind() NodeKind { _ = "STUB: not implemented"; return *new(NodeKind) }

func (n *AutoLink) URL(source []byte) []byte { _ = "STUB: not implemented"; return nil }

func (n *AutoLink) Label(source []byte) []byte { _ = "STUB: not implemented"; return nil }

func (n *AutoLink) Text(source []byte) []byte { _ = "STUB: not implemented"; return nil }

func NewAutoLink(typ AutoLinkType, value *Text) *AutoLink { _ = "STUB: not implemented"; return nil }

type RawHTML struct {
	BaseInline
	Segments *textm.Segments
}

func (n *RawHTML) Inline() { _ = "STUB: not implemented"; return }

func (n *RawHTML) Dump(source []byte, level int) { _ = "STUB: not implemented"; return }

var KindRawHTML = NewNodeKind("RawHTML")

func (n *RawHTML) Kind() NodeKind { _ = "STUB: not implemented"; return *new(NodeKind) }

func (n *RawHTML) Text(source []byte) []byte { _ = "STUB: not implemented"; return nil }

func NewRawHTML() *RawHTML { _ = "STUB: not implemented"; return nil }
