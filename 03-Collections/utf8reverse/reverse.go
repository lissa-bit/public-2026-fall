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
	for i := 0; i < n; i++ {
		runes[i], runes[n-i+1] = runes[n-i+1], runes[i]
	}

	return string(runes)
}
