package extension

import (
	"regexp"

	"github.com/yuin/goldmark"
	gast "github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/extension/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

var escapedPipeCellListKey = parser.NewContextKey()

type escapedPipeCell struct {
	Cell        *ast.TableCell
	Pos         []int
	Transformed bool
}

type TableCellAlignMethod int

const (
	TableCellAlignDefault TableCellAlignMethod = iota

	TableCellAlignAttribute

	TableCellAlignStyle

	TableCellAlignNone
)

type TableConfig struct {
	html.Config

	TableCellAlignMethod TableCellAlignMethod
}

type TableOption interface {
	renderer.Option

	SetTableOption(*TableConfig)
}

func NewTableConfig() TableConfig { _ = "STUB: not implemented"; return *new(TableConfig) }

func (c *TableConfig) SetOption(name renderer.OptionName, value any) {
	_ = "STUB: not implemented"
	return
}

type withTableHTMLOptions struct {
	value []html.Option
}

func (o *withTableHTMLOptions) SetConfig(c *renderer.Config) { _ = "STUB: not implemented"; return }

func (o *withTableHTMLOptions) SetTableOption(c *TableConfig) { _ = "STUB: not implemented"; return }

func WithTableHTMLOptions(opts ...html.Option) TableOption {
	_ = "STUB: not implemented"
	return *new(TableOption)
}

const optTableCellAlignMethod renderer.OptionName = "TableTableCellAlignMethod"

type withTableCellAlignMethod struct {
	value TableCellAlignMethod
}

func (o *withTableCellAlignMethod) SetConfig(c *renderer.Config) { _ = "STUB: not implemented"; return }

func (o *withTableCellAlignMethod) SetTableOption(c *TableConfig) {
	_ = "STUB: not implemented"
	return
}

func WithTableCellAlignMethod(a TableCellAlignMethod) TableOption {
	_ = "STUB: not implemented"
	return *new(TableOption)
}

func isTableDelim(bs []byte) bool { _ = "STUB: not implemented"; return false }

var tableDelimLeft = regexp.MustCompile(`^\s*\:\-+\s*$`)
var tableDelimRight = regexp.MustCompile(`^\s*\-+\:\s*$`)
var tableDelimCenter = regexp.MustCompile(`^\s*\:\-+\:\s*$`)
var tableDelimNone = regexp.MustCompile(`^\s*\-+\s*$`)

type tableParagraphTransformer struct {
}

var defaultTableParagraphTransformer = &tableParagraphTransformer{}

func NewTableParagraphTransformer() parser.ParagraphTransformer {
	_ = "STUB: not implemented"
	return *new(parser.ParagraphTransformer)
}

func (b *tableParagraphTransformer) Transform(node *gast.Paragraph, reader text.Reader, pc parser.Context) {
	_ = "STUB: not implemented"
	return
}

func (b *tableParagraphTransformer) parseRow(segment text.Segment,
	alignments []ast.Alignment, isHeader bool, reader text.Reader, pc parser.Context) *ast.TableRow {
	_ = "STUB: not implemented"
	return nil
}

func (b *tableParagraphTransformer) parseDelimiter(segment text.Segment, reader text.Reader) []ast.Alignment {
	_ = "STUB: not implemented"
	return nil
}

type tableASTTransformer struct {
}

var defaultTableASTTransformer = &tableASTTransformer{}

func NewTableASTTransformer() parser.ASTTransformer {
	_ = "STUB: not implemented"
	return *new(parser.ASTTransformer)
}

func (a *tableASTTransformer) Transform(node *gast.Document, reader text.Reader, pc parser.Context) {
	_ = "STUB: not implemented"
	return
}

type TableHTMLRenderer struct {
	TableConfig
}

func NewTableHTMLRenderer(opts ...TableOption) renderer.NodeRenderer {
	_ = "STUB: not implemented"
	return *new(renderer.NodeRenderer)
}

func (r *TableHTMLRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	_ = "STUB: not implemented"
	return
}

var TableAttributeFilter = html.GlobalAttributeFilter.ExtendString(`align,bgcolor,border,cellpadding,cellspacing,frame,rules,summary,width`)

func (r *TableHTMLRenderer) renderTable(
	w util.BufWriter, source []byte, n gast.Node, entering bool) (gast.WalkStatus, error) {
	_ = "STUB: not implemented"
	return *new(gast.WalkStatus), nil
}

var TableHeaderAttributeFilter = html.GlobalAttributeFilter.ExtendString(`align,bgcolor,char,charoff,valign`)

func (r *TableHTMLRenderer) renderTableHeader(
	w util.BufWriter, source []byte, n gast.Node, entering bool) (gast.WalkStatus, error) {
	_ = "STUB: not implemented"
	return *new(gast.WalkStatus), nil
}

var TableRowAttributeFilter = html.GlobalAttributeFilter.ExtendString(`align,bgcolor,char,charoff,valign`)

func (r *TableHTMLRenderer) renderTableRow(
	w util.BufWriter, source []byte, n gast.Node, entering bool) (gast.WalkStatus, error) {
	_ = "STUB: not implemented"
	return *new(gast.WalkStatus), nil
}

var TableThCellAttributeFilter = html.GlobalAttributeFilter.ExtendString(`abbr,align,axis,bgcolor,char,charoff,colspan,headers,height,rowspan,scope,valign,width`)

var TableTdCellAttributeFilter = html.GlobalAttributeFilter.ExtendString(`abbr,align,axis,bgcolor,char,charoff,colspan,headers,height,rowspan,scope,valign,width`)

func (r *TableHTMLRenderer) renderTableCell(
	w util.BufWriter, source []byte, node gast.Node, entering bool) (gast.WalkStatus, error) {
	_ = "STUB: not implemented"
	return *new(gast.WalkStatus), nil
}

type table struct {
	options []TableOption
}

var Table = &table{
	options: []TableOption{},
}

func NewTable(opts ...TableOption) goldmark.Extender {
	_ = "STUB: not implemented"
	return *new(goldmark.Extender)
}

func (e *table) Extend(m goldmark.Markdown) { _ = "STUB: not implemented"; return }
