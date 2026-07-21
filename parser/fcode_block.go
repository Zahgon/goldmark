package parser

import (
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/text"
)

type fencedCodeBlockParser struct {
}

var defaultFencedCodeBlockParser = &fencedCodeBlockParser{}

func NewFencedCodeBlockParser() BlockParser { _ = "STUB: not implemented"; return *new(BlockParser) }

type fenceData struct {
	char   byte
	indent int
	length int
	node   ast.Node
}

var fencedCodeBlockInfoKey = NewContextKey()

func (b *fencedCodeBlockParser) Trigger() []byte { _ = "STUB: not implemented"; return nil }

func (b *fencedCodeBlockParser) Open(parent ast.Node, reader text.Reader, pc Context) (ast.Node, State) {
	_ = "STUB: not implemented"
	return *new(ast.Node), *new(State)
}

func (b *fencedCodeBlockParser) Continue(node ast.Node, reader text.Reader, pc Context) State {
	_ = "STUB: not implemented"
	return *new(State)
}

func (b *fencedCodeBlockParser) Close(node ast.Node, reader text.Reader, pc Context) {
	_ = "STUB: not implemented"
	return
}

func (b *fencedCodeBlockParser) CanInterruptParagraph() bool {
	_ = "STUB: not implemented"
	return false
}

func (b *fencedCodeBlockParser) CanAcceptIndentedLine() bool {
	_ = "STUB: not implemented"
	return false
}
