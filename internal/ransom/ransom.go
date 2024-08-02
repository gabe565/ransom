package ransom

import "strings"

const Space = "blank"

func Default() *Replacer {
	return New(
		WithAlphabet(),
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
		once: make([]string, 0, len(opts)*2),
		loop: make([]string, 0, len(opts)*2),
	}
	for _, opt := range opts {
		opt(r)
	}
	return r
}

type Replacer struct {
	once []string
	loop []string
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

func (r *Replacer) Replace(args ...string) string {
	s := strings.Join(args, " ")
	if len(s) == 0 {
		return ""
	}

	// Emoji codes are case insensitive
	s = strings.ToLower(s)

	// Run the initial replacers
	if len(r.once) != 0 {
		s = strings.NewReplacer(r.once...).Replace(s)
	}

	// Surround with space since some replacers depend on it
	const prefixSuffix = ":" + Space + ":"
	s = prefixSuffix + s + prefixSuffix

	if len(r.loop) != 0 {
		// Run the loop replacers
		// This is required since strings.Replacer does not handle overlaps
		sr := strings.NewReplacer(r.loop...)
		for {
			prev := s
			s = sr.Replace(s)
			if s == prev {
				break
			}
		}
	}

	// Trim the spaces
	s = strings.TrimPrefix(s, prefixSuffix)
	s = strings.TrimSuffix(s, prefixSuffix)
	return s
}
