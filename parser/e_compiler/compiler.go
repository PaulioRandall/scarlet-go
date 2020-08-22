package compiler

import (
	"github.com/PaulioRandall/scarlet-go/inst"
	"github.com/PaulioRandall/scarlet-go/lexeme"
)

type Queue interface {
	More() bool
	Empty() bool
	Head() *lexeme.Lexeme
	Take() *lexeme.Lexeme
}

type output struct {
	out []inst.Instruction
}

func (out *output) len() int {
	return len(out.out)
}

func (out *output) emit(in inst.Instruction) {
	out.out = append(out.out, in)
}

func (out *output) emitSet(set *output) {
	out.out = append(out.out, set.out...)
}

type input struct {
	in Queue
}

func (in *input) more() bool {
	return in.in.More()
}

func (in *input) empty() bool {
	return in.in.Empty()
}

func (in *input) is(tk lexeme.Token) bool {
	return in.in.More() && in.in.Head().Tok == tk
}

func (in *input) tok() lexeme.Token {

	if in.in.More() {
		return in.in.Head().Tok
	}

	return lexeme.UNDEFINED
}

func (in *input) take() *lexeme.Lexeme {
	return in.in.Take()
}

func (in *input) discard() {
	if in.in.More() {
		in.in.Take()
	}
}

func (in *input) unexpected() {

	if in.in.Empty() {
		panic("Unexpected EOF")
	}

	panic("Unexpected token: " + in.in.Head().String())
}

func (in *input) println() {
	if in.in.More() {
		println(in.in.Head().String())
	}
}
