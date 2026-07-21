package parser

import (
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/text"
)

type linkReferenceParagraphTransformer struct {
}

var LinkReferenceParagraphTransformer = &linkReferenceParagraphTransformer{}

func (p *linkReferenceParagraphTransformer) Transform(node *ast.Paragraph, reader text.Reader, pc Context) {
	_ = "STUB: not implemented"
	return
}

func parseLinkReferenceDefinition(block text.Reader, pc Context) (ast.Node, int, int) {
	_ = "STUB: not implemented"
	return *new(ast.Node), 0, 0
}
