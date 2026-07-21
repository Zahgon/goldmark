package parser

import (
	"io"

	"github.com/yuin/goldmark/text"
)

var attrNameID = []byte("id")
var attrNameClass = []byte("class")

type Attribute struct {
	Name  []byte
	Value any
}

type Attributes []Attribute

func (as Attributes) Find(name []byte) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

func (as Attributes) findUpdate(name []byte, cb func(v any) any) bool {
	_ = "STUB: not implemented"
	return false
}

func ParseAttributes(reader text.Reader) (Attributes, bool) {
	_ = "STUB: not implemented"
	return *new(Attributes), false
}

func parseAttribute(reader text.Reader) (Attribute, bool) {
	_ = "STUB: not implemented"
	return *new(Attribute), false
}

func parseAttributeValue(reader text.Reader) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

func parseAttributeArray(reader text.Reader) ([]any, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func parseAttributeString(reader text.Reader) ([]byte, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func scanAttributeDecimal(reader text.Reader, w io.ByteWriter) { _ = "STUB: not implemented"; return }

func parseAttributeNumber(reader text.Reader) (float64, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

var bytesTrue = []byte("true")
var bytesFalse = []byte("false")
var bytesNull = []byte("null")

func parseAttributeOthers(reader text.Reader) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}
