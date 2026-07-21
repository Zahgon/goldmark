package ast

import (
	gast "github.com/yuin/goldmark/ast"
)

type Strikethrough struct {
	gast.BaseInline
}

func (n *Strikethrough) Dump(source []byte, level int) { _ = "STUB: not implemented"; return }

var KindStrikethrough = gast.NewNodeKind("Strikethrough")

func (n *Strikethrough) Kind() gast.NodeKind { _ = "STUB: not implemented"; return *new(gast.NodeKind) }

func NewStrikethrough() *Strikethrough { _ = "STUB: not implemented"; return nil }
