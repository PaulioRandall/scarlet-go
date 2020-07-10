package shunt

import (
	"github.com/PaulioRandall/scarlet-go/pkg/rincewind/shared/token"
	. "github.com/PaulioRandall/scarlet-go/pkg/rincewind/shared/token/types"
)

func next(shy *shuntingYard) (token.Token, error) {

	for !shy.empty() {

		switch {
		case shy.matchBuff(GEN_IDENTIFIER):
			return shy.next(), nil

		case shy.matchBuff(GEN_LITERAL):
			return shy.next(), nil

		case shy.matchBuff(SU_VALUE_DELIM):
			if shy.matchTop(SU_PAREN_OPEN) {
				shy.next()
				continue
			}

			return shy.Pop(), nil

		case shy.acceptPush(GEN_SPELL):
			if e := shy.expectPush(SU_PAREN_OPEN); e != nil {
				return nil, e
			}

			return retypeToken(shy.Top(), GEN_PARAMS, SU_UNDEFINED), nil

		case shy.matchBuff(SU_PAREN_CLOSE):
			if tk := shy.acceptPop(SU_PAREN_OPEN); tk != nil {
				shy.next()
			}

			return shy.Pop(), nil

		case shy.matchBuff(GEN_TERMINATOR):
			if shy.emptyStack() {
				return shy.next(), nil
			}

		default:
			return nil, errorUnexpectedToken(shy.peek())
		}
	}

	return nil, errorUnexpectedEOF(shy.Top())
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
