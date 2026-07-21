package html

import (
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
)

type Config struct {
	Writer              Writer
	HardWraps           bool
	EastAsianLineBreaks EastAsianLineBreaks
	XHTML               bool
	Unsafe              bool
}

func NewConfig() Config { _ = "STUB: not implemented"; return *new(Config) }

func (c *Config) SetOption(name renderer.OptionName, value any) { _ = "STUB: not implemented"; return }

type Option interface {
	SetHTMLOption(*Config)
}

const optTextWriter renderer.OptionName = "Writer"

type withWriter struct {
	value Writer
}

func (o *withWriter) SetConfig(c *renderer.Config) { _ = "STUB: not implemented"; return }

func (o *withWriter) SetHTMLOption(c *Config) { _ = "STUB: not implemented"; return }

func WithWriter(writer Writer) interface {
	renderer.Option
	Option
} {
	_ = "STUB: not implemented"
	return nil
}

const optHardWraps renderer.OptionName = "HardWraps"

type withHardWraps struct {
}

func (o *withHardWraps) SetConfig(c *renderer.Config) { _ = "STUB: not implemented"; return }

func (o *withHardWraps) SetHTMLOption(c *Config) { _ = "STUB: not implemented"; return }

func WithHardWraps() interface {
	renderer.Option
	Option
} {
	_ = "STUB: not implemented"
	return nil
}

const optEastAsianLineBreaks renderer.OptionName = "EastAsianLineBreaks"

type EastAsianLineBreaks int

const (
	EastAsianLineBreaksNone EastAsianLineBreaks = iota

	EastAsianLineBreaksSimple

	EastAsianLineBreaksCSS3Draft
)

func (b EastAsianLineBreaks) softLineBreak(thisLastRune rune, siblingFirstRune rune) bool {
	_ = "STUB: not implemented"
	return false
}

func eastAsianLineBreaksCSS3DraftSoftLineBreak(thisLastRune rune, siblingFirstRune rune) bool {
	_ = "STUB: not implemented"
	return false
}

type withEastAsianLineBreaks struct {
	eastAsianLineBreaksStyle EastAsianLineBreaks
}

func (o *withEastAsianLineBreaks) SetConfig(c *renderer.Config) { _ = "STUB: not implemented"; return }

func (o *withEastAsianLineBreaks) SetHTMLOption(c *Config) { _ = "STUB: not implemented"; return }

func WithEastAsianLineBreaks(e EastAsianLineBreaks) interface {
	renderer.Option
	Option
} {
	_ = "STUB: not implemented"
	return nil
}

const optXHTML renderer.OptionName = "XHTML"

type withXHTML struct {
}

func (o *withXHTML) SetConfig(c *renderer.Config) { _ = "STUB: not implemented"; return }

func (o *withXHTML) SetHTMLOption(c *Config) { _ = "STUB: not implemented"; return }

func WithXHTML() interface {
	Option
	renderer.Option
} {
	_ = "STUB: not implemented"
	return nil
}

const optUnsafe renderer.OptionName = "Unsafe"

type withUnsafe struct {
}

func (o *withUnsafe) SetConfig(c *renderer.Config) { _ = "STUB: not implemented"; return }

func (o *withUnsafe) SetHTMLOption(c *Config) { _ = "STUB: not implemented"; return }

func WithUnsafe() interface {
	renderer.Option
	Option
} {
	_ = "STUB: not implemented"
	return nil
}

type Renderer struct {
	Config
}

func NewRenderer(opts ...Option) renderer.NodeRenderer {
	_ = "STUB: not implemented"
	return *new(renderer.NodeRenderer)
}

func (r *Renderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	_ = "STUB: not implemented"
	return
}

func (r *Renderer) writeLines(w util.BufWriter, source []byte, n ast.Node) {
	_ = "STUB: not implemented"
	return
}

var GlobalAttributeFilter = util.NewBytesFilterString(`accesskey,autocapitalize,autofocus,class,contenteditable,dir,draggable,enterkeyhint,hidden,id,inert,inputmode,is,itemid,itemprop,itemref,itemscope,itemtype,lang,part,role,slot,spellcheck,style,tabindex,title,translate`)

func (r *Renderer) renderDocument(
	w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus), nil
}

var HeadingAttributeFilter = GlobalAttributeFilter

func (r *Renderer) renderHeading(
	w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus), nil
}

var BlockquoteAttributeFilter = GlobalAttributeFilter.ExtendString(`cite`)

func (r *Renderer) renderBlockquote(
	w util.BufWriter, source []byte, n ast.Node, entering bool) (ast.WalkStatus, error) {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus), nil
}

func (r *Renderer) renderCodeBlock(w util.BufWriter, source []byte, n ast.Node, entering bool) (ast.WalkStatus, error) {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus), nil
}

func (r *Renderer) renderFencedCodeBlock(
	w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus), nil
}

func (r *Renderer) renderHTMLBlock(
	w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus), nil
}

var ListAttributeFilter = GlobalAttributeFilter.ExtendString(`start,reversed,type`)

func (r *Renderer) renderList(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus), nil
}

var ListItemAttributeFilter = GlobalAttributeFilter.ExtendString(`value`)

func (r *Renderer) renderListItem(w util.BufWriter, source []byte, n ast.Node, entering bool) (ast.WalkStatus, error) {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus), nil
}

var ParagraphAttributeFilter = GlobalAttributeFilter

func (r *Renderer) renderParagraph(w util.BufWriter, source []byte, n ast.Node, entering bool) (ast.WalkStatus, error) {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus), nil
}

func (r *Renderer) renderTextBlock(w util.BufWriter, source []byte, n ast.Node, entering bool) (ast.WalkStatus, error) {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus), nil
}

var ThematicAttributeFilter = GlobalAttributeFilter.ExtendString(`align,color,noshade,size,width`)

func (r *Renderer) renderThematicBreak(
	w util.BufWriter, source []byte, n ast.Node, entering bool) (ast.WalkStatus, error) {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus), nil
}

var LinkAttributeFilter = GlobalAttributeFilter.ExtendString(`download,href,lang,media,ping,referrerpolicy,rel,shape,target`)

func (r *Renderer) renderAutoLink(
	w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus), nil
}

var CodeAttributeFilter = GlobalAttributeFilter

func (r *Renderer) renderCodeSpan(w util.BufWriter, source []byte, n ast.Node, entering bool) (ast.WalkStatus, error) {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus), nil
}

var EmphasisAttributeFilter = GlobalAttributeFilter

func (r *Renderer) renderEmphasis(
	w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus), nil
}

func (r *Renderer) renderLink(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus), nil
}

var ImageAttributeFilter = GlobalAttributeFilter.ExtendString(`align,border,crossorigin,decoding,height,importance,intrinsicsize,ismap,loading,referrerpolicy,sizes,srcset,usemap,width`)

func (r *Renderer) renderImage(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus), nil
}

func (r *Renderer) renderRawHTML(
	w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus), nil
}

func (r *Renderer) renderText(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus), nil
}

func (r *Renderer) renderString(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus), nil
}

func (r *Renderer) renderTexts(w util.BufWriter, source []byte, n ast.Node) {
	_ = "STUB: not implemented"
	return
}

var dataPrefix = []byte("data-")

func RenderAttributes(w util.BufWriter, node ast.Node, filter util.BytesFilter) {
	_ = "STUB: not implemented"
	return
}

type Writer interface {
	Write(writer util.BufWriter, source []byte)

	RawWrite(writer util.BufWriter, source []byte)

	SecureWrite(writer util.BufWriter, source []byte)
}

var replacementCharacter = []byte("\ufffd")

type WriterConfig struct {
	EscapedSpace bool
}

type WriterOption func(*WriterConfig)

func WithEscapedSpace() WriterOption { _ = "STUB: not implemented"; return *new(WriterOption) }

type defaultWriter struct {
	WriterConfig
}

func NewWriter(opts ...WriterOption) Writer { _ = "STUB: not implemented"; return *new(Writer) }

func escapeRune(writer util.BufWriter, r rune) { _ = "STUB: not implemented"; return }

func (d *defaultWriter) SecureWrite(writer util.BufWriter, source []byte) {
	_ = "STUB: not implemented"
	return
}

func (d *defaultWriter) RawWrite(writer util.BufWriter, source []byte) {
	_ = "STUB: not implemented"
	return
}

func (d *defaultWriter) Write(writer util.BufWriter, source []byte) {
	_ = "STUB: not implemented"
	return
}

var DefaultWriter = NewWriter()

var bDataImage = []byte("data:image/")
var bPng = []byte("png;")
var bGif = []byte("gif;")
var bJpeg = []byte("jpeg;")
var bWebp = []byte("webp;")
var bJs = []byte("javascript:")
var bVb = []byte("vbscript:")
var bFile = []byte("file:")
var bData = []byte("data:")

func hasPrefix(s, prefix []byte) bool { _ = "STUB: not implemented"; return false }

func IsDangerousURL(url []byte) bool { _ = "STUB: not implemented"; return false }
