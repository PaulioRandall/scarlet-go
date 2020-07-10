package shunt

import (
	"fmt"

	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/perror"
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/token"
	. "github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/token/types"
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

func (shy *shuntingYard) matchBuff(ty interface{}) bool {
	return matchToken(shy.buff, ty)
}

func (shy *shuntingYard) acceptPush(ty interface{}) bool {

	if matchToken(shy.buff, ty) {
		shy.Push(shy.next())
		return true
	}

	return false
}

func (shy *shuntingYard) expectPush(ty interface{}) error {

	e := expectToken(shy.buff, ty)
	if e != nil {
		return e
	}

	shy.Push(shy.next())
	return nil
}

func (shy *shuntingYard) matchTop(ty interface{}) bool {
	return matchToken(shy.Top(), ty)
}

func (shy *shuntingYard) acceptPop(ty interface{}) token.Token {

	if matchToken(shy.Top(), ty) {
		return shy.Pop()
	}

	return nil
}

func (shy *shuntingYard) expectPop(ty interface{}) (token.Token, error) {

	if shy.Top() == nil {
		return nil, expectToken(shy.buff, ty)
	}

	e := expectToken(shy.Top(), ty)
	if e != nil {
		return nil, e
	}

	return shy.Pop(), nil
}

func matchToken(tk token.Token, ty interface{}) bool {

	if tk == nil {
		return false
	}

	if pat, ok := ty.(GenType); ok {
		return pat == GEN_ANY || pat == tk.GenType()
	}

	if pat, ok := ty.(SubType); ok {
		return pat == SUB_ANY || pat == tk.SubType()
	}

	failNow("GenType or SubType required")
	return false
}

func expectToken(tk token.Token, ty interface{}) error {

	if tk == nil {
		return errorUnexpectedEOF(tk)
	}

	if matchToken(tk, ty) {
		return nil
	}

	return errorWrongToken(ty.(fmt.Stringer), tk)
}
