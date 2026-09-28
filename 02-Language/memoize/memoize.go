//go:build !solution

package memoize

// Memoize returns a lazy function that caches the result of source.
func Memoize(source Func) Func {
	if source == nil {
		panic("nil source")
	}

	var (
		res   bool
		value int
		err   error
	)

	return func() (int, error) {
		if !res {
			value, err = source()
			res = true
		}

		return value, err
	}
}
