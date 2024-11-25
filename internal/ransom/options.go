package ransom

import "gabe565.com/ransom/internal/util"

type Option func(r *Replacer)

func WithRaw(from, to string) Option {
	return func(r *Replacer) {
		r.WithRaw(from, to)
	}
}

func WithWord(s ...string) Option {
	return func(r *Replacer) {
		r.WithWord(s...)
	}
}

func With(s ...string) Option {
	return func(r *Replacer) {
		r.With(s...)
	}
}

func WithAlphabet(prefix string) Option {
	return func(r *Replacer) {
		for letter := range util.Alphabet() {
			r.pre = append(r.pre, letter, ":"+letter+":")
			if prefix != "" {
				r.post = append(r.post, ":"+letter+":", ":"+prefix+letter+":")
			}
		}
		r.clean = false
	}
}
