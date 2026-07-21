package parser

import (
	"sync"

	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

type Reference interface {
	String() string

	Label() []byte

	Destination() []byte

	Title() []byte
}

type reference struct {
	label       []byte
	destination []byte
	title       []byte
}

func NewReference(label, destination, title []byte) Reference {
	_ = "STUB: not implemented"
	return *new(Reference)
}

func newASTReference(v *ast.LinkReferenceDefinition) Reference {
	_ = "STUB: not implemented"
	return *new(Reference)
}

func (r *reference) Label() []byte { _ = "STUB: not implemented"; return nil }

func (r *reference) Destination() []byte { _ = "STUB: not implemented"; return nil }

func (r *reference) Title() []byte { _ = "STUB: not implemented"; return nil }

func (r *reference) String() string { _ = "STUB: not implemented"; return "" }

type astReference struct {
	v *ast.LinkReferenceDefinition
}

func (r *astReference) Label() []byte { _ = "STUB: not implemented"; return nil }

func (r *astReference) Destination() []byte { _ = "STUB: not implemented"; return nil }

func (r *astReference) Title() []byte { _ = "STUB: not implemented"; return nil }

func (r *astReference) String() string { _ = "STUB: not implemented"; return "" }

type IDs interface {
	Generate(value []byte, kind ast.NodeKind) []byte

	Put(value []byte)
}

type ids struct {
	values map[string]bool
}

func newIDs() IDs { _ = "STUB: not implemented"; return *new(IDs) }

func (s *ids) Generate(value []byte, kind ast.NodeKind) []byte {
	_ = "STUB: not implemented"
	return nil
}

func (s *ids) Put(value []byte) { _ = "STUB: not implemented"; return }

type ContextKey int

var ContextKeyMax ContextKey

func NewContextKey() ContextKey { _ = "STUB: not implemented"; return *new(ContextKey) }

type Context interface {
	String() string

	Get(ContextKey) any

	ComputeIfAbsent(ContextKey, func() any) any

	Set(ContextKey, any)

	AddReference(Reference)

	Reference(label string) (Reference, bool)

	References() []Reference

	IDs() IDs

	BlockOffset() int

	SetBlockOffset(int)

	BlockIndent() int

	SetBlockIndent(int)

	FirstDelimiter() *Delimiter

	LastDelimiter() *Delimiter

	PushDelimiter(delimiter *Delimiter)

	RemoveDelimiter(d *Delimiter)

	ClearDelimiters(bottom ast.Node)

	OpenedBlocks() []Block

	SetOpenedBlocks([]Block)

	LastOpenedBlock() Block

	IsInLinkLabel() bool
}

type ContextConfig struct {
	IDs IDs
}

type ContextOption func(*ContextConfig)

func WithIDs(ids IDs) ContextOption { _ = "STUB: not implemented"; return *new(ContextOption) }

type parseContext struct {
	store         []any
	ids           IDs
	refs          map[string]Reference
	blockOffset   int
	blockIndent   int
	delimiters    *Delimiter
	lastDelimiter *Delimiter
	openedBlocks  []Block
}

func NewContext(options ...ContextOption) Context { _ = "STUB: not implemented"; return *new(Context) }

func (p *parseContext) Get(key ContextKey) any { _ = "STUB: not implemented"; return *new(any) }

func (p *parseContext) ComputeIfAbsent(key ContextKey, f func() any) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (p *parseContext) Set(key ContextKey, value any) { _ = "STUB: not implemented"; return }

func (p *parseContext) IDs() IDs { _ = "STUB: not implemented"; return *new(IDs) }

func (p *parseContext) BlockOffset() int { _ = "STUB: not implemented"; return 0 }

func (p *parseContext) SetBlockOffset(v int) { _ = "STUB: not implemented"; return }

func (p *parseContext) BlockIndent() int { _ = "STUB: not implemented"; return 0 }

func (p *parseContext) SetBlockIndent(v int) { _ = "STUB: not implemented"; return }

func (p *parseContext) LastDelimiter() *Delimiter { _ = "STUB: not implemented"; return nil }

func (p *parseContext) FirstDelimiter() *Delimiter { _ = "STUB: not implemented"; return nil }

func (p *parseContext) PushDelimiter(d *Delimiter) { _ = "STUB: not implemented"; return }

func (p *parseContext) RemoveDelimiter(d *Delimiter) { _ = "STUB: not implemented"; return }

func (p *parseContext) ClearDelimiters(bottom ast.Node) { _ = "STUB: not implemented"; return }

func (p *parseContext) AddReference(ref Reference) { _ = "STUB: not implemented"; return }

func (p *parseContext) Reference(label string) (Reference, bool) {
	_ = "STUB: not implemented"
	return *new(Reference), false
}

func (p *parseContext) References() []Reference { _ = "STUB: not implemented"; return nil }

func (p *parseContext) String() string { _ = "STUB: not implemented"; return "" }

func (p *parseContext) OpenedBlocks() []Block { _ = "STUB: not implemented"; return nil }

func (p *parseContext) SetOpenedBlocks(v []Block) { _ = "STUB: not implemented"; return }

func (p *parseContext) LastOpenedBlock() Block { _ = "STUB: not implemented"; return *new(Block) }

func (p *parseContext) IsInLinkLabel() bool { _ = "STUB: not implemented"; return false }

type State int

const (
	None State = 1 << iota

	Continue

	Close

	HasChildren

	NoChildren

	RequireParagraph
)

type Config struct {
	Options               map[OptionName]any
	BlockParsers          util.PrioritizedSlice
	InlineParsers         util.PrioritizedSlice
	ParagraphTransformers util.PrioritizedSlice
	ASTTransformers       util.PrioritizedSlice
	EscapedSpace          bool
}

func NewConfig() *Config { _ = "STUB: not implemented"; return nil }

type Option interface {
	SetParserOption(*Config)
}

type OptionName string

const optAttribute OptionName = "Attribute"

type withAttribute struct {
}

func (o *withAttribute) SetParserOption(c *Config) { _ = "STUB: not implemented"; return }

func WithAttribute() Option { _ = "STUB: not implemented"; return *new(Option) }

type Parser interface {
	Parse(reader text.Reader, opts ...ParseOption) ast.Node

	AddOptions(...Option)
}

type SetOptioner interface {
	SetOption(name OptionName, value any)
}

type BlockParser interface {
	Trigger() []byte

	Open(parent ast.Node, reader text.Reader, pc Context) (ast.Node, State)

	Continue(node ast.Node, reader text.Reader, pc Context) State

	Close(node ast.Node, reader text.Reader, pc Context)

	CanInterruptParagraph() bool

	CanAcceptIndentedLine() bool
}

type InlineParser interface {
	Trigger() []byte

	Parse(parent ast.Node, block text.Reader, pc Context) ast.Node
}

type CloseBlocker interface {
	CloseBlock(parent ast.Node, block text.Reader, pc Context)
}

type ParagraphTransformer interface {
	Transform(node *ast.Paragraph, reader text.Reader, pc Context)
}

type ASTTransformer interface {
	Transform(node *ast.Document, reader text.Reader, pc Context)
}

func DefaultBlockParsers() []util.PrioritizedValue { _ = "STUB: not implemented"; return nil }

func DefaultInlineParsers() []util.PrioritizedValue { _ = "STUB: not implemented"; return nil }

func DefaultParagraphTransformers() []util.PrioritizedValue { _ = "STUB: not implemented"; return nil }

type Block struct {
	Node ast.Node

	Parser BlockParser
}

type parser struct {
	options               map[OptionName]any
	blockParsers          [256][]BlockParser
	freeBlockParsers      []BlockParser
	inlineParsers         [256][]InlineParser
	closeBlockers         []CloseBlocker
	paragraphTransformers []ParagraphTransformer
	astTransformers       []ASTTransformer
	escapedSpace          bool
	config                *Config
	initSync              sync.Once
}

type withBlockParsers struct {
	value []util.PrioritizedValue
}

func (o *withBlockParsers) SetParserOption(c *Config) { _ = "STUB: not implemented"; return }

func WithBlockParsers(bs ...util.PrioritizedValue) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

type withInlineParsers struct {
	value []util.PrioritizedValue
}

func (o *withInlineParsers) SetParserOption(c *Config) { _ = "STUB: not implemented"; return }

func WithInlineParsers(bs ...util.PrioritizedValue) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

type withParagraphTransformers struct {
	value []util.PrioritizedValue
}

func (o *withParagraphTransformers) SetParserOption(c *Config) { _ = "STUB: not implemented"; return }

func WithParagraphTransformers(ps ...util.PrioritizedValue) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

type withASTTransformers struct {
	value []util.PrioritizedValue
}

func (o *withASTTransformers) SetParserOption(c *Config) { _ = "STUB: not implemented"; return }

func WithASTTransformers(ps ...util.PrioritizedValue) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

type withEscapedSpace struct {
}

func (o *withEscapedSpace) SetParserOption(c *Config) { _ = "STUB: not implemented"; return }

func WithEscapedSpace() Option { _ = "STUB: not implemented"; return *new(Option) }

type withOption struct {
	name  OptionName
	value any
}

func (o *withOption) SetParserOption(c *Config) { _ = "STUB: not implemented"; return }

func WithOption(name OptionName, value any) Option { _ = "STUB: not implemented"; return *new(Option) }

func NewParser(options ...Option) Parser { _ = "STUB: not implemented"; return *new(Parser) }

func (p *parser) AddOptions(opts ...Option) { _ = "STUB: not implemented"; return }

func (p *parser) addBlockParser(v util.PrioritizedValue, options map[OptionName]any) {
	_ = "STUB: not implemented"
	return
}

func (p *parser) addInlineParser(v util.PrioritizedValue, options map[OptionName]any) {
	_ = "STUB: not implemented"
	return
}

func (p *parser) addParagraphTransformer(v util.PrioritizedValue, options map[OptionName]any) {
	_ = "STUB: not implemented"
	return
}

func (p *parser) addASTTransformer(v util.PrioritizedValue, options map[OptionName]any) {
	_ = "STUB: not implemented"
	return
}

type ParseConfig struct {
	Context Context
}

type ParseOption func(c *ParseConfig)

func WithContext(context Context) ParseOption { _ = "STUB: not implemented"; return *new(ParseOption) }

func (p *parser) Parse(reader text.Reader, opts ...ParseOption) ast.Node {
	_ = "STUB: not implemented"
	return *new(ast.Node)
}

func (p *parser) transformParagraph(node *ast.Paragraph, reader text.Reader, pc Context) bool {
	_ = "STUB: not implemented"
	return false
}

func (p *parser) closeBlocks(from, to int, reader text.Reader, pc Context) {
	_ = "STUB: not implemented"
	return
}

type blockOpenResult int

const (
	paragraphContinuation blockOpenResult = iota + 1
	newBlocksOpened
	noBlocksOpened
)

func (p *parser) openBlocks(parent ast.Node, blankLine bool, reader text.Reader, pc Context) blockOpenResult {
	_ = "STUB: not implemented"
	return *new(blockOpenResult)
}

type lineStat struct {
	lineNum int
	level   int
	isBlank bool
}

func isBlankLine(lineNum, level int, stats []lineStat) bool {
	_ = "STUB: not implemented"
	return false
}

func (p *parser) parseBlocks(parent ast.Node, reader text.Reader, pc Context) {
	_ = "STUB: not implemented"
	return
}

func (p *parser) walkBlock(block ast.Node, cb func(node ast.Node)) {
	_ = "STUB: not implemented"
	return
}

const (
	lineBreakHard uint8 = 1 << iota
	lineBreakSoft
	lineBreakVisible
)

func (p *parser) parseBlock(block text.BlockReader, parent ast.Node, pc Context) {
	_ = "STUB: not implemented"
	return
}
