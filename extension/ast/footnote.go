package ast

import (
	gast "github.com/yuin/goldmark/ast"
)

type FootnoteLink struct {
	gast.BaseInline
	Index    int
	RefCount int
	RefIndex int
}

func (n *FootnoteLink) Dump(source []byte, level int) { _ = "STUB: not implemented"; return }

var KindFootnoteLink = gast.NewNodeKind("FootnoteLink")

func (n *FootnoteLink) Kind() gast.NodeKind { _ = "STUB: not implemented"; return *new(gast.NodeKind) }

func NewFootnoteLink(index int) *FootnoteLink { _ = "STUB: not implemented"; return nil }

type FootnoteBacklink struct {
	gast.BaseInline
	Index    int
	RefCount int
	RefIndex int
}

func (n *FootnoteBacklink) Dump(source []byte, level int) { _ = "STUB: not implemented"; return }

var KindFootnoteBacklink = gast.NewNodeKind("FootnoteBacklink")

func (n *FootnoteBacklink) Kind() gast.NodeKind {
	_ = "STUB: not implemented"
	return *new(gast.NodeKind)
}

func NewFootnoteBacklink(index int) *FootnoteBacklink { _ = "STUB: not implemented"; return nil }

type Footnote struct {
	gast.BaseBlock
	Ref   []byte
	Index int
}

func (n *Footnote) Dump(source []byte, level int) { _ = "STUB: not implemented"; return }

var KindFootnote = gast.NewNodeKind("Footnote")

func (n *Footnote) Kind() gast.NodeKind { _ = "STUB: not implemented"; return *new(gast.NodeKind) }

func NewFootnote(ref []byte) *Footnote { _ = "STUB: not implemented"; return nil }

type FootnoteList struct {
	gast.BaseBlock
	Count int
}

func (n *FootnoteList) Dump(source []byte, level int) { _ = "STUB: not implemented"; return }

var KindFootnoteList = gast.NewNodeKind("FootnoteList")

func (n *FootnoteList) Kind() gast.NodeKind { _ = "STUB: not implemented"; return *new(gast.NodeKind) }

func NewFootnoteList() *FootnoteList { _ = "STUB: not implemented"; return nil }
