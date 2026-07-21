package extension

import (
	"github.com/yuin/goldmark"
	gast "github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
)

var uncloseCounterKey = parser.NewContextKey()

type unclosedCounter struct {
	Single int
	Double int
}

func (u *unclosedCounter) Reset() { _ = "STUB: not implemented"; return }

func getUnclosedCounter(pc parser.Context) *unclosedCounter { _ = "STUB: not implemented"; return nil }

type TypographicPunctuation int

const (
	LeftSingleQuote TypographicPunctuation = iota + 1

	RightSingleQuote

	LeftDoubleQuote

	RightDoubleQuote

	EnDash

	EmDash

	Ellipsis

	LeftAngleQuote

	RightAngleQuote

	Apostrophe

	typographicPunctuationMax
)

type TypographerConfig struct {
	Substitutions [][]byte
}

func newDefaultSubstitutions() [][]byte { _ = "STUB: not implemented"; return nil }

func (b *TypographerConfig) SetOption(name parser.OptionName, value any) {
	_ = "STUB: not implemented"
	return
}

type TypographerOption interface {
	parser.Option
	SetTypographerOption(*TypographerConfig)
}

const optTypographicSubstitutions parser.OptionName = "TypographicSubstitutions"

type TypographicSubstitutions map[TypographicPunctuation][]byte

type withTypographicSubstitutions struct {
	value [][]byte
}

func (o *withTypographicSubstitutions) SetParserOption(c *parser.Config) {
	_ = "STUB: not implemented"
	return
}

func (o *withTypographicSubstitutions) SetTypographerOption(p *TypographerConfig) {
	_ = "STUB: not implemented"
	return
}

func WithTypographicSubstitutions[T []byte | string](values map[TypographicPunctuation]T) TypographerOption {
	_ = "STUB: not implemented"
	return *new(TypographerOption)
}

type typographerDelimiterProcessor struct {
}

func (p *typographerDelimiterProcessor) IsDelimiter(b byte) bool {
	_ = "STUB: not implemented"
	return false
}

func (p *typographerDelimiterProcessor) CanOpenCloser(opener, closer *parser.Delimiter) bool {
	_ = "STUB: not implemented"
	return false
}

func (p *typographerDelimiterProcessor) OnMatch(consumes int) gast.Node {
	_ = "STUB: not implemented"
	return *new(gast.Node)
}

var defaultTypographerDelimiterProcessor = &typographerDelimiterProcessor{}

type typographerParser struct {
	TypographerConfig
}

func NewTypographerParser(opts ...TypographerOption) parser.InlineParser {
	_ = "STUB: not implemented"
	return *new(parser.InlineParser)
}

func (s *typographerParser) Trigger() []byte { _ = "STUB: not implemented"; return nil }

func (s *typographerParser) Parse(parent gast.Node, block text.Reader, pc parser.Context) gast.Node {
	_ = "STUB: not implemented"
	return *new(gast.Node)
}

func (s *typographerParser) CloseBlock(parent gast.Node, pc parser.Context) {
	_ = "STUB: not implemented"
	return
}

type typographer struct {
	options []TypographerOption
}

var Typographer = &typographer{}

func NewTypographer(opts ...TypographerOption) goldmark.Extender {
	_ = "STUB: not implemented"
	return *new(goldmark.Extender)
}

func (e *typographer) Extend(m goldmark.Markdown) { _ = "STUB: not implemented"; return }
