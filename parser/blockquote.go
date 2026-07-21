package parser

import (
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/text"
)

type blockquoteParser struct {
}

var defaultBlockquoteParser = &blockquoteParser{}

func NewBlockquoteParser() BlockParser { _ = "STUB: not implemented"; return *new(BlockParser) }

func (b *blockquoteParser) process(reader text.Reader) bool {
	_ = "STUB: not implemented"
	return false
}

func (b *blockquoteParser) Trigger() []byte { _ = "STUB: not implemented"; return nil }

func (b *blockquoteParser) Open(parent ast.Node, reader text.Reader, pc Context) (ast.Node, State) {
	_ = "STUB: not implemented"
	return *new(ast.Node), *new(State)
}

func (b *blockquoteParser) Continue(node ast.Node, reader text.Reader, pc Context) State {
	_ = "STUB: not implemented"
	return *new(State)
}

func (b *blockquoteParser) Close(node ast.Node, reader text.Reader, pc Context) {
	_ = "STUB: not implemented"
	return
}

func (b *blockquoteParser) CanInterruptParagraph() bool { _ = "STUB: not implemented"; return false }

func (b *blockquoteParser) CanAcceptIndentedLine() bool { _ = "STUB: not implemented"; return false }
