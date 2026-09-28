//go:build !solution

package utf8reverse

import "unicode/utf8"

// Reverse returns input runes in reverse order.
func Reverse(input string) string {
	runes := make([]rune, 0, utf8.RuneCountInString(input))

	for len(input) > 0 {
		r, size := utf8.DecodeRuneInString(input)
		runes = append(runes, r)
		input = input[size:]
	}

	n := len(runes)
	for j := 0; j < n; j++ {
		runes[j], runes[n-j-1] = runes[n-j-1], runes[j]
	}

	return string(runes)
}
