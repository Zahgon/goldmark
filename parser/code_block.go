package parser

import (
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/text"
)

type codeBlockParser struct {
}

var defaultCodeBlockParser = &codeBlockParser{}

func NewCodeBlockParser() BlockParser { _ = "STUB: not implemented"; return *new(BlockParser) }

func (b *codeBlockParser) Trigger() []byte { _ = "STUB: not implemented"; return nil }

func (b *codeBlockParser) Open(parent ast.Node, reader text.Reader, pc Context) (ast.Node, State) {
	_ = "STUB: not implemented"
	return *new(ast.Node), *new(State)
}

func (b *codeBlockParser) Continue(node ast.Node, reader text.Reader, pc Context) State {
	_ = "STUB: not implemented"
	return *new(State)
}

func (b *codeBlockParser) Close(node ast.Node, reader text.Reader, pc Context) {
	_ = "STUB: not implemented"
	return
}

func (b *codeBlockParser) CanInterruptParagraph() bool { _ = "STUB: not implemented"; return false }

func (b *codeBlockParser) CanAcceptIndentedLine() bool { _ = "STUB: not implemented"; return false }

func preserveLeadingTabInCodeBlock(segment *text.Segment, reader text.Reader, indent int) {
	_ = "STUB: not implemented"
	return
}
