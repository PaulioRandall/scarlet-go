package refix

import (
	"github.com/PaulioRandall/scarlet-go/pkg/rincewind/token"
	. "github.com/PaulioRandall/scarlet-go/pkg/rincewind/token/types"
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
		println("\t " + data.(token.Token).String())
	})
}

func next(rfx *refixer) (token.Token, error) {

CONTINUE:

	rfx.println()

	switch {
	case rfx.Empty():
		return nil, errorMissingExpression(rfx.PeekNext().(token.Token))

	case rfx.AcceptPush(GE_SPELL):
		if e := rfx.ExpectPush(SU_PAREN_OPEN); e != nil {
			return nil, e
		}

		return token.MagicToken(GE_PARAMS, SU_UNDEFINED, rfx.PeekTop()), nil

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
		if rfx.MatchTop(SU_PAREN_OPEN) {
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
