package util

import "iter"

func Alphabet() iter.Seq[string] {
	return func(yield func(string) bool) {
		for b := 'a'; b <= 'z'; b++ {
			if !yield(string(b)) {
				return
			}
		}
	}
}
