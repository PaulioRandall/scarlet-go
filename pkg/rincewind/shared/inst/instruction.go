package inst

import (
	"fmt"
)

type Instruction interface {
	Code() Code
	Data() interface{}
	Begin() (int, int)
	End() (int, int)
	String() string
}

type instruction struct {
	Instruction
	code   Code
	data   interface{}
	opener Snippet
	closer Snippet
}

type Snippet interface {
	Begin() (int, int)
	End() (int, int)
}

func New(code Code, data interface{}, opener, closer Snippet) instruction {
	return instruction{
		code:   code,
		data:   data,
		opener: opener,
		closer: closer,
	}
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
