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
	case chk.matchAny(lexeme.SPELL):
		e = spell(chk)

	case chk.matchAny(lexeme.IDENT):
		e = assignmentOrExpression(chk)

	case chk.tok().IsTerm():
		e = expressionStatement(chk)

	default:
		return chk.unexpected("<STATEMENT>")
	}

	if e != nil {
		return e
	}

	return chk.expect("<TERMINATOR>", chk.tok().IsTerminator())
}

func spell(chk *checker) error {
	// @Println(?, ?, ...)

	e := chk.expectAny(lexeme.SPELL)
	if e != nil {
		return e
	}

	e = chk.expectAny(lexeme.L_PAREN)
	if e != nil {
		return e
	}

	if chk.acceptAny(lexeme.R_PAREN) {
		return nil
	}

	e = parameters(chk)
	if e != nil {
		return e
	}

	return chk.expectAny(lexeme.R_PAREN)
}

func parameters(chk *checker) error {

	for more := true; more; {

		if e := expression(chk); e != nil {
			return e
		}

		more = chk.acceptAny(lexeme.DELIM)
	}

	return nil
}

func assignmentOrExpression(chk *checker) error {

	if e := chk.expectAny(lexeme.IDENT); e != nil {
		return e
	}

	if chk.matchAny(lexeme.ASSIGN, lexeme.DELIM) {
		return assignment(chk)
	}

	return expressionStatement(chk)
}

func assignment(chk *checker) error {

	count := 1
	for chk.acceptAny(lexeme.DELIM) {
		if e := chk.expectAny(lexeme.IDENT); e != nil {
			return e
		}
		count++
	}

	if e := chk.expectAny(lexeme.ASSIGN); e != nil {
		return e
	}

	for count > 0 {
		if e := expression(chk); e != nil {
			return e
		}

		count--

		if !chk.acceptAny(lexeme.DELIM) && count != 0 {
			return chk.unexpected(lexeme.DELIM.String())
		}
	}

	return nil
}

func expressionStatement(chk *checker) error {

	e := term(chk)
	if e != nil {
		return e
	}

	for !chk.tok().IsTerminator() {

		e = chk.expect("<OPERATOR>", chk.tok().IsOperator())
		if e != nil {
			return e
		}

		if e = term(chk); e != nil {
			return e
		}
	}

	return nil
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

func term(chk *checker) error {
	switch {
	//case chk.matchAny(lexeme.SPELL):
	//return spell(chk)

	case chk.accept(chk.tok().IsTerm()):
		return nil

	default:
		return chk.unexpected("<TERM>")
	}

	return nil
}
