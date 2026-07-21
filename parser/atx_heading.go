package parser

import (
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/text"
)

type HeadingConfig struct {
	AutoHeadingID bool
	Attribute     bool
}

func (b *HeadingConfig) SetOption(name OptionName, _ any) { _ = "STUB: not implemented"; return }

type HeadingOption interface {
	Option
	SetHeadingOption(*HeadingConfig)
}

const optAutoHeadingID OptionName = "AutoHeadingID"

type withAutoHeadingID struct {
}

func (o *withAutoHeadingID) SetParserOption(c *Config) { _ = "STUB: not implemented"; return }

func (o *withAutoHeadingID) SetHeadingOption(p *HeadingConfig) { _ = "STUB: not implemented"; return }

func WithAutoHeadingID() HeadingOption { _ = "STUB: not implemented"; return *new(HeadingOption) }

type withHeadingAttribute struct {
	Option
}

func (o *withHeadingAttribute) SetHeadingOption(p *HeadingConfig) {
	_ = "STUB: not implemented"
	return
}

func WithHeadingAttribute() HeadingOption { _ = "STUB: not implemented"; return *new(HeadingOption) }

type atxHeadingParser struct {
	HeadingConfig
}

func NewATXHeadingParser(opts ...HeadingOption) BlockParser {
	_ = "STUB: not implemented"
	return *new(BlockParser)
}

func (b *atxHeadingParser) Trigger() []byte { _ = "STUB: not implemented"; return nil }

func (b *atxHeadingParser) Open(parent ast.Node, reader text.Reader, pc Context) (ast.Node, State) {
	_ = "STUB: not implemented"
	return *new(ast.Node), *new(State)
}

func (b *atxHeadingParser) Continue(node ast.Node, reader text.Reader, pc Context) State {
	_ = "STUB: not implemented"
	return *new(State)
}

func (b *atxHeadingParser) Close(node ast.Node, reader text.Reader, pc Context) {
	_ = "STUB: not implemented"
	return
}

func (b *atxHeadingParser) CanInterruptParagraph() bool { _ = "STUB: not implemented"; return false }

func (b *atxHeadingParser) CanAcceptIndentedLine() bool { _ = "STUB: not implemented"; return false }

func generateAutoHeadingID(node *ast.Heading, reader text.Reader, pc Context) {
	_ = "STUB: not implemented"
	return
}

func parseLastLineAttributes(node ast.Node, reader text.Reader, _ Context) {
	_ = "STUB: not implemented"
	return
}
