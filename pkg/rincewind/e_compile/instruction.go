package compile

import (
	"fmt"

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

func (in instruction) Code() Code {
	return in.code
}

func (in instruction) Data() interface{} {
	return in.data
}

func (in instruction) Begin() (int, int) {
	return in.opener.Begin()
}

func (in instruction) End() (int, int) {
	return in.closer.End()
}

func (in instruction) String() string {

	lineBegin, colBegin := in.opener.Begin()
	lineEnd, colEnd := in.closer.End()

	return fmt.Sprintf("%d:%d %d:%d %v %v",
		lineBegin, colBegin,
		lineEnd, colEnd,
		in.code.String(),
		in.data,
	)
}
