package inst

import (
	"fmt"

	. "github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/codes"
)

type Instruction interface {
	Code() Code
	Data() interface{}
	Snippet
	String() string
}

type Inst struct {
	Instruction
	InstCode Code
	InstData interface{}
	Opener   Snippet
	Closer   Snippet
}

type Snippet interface {
	Begin() (int, int)
	End() (int, int)
}

func (in Inst) Code() Code {
	return in.InstCode
}

func (in Inst) Data() interface{} {
	return in.InstData
}

func (in Inst) Begin() (int, int) {
	return in.Opener.Begin()
}

func (in Inst) End() (int, int) {
	return in.Closer.End()
}

func (in Inst) String() string {

	lineBegin, colBegin := in.Opener.Begin()
	lineEnd, colEnd := in.Closer.End()

	return fmt.Sprintf("%d:%d %d:%d %v %v",
		lineBegin, colBegin,
		lineEnd, colEnd,
		in.InstCode.String(),
		in.InstData,
	)
}
