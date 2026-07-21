package parser

import (
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/text"
)

type emphasisDelimiterProcessor struct {
}

func (p *emphasisDelimiterProcessor) IsDelimiter(b byte) bool {
	_ = "STUB: not implemented"
	return false
}

func (p *emphasisDelimiterProcessor) CanOpenCloser(opener, closer *Delimiter) bool {
	_ = "STUB: not implemented"
	return false
}

func (p *emphasisDelimiterProcessor) OnMatch(consumes int) ast.Node {
	_ = "STUB: not implemented"
	return *new(ast.Node)
}

var defaultEmphasisDelimiterProcessor = &emphasisDelimiterProcessor{}

type emphasisParser struct {
}

var defaultEmphasisParser = &emphasisParser{}

func NewEmphasisParser() InlineParser { _ = "STUB: not implemented"; return *new(InlineParser) }

func (s *emphasisParser) Trigger() []byte { _ = "STUB: not implemented"; return nil }

func (s *emphasisParser) Parse(parent ast.Node, block text.Reader, pc Context) ast.Node {
	_ = "STUB: not implemented"
	return *new(ast.Node)
}
