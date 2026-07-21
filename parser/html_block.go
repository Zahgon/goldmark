package parser

import (
	"regexp"

	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/text"
)

var allowedBlockTags = map[string]bool{
	"address":    true,
	"article":    true,
	"aside":      true,
	"base":       true,
	"basefont":   true,
	"blockquote": true,
	"body":       true,
	"caption":    true,
	"center":     true,
	"col":        true,
	"colgroup":   true,
	"dd":         true,
	"details":    true,
	"dialog":     true,
	"dir":        true,
	"div":        true,
	"dl":         true,
	"dt":         true,
	"fieldset":   true,
	"figcaption": true,
	"figure":     true,
	"footer":     true,
	"form":       true,
	"frame":      true,
	"frameset":   true,
	"h1":         true,
	"h2":         true,
	"h3":         true,
	"h4":         true,
	"h5":         true,
	"h6":         true,
	"head":       true,
	"header":     true,
	"hr":         true,
	"html":       true,
	"iframe":     true,
	"legend":     true,
	"li":         true,
	"link":       true,
	"main":       true,
	"menu":       true,
	"menuitem":   true,
	"meta":       true,
	"nav":        true,
	"noframes":   true,
	"ol":         true,
	"optgroup":   true,
	"option":     true,
	"p":          true,
	"param":      true,
	"search":     true,
	"section":    true,
	"summary":    true,
	"table":      true,
	"tbody":      true,
	"td":         true,
	"tfoot":      true,
	"th":         true,
	"thead":      true,
	"title":      true,
	"tr":         true,
	"track":      true,
	"ul":         true,
}

var htmlBlockType1OpenRegexp = regexp.MustCompile(`(?i)^[ ]{0,3}<(script|pre|style|textarea)(?:\s.*|>.*|/>.*|)(?:\r\n|\n)?$`) //nolint:golint,lll
var htmlBlockType1CloseRegexp = regexp.MustCompile(`(?i)^.*</(?:script|pre|style|textarea)>.*`)

var htmlBlockType2OpenRegexp = regexp.MustCompile(`^[ ]{0,3}<!\-\-`)
var htmlBlockType2Close = []byte{'-', '-', '>'}

var htmlBlockType3OpenRegexp = regexp.MustCompile(`^[ ]{0,3}<\?`)
var htmlBlockType3Close = []byte{'?', '>'}

var htmlBlockType4OpenRegexp = regexp.MustCompile(`^[ ]{0,3}<![A-Z]+.*(?:\r\n|\n)?$`)
var htmlBlockType4Close = []byte{'>'}

var htmlBlockType5OpenRegexp = regexp.MustCompile(`^[ ]{0,3}<\!\[CDATA\[`)
var htmlBlockType5Close = []byte{']', ']', '>'}

var htmlBlockType6Regexp = regexp.MustCompile(`^[ ]{0,3}<(?:/[ ]*)?([a-zA-Z]+[a-zA-Z0-9\-]*)(?:[ ].*|>.*|/>.*|)(?:\r\n|\n)?$`) //nolint:golint,lll

var htmlBlockType7Regexp = regexp.MustCompile(`^[ ]{0,3}<(/[ ]*)?([a-zA-Z]+[a-zA-Z0-9\-]*)(` + attributePattern + `*)[ ]*(?:>|/>)[ ]*(?:\r\n|\n)?$`) //nolint:golint,lll

type htmlBlockParser struct {
}

var defaultHTMLBlockParser = &htmlBlockParser{}

func NewHTMLBlockParser() BlockParser { _ = "STUB: not implemented"; return *new(BlockParser) }

func (b *htmlBlockParser) Trigger() []byte { _ = "STUB: not implemented"; return nil }

func (b *htmlBlockParser) Open(parent ast.Node, reader text.Reader, pc Context) (ast.Node, State) {
	_ = "STUB: not implemented"
	return *new(ast.Node), *new(State)
}

func (b *htmlBlockParser) Continue(node ast.Node, reader text.Reader, pc Context) State {
	_ = "STUB: not implemented"
	return *new(State)
}

func (b *htmlBlockParser) Close(node ast.Node, reader text.Reader, pc Context) {
	_ = "STUB: not implemented"
	return
}

func (b *htmlBlockParser) CanInterruptParagraph() bool { _ = "STUB: not implemented"; return false }

func (b *htmlBlockParser) CanAcceptIndentedLine() bool { _ = "STUB: not implemented"; return false }
