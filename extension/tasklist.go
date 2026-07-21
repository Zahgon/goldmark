package extension

import (
	"regexp"

	"github.com/yuin/goldmark"
	gast "github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

var taskListRegexp = regexp.MustCompile(`^\[([\sxX])\]\s*`)

type taskCheckBoxParser struct {
}

var defaultTaskCheckBoxParser = &taskCheckBoxParser{}

func NewTaskCheckBoxParser() parser.InlineParser {
	_ = "STUB: not implemented"
	return *new(parser.InlineParser)
}

func (s *taskCheckBoxParser) Trigger() []byte { _ = "STUB: not implemented"; return nil }

func (s *taskCheckBoxParser) Parse(parent gast.Node, block text.Reader, pc parser.Context) gast.Node {
	_ = "STUB: not implemented"
	return *new(gast.Node)
}

func (s *taskCheckBoxParser) CloseBlock(parent gast.Node, pc parser.Context) {
	_ = "STUB: not implemented"
	return
}

type TaskCheckBoxHTMLRenderer struct {
	html.Config
}

func NewTaskCheckBoxHTMLRenderer(opts ...html.Option) renderer.NodeRenderer {
	_ = "STUB: not implemented"
	return *new(renderer.NodeRenderer)
}

func (r *TaskCheckBoxHTMLRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	_ = "STUB: not implemented"
	return
}

func (r *TaskCheckBoxHTMLRenderer) renderTaskCheckBox(
	w util.BufWriter, source []byte, node gast.Node, entering bool) (gast.WalkStatus, error) {
	_ = "STUB: not implemented"
	return *new(gast.WalkStatus), nil
}

type taskList struct {
}

var TaskList = &taskList{}

func (e *taskList) Extend(m goldmark.Markdown) { _ = "STUB: not implemented"; return }
