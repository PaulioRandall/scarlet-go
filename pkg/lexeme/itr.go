package lexeme

// TokenIterator is a ...token ...iterator.
type TokenIterator struct {
	tks   []Token
	size  int
	index int
}

// NewIterator creates a new iterator from the token slice.
func NewIterator(tks []Token) *TokenIterator {
	return &TokenIterator{tks, len(tks), 0}
}

// Index returns the current index of the iterator.
func (itr *TokenIterator) Index() int {
	return itr.index - 1
}

func (itr *TokenIterator) Empty() bool {
	return itr.index >= itr.size
}

// Peek returns the next token in the iterator without removing it.
func (itr *TokenIterator) Peek() Token {

	if itr.Empty() {
		return Token{}
	}

	return itr.tks[itr.index]
}

// Next returns the next token in the iterator.
func (itr *TokenIterator) Next() Token {

	if itr.Empty() {
		return Token{}
	}

	tk := itr.Peek()
	itr.index++
	return tk
}

// Skip sjips the next token in the iterator.
func (itr *TokenIterator) Skip() {
	if !itr.Empty() {
		itr.index++
	}
}
