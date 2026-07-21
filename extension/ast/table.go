package ast

import (
	gast "github.com/yuin/goldmark/ast"
)

type Alignment int

const (
	AlignLeft Alignment = iota + 1

	AlignRight

	AlignCenter

	AlignNone
)

func (a Alignment) String() string { _ = "STUB: not implemented"; return "" }

type Table struct {
	gast.BaseBlock

	Alignments []Alignment
}

func (n *Table) Dump(source []byte, level int) { _ = "STUB: not implemented"; return }

var KindTable = gast.NewNodeKind("Table")

func (n *Table) Kind() gast.NodeKind { _ = "STUB: not implemented"; return *new(gast.NodeKind) }

func NewTable() *Table { _ = "STUB: not implemented"; return nil }

type TableRow struct {
	gast.BaseBlock
	Alignments []Alignment
}

func (n *TableRow) Dump(source []byte, level int) { _ = "STUB: not implemented"; return }

var KindTableRow = gast.NewNodeKind("TableRow")

func (n *TableRow) Kind() gast.NodeKind { _ = "STUB: not implemented"; return *new(gast.NodeKind) }

func NewTableRow(alignments []Alignment) *TableRow { _ = "STUB: not implemented"; return nil }

type TableHeader struct {
	gast.BaseBlock
	Alignments []Alignment
}

var KindTableHeader = gast.NewNodeKind("TableHeader")

func (n *TableHeader) Kind() gast.NodeKind { _ = "STUB: not implemented"; return *new(gast.NodeKind) }

func (n *TableHeader) Dump(source []byte, level int) { _ = "STUB: not implemented"; return }

func NewTableHeader(row *TableRow) *TableHeader { _ = "STUB: not implemented"; return nil }

type TableCell struct {
	gast.BaseBlock
	Alignment Alignment
}

func (n *TableCell) Dump(source []byte, level int) { _ = "STUB: not implemented"; return }

var KindTableCell = gast.NewNodeKind("TableCell")

func (n *TableCell) Kind() gast.NodeKind { _ = "STUB: not implemented"; return *new(gast.NodeKind) }

func NewTableCell() *TableCell { _ = "STUB: not implemented"; return nil }
