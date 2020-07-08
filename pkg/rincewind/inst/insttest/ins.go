package insttest

import (
	"fmt"

	. "github.com/PaulioRandall/scarlet-go/pkg/rincewind/inst"
)

func Ins(code Code, data interface{}, line, colBegin, colEnd int) Instruction {
	return ins{
		code:      code,
		data:      data,
		lineBegin: line,
		colBegin:  colBegin,
		lineEnd:   line,
		colEnd:    colEnd,
	}
}

func HalfIns(code Code, data interface{}) Instruction {
	return ins{
		code: code,
		data: data,
	}
}

type ins struct {
	Instruction
	code                Code
	data                interface{}
	lineBegin, colBegin int
	lineEnd, colEnd     int
}

func (in ins) Code() Code {
	return in.code
}

func (in ins) Data() interface{} {
	return in.data
}

func (in ins) Begin() (int, int) {
	return in.lineBegin, in.lineEnd
}

func (in ins) End() (int, int) {
	return in.lineEnd, in.colEnd
}

func (in ins) String() string {

	return fmt.Sprintf("%d:%d %d:%d %v %v",
		in.lineBegin, in.colBegin,
		in.lineEnd, in.colEnd,
		in.code.String(),
		in.data,
	)
}
