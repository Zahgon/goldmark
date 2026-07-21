package ast

import (
	gast "github.com/yuin/goldmark/ast"
)

type TaskCheckBox struct {
	gast.BaseInline
	IsChecked bool
}

func (n *TaskCheckBox) Dump(source []byte, level int) { _ = "STUB: not implemented"; return }

var KindTaskCheckBox = gast.NewNodeKind("TaskCheckBox")

func (n *TaskCheckBox) Kind() gast.NodeKind { _ = "STUB: not implemented"; return *new(gast.NodeKind) }

func NewTaskCheckBox(checked bool) *TaskCheckBox { _ = "STUB: not implemented"; return nil }
