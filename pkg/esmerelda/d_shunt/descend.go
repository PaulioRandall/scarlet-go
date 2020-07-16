package shunt

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/prop"
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/token"
	. "github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/token/types"
)

func next(shy *shuntingYard) (token.Token, error) {

	for !shy.empty() {

		switch {
		case shy.matchBuff(PR_TERM):
			return shy.next(), nil

		case shy.matchBuff(PR_SEPARATOR):
			if shy.matchTop(PR_PARENTHESIS, PR_OPENER) {
				shy.next()
				continue
			}

			return shy.Pop(), nil

		case shy.acceptPush(PR_SPELL):
			if e := shy.expectPush(PR_PARENTHESIS, PR_OPENER); e != nil {
				return nil, e
			}

			return addProp(shy.Top(), PR_PARAMETERS), nil

		case shy.matchBuff(PR_PARENTHESIS, PR_CLOSER):
			if tk := shy.acceptPop(PR_PARENTHESIS, PR_OPENER); tk != nil {
				shy.next()
			}

			return shy.Pop(), nil

		case shy.matchBuff(PR_TERMINATOR):
			if shy.emptyStack() {
				return shy.next(), nil
			}

		default:
			return nil, errorUnexpectedToken(shy.peek())
		}
	}

	return nil, errorUnexpectedEOF(shy.Top())
}

func addProp(tk token.Token, p Prop) token.Token {

	r := token.Tok{
		Gen:      GEN_PARAMS,
		Sub:      SUB_UNDEFINED,
		RawProps: append(tk.Props(), p),
		RawStr:   tk.Raw(),
	}

	r.Line, r.ColBegin = tk.Begin()
	_, r.ColEnd = tk.End()

	return r
}
