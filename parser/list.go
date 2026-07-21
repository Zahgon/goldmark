package parser

import (
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/text"
)

type listItemType int

const (
	notList listItemType = iota
	bulletList
	orderedList
)

var skipListParserKey = NewContextKey()
var emptyListItemWithBlankLines = NewContextKey()
var listItemFlagValue any = true

func parseListItem(line []byte) ([6]int, listItemType) {
	_ = "STUB: not implemented"
	return [6]int{}, *new(listItemType)
}

func calcListOffset(source []byte, match [6]int) int { _ = "STUB: not implemented"; return 0 }

func lastOffset(node ast.Node) int { _ = "STUB: not implemented"; return 0 }

type listParser struct {
}

var defaultListParser = &listParser{}

func NewListParser() BlockParser { _ = "STUB: not implemented"; return *new(BlockParser) }

func (b *listParser) Trigger() []byte { _ = "STUB: not implemented"; return nil }

func (b *listParser) Open(parent ast.Node, reader text.Reader, pc Context) (ast.Node, State) {
	_ = "STUB: not implemented"
	return *new(ast.Node), *new(State)
}

func (b *listParser) Continue(node ast.Node, reader text.Reader, pc Context) State {
	_ = "STUB: not implemented"
	return *new(State)
}

func (b *listParser) Close(node ast.Node, reader text.Reader, pc Context) {
	_ = "STUB: not implemented"
	return
}

func (b *listParser) CanInterruptParagraph() bool { _ = "STUB: not implemented"; return false }

func (b *listParser) CanAcceptIndentedLine() bool { _ = "STUB: not implemented"; return false }
