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

func (chk *checker) match(props ...lexeme.Prop) bool {

	if chk.lex == nil {
		return false
	}

	return chk.lex.Has(props...)
}

func (chk *checker) accept(props ...lexeme.Prop) bool {

	if chk.match(props...) {
		chk.lex = chk.lex.Next
		return true
	}

	return false
}

func (chk *checker) expect(props ...lexeme.Prop) error {

	if chk.lex == nil {
		return perror.New(
			"Unexpected token\nWant: %s\nHave: EOF",
			lexeme.JoinProps(" & ", props...),
		)
	}

	if chk.accept(props...) {
		return nil
	}

	return perror.New(
		"Unexpected token\nWant: %s\nHave: %s",
		lexeme.JoinProps(" & ", props...),
		lexeme.JoinProps(" & ", chk.lex.Props...),
	)
}
