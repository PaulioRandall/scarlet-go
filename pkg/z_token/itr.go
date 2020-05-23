package z_token

func NewIterator(tks []Token) *TokenIterator {
	return &TokenIterator{tks, len(tks), 0}
}

type TokenIterator struct {
	tks   []Token
	size  int
	index int
}

/*
func (itr *TokenIterator) Index() int {
	return itr.index - 1
}

func (itr *TokenIterator) Empty() bool {
	return itr.index >= itr.size
}

func (itr *TokenIterator) Peek() Token {

	if itr.Empty() {
		return nil
	}

	return itr.tks[itr.index]
}

func (itr *TokenIterator) Next() Token {

	if itr.Empty() {
		return nil
	}

	tk := itr.Peek()
	itr.index++
	return tk
}

func (itr *TokenIterator) Skip() {
	if !itr.Empty() {
		itr.index++
	}
}

func (itr *TokenIterator) Past() Token {
	if itr.index > 0 {
		return itr.tks[itr.index-1]
	}

	return nil
}

func (itr *TokenIterator) Back() {
	if itr.index != 0 {
		itr.index--
	}
}
*/
