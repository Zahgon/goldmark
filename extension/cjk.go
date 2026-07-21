package extension

import (
	"github.com/yuin/goldmark"
)

type CJKOption func(*cjk)

type EastAsianLineBreaks int

const (
	EastAsianLineBreaksNone EastAsianLineBreaks = iota

	EastAsianLineBreaksSimple

	EastAsianLineBreaksCSS3Draft
)

func WithEastAsianLineBreaks(style ...EastAsianLineBreaks) CJKOption {
	_ = "STUB: not implemented"
	return *new(CJKOption)
}

func WithEscapedSpace() CJKOption { _ = "STUB: not implemented"; return *new(CJKOption) }

type cjk struct {
	EastAsianLineBreaks EastAsianLineBreaks
	EscapedSpace        bool
}

var CJK = NewCJK(WithEastAsianLineBreaks(), WithEscapedSpace())

func NewCJK(opts ...CJKOption) goldmark.Extender {
	_ = "STUB: not implemented"
	return *new(goldmark.Extender)
}

func (e *cjk) Extend(m goldmark.Markdown) { _ = "STUB: not implemented"; return }
