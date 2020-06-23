package scanner2

import (
	"unicode"
)

type buffer struct {
	SymItr
	buff rune
}

func (b *buffer) bufferNext() {

	var ok bool
	b.buff, ok = b.Next()

	if !ok {
		b.buff = rune(0)
	}
}

func (b *buffer) hasNext() bool {
	return b.buff != rune(0)
}

func (b *buffer) empty() bool {
	return b.buff == rune(0)
}

func (b *buffer) next() rune {

	if b.empty() {
		progError("No symbols remaining")
	}

	r := b.buff
	b.bufferNext()

	return r
}

func (b *buffer) peek() rune {
	return b.buff
}

func (b *buffer) expect(exp rune) (rune, bool) {

	if b.match(exp) {
		return b.next(), true
	}

	return rune(0), false
}

func (b *buffer) match(ru rune) bool {

	if b.peek() != ru {
		return false
	}

	return true
}

func (b *buffer) notMatch(ru rune) bool {
	return !b.match(ru)
}

func (b *buffer) matchNewline() bool {
	ru := b.peek()
	return ru == '\r' || ru == '\n'
}

func (b *buffer) matchSpace() bool {

	if unicode.IsSpace(b.peek()) {
		return true
	}

	return false
}
