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

type definitionListParser struct {
}

var defaultDefinitionListParser = &definitionListParser{}

func NewDefinitionListParser() parser.BlockParser {
	_ = "STUB: not implemented"
	return *new(parser.BlockParser)
}

func (b *definitionListParser) Trigger() []byte { _ = "STUB: not implemented"; return nil }

func (b *definitionListParser) Open(parent gast.Node, reader text.Reader, pc parser.Context) (gast.Node, parser.State) {
	_ = "STUB: not implemented"
	return *new(gast.Node), *new(parser.State)
}

func (b *definitionListParser) Continue(node gast.Node, reader text.Reader, pc parser.Context) parser.State {
	_ = "STUB: not implemented"
	return *new(parser.State)
}

func (b *definitionListParser) Close(node gast.Node, reader text.Reader, pc parser.Context) {
	_ = "STUB: not implemented"
	return
}

func (b *definitionListParser) CanInterruptParagraph() bool {
	_ = "STUB: not implemented"
	return false
}

func (b *definitionListParser) CanAcceptIndentedLine() bool {
	_ = "STUB: not implemented"
	return false
}

type definitionDescriptionParser struct {
}

var defaultDefinitionDescriptionParser = &definitionDescriptionParser{}

func NewDefinitionDescriptionParser() parser.BlockParser {
	_ = "STUB: not implemented"
	return *new(parser.BlockParser)
}

func (b *definitionDescriptionParser) Trigger() []byte { _ = "STUB: not implemented"; return nil }

func (b *definitionDescriptionParser) Open(
	parent gast.Node, reader text.Reader, pc parser.Context) (gast.Node, parser.State) {
	_ = "STUB: not implemented"
	return *new(gast.Node), *new(parser.State)
}

func (b *definitionDescriptionParser) Continue(node gast.Node, reader text.Reader, pc parser.Context) parser.State {
	_ = "STUB: not implemented"
	return *new(parser.State)
}

func (b *definitionDescriptionParser) Close(node gast.Node, reader text.Reader, pc parser.Context) {
	_ = "STUB: not implemented"
	return
}

func (b *definitionDescriptionParser) CanInterruptParagraph() bool {
	_ = "STUB: not implemented"
	return false
}

func (b *definitionDescriptionParser) CanAcceptIndentedLine() bool {
	_ = "STUB: not implemented"
	return false
}

type DefinitionListHTMLRenderer struct {
	html.Config
}

func NewDefinitionListHTMLRenderer(opts ...html.Option) renderer.NodeRenderer {
	_ = "STUB: not implemented"
	return *new(renderer.NodeRenderer)
}

func (r *DefinitionListHTMLRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	_ = "STUB: not implemented"
	return
}

var DefinitionListAttributeFilter = html.GlobalAttributeFilter

func (r *DefinitionListHTMLRenderer) renderDefinitionList(
	w util.BufWriter, source []byte, n gast.Node, entering bool) (gast.WalkStatus, error) {
	_ = "STUB: not implemented"
	return *new(gast.WalkStatus), nil
}

var DefinitionTermAttributeFilter = html.GlobalAttributeFilter

func (r *DefinitionListHTMLRenderer) renderDefinitionTerm(
	w util.BufWriter, source []byte, n gast.Node, entering bool) (gast.WalkStatus, error) {
	_ = "STUB: not implemented"
	return *new(gast.WalkStatus), nil
}

var DefinitionDescriptionAttributeFilter = html.GlobalAttributeFilter

func (r *DefinitionListHTMLRenderer) renderDefinitionDescription(
	w util.BufWriter, source []byte, node gast.Node, entering bool) (gast.WalkStatus, error) {
	_ = "STUB: not implemented"
	return *new(gast.WalkStatus), nil
}

type definitionList struct {
}

var DefinitionList = &definitionList{}

func (e *definitionList) Extend(m goldmark.Markdown) { _ = "STUB: not implemented"; return }
