package token

func NewIterator(tks []Token) *TokenIterator {
	return &TokenIterator{tks, len(tks), 0}
}

type TokenIterator struct {
	tks  []Token
	size int
	pos  int
}

func (itr *TokenIterator) Index() int {
	return itr.pos - 1
}

func (itr *TokenIterator) Empty() bool {
	return itr.pos >= itr.size
}

func (itr *TokenIterator) Peek() Token {

	if itr.Empty() {
		return nil
	}

	return itr.tks[itr.pos]
}

func (itr *TokenIterator) Next() Token {

	if itr.Empty() {
		return nil
	}

	tk := itr.Peek()
	itr.pos++
	return tk
}

func (itr *TokenIterator) Skip() {
	if !itr.Empty() {
		itr.pos++
	}
}

func (itr *TokenIterator) Past() Token {
	if itr.pos > 0 {
		return itr.tks[itr.pos-1]
	}

	return nil
}

func (itr *TokenIterator) Back() {
	if itr.pos != 0 {
		itr.pos--
	}
}
