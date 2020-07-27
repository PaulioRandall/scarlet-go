package checker

import (
	"github.com/PaulioRandall/scarlet-go/eskarina/shared/lexeme"
	"github.com/PaulioRandall/scarlet-go/eskarina/shared/perror"
)

type checker struct {
	lex *lexeme.Lexeme
}

func (chk *checker) more() bool {
	return chk.lex != nil
}

func (chk *checker) unexpected(want string) error {
	return perror.New(
		"Unexpected token\nHave: %s\nWant: %s",
		chk.lex.Tok.String(),
		want,
	)
}

func (chk *checker) matchAny(tks ...lexeme.Token) bool {

	if chk.lex == nil {
		return false
	}

	return chk.lex.Tok.IsAny(tks...)
}

func (chk *checker) acceptAny(tks ...lexeme.Token) bool {

	if chk.matchAny(tks...) {
		chk.lex = chk.lex.Next
		return true
	}

	return false
}

func (chk *checker) expectAny(tks ...lexeme.Token) error {

	if chk.lex == nil {
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

func (chk *checker) accept(f func() bool) bool {

	if f != nil && f() {
		chk.lex = chk.lex.Next
		return true
	}

	return false
}

func (chk *checker) expect(want string, f func() bool) error {

	if f == nil {
		return perror.New("Unexpected token\nHave: EOF\nWant: %s", want)
	}

	if chk.accept(f) {
		return nil
	}

	return chk.unexpected(want)
}

func (chk *checker) terminatorPredicate() func() bool {

	if chk.lex == nil {
		return nil
	}

	return chk.lex.Tok.IsTerminator
}

func (chk *checker) termPredicate() func() bool {

	if chk.lex == nil {
		return nil
	}

	return chk.lex.Tok.IsTerm
}
