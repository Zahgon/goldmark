package util

import "unicode"

var cjkRadicalsSupplement = &unicode.RangeTable{
	R16: []unicode.Range16{
		{0x2E80, 0x2EFF, 1},
	},
}

var kangxiRadicals = &unicode.RangeTable{
	R16: []unicode.Range16{
		{0x2F00, 0x2FDF, 1},
	},
}

var ideographicDescriptionCharacters = &unicode.RangeTable{
	R16: []unicode.Range16{
		{0x2FF0, 0x2FFF, 1},
	},
}

var cjkSymbolsAndPunctuation = &unicode.RangeTable{
	R16: []unicode.Range16{
		{0x3000, 0x303F, 1},
	},
}

var hiragana = &unicode.RangeTable{
	R16: []unicode.Range16{
		{0x3040, 0x309F, 1},
	},
}

var katakana = &unicode.RangeTable{
	R16: []unicode.Range16{
		{0x30A0, 0x30FF, 1},
	},
}

var kanbun = &unicode.RangeTable{
	R16: []unicode.Range16{
		{0x3130, 0x318F, 1},
		{0x3190, 0x319F, 1},
	},
}

var cjkStrokes = &unicode.RangeTable{
	R16: []unicode.Range16{
		{0x31C0, 0x31EF, 1},
	},
}

var katakanaPhoneticExtensions = &unicode.RangeTable{
	R16: []unicode.Range16{
		{0x31F0, 0x31FF, 1},
	},
}

var cjkCompatibility = &unicode.RangeTable{
	R16: []unicode.Range16{
		{0x3300, 0x33FF, 1},
	},
}

var cjkUnifiedIdeographsExtensionA = &unicode.RangeTable{
	R16: []unicode.Range16{
		{0x3400, 0x4DBF, 1},
	},
}

var cjkUnifiedIdeographs = &unicode.RangeTable{
	R16: []unicode.Range16{
		{0x4E00, 0x9FFF, 1},
	},
}

var yiSyllables = &unicode.RangeTable{
	R16: []unicode.Range16{
		{0xA000, 0xA48F, 1},
	},
}

var yiRadicals = &unicode.RangeTable{
	R16: []unicode.Range16{
		{0xA490, 0xA4CF, 1},
	},
}

var cjkCompatibilityIdeographs = &unicode.RangeTable{
	R16: []unicode.Range16{
		{0xF900, 0xFAFF, 1},
	},
}

var verticalForms = &unicode.RangeTable{
	R16: []unicode.Range16{
		{0xFE10, 0xFE1F, 1},
	},
}

var cjkCompatibilityForms = &unicode.RangeTable{
	R16: []unicode.Range16{
		{0xFE30, 0xFE4F, 1},
	},
}

var smallFormVariants = &unicode.RangeTable{
	R16: []unicode.Range16{
		{0xFE50, 0xFE6F, 1},
	},
}

var halfwidthAndFullwidthForms = &unicode.RangeTable{
	R16: []unicode.Range16{
		{0xFF00, 0xFFEF, 1},
	},
}

var kanaSupplement = &unicode.RangeTable{
	R32: []unicode.Range32{
		{0x1B000, 0x1B0FF, 1},
	},
}

var kanaExtendedA = &unicode.RangeTable{
	R32: []unicode.Range32{
		{0x1B100, 0x1B12F, 1},
	},
}

var smallKanaExtension = &unicode.RangeTable{
	R32: []unicode.Range32{
		{0x1B130, 0x1B16F, 1},
	},
}

var cjkUnifiedIdeographsExtensionB = &unicode.RangeTable{
	R32: []unicode.Range32{
		{0x20000, 0x2A6DF, 1},
	},
}

var cjkUnifiedIdeographsExtensionC = &unicode.RangeTable{
	R32: []unicode.Range32{
		{0x2A700, 0x2B73F, 1},
	},
}

var cjkUnifiedIdeographsExtensionD = &unicode.RangeTable{
	R32: []unicode.Range32{
		{0x2B740, 0x2B81F, 1},
	},
}

var cjkUnifiedIdeographsExtensionE = &unicode.RangeTable{
	R32: []unicode.Range32{
		{0x2B820, 0x2CEAF, 1},
	},
}

var cjkUnifiedIdeographsExtensionF = &unicode.RangeTable{
	R32: []unicode.Range32{
		{0x2CEB0, 0x2EBEF, 1},
	},
}

var cjkCompatibilityIdeographsSupplement = &unicode.RangeTable{
	R32: []unicode.Range32{
		{0x2F800, 0x2FA1F, 1},
	},
}

var cjkUnifiedIdeographsExtensionG = &unicode.RangeTable{
	R32: []unicode.Range32{
		{0x30000, 0x3134F, 1},
	},
}

func IsEastAsianWideRune(r rune) bool { _ = "STUB: not implemented"; return false }

func IsSpaceDiscardingUnicodeRune(r rune) bool { _ = "STUB: not implemented"; return false }

func EastAsianWidth(r rune) string { _ = "STUB: not implemented"; return "" }
