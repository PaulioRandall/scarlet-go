package token

type Stream struct {
	tks  []Token
	size int
	idx  int
}

func NewStream(tks []Token) *Stream {
	return &Stream{
		tks:  tks,
		size: len(tks),
	}
}

func (stm *Stream) Next() Token {

	if stm.idx >= stm.size {
		return nil
	}

	tk := stm.tks[stm.idx]
	stm.idx++
	return tk
}
