package checker

import (
	"github.com/PaulioRandall/scarlet-go/shared/lexeme"
)

func CheckAll(con *lexeme.Container) error {

	chk := &checker{
		it: con.Iterator(),
	}

	for chk.it.Next(); chk.it.HasNext(); {
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
	case chk.matchAny(lexeme.IDENT):
		e = assignmentOrExpression(chk)

	case chk.tok().IsTerm(), chk.matchAny(lexeme.SPELL, lexeme.L_PAREN):
		e = expression(chk)

	case chk.matchAny(lexeme.L_SQUARE):
		e = guard(chk)

	default:
		return chk.unexpected("<STATEMENT>")
	}

	if e != nil {
		return e
	}

	return chk.expect("<TERMINATOR>", chk.tok().IsTerminator())
}

func assignmentOrExpression(chk *checker) error {

	if e := chk.expectAny(lexeme.IDENT); e != nil {
		return e
	}

	if chk.matchAny(lexeme.ASSIGN, lexeme.DELIM) {
		chk.undo()
		return assignment(chk)
	}

	return expression(chk)
}

func assignment(chk *checker) error {

	var count int

	for count == 0 || chk.acceptAny(lexeme.DELIM) {
		if e := chk.expectAny(lexeme.IDENT); e != nil {
			return e
		}

		count++
	}

	if e := chk.expectAny(lexeme.ASSIGN); e != nil {
		return e
	}

	i, e := expressions(chk)
	if e != nil {
		return e
	}

	if i != count {
		return chk.unexpected(lexeme.DELIM.String())
	}

	return nil
}

func spell(chk *checker) error {
	// @Println(?, ?, ...)

	e := chk.expectAny(lexeme.SPELL)
	if e != nil {
		return e
	}

	return parameters(chk)
}

func parameters(chk *checker) error {

	if e := chk.expectAny(lexeme.L_PAREN); e != nil {
		return e
	}

	if chk.acceptAny(lexeme.R_PAREN) {
		return nil
	}

	if _, e := expressions(chk); e != nil {
		return e
	}

	return chk.expectAny(lexeme.R_PAREN)
}

func expressions(chk *checker) (int, error) {

	var i int

	for more := true; more; more = chk.acceptAny(lexeme.DELIM) {
		if e := expression(chk); e != nil {
			return 0, e
		}

		i++
	}

	return i, nil
}

func expression(chk *checker) error {

	e := term(chk)
	if e != nil {
		return e
	}

	for chk.accept(chk.tok().IsOperator()) {
		if e = term(chk); e != nil {
			return e
		}
	}

	return nil
}

func group(chk *checker) error {

	if e := chk.expectAny(lexeme.L_PAREN); e != nil {
		return e
	}

	if e := expression(chk); e != nil {
		return e
	}

	return chk.expectAny(lexeme.R_PAREN)
}

func term(chk *checker) error {

	switch {
	case chk.matchAny(lexeme.SPELL):
		return spell(chk)

	case chk.matchAny(lexeme.L_PAREN):
		return group(chk)

	case chk.accept(chk.tok().IsTerm()):
		return nil

	default:
		return chk.unexpected("<TERM>")
	}

	return nil
}

func guard(chk *checker) error {

	if e := chk.expectAny(lexeme.L_SQUARE); e != nil {
		return e
	}

	if e := expression(chk); e != nil {
		return e
	}

	if e := chk.expectAny(lexeme.R_SQUARE); e != nil {
		return e
	}

	return nil
}
