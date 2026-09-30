//go:build !solution

package speller

import "strings"

// Spell returns the American English spelling of n.
func Spell(n int64) string {
	if n == 0 {
		return "zero"
	}

	var prefix string
	if n < 0 {
		prefix = "minus "
		n = -n
	}

	var parts []string

	for i := 0; i < len(scales); i++ {
		if n >= scales[i].value {
			count := n / scales[i].value
			parts = append(parts, spellUnderThousand(count)+" "+scales[i].word)
			n %= scales[i].value
		}
	}

	if n > 0 {
		parts = append(parts, spellUnderThousand(n))
	}

	return prefix + strings.Join(parts, " ")
}

func spellUnderThousand(n int64) string {
	if n == 0 {
		return ""
	}

	if n < 20 {
		return smallNumbers[n]
	}

	if n < 100 {
		ten := n / 10
		ost := n % 10

		if ost == 0 {
			return tens[ten]
		}

		return tens[ten] + "-" + smallNumbers[ost]
	}

	hundred := n / 100
	ost := n % 100

	if ost == 0 {
		return smallNumbers[hundred] + " hundred"
	}

	return smallNumbers[hundred] + " hundred " + spellUnderThousand(ost)
}
