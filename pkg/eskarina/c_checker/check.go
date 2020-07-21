package checker

import (
	"github.com/PaulioRandall/scarlet-go/pkg/eskarina/lexeme"
	"github.com/PaulioRandall/scarlet-go/pkg/eskarina/perror"
	"github.com/PaulioRandall/scarlet-go/pkg/eskarina/prop"
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
	case chk.match(prop.PR_SPELL):
		e = spell(chk)

	default:
		return perror.New(
			"Unexpected token\nWant: %s\nHave: %s",
			prop.Join(" & ", prop.PR_SPELL),
			prop.Join(" & ", chk.lex.Props...),
		)
	}

	if e != nil {
		return e
	}

	return chk.expect(prop.PR_TERMINATOR)
}

func spell(chk *checker) error {
	// @Println(a, 1)

	e := chk.expect(prop.PR_SPELL)
	if e != nil {
		return e
	}

	e = chk.expect(prop.PR_PARENTHESIS, prop.PR_OPENER)
	if e != nil {
		return e
	}

	if chk.accept(prop.PR_PARENTHESIS, prop.PR_CLOSER) {
		return nil
	}

	e = parameters(chk)
	if e != nil {
		return e
	}

	return chk.expect(prop.PR_PARENTHESIS, prop.PR_CLOSER)
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
	case chk.accept(prop.PR_TERM):
	default:
		return false, perror.New(
			"Unexpected token\nWant: %s\nHave: %s",
			prop.Join(" & ", prop.PR_TERM),
			prop.Join(" & ", chk.lex.Props...),
		)
	}

	return chk.accept(prop.PR_SEPARATOR), nil
}
