package checker

import (
	"github.com/PaulioRandall/scarlet-go/eskarina/shared/lexeme"
	"github.com/PaulioRandall/scarlet-go/eskarina/shared/perror"
)

func CheckAll(first *lexeme.Lexeme) error {

	chk := &checker{
		lex: first,
	}

	for chk.more() {
		e := check(chk)
		if e != nil {
			return e
		}
	}

	return nil
}

func check(chk *checker) error {

	var e error

	switch {
	case chk.match(lexeme.PR_SPELL):
		e = spell(chk)

	default:
		return perror.New(
			"Unexpected token\nWant: %s\nHave: %s",
			lexeme.JoinProps(" & ", lexeme.PR_SPELL),
			lexeme.JoinProps(" & ", chk.lex.Props...),
		)
	}

	if e != nil {
		return e
	}

	return chk.expect(lexeme.PR_TERMINATOR)
}

func spell(chk *checker) error {
	// @Println(?, ?, ...)

	e := chk.expect(lexeme.PR_SPELL)
	if e != nil {
		return e
	}

	e = chk.expect(lexeme.PR_PARENTHESIS, lexeme.PR_OPENER)
	if e != nil {
		return e
	}

	if chk.accept(lexeme.PR_PARENTHESIS, lexeme.PR_CLOSER) {
		return nil
	}

	e = parameters(chk)
	if e != nil {
		return e
	}

	return chk.expect(lexeme.PR_PARENTHESIS, lexeme.PR_CLOSER)
}

func parameters(chk *checker) error {

	var e error

	for more := true; more; {
		more, e = parameter(chk)
		if e != nil {
			return e
		}
	}

	return nil
}

func parameter(chk *checker) (bool, error) {

	switch {
	case chk.accept(lexeme.PR_TERM):
	default:
		return false, perror.New(
			"Unexpected token\nWant: %s\nHave: %s",
			lexeme.JoinProps(" & ", lexeme.PR_TERM),
			lexeme.JoinProps(" & ", chk.lex.Props...),
		)
	}

	return chk.accept(lexeme.PR_SEPARATOR), nil
}
