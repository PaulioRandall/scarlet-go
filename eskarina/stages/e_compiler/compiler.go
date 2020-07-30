package compiler

import (
	"github.com/PaulioRandall/scarlet-go/eskarina/shared/inst"
	"github.com/PaulioRandall/scarlet-go/eskarina/shared/lexeme"
)

type Queue interface {
	More() bool
	Empty() bool
	Head() *lexeme.Lexeme
	Take() *lexeme.Lexeme
}

type compiler struct {
	input Queue
	out   inst.Queue
}

func (com *compiler) more() bool {
	return com.input.More()
}

func (com *compiler) empty() bool {
	return com.input.Empty()
}

func (com *compiler) is(tk lexeme.Token) bool {
	return com.input.More() && com.input.Head().Tok == tk
}

func (com *compiler) tok() lexeme.Token {

	if com.input.More() {
		return com.input.Head().Tok
	}

	return lexeme.UNDEFINED
}

func (com *compiler) take() *lexeme.Lexeme {
	return com.input.Take()
}

func (com *compiler) reject() {
	if com.input.More() {
		com.input.Take()
	}
}

func (com *compiler) output(in *inst.Instruction) {
	com.out.Put(in)
}

func (com *compiler) unexpected() {

	if com.input.Empty() {
		panic("Unexpected EOF")
	}

	panic("Unexpected token: " + com.input.Head().String())
}
