package parser

import (
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/text"
)

type listItemParser struct {
}

var defaultListItemParser = &listItemParser{}

func NewListItemParser() BlockParser { _ = "STUB: not implemented"; return *new(BlockParser) }

func (b *listItemParser) Trigger() []byte { _ = "STUB: not implemented"; return nil }

func (b *listItemParser) Open(parent ast.Node, reader text.Reader, pc Context) (ast.Node, State) {
	_ = "STUB: not implemented"
	return *new(ast.Node), *new(State)
}

func (b *listItemParser) Continue(node ast.Node, reader text.Reader, pc Context) State {
	_ = "STUB: not implemented"
	return *new(State)
}

func (b *listItemParser) Close(node ast.Node, reader text.Reader, pc Context) {
	_ = "STUB: not implemented"
	return
}

func (b *listItemParser) CanInterruptParagraph() bool { _ = "STUB: not implemented"; return false }

func (b *listItemParser) CanAcceptIndentedLine() bool { _ = "STUB: not implemented"; return false }
