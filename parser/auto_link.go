package parser

import (
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/text"
)

type autoLinkParser struct {
}

var defaultAutoLinkParser = &autoLinkParser{}

func NewAutoLinkParser() InlineParser { _ = "STUB: not implemented"; return *new(InlineParser) }

func (s *autoLinkParser) Trigger() []byte { _ = "STUB: not implemented"; return nil }

func (s *autoLinkParser) Parse(parent ast.Node, block text.Reader, pc Context) ast.Node {
	_ = "STUB: not implemented"
	return *new(ast.Node)
}
