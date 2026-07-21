package testutil

import (
	"regexp"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
)

type TestingT interface {
	Logf(string, ...any)
	Skipf(string, ...any)
	Errorf(string, ...any)
	FailNow()
}

type MarkdownTestCase struct {
	No          int
	Description string
	Options     MarkdownTestCaseOptions
	Markdown    string
	Expected    string
}

func source(t *MarkdownTestCase) string { _ = "STUB: not implemented"; return "" }

func expected(t *MarkdownTestCase) string { _ = "STUB: not implemented"; return "" }

type MarkdownTestCaseOptions struct {
	EnableEscape bool
	Trim         bool
}

const attributeSeparator = "//- - - - - - - - -//"
const caseSeparator = "//= = = = = = = = = = = = = = = = = = = = = = = =//"

var optionsRegexp = regexp.MustCompile(`(?i)\s*options:(.*)`)

func ParseCliCaseArg() []int { _ = "STUB: not implemented"; return nil }

func DoTestCaseFile(m goldmark.Markdown, filename string, t TestingT, no ...int) {
	_ = "STUB: not implemented"
	return
}

func DoTestCases(m goldmark.Markdown, cases []MarkdownTestCase, t TestingT, opts ...parser.ParseOption) {
	_ = "STUB: not implemented"
	return
}

func DoTestCase(m goldmark.Markdown, testCase MarkdownTestCase, t TestingT, opts ...parser.ParseOption) {
	_ = "STUB: not implemented"
	return
}

type diffType int

const (
	diffRemoved diffType = iota
	diffAdded
	diffNone
)

type diff struct {
	Type  diffType
	Lines [][]byte
}

func simpleDiff(v1, v2 []byte) []diff { _ = "STUB: not implemented"; return nil }

func simpleDiffAux(v1lines, v2lines [][]byte) []diff { _ = "STUB: not implemented"; return nil }

func DiffPretty(v1, v2 []byte) []byte { _ = "STUB: not implemented"; return nil }

func applyEscapeSequence(b []byte) []byte { _ = "STUB: not implemented"; return nil }
