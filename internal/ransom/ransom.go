package ransom

import "strings"

const Space = "blank"

func Default(prefix string) *Replacer {
	return New(
		WithAlphabet(prefix),
		With(" ", Space),
		WithWord("on"),
		WithWord("back"),
		WithWord("end"),
		WithWord("soon"),
		WithWord("top"),
		With(".", "hole"),
		With("!?", "interrobang"),
		With("!!", "bangbang"),
		With("!", "exclamation"),
		With("?", "question"),
		With("0", "zero"),
		With("1", "one"),
		With("2", "two"),
		With("3", "three"),
		With("4", "four"),
		With("5", "five"),
		With("6", "six"),
		With("7", "seven"),
		With("8", "eight"),
		With("9", "nine"),
		With("#", "hash"),
		With("*", "keycap_star"),
		With("+", "heavy_plus_sign"),
		With("-", "heavy_minus_sign"),
		With("$", "heavy_dollar_sign"),
		With("=", "heavy_equals_sign"),
		With("^", "this"),
	)
}

func New(opts ...Option) *Replacer {
	r := &Replacer{
		pre:  make([]string, 0, len(opts)*2),
		loop: make([]string, 0, len(opts)*2),
	}
	for _, opt := range opts {
		opt(r)
	}
	return r
}

type Replacer struct {
	pre          []string
	preReplacer  *strings.Replacer
	loop         []string
	loopReplacer *strings.Replacer
	post         []string
	postReplacer *strings.Replacer
	clean        bool
}

func sep(s string) string {
	if len(s) == 0 {
		return ""
	}
	s = strings.ReplaceAll(s, "", "::")
	return s[1 : len(s)-1]
}

func (r *Replacer) WithRaw(from, to string) *Replacer {
	r.loop = append(r.loop, from, to)
	r.clean = false
	return r
}

func (r *Replacer) WithWord(s ...string) *Replacer {
	var from, to string
	switch len(s) {
	case 0:
		return r
	case 1:
		from = Space + ":" + sep(s[0]) + ":" + Space
		to = Space + "::" + s[0] + "::" + Space
	case 2:
		from = Space + ":" + sep(s[0]) + ":" + Space
		to = Space + "::" + s[1] + "::" + Space
	}
	return r.WithRaw(from, to)
}

func (r *Replacer) With(s ...string) *Replacer {
	var from, to string
	switch len(s) {
	case 0:
		return r
	case 1:
		from = s[0]
		to = ":" + s[0] + ":"
	case 2:
		from = s[0]
		to = ":" + s[1] + ":"
	}
	return r.WithRaw(from, to)
}

func (r *Replacer) Build() {
	if len(r.pre) == 0 {
		r.preReplacer = nil
	} else {
		r.preReplacer = strings.NewReplacer(r.pre...)
	}

	if len(r.loop) == 0 {
		r.loopReplacer = nil
	} else {
		r.loopReplacer = strings.NewReplacer(r.loop...)
	}

	if len(r.post) == 0 {
		r.postReplacer = nil
	} else {
		r.postReplacer = strings.NewReplacer(r.post...)
	}

	r.clean = true
}

func (r *Replacer) Replace(args ...string) string {
	if !r.clean {
		r.Build()
	}

	s := strings.Join(args, " ")
	if len(s) == 0 {
		return ""
	}

	// Emoji codes are case insensitive
	s = strings.ToLower(s)

	// Run the initial replacers
	if r.pre != nil {
		s = r.preReplacer.Replace(s)
	}

	// Surround with space since some replacers depend on it
	const prefixSuffix = ":" + Space + ":"
	s = prefixSuffix + s + prefixSuffix

	if r.loopReplacer != nil {
		// Run the loop replacers
		// This is required since strings.Replacer does not handle overlaps
		for {
			prev := s
			s = r.loopReplacer.Replace(s)
			if s == prev {
				break
			}
		}
	}

	// Run the final replacers
	if r.postReplacer != nil {
		s = r.postReplacer.Replace(s)
	}

	// Trim the spaces
	s = strings.TrimPrefix(s, prefixSuffix)
	s = strings.TrimSuffix(s, prefixSuffix)
	return s
}
