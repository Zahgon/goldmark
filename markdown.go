package goldmark

import (
	"io"

	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
)

func DefaultParser() parser.Parser { _ = "STUB: not implemented"; return *new(parser.Parser) }

func DefaultRenderer() renderer.Renderer { _ = "STUB: not implemented"; return *new(renderer.Renderer) }

var defaultMarkdown = New()

func Convert(source []byte, w io.Writer, opts ...parser.ParseOption) error {
	_ = "STUB: not implemented"
	return nil
}

type Markdown interface {
	Convert(source []byte, writer io.Writer, opts ...parser.ParseOption) error

	Parser() parser.Parser

	SetParser(parser.Parser)

	Renderer() renderer.Renderer

	SetRenderer(renderer.Renderer)
}

type Option func(*markdown)

func WithExtensions(ext ...Extender) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithParser(p parser.Parser) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithParserOptions(opts ...parser.Option) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithRenderer(r renderer.Renderer) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithRendererOptions(opts ...renderer.Option) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

type markdown struct {
	parser     parser.Parser
	renderer   renderer.Renderer
	extensions []Extender
}

func New(options ...Option) Markdown { _ = "STUB: not implemented"; return *new(Markdown) }

func (m *markdown) Convert(source []byte, writer io.Writer, opts ...parser.ParseOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *markdown) Parser() parser.Parser { _ = "STUB: not implemented"; return *new(parser.Parser) }

func (m *markdown) SetParser(v parser.Parser) { _ = "STUB: not implemented"; return }

func (m *markdown) Renderer() renderer.Renderer {
	_ = "STUB: not implemented"
	return *new(renderer.Renderer)
}

func (m *markdown) SetRenderer(v renderer.Renderer) { _ = "STUB: not implemented"; return }

type Extender interface {
	Extend(Markdown)
}
