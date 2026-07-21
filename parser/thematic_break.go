package parser

import (
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/text"
)

type thematicBreakPraser struct {
}

var defaultThematicBreakPraser = &thematicBreakPraser{}

func NewThematicBreakParser() BlockParser { _ = "STUB: not implemented"; return *new(BlockParser) }

func isThematicBreak(line []byte, offset int) bool { _ = "STUB: not implemented"; return false }

func (b *thematicBreakPraser) Trigger() []byte { _ = "STUB: not implemented"; return nil }

func (b *thematicBreakPraser) Open(parent ast.Node, reader text.Reader, pc Context) (ast.Node, State) {
	_ = "STUB: not implemented"
	return *new(ast.Node), *new(State)
}

func (b *thematicBreakPraser) Continue(node ast.Node, reader text.Reader, pc Context) State {
	_ = "STUB: not implemented"
	return *new(State)
}

func (b *thematicBreakPraser) Close(node ast.Node, reader text.Reader, pc Context) {
	_ = "STUB: not implemented"
	return
}

func (b *thematicBreakPraser) CanInterruptParagraph() bool { _ = "STUB: not implemented"; return false }

func (b *thematicBreakPraser) CanAcceptIndentedLine() bool { _ = "STUB: not implemented"; return false }
