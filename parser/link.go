package parser

import (
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/text"
)

var linkLabelStateKey = NewContextKey()

type linkLabelState struct {
	ast.BaseInline

	Segment text.Segment

	IsImage bool

	Prev *linkLabelState

	Next *linkLabelState

	First *linkLabelState

	Last *linkLabelState
}

func newLinkLabelState(segment text.Segment, isImage bool) *linkLabelState {
	_ = "STUB: not implemented"
	return nil
}

func (s *linkLabelState) Text(source []byte) []byte { _ = "STUB: not implemented"; return nil }

func (s *linkLabelState) Dump(source []byte, level int) { _ = "STUB: not implemented"; return }

var kindLinkLabelState = ast.NewNodeKind("LinkLabelState")

func (s *linkLabelState) Kind() ast.NodeKind { _ = "STUB: not implemented"; return *new(ast.NodeKind) }

func linkLabelStateLength(v *linkLabelState) int { _ = "STUB: not implemented"; return 0 }

func pushLinkLabelState(pc Context, v *linkLabelState) { _ = "STUB: not implemented"; return }

func removeLinkLabelState(pc Context, d *linkLabelState) { _ = "STUB: not implemented"; return }

type linkParser struct {
}

var defaultLinkParser = &linkParser{}

func NewLinkParser() InlineParser { _ = "STUB: not implemented"; return *new(InlineParser) }

func (s *linkParser) Trigger() []byte { _ = "STUB: not implemented"; return nil }

var linkBottom = NewContextKey()

func (s *linkParser) Parse(parent ast.Node, block text.Reader, pc Context) ast.Node {
	_ = "STUB: not implemented"
	return *new(ast.Node)
}

func (s *linkParser) containsLink(n ast.Node) bool { _ = "STUB: not implemented"; return false }

func processLinkLabelOpen(block text.Reader, pos int, isImage bool, pc Context) *linkLabelState {
	_ = "STUB: not implemented"
	return nil
}

func (s *linkParser) processLinkLabel(parent ast.Node, link *ast.Link, last *linkLabelState, pc Context) {
	_ = "STUB: not implemented"
	return
}

var linkFindClosureOptions text.FindClosureOptions = text.FindClosureOptions{
	Nesting: false,
	Newline: true,
	Advance: true,
}

func (s *linkParser) parseReferenceLink(parent ast.Node, last *linkLabelState,
	block text.Reader, pc Context) (*ast.Link, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (s *linkParser) parseLink(parent ast.Node, last *linkLabelState, block text.Reader, pc Context) *ast.Link {
	_ = "STUB: not implemented"
	return nil
}

func parseLinkDestination(block text.Reader) ([]byte, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func parseLinkTitle(block text.Reader) ([]byte, bool) { _ = "STUB: not implemented"; return nil, false }

func pushLinkBottom(pc Context) { _ = "STUB: not implemented"; return }

func popLinkBottom(pc Context) ast.Node { _ = "STUB: not implemented"; return *new(ast.Node) }

func (s *linkParser) CloseBlock(parent ast.Node, block text.Reader, pc Context) {
	_ = "STUB: not implemented"
	return
}
