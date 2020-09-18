package checker

import (
	"github.com/PaulioRandall/scarlet-go/lexeme"
)

type checker struct {
	it *lexeme.Iterator
}

func (chk *checker) unexpected(want string) error {
	return newErr(chk.it.Curr(),
		"Unexpected token, want %s, have %s", want, chk.it.Curr().Tok.String())
}

func (chk *checker) matchAny(tks ...lexeme.Token) bool {
	return chk.it.Curr() != nil && chk.it.Curr().Tok.IsAny(tks...)
}

func (chk *checker) acceptAny(tks ...lexeme.Token) bool {

	if chk.matchAny(tks...) {
		chk.it.Next()
		return true
	}

	return false
}

func (chk *checker) expectAny(tks ...lexeme.Token) error {

	if chk.it.Curr() == nil {
		s := lexeme.JoinTokens(" & ", tks...)
		return newEofErr("Unexpected token, want %s, have EOF", s)
	}

	if chk.acceptAny(tks...) {
		return nil
	}

	s := lexeme.JoinTokens(" & ", tks...)
	return chk.unexpected(s)
}

func (chk *checker) tok() lexeme.Token {

	if chk.it.Curr() != nil {
		return chk.it.Curr().Tok
	}

	return lexeme.UNDEFINED
}

func (chk *checker) next() {
	chk.it.Next()
}

func (chk *checker) accept(ok bool) bool {

	if chk.it.Curr() != nil && ok {
		chk.it.Next()
		return true
	}

	return false
}

func (chk *checker) expect(want string, ok bool) error {

	if chk.it.Curr() == nil {
		return newEofErr("Unexpected token, want %s, have EOF", want)
	}

	if chk.accept(ok) {
		return nil
	}

	return chk.unexpected(want)
}

func (chk *checker) undo() {
	chk.it.Prev()
}
