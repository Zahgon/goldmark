package parser

import (
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/text"
)

type codeSpanParser struct {
}

var defaultCodeSpanParser = &codeSpanParser{}

func NewCodeSpanParser() InlineParser { _ = "STUB: not implemented"; return *new(InlineParser) }

func (s *codeSpanParser) Trigger() []byte { _ = "STUB: not implemented"; return nil }

func (s *codeSpanParser) Parse(parent ast.Node, block text.Reader, pc Context) ast.Node {
	_ = "STUB: not implemented"
	return *new(ast.Node)
}

func isSpaceOrNewline(c byte) bool { _ = "STUB: not implemented"; return false }
