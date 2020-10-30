package token

// LexList represents a list of Lexemes. There is no way to remove items,
// instead a new LexList should be created accomodate any list changes.
type LexList []Lexeme

// LexListFrom returns a new LexList initialised with the supplied items. The
// initial capacity will be input size.
func LexListFrom(lexs ...Lexeme) LexList {
	items := make([]Lexeme, 0, len(lexs))
	items = append(items, lexs...)
	return LexList(items)
}

// LexItr represents an iterator of Lexemes.
type LexItr struct {
	Items []Lexeme
	Idx   int
}

// NewLexItr returns a new initialised LexItr.
func NewLexItr(items []Lexeme) *LexItr {
	return &LexItr{
		Items: items,
		Idx:   -1,
	}
}

// More returns true if the end of iterator has not been reached yet.
func (itr *LexItr) More() bool {
	return itr.Idx+1 < len(itr.Items)
}

// More returns true if the start of iterator has not been reached yet.
func (itr *LexItr) Less() bool {
	return itr.Idx > 0
}

// Next returns the next lexeme in the iterator incrementing the iterators
// index accordingly. If the end of the iterator has already been reached then
// a panic insues.
func (itr *LexItr) Next() Lexeme {
	if !itr.More() {
		panic("End of iterator reached, check using LexItr.More first")
	}
	itr.Idx++
	return itr.Items[itr.Idx]
}

// Prev returns the previous lexeme in the iterator decrementing the iterators
// index accordingly. If the start of the iterator has already been reached
// then a panic insues.
func (itr *LexItr) Prev() Lexeme {
	if !itr.Less() {
		panic("Start of iterator reached, check using LexItr.Less first")
	}
	itr.Idx--
	return itr.Items[itr.Idx]
}

// Peek returns the next lexeme in the iterator or the zero lexeme if the end
// of the iterator has been reached.
func (itr *LexItr) Peek() Lexeme {
	if !itr.More() {
		return Lexeme{}
	}
	return itr.Items[itr.Idx]
}

// Window returns the lexeme indicated by the iterators pointer along with the
// lexemes before and after it. If the next or previous do not exist then the
// zero Lexeme is returned in their place.
func (itr *LexItr) Window() (prev, curr, next Lexeme) {
	if itr.Less() {
		prev = itr.Items[itr.Idx-1]
	}
	curr = itr.Items[itr.Idx]
	if itr.More() {
		next = itr.Items[itr.Idx+1]
	}
	return
}
