package parser

import (
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/text"
)

var temporaryParagraphKey = NewContextKey()

type setextHeadingParser struct {
	HeadingConfig
}

func matchesSetextHeadingBar(line []byte) (byte, bool) { _ = "STUB: not implemented"; return 0, false }

func NewSetextHeadingParser(opts ...HeadingOption) BlockParser {
	_ = "STUB: not implemented"
	return *new(BlockParser)
}

func (b *setextHeadingParser) Trigger() []byte { _ = "STUB: not implemented"; return nil }

func (b *setextHeadingParser) Open(parent ast.Node, reader text.Reader, pc Context) (ast.Node, State) {
	_ = "STUB: not implemented"
	return *new(ast.Node), *new(State)
}

func (b *setextHeadingParser) Continue(node ast.Node, reader text.Reader, pc Context) State {
	_ = "STUB: not implemented"
	return *new(State)
}

func (b *setextHeadingParser) Close(node ast.Node, reader text.Reader, pc Context) {
	_ = "STUB: not implemented"
	return
}

func (b *setextHeadingParser) CanInterruptParagraph() bool { _ = "STUB: not implemented"; return false }

func (b *setextHeadingParser) CanAcceptIndentedLine() bool { _ = "STUB: not implemented"; return false }
