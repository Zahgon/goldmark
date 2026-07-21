package extension

import (
	"github.com/yuin/goldmark"
	gast "github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

type strikethroughDelimiterProcessor struct {
}

func (p *strikethroughDelimiterProcessor) IsDelimiter(b byte) bool {
	_ = "STUB: not implemented"
	return false
}

func (p *strikethroughDelimiterProcessor) CanOpenCloser(opener, closer *parser.Delimiter) bool {
	_ = "STUB: not implemented"
	return false
}

func (p *strikethroughDelimiterProcessor) OnMatch(consumes int) gast.Node {
	_ = "STUB: not implemented"
	return *new(gast.Node)
}

var defaultStrikethroughDelimiterProcessor = &strikethroughDelimiterProcessor{}

type strikethroughParser struct {
}

var defaultStrikethroughParser = &strikethroughParser{}

func NewStrikethroughParser() parser.InlineParser {
	_ = "STUB: not implemented"
	return *new(parser.InlineParser)
}

func (s *strikethroughParser) Trigger() []byte { _ = "STUB: not implemented"; return nil }

func (s *strikethroughParser) Parse(parent gast.Node, block text.Reader, pc parser.Context) gast.Node {
	_ = "STUB: not implemented"
	return *new(gast.Node)
}

func (s *strikethroughParser) CloseBlock(parent gast.Node, pc parser.Context) {
	_ = "STUB: not implemented"
	return
}

type StrikethroughHTMLRenderer struct {
	html.Config
}

func NewStrikethroughHTMLRenderer(opts ...html.Option) renderer.NodeRenderer {
	_ = "STUB: not implemented"
	return *new(renderer.NodeRenderer)
}

func (r *StrikethroughHTMLRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	_ = "STUB: not implemented"
	return
}

var StrikethroughAttributeFilter = html.GlobalAttributeFilter

func (r *StrikethroughHTMLRenderer) renderStrikethrough(
	w util.BufWriter, source []byte, n gast.Node, entering bool) (gast.WalkStatus, error) {
	_ = "STUB: not implemented"
	return *new(gast.WalkStatus), nil
}

type strikethrough struct {
}

var Strikethrough = &strikethrough{}

func (e *strikethrough) Extend(m goldmark.Markdown) { _ = "STUB: not implemented"; return }
