package scanner

type SymItr interface {
	HasNext() bool
	Next() rune
}

type sym struct {
	SymItr
	buf rune
}

func (s sym) HasNext() bool {
	return s.buf != rune(0) || s.SymItr.HasNext()
}

func (s sym) Next() rune {

	if !s.HasNext() {
		m := "No more symbols to return, you should call sym.HasNext() first"
		panic("PROGRAMMERS ERROR! " + m)
	}

	getNext := func() rune {
		if s.SymItr.HasNext() {
			return s.SymItr.Next()
		}

		return rune(0)
	}

	r := s.buf
	if r == rune(0) {
		r = getNext()
	}

	s.buf = getNext()
	return r
}
