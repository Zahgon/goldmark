package text

import (
	"io"
	"regexp"
)

const invalidValue = -1

const EOF = byte(0xff)

type Reader interface {
	io.RuneReader

	Source() []byte

	ResetPosition()

	Peek() byte

	PeekLine() ([]byte, Segment)

	PrecendingCharacter() rune

	Value(Segment) []byte

	LineOffset() int

	Position() (int, Segment)

	SetPosition(int, Segment)

	SetPadding(int)

	Advance(int)

	AdvanceAndSetPadding(int, int)

	AdvanceToEOL()

	AdvanceLine()

	SkipSpaces() (Segment, int, bool)

	SkipBlankLines() (Segment, int, bool)

	Match(reg *regexp.Regexp) bool

	FindSubMatch(reg *regexp.Regexp) [][]byte

	FindClosure(opener, closer byte, options FindClosureOptions) (*Segments, bool)
}

type FindClosureOptions struct {
	CodeSpan bool

	Nesting bool

	Newline bool

	Advance bool
}

type reader struct {
	source       []byte
	sourceLength int
	line         int
	peekedLine   []byte
	pos          Segment
	head         int
	lineOffset   int
}

func NewReader(source []byte) Reader { _ = "STUB: not implemented"; return *new(Reader) }

func (r *reader) FindClosure(opener, closer byte, options FindClosureOptions) (*Segments, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (r *reader) ResetPosition() { _ = "STUB: not implemented"; return }

func (r *reader) Source() []byte { _ = "STUB: not implemented"; return nil }

func (r *reader) Value(seg Segment) []byte { _ = "STUB: not implemented"; return nil }

func (r *reader) Peek() byte { _ = "STUB: not implemented"; return 0 }

func (r *reader) PeekLine() ([]byte, Segment) { _ = "STUB: not implemented"; return nil, *new(Segment) }

func (r *reader) ReadRune() (rune, int, error) { _ = "STUB: not implemented"; return 0, 0, nil }

func (r *reader) LineOffset() int { _ = "STUB: not implemented"; return 0 }

func (r *reader) PrecendingCharacter() rune { _ = "STUB: not implemented"; return 0 }

func (r *reader) Advance(n int) { _ = "STUB: not implemented"; return }

func (r *reader) AdvanceAndSetPadding(n, padding int) { _ = "STUB: not implemented"; return }

func (r *reader) AdvanceToEOL() { _ = "STUB: not implemented"; return }

func (r *reader) AdvanceLine() { _ = "STUB: not implemented"; return }

func (r *reader) Position() (int, Segment) { _ = "STUB: not implemented"; return 0, *new(Segment) }

func (r *reader) SetPosition(line int, pos Segment) { _ = "STUB: not implemented"; return }

func (r *reader) SetPadding(v int) { _ = "STUB: not implemented"; return }

func (r *reader) SkipSpaces() (Segment, int, bool) {
	_ = "STUB: not implemented"
	return *new(Segment), 0, false
}

func (r *reader) SkipBlankLines() (Segment, int, bool) {
	_ = "STUB: not implemented"
	return *new(Segment), 0, false
}

func (r *reader) Match(reg *regexp.Regexp) bool { _ = "STUB: not implemented"; return false }

func (r *reader) FindSubMatch(reg *regexp.Regexp) [][]byte { _ = "STUB: not implemented"; return nil }

type BlockReader interface {
	Reader

	Reset(segment *Segments)
}

type blockReader struct {
	source         []byte
	segments       *Segments
	segmentsLength int
	line           int
	pos            Segment
	head           int
	last           int
	lineOffset     int
}

func NewBlockReader(source []byte, segments *Segments) BlockReader {
	_ = "STUB: not implemented"
	return *new(BlockReader)
}

func (r *blockReader) FindClosure(opener, closer byte, options FindClosureOptions) (*Segments, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (r *blockReader) ResetPosition() { _ = "STUB: not implemented"; return }

func (r *blockReader) Reset(segments *Segments) { _ = "STUB: not implemented"; return }

func (r *blockReader) Source() []byte { _ = "STUB: not implemented"; return nil }

func (r *blockReader) Value(seg Segment) []byte { _ = "STUB: not implemented"; return nil }

func (r *blockReader) ReadRune() (rune, int, error) { _ = "STUB: not implemented"; return 0, 0, nil }

func (r *blockReader) PrecendingCharacter() rune { _ = "STUB: not implemented"; return 0 }

func (r *blockReader) LineOffset() int { _ = "STUB: not implemented"; return 0 }

func (r *blockReader) Peek() byte { _ = "STUB: not implemented"; return 0 }

func (r *blockReader) PeekLine() ([]byte, Segment) {
	_ = "STUB: not implemented"
	return nil, *new(Segment)
}

func (r *blockReader) Advance(n int) { _ = "STUB: not implemented"; return }

func (r *blockReader) AdvanceAndSetPadding(n, padding int) { _ = "STUB: not implemented"; return }

func (r *blockReader) AdvanceToEOL() { _ = "STUB: not implemented"; return }

func (r *blockReader) AdvanceLine() { _ = "STUB: not implemented"; return }

func (r *blockReader) Position() (int, Segment) { _ = "STUB: not implemented"; return 0, *new(Segment) }

func (r *blockReader) SetPosition(line int, pos Segment) { _ = "STUB: not implemented"; return }

func (r *blockReader) SetPadding(v int) { _ = "STUB: not implemented"; return }

func (r *blockReader) SkipSpaces() (Segment, int, bool) {
	_ = "STUB: not implemented"
	return *new(Segment), 0, false
}

func (r *blockReader) SkipBlankLines() (Segment, int, bool) {
	_ = "STUB: not implemented"
	return *new(Segment), 0, false
}

func (r *blockReader) Match(reg *regexp.Regexp) bool { _ = "STUB: not implemented"; return false }

func (r *blockReader) FindSubMatch(reg *regexp.Regexp) [][]byte {
	_ = "STUB: not implemented"
	return nil
}

func skipBlankLinesReader(r Reader) (Segment, int, bool) {
	_ = "STUB: not implemented"
	return *new(Segment), 0, false
}

func skipSpacesReader(r Reader) (Segment, int, bool) {
	_ = "STUB: not implemented"
	return *new(Segment), 0, false
}

func matchReader(r Reader, reg *regexp.Regexp) bool { _ = "STUB: not implemented"; return false }

func findSubMatchReader(r Reader, reg *regexp.Regexp) [][]byte {
	_ = "STUB: not implemented"
	return nil
}

func readRuneReader(r Reader) (rune, int, error) { _ = "STUB: not implemented"; return 0, 0, nil }

func findClosureReader(r Reader, opener, closer byte, opts FindClosureOptions) (*Segments, bool) {
	_ = "STUB: not implemented"
	return nil, false
}
