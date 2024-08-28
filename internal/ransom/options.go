package ransom

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
		for b := 'a'; b <= 'z'; b++ {
			r.pre = append(r.pre, string(b), ":"+string(b)+":")
			if prefix != "" {
				r.post = append(r.post, ":"+string(b)+":", ":"+prefix+string(b)+":")
			}
		}
	}
}
