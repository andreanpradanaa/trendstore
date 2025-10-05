package shared

import (
	"strings"
	"unicode"
)

// GenerateSlug will create a URL-friendly slug
// by converting the input string to lowercase,
// replacing spaces with hyphens, and removing
// special characters.
// e.g. "Sample Product Name" -> "sample-product-name"
// https://goplay.tools/snippet/3c0Pbu0yrCf
func GenerateSlug(name string) string {
	if name == "" {
		return ""
	}

	var result strings.Builder
	writeHyphen := false

	for _, r := range name {
		switch {
		case unicode.IsLetter(r) || unicode.IsDigit(r):
			if writeHyphen {
				result.WriteByte('-')
				writeHyphen = false
			}
			result.WriteRune(unicode.ToLower(r))
		case unicode.IsSpace(r) || r == '-' || r == '_':
			writeHyphen = result.Len() > 0
		}
	}

	return result.String()
}
