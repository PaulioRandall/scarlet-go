package refix

import (
	"github.com/PaulioRandall/scarlet-go/pkg/rincewind/shared/token"
	. "github.com/PaulioRandall/scarlet-go/pkg/rincewind/shared/token/types"
)

func (rfx *refixer) println() {

	println("*************************************")

	if rfx.PeekBuff() == nil {
		println("nxt: ---")
	} else {
		println("nxt:    " + rfx.PeekBuff().String())
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
		return nil, errorMissingExpression(rfx.PeekBuff())

	case rfx.AcceptPush(GE_SPELL):
		if e := rfx.ExpectPush(SU_PAREN_OPEN); e != nil {
			return nil, e
		}

		return retypeToken(rfx.PeekTop(), GE_PARAMS, SU_UNDEFINED), nil

	case rfx.MatchBuff(SU_PAREN_CLOSE):
		if tk := rfx.AcceptPop(SU_PAREN_OPEN); tk != nil {
			rfx.Next()
		}

		return rfx.Pop(), nil

	case rfx.MatchBuff(GE_IDENTIFIER):
		return rfx.Next(), nil

	case rfx.MatchBuff(GE_LITERAL):
		return rfx.Next(), nil

	case rfx.MatchBuff(SU_VALUE_DELIM):
		if rfx.MatchTop(SU_PAREN_OPEN) {
			rfx.Next()
			goto CONTINUE
		}

		return rfx.Pop(), nil

	case rfx.MatchBuff(GE_TERMINATOR):
		if rfx.EmptyStack() {
			return rfx.Next(), nil
		}
	}

	return nil, errorUnexpectedToken(rfx.PeekBuff())
}

func retypeToken(tk token.Token, gen GenType, sub SubType) token.Token {

	r := token.Tok{
		Gen:    gen,
		Sub:    sub,
		RawStr: tk.Raw(),
	}

	r.Line, r.ColBegin = tk.Begin()
	_, r.ColEnd = tk.End()

	return r
}
