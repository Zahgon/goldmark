package parser

import (
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/text"
)

type paragraphParser struct {
}

var defaultParagraphParser = &paragraphParser{}

func NewParagraphParser() BlockParser { _ = "STUB: not implemented"; return *new(BlockParser) }

func (b *paragraphParser) Trigger() []byte { _ = "STUB: not implemented"; return nil }

func (b *paragraphParser) Open(parent ast.Node, reader text.Reader, pc Context) (ast.Node, State) {
	_ = "STUB: not implemented"
	return *new(ast.Node), *new(State)
}

func (b *paragraphParser) Continue(node ast.Node, reader text.Reader, pc Context) State {
	_ = "STUB: not implemented"
	return *new(State)
}

func (b *paragraphParser) Close(node ast.Node, reader text.Reader, pc Context) {
	_ = "STUB: not implemented"
	return
}

func (b *paragraphParser) CanInterruptParagraph() bool { _ = "STUB: not implemented"; return false }

func (b *paragraphParser) CanAcceptIndentedLine() bool { _ = "STUB: not implemented"; return false }
