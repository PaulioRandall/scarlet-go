package checker

import (
	"github.com/PaulioRandall/scarlet-go/pkg/eskarina/lexeme"
	"github.com/PaulioRandall/scarlet-go/pkg/eskarina/perror"
	"github.com/PaulioRandall/scarlet-go/pkg/eskarina/prop"
)

type checker struct {
	lex *lexeme.Lexeme
}

func (chk *checker) more() bool {
	return chk.lex != nil
}

func (chk *checker) match(props ...prop.Prop) bool {

	if chk.lex == nil {
		return false
	}

	return chk.lex.Is(props...)
}

func (chk *checker) accept(props ...prop.Prop) bool {

	if chk.match(props...) {
		chk.lex = chk.lex.Next
		return true
	}

	return false
}

func (chk *checker) expect(props ...prop.Prop) error {

	if chk.lex == nil {
		return perror.New(
			"Unexpected token\nWant: %s\nHave: EOF",
			prop.Join(" & ", props...),
		)
	}

	if chk.accept(props...) {
		return nil
	}

	return perror.New(
		"Unexpected token\nWant: %s\nHave: %s",
		prop.Join(" & ", props...),
		prop.Join(" & ", chk.lex.Props...),
	)
}
