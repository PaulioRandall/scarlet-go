package token

type Stream interface {
	Next() Token
}

type tkStream struct {
	tks  []Token
	size int
	idx  int
}

func NewStream(tks []Token) Stream {
	return &tkStream{
		tks:  tks,
		size: len(tks),
	}
}

func (stm *tkStream) Next() Token {

	if stm.idx >= stm.size {
		return nil
	}

	tk := stm.tks[stm.idx]
	stm.idx++
	return tk
}
