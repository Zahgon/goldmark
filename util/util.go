package util

import (
	"io"
	"regexp"
)

type CopyOnWriteBuffer struct {
	buffer []byte
	copied bool
}

func NewCopyOnWriteBuffer(buffer []byte) CopyOnWriteBuffer {
	_ = "STUB: not implemented"
	return *new(CopyOnWriteBuffer)
}

func (b *CopyOnWriteBuffer) Write(value []byte) { _ = "STUB: not implemented"; return }

func (b *CopyOnWriteBuffer) WriteString(value string) { _ = "STUB: not implemented"; return }

func (b *CopyOnWriteBuffer) Append(value []byte) { _ = "STUB: not implemented"; return }

func (b *CopyOnWriteBuffer) AppendString(value string) { _ = "STUB: not implemented"; return }

func (b *CopyOnWriteBuffer) WriteByte(c byte) error { _ = "STUB: not implemented"; return nil }

func (b *CopyOnWriteBuffer) AppendByte(c byte) { _ = "STUB: not implemented"; return }

func (b *CopyOnWriteBuffer) Bytes() []byte { _ = "STUB: not implemented"; return nil }

func (b *CopyOnWriteBuffer) IsCopied() bool { _ = "STUB: not implemented"; return false }

func IsEscapedPunctuation(source []byte, i int) bool { _ = "STUB: not implemented"; return false }

func ReadWhile(source []byte, index [2]int, pred func(byte) bool) (int, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

func IsBlank(bs []byte) bool { _ = "STUB: not implemented"; return false }

func VisualizeSpaces(bs []byte) []byte { _ = "STUB: not implemented"; return nil }

func TabWidth(currentPos int) int { _ = "STUB: not implemented"; return 0 }

func IndentPosition(bs []byte, currentPos, width int) (pos, padding int) {
	_ = "STUB: not implemented"
	return 0, 0
}

func IndentPositionPadding(bs []byte, currentPos, paddingv, width int) (pos, padding int) {
	_ = "STUB: not implemented"
	return 0, 0
}

func DedentPosition(bs []byte, currentPos, width int) (pos, padding int) {
	_ = "STUB: not implemented"
	return 0, 0
}

func DedentPositionPadding(bs []byte, currentPos, paddingv, width int) (pos, padding int) {
	_ = "STUB: not implemented"
	return 0, 0
}

func IndentWidth(bs []byte, currentPos int) (width, pos int) {
	_ = "STUB: not implemented"
	return 0, 0
}

func FirstNonSpacePosition(bs []byte) int { _ = "STUB: not implemented"; return 0 }

func FindClosure(bs []byte, opener, closure byte, codeSpan, allowNesting bool) int {
	_ = "STUB: not implemented"
	return 0
}

func TrimLeft(source, b []byte) []byte { _ = "STUB: not implemented"; return nil }

func TrimRight(source, b []byte) []byte { _ = "STUB: not implemented"; return nil }

func TrimLeftLength(source, s []byte) int { _ = "STUB: not implemented"; return 0 }

func TrimRightLength(source, s []byte) int { _ = "STUB: not implemented"; return 0 }

func TrimLeftSpaceLength(source []byte) int { _ = "STUB: not implemented"; return 0 }

func TrimRightSpaceLength(source []byte) int { _ = "STUB: not implemented"; return 0 }

func TrimLeftSpace(source []byte) []byte { _ = "STUB: not implemented"; return nil }

func TrimRightSpace(source []byte) []byte { _ = "STUB: not implemented"; return nil }

func DoFullUnicodeCaseFolding(v []byte) []byte { _ = "STUB: not implemented"; return nil }

func ReplaceSpaces(source []byte, repl byte) []byte { _ = "STUB: not implemented"; return nil }

func ToRune(source []byte, pos int) rune { _ = "STUB: not implemented"; return 0 }

func ToValidRune(v rune) rune { _ = "STUB: not implemented"; return 0 }

func ToLinkReference(v []byte) string { _ = "STUB: not implemented"; return "" }

var htmlQuote = []byte("&quot;")
var htmlAmp = []byte("&amp;")
var htmlLess = []byte("&lt;")
var htmlGreater = []byte("&gt;")
var htmlNull = []byte("\ufffd")

var htmlEscapeTable = [256]*[]byte{&htmlNull, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, &htmlQuote, nil, nil, nil, &htmlAmp, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, &htmlLess, nil, &htmlGreater, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil} //nolint:golint,lll

func EscapeHTMLByte(b byte) []byte { _ = "STUB: not implemented"; return nil }

func EscapeHTML(v []byte) []byte { _ = "STUB: not implemented"; return nil }

func UnescapePunctuations(source []byte) []byte { _ = "STUB: not implemented"; return nil }

func ResolveNumericReferences(source []byte) []byte { _ = "STUB: not implemented"; return nil }

func ResolveEntityNames(source []byte) []byte { _ = "STUB: not implemented"; return nil }

var htmlSpace = []byte("%20")

func URLEscape(v []byte, resolveReference bool) []byte { _ = "STUB: not implemented"; return nil }

func FindURLIndex(b []byte) int { _ = "STUB: not implemented"; return 0 }

var emailDomainRegexp = regexp.MustCompile(`^[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*`) //nolint:golint,lll

func FindEmailIndex(b []byte) int { _ = "STUB: not implemented"; return 0 }

var spaces = []byte(" \t\n\x0b\x0c\x0d")

var spaceTable = [256]int8{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0} //nolint:golint,lll

var punctTable = [256]int8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0} //nolint:golint,lll

var urlEscapeTable = [256]int8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0} //nolint:golint,lll

var utf8lenTable = [256]int8{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4, 4, 4, 99, 99, 99, 99, 99, 99, 99, 99} //nolint:golint,lll

var urlTable = [256]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 5, 5, 1, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 1, 1, 0, 1, 0, 1, 1, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 1, 1, 1, 1, 1, 1, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1} //nolint:golint,lll

var emailTable = [256]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, 1, 1, 1, 1, 0, 0, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0} //nolint:golint,lll

func UTF8Len(b byte) int8 { _ = "STUB: not implemented"; return 0 }

func IsPunct(c byte) bool { _ = "STUB: not implemented"; return false }

func IsPunctRune(r rune) bool { _ = "STUB: not implemented"; return false }

func IsSpace(c byte) bool { _ = "STUB: not implemented"; return false }

func IsSpaceRune(r rune) bool { _ = "STUB: not implemented"; return false }

func IsNumeric(c byte) bool { _ = "STUB: not implemented"; return false }

func IsHexDecimal(c byte) bool { _ = "STUB: not implemented"; return false }

func IsAlphaNumeric(c byte) bool { _ = "STUB: not implemented"; return false }

type BufWriter interface {
	io.Writer
	Available() int
	Buffered() int
	Flush() error
	WriteByte(c byte) error
	WriteRune(r rune) (size int, err error)
	WriteString(s string) (int, error)
}

type PrioritizedValue struct {
	Value any

	Priority int
}

type PrioritizedSlice []PrioritizedValue

func (s PrioritizedSlice) Sort() { _ = "STUB: not implemented"; return }

func (s PrioritizedSlice) Remove(v any) PrioritizedSlice {
	_ = "STUB: not implemented"
	return *new(PrioritizedSlice)
}

func Prioritized(v any, priority int) PrioritizedValue {
	_ = "STUB: not implemented"
	return *new(PrioritizedValue)
}

func bytesHash(b []byte) uint64 { _ = "STUB: not implemented"; return 0 }

type BytesFilter interface {
	Add([]byte)

	Contains([]byte) bool

	Extend(...[]byte) BytesFilter

	ExtendString(string) BytesFilter
}

type bytesFilter struct {
	chars     [256]uint8
	threshold int
	slots     [][][]byte
}

func NewBytesFilter(elements ...[]byte) BytesFilter {
	_ = "STUB: not implemented"
	return *new(BytesFilter)
}

func NewBytesFilterString(elements string) BytesFilter {
	_ = "STUB: not implemented"
	return *new(BytesFilter)
}

func (s *bytesFilter) Add(b []byte) { _ = "STUB: not implemented"; return }

func (s *bytesFilter) Extend(bs ...[]byte) BytesFilter {
	_ = "STUB: not implemented"
	return *new(BytesFilter)
}

func (s *bytesFilter) ExtendString(elements string) BytesFilter {
	_ = "STUB: not implemented"
	return *new(BytesFilter)
}

func (s *bytesFilter) Contains(b []byte) bool { _ = "STUB: not implemented"; return false }
