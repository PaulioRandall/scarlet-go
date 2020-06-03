package matcher

type PatternMatcher func(*Symbols) (size int, e error)

type Pattern interface {
	Matcher() PatternMatcher
	OnMatch(value string, line, col int) interface{}
}

type Matcher struct {
	ps []Pattern
	s  *Symbols
}

func New(text string, patterns []Pattern) Matcher {
	return Matcher{
		ps: patterns,
		s: &Symbols{
			runes: []rune(text),
		},
	}
}

func (m Matcher) Next() (token interface{}, e error) {

	if m.s.Empty() {
		return nil, nil
	}

	for _, p := range m.ps {

		n, e := p.Matcher()(m.s)
		if e != nil {
			return nil, e
		}

		if n > 0 {
			return m.onMatch(p, n)
		}
	}

	return nil, nil
}

func (m Matcher) onMatch(p Pattern, n int) (token interface{}, e error) {

	line, col := m.s.line, m.s.col
	value, e := m.s.read(n)
	if e != nil {
		return nil, e
	}

	return p.OnMatch(value, line, col), nil
}
