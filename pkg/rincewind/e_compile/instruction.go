package compile

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/rincewind/inst"
	. "github.com/PaulioRandall/scarlet-go/pkg/rincewind/token"
)

type instruction struct {
	Instruction
	code   Code
	data   interface{}
	opener Token
	closer Token
}

func (ins instruction) Code() Code {
	return ins.code
}

func (ins instruction) Data() interface{} {
	return ins.data
}

func (ins instruction) Begin() (int, int) {
	return ins.opener.Begin()
}

func (ins instruction) End() (int, int) {
	return ins.closer.End()
}
