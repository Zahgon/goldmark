package extension

import (
	"regexp"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
)

var wwwURLRegxp = regexp.MustCompile(`^www\.[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-z]+(?:[/#?][-a-zA-Z0-9@:%_\+.~#!?&/=\(\);,'">\^{}\[\]` + "`" + `]*)?`) //nolint:golint,lll

var urlRegexp = regexp.MustCompile(`^(?:http|https|ftp)://[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-z]+(?::\d+)?(?:[/#?][-a-zA-Z0-9@:%_+.~#$!?&/=\(\);,'">\^{}\[\]` + "`" + `]*)?`) //nolint:golint,lll

type LinkifyConfig struct {
	AllowedProtocols [][]byte
	URLRegexp        *regexp.Regexp
	WWWRegexp        *regexp.Regexp
	EmailRegexp      *regexp.Regexp
}

const (
	optLinkifyAllowedProtocols parser.OptionName = "LinkifyAllowedProtocols"
	optLinkifyURLRegexp        parser.OptionName = "LinkifyURLRegexp"
	optLinkifyWWWRegexp        parser.OptionName = "LinkifyWWWRegexp"
	optLinkifyEmailRegexp      parser.OptionName = "LinkifyEmailRegexp"
)

func (c *LinkifyConfig) SetOption(name parser.OptionName, value any) {
	_ = "STUB: not implemented"
	return
}

type LinkifyOption interface {
	parser.Option
	SetLinkifyOption(*LinkifyConfig)
}

type withLinkifyAllowedProtocols struct {
	value [][]byte
}

func (o *withLinkifyAllowedProtocols) SetParserOption(c *parser.Config) {
	_ = "STUB: not implemented"
	return
}

func (o *withLinkifyAllowedProtocols) SetLinkifyOption(p *LinkifyConfig) {
	_ = "STUB: not implemented"
	return
}

func WithLinkifyAllowedProtocols[T []byte | string](value []T) LinkifyOption {
	_ = "STUB: not implemented"
	return *new(LinkifyOption)
}

type withLinkifyURLRegexp struct {
	value *regexp.Regexp
}

func (o *withLinkifyURLRegexp) SetParserOption(c *parser.Config) { _ = "STUB: not implemented"; return }

func (o *withLinkifyURLRegexp) SetLinkifyOption(p *LinkifyConfig) {
	_ = "STUB: not implemented"
	return
}

func WithLinkifyURLRegexp(value *regexp.Regexp) LinkifyOption {
	_ = "STUB: not implemented"
	return *new(LinkifyOption)
}

type withLinkifyWWWRegexp struct {
	value *regexp.Regexp
}

func (o *withLinkifyWWWRegexp) SetParserOption(c *parser.Config) { _ = "STUB: not implemented"; return }

func (o *withLinkifyWWWRegexp) SetLinkifyOption(p *LinkifyConfig) {
	_ = "STUB: not implemented"
	return
}

func WithLinkifyWWWRegexp(value *regexp.Regexp) LinkifyOption {
	_ = "STUB: not implemented"
	return *new(LinkifyOption)
}

type withLinkifyEmailRegexp struct {
	value *regexp.Regexp
}

func (o *withLinkifyEmailRegexp) SetParserOption(c *parser.Config) {
	_ = "STUB: not implemented"
	return
}

func (o *withLinkifyEmailRegexp) SetLinkifyOption(p *LinkifyConfig) {
	_ = "STUB: not implemented"
	return
}

func WithLinkifyEmailRegexp(value *regexp.Regexp) LinkifyOption {
	_ = "STUB: not implemented"
	return *new(LinkifyOption)
}

type linkifyParser struct {
	LinkifyConfig
}

func NewLinkifyParser(opts ...LinkifyOption) parser.InlineParser {
	_ = "STUB: not implemented"
	return *new(parser.InlineParser)
}

func (s *linkifyParser) Trigger() []byte { _ = "STUB: not implemented"; return nil }

var (
	protoHTTP  = []byte("http:")
	protoHTTPS = []byte("https:")
	protoFTP   = []byte("ftp:")
	domainWWW  = []byte("www.")
)

func (s *linkifyParser) Parse(parent ast.Node, block text.Reader, pc parser.Context) ast.Node {
	_ = "STUB: not implemented"
	return *new(ast.Node)
}

func (s *linkifyParser) CloseBlock(parent ast.Node, pc parser.Context) {
	_ = "STUB: not implemented"
	return
}

type linkify struct {
	options []LinkifyOption
}

var Linkify = &linkify{}

func NewLinkify(opts ...LinkifyOption) goldmark.Extender {
	_ = "STUB: not implemented"
	return *new(goldmark.Extender)
}

func (e *linkify) Extend(m goldmark.Markdown) { _ = "STUB: not implemented"; return }
