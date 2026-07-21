package renderer

import (
	"io"
	"sync"

	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/util"
)

type Config struct {
	Options       map[OptionName]any
	NodeRenderers util.PrioritizedSlice
}

func NewConfig() *Config { _ = "STUB: not implemented"; return nil }

type OptionName string

type Option interface {
	SetConfig(*Config)
}

type withNodeRenderers struct {
	value []util.PrioritizedValue
}

func (o *withNodeRenderers) SetConfig(c *Config) { _ = "STUB: not implemented"; return }

func WithNodeRenderers(ps ...util.PrioritizedValue) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

type withOption struct {
	name  OptionName
	value any
}

func (o *withOption) SetConfig(c *Config) { _ = "STUB: not implemented"; return }

func WithOption(name OptionName, value any) Option { _ = "STUB: not implemented"; return *new(Option) }

type SetOptioner interface {
	SetOption(name OptionName, value any)
}

type NodeRendererFunc func(writer util.BufWriter, source []byte, n ast.Node, entering bool) (ast.WalkStatus, error)

type NodeRenderer interface {
	RegisterFuncs(NodeRendererFuncRegisterer)
}

type NodeRendererFuncRegisterer interface {
	Register(ast.NodeKind, NodeRendererFunc)
}

type Renderer interface {
	Render(w io.Writer, source []byte, n ast.Node) error

	AddOptions(...Option)
}

type renderer struct {
	config               *Config
	options              map[OptionName]any
	nodeRendererFuncsTmp map[ast.NodeKind]NodeRendererFunc
	maxKind              int
	nodeRendererFuncs    []NodeRendererFunc
	initSync             sync.Once
}

func NewRenderer(options ...Option) Renderer { _ = "STUB: not implemented"; return *new(Renderer) }

func (r *renderer) AddOptions(opts ...Option) { _ = "STUB: not implemented"; return }

func (r *renderer) Register(kind ast.NodeKind, v NodeRendererFunc) {
	_ = "STUB: not implemented"
	return
}

func (r *renderer) Render(w io.Writer, source []byte, n ast.Node) error {
	_ = "STUB: not implemented"
	return nil
}
