package parser

import (
	"regexp"

	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/text"
)

type rawHTMLParser struct {
}

var defaultRawHTMLParser = &rawHTMLParser{}

func NewRawHTMLParser() InlineParser { _ = "STUB: not implemented"; return *new(InlineParser) }

func (s *rawHTMLParser) Trigger() []byte { _ = "STUB: not implemented"; return nil }

func (s *rawHTMLParser) Parse(parent ast.Node, block text.Reader, pc Context) ast.Node {
	_ = "STUB: not implemented"
	return *new(ast.Node)
}

var tagnamePattern = `([A-Za-z][A-Za-z0-9-]*)`
var spaceOrOneNewline = `(?:[ \t]|(?:\r\n|\n){0,1})`
var attributePattern = `(?:[\r\n \t]+[a-zA-Z_:][a-zA-Z0-9:._-]*(?:[\r\n \t]*=[\r\n \t]*(?:[^\"'=<>` + "`" + `\x00-\x20]+|'[^']*'|"[^"]*"))?)` //nolint:golint,lll
var openTagRegexp = regexp.MustCompile("^<" + tagnamePattern + attributePattern + `*` + spaceOrOneNewline + `*/?>`)
var closeTagRegexp = regexp.MustCompile("^</" + tagnamePattern + spaceOrOneNewline + `*>`)

var openProcessingInstruction = []byte("<?")
var closeProcessingInstruction = []byte("?>")
var openCDATA = []byte("<![CDATA[")
var closeCDATA = []byte("]]>")
var closeDecl = []byte(">")
var emptyComment1 = []byte("<!-->")
var emptyComment2 = []byte("<!--->")
var openComment = []byte("<!--")
var closeComment = []byte("-->")

func (s *rawHTMLParser) parseComment(block text.Reader, _ Context) ast.Node {
	_ = "STUB: not implemented"
	return *new(ast.Node)
}

func (s *rawHTMLParser) parseUntil(block text.Reader, closer []byte, _ Context) ast.Node {
	_ = "STUB: not implemented"
	return *new(ast.Node)
}

func (s *rawHTMLParser) parseMultiLineRegexp(reg *regexp.Regexp, block text.Reader, _ Context) ast.Node {
	_ = "STUB: not implemented"
	return *new(ast.Node)
}
