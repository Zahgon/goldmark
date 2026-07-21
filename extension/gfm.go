package extension

import (
	"github.com/yuin/goldmark"
)

type gfm struct {
}

var GFM = &gfm{}

func (e *gfm) Extend(m goldmark.Markdown) { _ = "STUB: not implemented"; return }
