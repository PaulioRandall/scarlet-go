package compile

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/rincewind/instru"
)

type instruction struct {
	code Code
	data interface{}
}

func (ins instruction) Code() Code {
	return ins.code
}

func (ins instruction) Data() interface{} {
	return ins.data
}
