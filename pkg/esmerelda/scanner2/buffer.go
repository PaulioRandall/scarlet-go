package scanner2

type buffer struct {
	itr  SymItr
	buff rune
}

func (b buffer) bufferNext() {
	var ok bool
	if b.buff, ok = b.itr.Next(); !ok {
		b.buff = rune(0)
	}
}

func (b buffer) hasNext() bool {
	return b.buff != rune(0)
}

func (b buffer) empty() bool {
	return b.buff == rune(0)
}

func (b buffer) nextSym() rune {

	if b.empty() {
		panic("PROGRAMMERS ERROR! No more symbols left")
	}

	r := b.buff
	b.bufferNext()

	return r
}

func (b buffer) peekSym() rune {
	return b.buff
}
