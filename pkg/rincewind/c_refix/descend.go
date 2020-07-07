package refix

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/rincewind/token"
)

func (rfx *refixer) println() {

	println("*************************************")

	if rfx.PeekNext() == nil {
		println("nxt: ---")
	} else {
		println("nxt:    " + rfx.PeekNext().String())
	}

	println("stk:")
	if rfx.EmptyStack() {
		println("\t ---")
		return
	}

	rfx.DescendStack(func(data interface{}) {
		println("\t " + data.(Token).String())
	})
}

func next(rfx *refixer) (Token, error) {

CONTINUE:

	rfx.println()

	switch {
	case rfx.Empty():
		return nil, errorMissingExpression(rfx.PeekNext().(Token))

	case rfx.AcceptPush(GE_SPELL):
		if e := rfx.ExpectPush(SU_PAREN_OPEN); e != nil {
			return nil, e
		}

		goto CONTINUE

	case rfx.MatchNext(SU_PAREN_CLOSE):
		if tk := rfx.AcceptPop(SU_PAREN_OPEN); tk != nil {
			rfx.Next()
		}

		return rfx.Pop(), nil

	case rfx.MatchNext(GE_IDENTIFIER):
		return rfx.Next(), nil

	case rfx.MatchNext(GE_LITERAL):
		return rfx.Next(), nil

	case rfx.MatchNext(SU_VALUE_DELIM):
		if rfx.MatchStack(SU_PAREN_OPEN) {
			rfx.Next()
			goto CONTINUE
		}

		return rfx.Pop(), nil

	case rfx.MatchNext(GE_TERMINATOR):
		if rfx.EmptyStack() {
			return rfx.Next(), nil
		}
	}

	return nil, errorUnexpectedToken(rfx.PeekNext())
}
