//go:build !solution

package postfix

import (
	"unicode"
	"unicode/utf8"
)

// Eval evaluates an arithmetic expression in postfix notation.
func Eval(expression string) (int, error) {
	var values []int

	for len(expression) > 0 {
		r, size := utf8.DecodeRuneInString(expression)

		neg_num := false

		if r == '-' && len(expression) > size {
			r1, _ := utf8.DecodeRuneInString(expression[size:])
			neg_num = unicode.IsDigit(r1)
		}

		switch {
		case unicode.IsSpace(r):
			expression = expression[size:]
		case unicode.IsDigit(r) || neg_num:
			num := 0

			if neg_num {
				expression = expression[size:]
			}

			i := 0
			for i < len(expression) {
				r2, size2 := utf8.DecodeRuneInString(expression[i:])
				if !unicode.IsDigit(r2) {
					break
				}

				num = num*10 + int(r2-'0')
				i += size2
			}

			if neg_num {
				num = -num
			}

			values = append(values, num)
			expression = expression[i:]
		case r == '+' || r == '*' || r == '-' || r == '/':
			if len(values) < 2 {
				return 0, ErrInvalidExpression
			}

			n1 := values[len(values)-1]
			n2 := values[len(values)-2]
			values = values[:len(values)-2]

			var res int

			switch r {
			case '+':
				res = n1 + n2
			case '*':
				res = n1 * n2
			case '-':
				res = n2 - n1
			case '/':
				if n1 == 0 {
					return 0, ErrDivisionByZero
				}

				res = n2 / n1
			}

			values = append(values, res)
			expression = expression[size:]

		default:
			return 0, ErrUnknownToken
		}
	}

	if len(values) != 1 {
		return 0, ErrInvalidExpression
	}

	return values[0], nil
}
