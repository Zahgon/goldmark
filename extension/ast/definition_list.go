package ast

import (
	gast "github.com/yuin/goldmark/ast"
)

type DefinitionList struct {
	gast.BaseBlock
	Offset             int
	TemporaryParagraph *gast.Paragraph
}

func (n *DefinitionList) Dump(source []byte, level int) { _ = "STUB: not implemented"; return }

func (n *DefinitionList) Pos() int { _ = "STUB: not implemented"; return 0 }

var KindDefinitionList = gast.NewNodeKind("DefinitionList")

func (n *DefinitionList) Kind() gast.NodeKind {
	_ = "STUB: not implemented"
	return *new(gast.NodeKind)
}

func NewDefinitionList(offset int, para *gast.Paragraph) *DefinitionList {
	_ = "STUB: not implemented"
	return nil
}

type DefinitionTerm struct {
	gast.BaseBlock
}

func (n *DefinitionTerm) Dump(source []byte, level int) { _ = "STUB: not implemented"; return }

func (n *DefinitionTerm) Pos() int { _ = "STUB: not implemented"; return 0 }

var KindDefinitionTerm = gast.NewNodeKind("DefinitionTerm")

func (n *DefinitionTerm) Kind() gast.NodeKind {
	_ = "STUB: not implemented"
	return *new(gast.NodeKind)
}

func NewDefinitionTerm() *DefinitionTerm { _ = "STUB: not implemented"; return nil }

type DefinitionDescription struct {
	gast.BaseBlock
	IsTight bool
}

func (n *DefinitionDescription) Dump(source []byte, level int) { _ = "STUB: not implemented"; return }

var KindDefinitionDescription = gast.NewNodeKind("DefinitionDescription")

func (n *DefinitionDescription) Kind() gast.NodeKind {
	_ = "STUB: not implemented"
	return *new(gast.NodeKind)
}

func NewDefinitionDescription() *DefinitionDescription { _ = "STUB: not implemented"; return nil }
