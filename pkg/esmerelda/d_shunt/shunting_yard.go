package shunt

import (
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/perror"
	. "github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/prop"
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/token"
)

type shuntingYard struct {
	token.Stack
	ts   TokenStream
	buff token.Token
}

func (shy *shuntingYard) shunt() (token.Token, RefixFunc, error) {

	tk, e := next(shy)
	if e != nil {
		return nil, nil, e
	}

	if shy.empty() {
		return tk, nil, nil
	}

	return tk, shy.shunt, nil
}

func (shy *shuntingYard) empty() bool {
	return shy.emptyBuff() && shy.emptyStack()
}

func (shy *shuntingYard) emptyBuff() bool {
	return shy.buff == nil
}

func (shy *shuntingYard) emptyStack() bool {
	return shy.Top() == nil
}

func (shy *shuntingYard) next() token.Token {

	if shy.empty() {
		perror.Panic("No data remaining, check first")
	}

	data := shy.buff
	shy.buff = shy.ts.Next()
	return data
}

func (shy *shuntingYard) peek() token.Token {
	return shy.buff
}

func (shy *shuntingYard) matchBuff(props ...Prop) bool {
	return matchToken(shy.buff, props...)
}

func (shy *shuntingYard) acceptPush(props ...Prop) bool {

	if matchToken(shy.buff, props...) {
		shy.Push(shy.next())
		return true
	}

	return false
}

func (shy *shuntingYard) expectPush(props ...Prop) error {

	e := expectToken(shy.buff, props...)
	if e != nil {
		return e
	}

	shy.Push(shy.next())
	return nil
}

func (shy *shuntingYard) matchTop(props ...Prop) bool {
	return matchToken(shy.Top(), props...)
}

func (shy *shuntingYard) acceptPop(props ...Prop) token.Token {

	if matchToken(shy.Top(), props...) {
		return shy.Pop()
	}

	return nil
}

func (shy *shuntingYard) expectPop(props ...Prop) (token.Token, error) {

	if shy.Top() == nil {
		return nil, expectToken(shy.buff, props...)
	}

	e := expectToken(shy.Top(), props...)
	if e != nil {
		return nil, e
	}

	return shy.Pop(), nil
}

func matchToken(tk token.Token, props ...Prop) bool {

	if tk == nil {
		return false
	}

	return tk.Is(props...)
}

func expectToken(tk token.Token, props ...Prop) error {

	if tk == nil {
		return errorUnexpectedEOF(tk)
	}

	if matchToken(tk, props...) {
		return nil
	}

	return errorWrongToken(props, tk)
}
