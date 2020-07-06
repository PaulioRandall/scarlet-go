package refix

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/rincewind/token"
)

func next(rfx *refixer) (Token, error) {

CONTINUE:

	switch {
	case rfx.empty():
		return nil, errorMissingExpression(rfx)

	case rfx.match(GE_SPELL, SU_ANY):
		rfx.pushNext()

		tk, e := rfx.expect(GE_PARENTHESIS, SU_PAREN_OPEN)
		if e != nil {
			return nil, e
		}

		rfx.stk.push(tk)
		goto CONTINUE

	case rfx.match(GE_PARENTHESIS, SU_PAREN_CLOSE):
		if rfx.stk.peek().SubType() == SU_PAREN_OPEN {
			rfx.stk.pop()
			rfx.next()
		}

		return rfx.stk.pop(), nil

	case rfx.match(GE_IDENTIFIER, SU_ANY):
		return rfx.next(), nil

	case rfx.match(GE_LITERAL, SU_ANY):
		return rfx.next(), nil

	case rfx.match(GE_DELIMITER, SU_VALUE_DELIM):
		if rfx.stk.peek().SubType() != SU_PAREN_OPEN {
			return rfx.stk.pop(), nil
		}

		rfx.next()
		goto CONTINUE

	case rfx.match(GE_TERMINATOR, SU_ANY):
		if rfx.stk.empty() {
			return rfx.next(), nil
		}
	}

	return nil, errorUnexpectedToken(rfx)
}
