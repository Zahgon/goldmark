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

var footnoteListKey = parser.NewContextKey()
var footnoteLinkListKey = parser.NewContextKey()

type footnoteBlockParser struct {
}

var defaultFootnoteBlockParser = &footnoteBlockParser{}

func NewFootnoteBlockParser() parser.BlockParser {
	_ = "STUB: not implemented"
	return *new(parser.BlockParser)
}

func (b *footnoteBlockParser) Trigger() []byte { _ = "STUB: not implemented"; return nil }

func (b *footnoteBlockParser) Open(parent gast.Node, reader text.Reader, pc parser.Context) (gast.Node, parser.State) {
	_ = "STUB: not implemented"
	return *new(gast.Node), *new(parser.State)
}

//nolint:staticcheck

func (b *footnoteBlockParser) Continue(node gast.Node, reader text.Reader, pc parser.Context) parser.State {
	_ = "STUB: not implemented"
	return *new(parser.State)
}

func (b *footnoteBlockParser) Close(node gast.Node, reader text.Reader, pc parser.Context) {
	_ = "STUB: not implemented"
	return
}

func (b *footnoteBlockParser) CanInterruptParagraph() bool { _ = "STUB: not implemented"; return false }

func (b *footnoteBlockParser) CanAcceptIndentedLine() bool { _ = "STUB: not implemented"; return false }

type footnoteParser struct {
}

var defaultFootnoteParser = &footnoteParser{}

func NewFootnoteParser() parser.InlineParser {
	_ = "STUB: not implemented"
	return *new(parser.InlineParser)
}

func (s *footnoteParser) Trigger() []byte { _ = "STUB: not implemented"; return nil }

func (s *footnoteParser) Parse(parent gast.Node, block text.Reader, pc parser.Context) gast.Node {
	_ = "STUB: not implemented"
	return *new(gast.Node)
}

//nolint:staticcheck

type footnoteASTTransformer struct {
}

var defaultFootnoteASTTransformer = &footnoteASTTransformer{}

func NewFootnoteASTTransformer() parser.ASTTransformer {
	_ = "STUB: not implemented"
	return *new(parser.ASTTransformer)
}

func (a *footnoteASTTransformer) Transform(node *gast.Document, reader text.Reader, pc parser.Context) {
	_ = "STUB: not implemented"
	return
}

type FootnoteConfig struct {
	html.Config

	IDPrefix []byte

	IDPrefixFunction func(gast.Node) []byte

	LinkTitle []byte

	BacklinkTitle []byte

	LinkClass []byte

	BacklinkClass []byte

	BacklinkHTML []byte
}

type FootnoteOption interface {
	renderer.Option

	SetFootnoteOption(*FootnoteConfig)
}

func NewFootnoteConfig() FootnoteConfig { _ = "STUB: not implemented"; return *new(FootnoteConfig) }

func (c *FootnoteConfig) SetOption(name renderer.OptionName, value any) {
	_ = "STUB: not implemented"
	return
}

type withFootnoteHTMLOptions struct {
	value []html.Option
}

func (o *withFootnoteHTMLOptions) SetConfig(c *renderer.Config) { _ = "STUB: not implemented"; return }

func (o *withFootnoteHTMLOptions) SetFootnoteOption(c *FootnoteConfig) {
	_ = "STUB: not implemented"
	return
}

func WithFootnoteHTMLOptions(opts ...html.Option) FootnoteOption {
	_ = "STUB: not implemented"
	return *new(FootnoteOption)
}

const optFootnoteIDPrefix renderer.OptionName = "FootnoteIDPrefix"

type withFootnoteIDPrefix struct {
	value []byte
}

func (o *withFootnoteIDPrefix) SetConfig(c *renderer.Config) { _ = "STUB: not implemented"; return }

func (o *withFootnoteIDPrefix) SetFootnoteOption(c *FootnoteConfig) {
	_ = "STUB: not implemented"
	return
}

func WithFootnoteIDPrefix[T []byte | string](a T) FootnoteOption {
	_ = "STUB: not implemented"
	return *new(FootnoteOption)
}

const optFootnoteIDPrefixFunction renderer.OptionName = "FootnoteIDPrefixFunction"

type withFootnoteIDPrefixFunction struct {
	value func(gast.Node) []byte
}

func (o *withFootnoteIDPrefixFunction) SetConfig(c *renderer.Config) {
	_ = "STUB: not implemented"
	return
}

func (o *withFootnoteIDPrefixFunction) SetFootnoteOption(c *FootnoteConfig) {
	_ = "STUB: not implemented"
	return
}

func WithFootnoteIDPrefixFunction(a func(gast.Node) []byte) FootnoteOption {
	_ = "STUB: not implemented"
	return *new(FootnoteOption)
}

const optFootnoteLinkTitle renderer.OptionName = "FootnoteLinkTitle"

type withFootnoteLinkTitle struct {
	value []byte
}

func (o *withFootnoteLinkTitle) SetConfig(c *renderer.Config) { _ = "STUB: not implemented"; return }

func (o *withFootnoteLinkTitle) SetFootnoteOption(c *FootnoteConfig) {
	_ = "STUB: not implemented"
	return
}

func WithFootnoteLinkTitle[T []byte | string](a T) FootnoteOption {
	_ = "STUB: not implemented"
	return *new(FootnoteOption)
}

const optFootnoteBacklinkTitle renderer.OptionName = "FootnoteBacklinkTitle"

type withFootnoteBacklinkTitle struct {
	value []byte
}

func (o *withFootnoteBacklinkTitle) SetConfig(c *renderer.Config) {
	_ = "STUB: not implemented"
	return
}

func (o *withFootnoteBacklinkTitle) SetFootnoteOption(c *FootnoteConfig) {
	_ = "STUB: not implemented"
	return
}

func WithFootnoteBacklinkTitle[T []byte | string](a T) FootnoteOption {
	_ = "STUB: not implemented"
	return *new(FootnoteOption)
}

const optFootnoteLinkClass renderer.OptionName = "FootnoteLinkClass"

type withFootnoteLinkClass struct {
	value []byte
}

func (o *withFootnoteLinkClass) SetConfig(c *renderer.Config) { _ = "STUB: not implemented"; return }

func (o *withFootnoteLinkClass) SetFootnoteOption(c *FootnoteConfig) {
	_ = "STUB: not implemented"
	return
}

func WithFootnoteLinkClass[T []byte | string](a T) FootnoteOption {
	_ = "STUB: not implemented"
	return *new(FootnoteOption)
}

const optFootnoteBacklinkClass renderer.OptionName = "FootnoteBacklinkClass"

type withFootnoteBacklinkClass struct {
	value []byte
}

func (o *withFootnoteBacklinkClass) SetConfig(c *renderer.Config) {
	_ = "STUB: not implemented"
	return
}

func (o *withFootnoteBacklinkClass) SetFootnoteOption(c *FootnoteConfig) {
	_ = "STUB: not implemented"
	return
}

func WithFootnoteBacklinkClass[T []byte | string](a T) FootnoteOption {
	_ = "STUB: not implemented"
	return *new(FootnoteOption)
}

const optFootnoteBacklinkHTML renderer.OptionName = "FootnoteBacklinkHTML"

type withFootnoteBacklinkHTML struct {
	value []byte
}

func (o *withFootnoteBacklinkHTML) SetConfig(c *renderer.Config) { _ = "STUB: not implemented"; return }

func (o *withFootnoteBacklinkHTML) SetFootnoteOption(c *FootnoteConfig) {
	_ = "STUB: not implemented"
	return
}

func WithFootnoteBacklinkHTML[T []byte | string](a T) FootnoteOption {
	_ = "STUB: not implemented"
	return *new(FootnoteOption)
}

type FootnoteHTMLRenderer struct {
	FootnoteConfig
}

func NewFootnoteHTMLRenderer(opts ...FootnoteOption) renderer.NodeRenderer {
	_ = "STUB: not implemented"
	return *new(renderer.NodeRenderer)
}

func (r *FootnoteHTMLRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	_ = "STUB: not implemented"
	return
}

func (r *FootnoteHTMLRenderer) renderFootnoteLink(
	w util.BufWriter, source []byte, node gast.Node, entering bool) (gast.WalkStatus, error) {
	_ = "STUB: not implemented"
	return *new(gast.WalkStatus), nil
}

func (r *FootnoteHTMLRenderer) renderFootnoteBacklink(
	w util.BufWriter, source []byte, node gast.Node, entering bool) (gast.WalkStatus, error) {
	_ = "STUB: not implemented"
	return *new(gast.WalkStatus), nil
}

func (r *FootnoteHTMLRenderer) renderFootnote(
	w util.BufWriter, source []byte, node gast.Node, entering bool) (gast.WalkStatus, error) {
	_ = "STUB: not implemented"
	return *new(gast.WalkStatus), nil
}

func (r *FootnoteHTMLRenderer) renderFootnoteList(
	w util.BufWriter, source []byte, node gast.Node, entering bool) (gast.WalkStatus, error) {
	_ = "STUB: not implemented"
	return *new(gast.WalkStatus), nil
}

func (r *FootnoteHTMLRenderer) idPrefix(node gast.Node) []byte {
	_ = "STUB: not implemented"
	return nil
}

func applyFootnoteTemplate(b []byte, index, refCount int) []byte {
	_ = "STUB: not implemented"
	return nil
}

type footnote struct {
	options []FootnoteOption
}

var Footnote = &footnote{
	options: []FootnoteOption{},
}

func NewFootnote(opts ...FootnoteOption) goldmark.Extender {
	_ = "STUB: not implemented"
	return *new(goldmark.Extender)
}

func (e *footnote) Extend(m goldmark.Markdown) { _ = "STUB: not implemented"; return }
