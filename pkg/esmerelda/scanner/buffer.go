package scanner

import (
	"unicode"
)

type SymItr interface {
	Next() (rune, bool)
}

type buffer struct {
	symItr SymItr
	buff   rune
}

func (b *buffer) bufferNext() {

	var ok bool
	b.buff, ok = b.symItr.Next()

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

func (b *buffer) match(ru rune) bool {
	return b.buff == ru
}

func (b *buffer) notMatch(ru rune) bool {
	return !b.match(ru)
}

func (b *buffer) matchNewline() bool {
	return b.buff == '\r' || b.buff == '\n'
}

func (b *buffer) matchSpace() bool {
	return unicode.IsSpace(b.buff)
}

func (b *buffer) matchLetter() bool {
	return unicode.IsLetter(b.buff)
}

func (b *buffer) matchDigit() bool {
	return unicode.IsDigit(b.buff)
}
