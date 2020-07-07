package refix

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/rincewind/token"
)

func (rfx *refixer) println() {

	println("*************************************")

	if rfx.buff == nil {
		println("rfx.buff: ---")
	} else {
		println("rfx.buff:    " + rfx.buff.String())
	}

	println("rfx.stk:")
	if rfx.stk.top == nil {
		println("\t ---")
		return
	}

	for n := rfx.stk.top; n != nil; n = n.next {
		println("\t" + n.tk.String())
	}
}

func next(rfx *refixer) (Token, error) {

CONTINUE:

	rfx.println()

	switch {
	case rfx.empty():
		return nil, errorMissingExpression(rfx)

	case rfx.matchPush(GE_SPELL, SU_ANY):

		e := rfx.expectPush(GE_PARENTHESIS, SU_PAREN_OPEN)
		if e != nil {
			return nil, e
		}

		goto CONTINUE

	case rfx.match(GE_PARENTHESIS, SU_PAREN_CLOSE):

		if tk := rfx.matchPop(GE_ANY, SU_PAREN_OPEN); tk != nil {
			rfx.next()
		}

		return rfx.pop(), nil

	case rfx.match(GE_IDENTIFIER, SU_ANY):
		return rfx.next(), nil

	case rfx.match(GE_LITERAL, SU_ANY):
		return rfx.next(), nil

	case rfx.match(GE_DELIMITER, SU_VALUE_DELIM):

		if rfx.matchStk(GE_PARENTHESIS, SU_PAREN_OPEN) {
			rfx.next()
			goto CONTINUE
		}

		return rfx.pop(), nil

	case rfx.match(GE_TERMINATOR, SU_ANY):
		if rfx.emptyStk() {
			return rfx.next(), nil
		}
	}

	return nil, errorUnexpectedToken(rfx)
}
