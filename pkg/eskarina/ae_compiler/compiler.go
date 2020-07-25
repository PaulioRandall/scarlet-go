package compiler

import (
	"github.com/PaulioRandall/scarlet-go/pkg/eskarina/inst"
	"github.com/PaulioRandall/scarlet-go/pkg/eskarina/lexeme"
	"github.com/PaulioRandall/scarlet-go/pkg/eskarina/prop"
)

type compiler struct {
	input lexeme.Queue
	out   inst.Queue
}

func (com *compiler) more() bool {
	return com.input.More()
}

func (com *compiler) empty() bool {
	return com.input.Empty()
}

func (com *compiler) has(props ...prop.Prop) bool {
	return com.input.More() && com.input.Head().Has(props...)
}

func (com *compiler) take() *lexeme.Lexeme {
	return com.input.Take()
}

func (com *compiler) reject(props ...prop.Prop) {
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
