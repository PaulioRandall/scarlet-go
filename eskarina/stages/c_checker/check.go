package checker

import (
	"github.com/PaulioRandall/scarlet-go/eskarina/shared/lexeme"
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
	case chk.matchAny(lexeme.SPELL):
		e = spell(chk)

	default:
		return chk.unexpected(lexeme.SPELL.String())
	}

	if e != nil {
		return e
	}

	return chk.expect("<TERMINATOR>", chk.terminatorPredicate())
}

func spell(chk *checker) error {
	// @Println(?, ?, ...)

	e := chk.expectAny(lexeme.SPELL)
	if e != nil {
		return e
	}

	e = chk.expectAny(lexeme.LEFT_PAREN)
	if e != nil {
		return e
	}

	if chk.acceptAny(lexeme.RIGHT_PAREN) {
		return nil
	}

	e = parameters(chk)
	if e != nil {
		return e
	}

	return chk.expectAny(lexeme.RIGHT_PAREN)
}

func parameters(chk *checker) error {

	for more := true; more; {

		switch {
		case chk.accept(chk.termPredicate()):
		default:
			return chk.unexpected("parameter")
		}

		more = chk.acceptAny(lexeme.SEPARATOR)
	}

	return nil
}
