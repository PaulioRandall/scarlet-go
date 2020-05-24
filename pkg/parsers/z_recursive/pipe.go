package z_recursive

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/z_token"
)

// pipe encapsulates a TokenIterator providing general iterator functionality
// useful when parsing.
type pipe struct {
	itr *TokenIterator
}

func (p *pipe) next() Token {
	return p.itr.Next()
}

func (p *pipe) peek() Token {
	return p.itr.Peek()
}

func (p *pipe) past() Token {
	return p.itr.Past() // Previous token, no iteration
}

func (p *pipe) back() {
	p.itr.Back()
}

func (p *pipe) match(m Morpheme) bool {
	o := p.itr.Peek().Morpheme()
	return o == ANY || o == m
}

func (p *pipe) matchAny(ms ...Morpheme) bool {
	for _, t := range ms {
		if p.match(t) {
			return true
		}
	}

	return false
}

func (p *pipe) matchSequence(ms ...Morpheme) bool {

	count := 0

	defer func() { // Undo all calls to p.itr.Next
		for ; count > 0; count-- {
			p.itr.Back()
		}
	}()

	for _, m := range ms {

		if m == ANY || p.match(m) {
			p.itr.Next()
			count++
			continue
		}

		return false
	}

	return true
}

func (p *pipe) accept(m Morpheme) bool {

	if m == UNDEFINED {
		return false
	}

	if m == ANY || p.match(m) {
		p.itr.Skip()
		return true
	}

	return false
}

func (p *pipe) expect(tag string, m Morpheme) Token {
	if !p.accept(m) {
		panic(unexpected(tag, p.itr.Peek(), m.String()))
	}

	return p.itr.Past()
}
