package compile

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/rincewind/token"
)

type TokenStream interface {
	Next() Token
}

func Compile(ts TokenStream) ([]instruction, error) {
	return nil, nil
}

type compiler struct {
	ts   TokenStream
	ins  stack
	buff Token
}

func (com *compiler) bufferNext() {
	com.buff = com.ts.Next()
}

func (com *compiler) empty() bool {
	return com.buff == nil
}

func (com *compiler) match(gen GenType, sub SubType) bool {

	if com.empty() {
		return false
	}

	g := gen == GE_ANY || gen == com.buff.GenType()
	s := sub == SU_ANY || sub == com.buff.SubType()

	return g && s
}

func (com *compiler) notMatch(gen GenType, sub SubType) bool {
	return !com.match(gen, sub)
}

func (com *compiler) expect(gen GenType, sub SubType) (Token, error) {

	if com.empty() {
		return nil, errorUnexpectedEOF(com)
	}

	if com.match(gen, sub) {
		return com.next(), nil
	}

	return nil, errorWrongToken(com, com.buff)
}

func (com *compiler) next() Token {

	if com.empty() {
		failNow("No tokens remaining, call `match` or `empty` first")
	}

	r := com.buff
	com.bufferNext()

	return r
}
