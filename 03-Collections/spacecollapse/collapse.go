//go:build !solution

package spacecollapse

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

// CollapseSpaces replaces each non-empty run of Unicode spaces with one ASCII space.
func CollapseSpaces(input string) string {
	var builder strings.Builder

	space := false

	for len(input) > 0 {
		r, size := utf8.DecodeRuneInString(input)
		input = input[size:]

		if unicode.IsSpace(r) {
			if !space {
				space = true

				builder.WriteByte(' ')
			}

			continue
		}

		builder.WriteRune(r)

		space = false
	}

	return builder.String()
}
