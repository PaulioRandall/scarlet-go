package parser

import (
	"github.com/PaulioRandall/scarlet-go/mmxxi/scarlet/token"
)

// LexIterator allows iteration over a lexeme slice.
type LexIterator interface {
	Line() int
	Accept(tk token.Token) bool
	Read() token.Lexeme
	Peek() token.Lexeme
	More() bool
	Index() int
	At(i int) token.Lexeme
	InRange(i int) bool
	Match(tk token.Token) bool
	MatchAny(tks ...token.Token) bool
	MatchPat(tks ...token.Token) bool
}

type lxIterator struct {
	tks  []token.Lexeme
	size int
	idx  int
}

// NewIterator returns a new lexeme iterator.
func NewIterator(tks []token.Lexeme) *lxIterator {
	return &lxIterator{
		tks:  tks,
		size: len(tks),
	}
}

// Line returns the current line number.
func (itr *lxIterator) Line() int {
	switch {
	case itr.More():
		return itr.tks[itr.idx].Snippet.Start.Line + 1
	case itr.idx != 0:
		return itr.tks[itr.idx-1].Snippet.Start.Line + 1
	default:
		return 0
	}
}

// Accept returns the next lexeme before incrementing.
func (itr *lxIterator) Accept(tk token.Token) bool {
	if itr.Match(tk) {
		itr.idx++
		return true
	}
	return false
}

// Read returns the next lexeme before incrementing.
func (itr *lxIterator) Read() token.Lexeme {
	lx := itr.tks[itr.idx]
	itr.idx++
	return lx
}

// Peek returns the next lexeme without incrementing.
func (itr *lxIterator) Peek() token.Lexeme {
	return itr.tks[itr.idx]
}

// More returns true if there are more tokens to be read.
func (itr *lxIterator) More() bool {
	return itr.size > itr.idx
}

// Index returns the current item mark.
func (itr *lxIterator) Index() int {
	return itr.idx
}

// At returns the lexeme at index 'i'.
func (itr *lxIterator) At(i int) token.Lexeme {
	return itr.tks[i]
}

// InRange returns true if index 'i' is within the remaining tokens.
func (itr *lxIterator) InRange(i int) bool {
	return i < itr.size
}

// Match returns true if the next token starts with 'tk'.
func (itr *lxIterator) Match(tk token.Token) bool {
	return itr.More() && itr.tks[itr.idx].Token == tk
}

// MatchAny returns true if the next token matches any in 'tks'.
func (itr *lxIterator) MatchAny(tks ...token.Token) bool {

	subject := itr.tks[itr.idx].Token
	for _, tk := range tks {
		if subject == tk {
			return true
		}
	}

	return false
}

// MatchPat returns true if the upcoming tokens starts with 'tks'.
func (itr *lxIterator) MatchPat(tks ...token.Token) bool {

	if itr.idx+len(tks) > itr.size {
		return false
	}

	i := itr.idx
	for _, n := range tks {
		if itr.tks[i].Token != n {
			return false
		}
		i++
	}

	return true
}
