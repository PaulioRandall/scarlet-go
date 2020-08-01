package checker

import (
	"github.com/PaulioRandall/scarlet-go/shared/lexeme"
	"github.com/PaulioRandall/scarlet-go/shared/perror"
)

type checker struct {
	it *lexeme.Iterator
}

func (chk *checker) unexpected(want string) error {
	return perror.New(
		"Unexpected token\nHave: %s\nWant: %s",
		chk.it.Curr().Tok.String(),
		want,
	)
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
		return perror.New(
			"Unexpected token\nHave: EOF\nWant: %s",
			lexeme.JoinTokens(" & ", tks...),
		)
	}

	if chk.acceptAny(tks...) {
		return nil
	}

	return chk.unexpected(lexeme.JoinTokens(" & ", tks...))
}

func (chk *checker) tok() lexeme.Token {

	if chk.it.Curr() != nil {
		return chk.it.Curr().Tok
	}

	return lexeme.UNDEFINED
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
		return perror.New("Unexpected token\nHave: EOF\nWant: %s", want)
	}

	if chk.accept(ok) {
		return nil
	}

	return chk.unexpected(want)
}
