package parser

import (
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/text"
)

type DelimiterProcessor interface {
	IsDelimiter(byte) bool

	CanOpenCloser(opener, closer *Delimiter) bool

	OnMatch(consumes int) ast.Node
}

type Delimiter struct {
	ast.BaseInline

	Segment text.Segment

	CanOpen bool

	CanClose bool

	Length int

	OriginalLength int

	Char byte

	PreviousDelimiter *Delimiter

	NextDelimiter *Delimiter

	Processor DelimiterProcessor
}

func (d *Delimiter) Inline() { _ = "STUB: not implemented"; return }

func (d *Delimiter) Dump(source []byte, level int) { _ = "STUB: not implemented"; return }

var kindDelimiter = ast.NewNodeKind("Delimiter")

func (d *Delimiter) Kind() ast.NodeKind { _ = "STUB: not implemented"; return *new(ast.NodeKind) }

func (d *Delimiter) Text(source []byte) []byte { _ = "STUB: not implemented"; return nil }

func (d *Delimiter) ConsumeCharacters(n int) { _ = "STUB: not implemented"; return }

func (d *Delimiter) CalcComsumption(closer *Delimiter) int { _ = "STUB: not implemented"; return 0 }

func NewDelimiter(canOpen, canClose bool, length int, char byte, processor DelimiterProcessor) *Delimiter {
	_ = "STUB: not implemented"
	return nil
}

func ScanDelimiter(line []byte, before rune, minimum int, processor DelimiterProcessor) *Delimiter {
	_ = "STUB: not implemented"
	return nil
}

func ProcessDelimiters(bottom ast.Node, pc Context) { _ = "STUB: not implemented"; return }
