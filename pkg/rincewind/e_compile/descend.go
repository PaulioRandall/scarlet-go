package compile

import (
	"github.com/PaulioRandall/scarlet-go/pkg/rincewind/shared/inst"
	"github.com/PaulioRandall/scarlet-go/pkg/rincewind/shared/number"

	. "github.com/PaulioRandall/scarlet-go/pkg/rincewind/shared/token/types"
)

func (com *compiler) println() {

	println("*************************************")

	com.Queue.Descend(func(data interface{}) {
		println(data.(inst.Instruction).String())
	})
}

func next(com *compiler) error {

	defer com.discard() // GE_TERMINATOR, now redundant

	if com.match(GE_PARAMS) {
		return call(com)
	}

	return errorUnexpectedToken(com.buff)
}

func call(com *compiler) error {

	com.discard() // GE_PARAMS, now redundant
	argCount := 0

	for !com.match(GE_SPELL) {
		argCount++

		e := expression(com)
		if e != nil {
			return e
		}
	}

	tk := com.next()
	in := inst.New(inst.IN_SPELL, []interface{}{argCount, tk.Value()}, tk, tk)
	com.Put(in)

	return nil
}

func expression(com *compiler) error {

	switch {
	case com.match(SU_IDENTIFIER):
		identifier(com)

	case com.match(GE_LITERAL):
		literal(com)

	default:
		return errorUnexpectedToken(com.buff)
	}

	return nil
}

func identifier(com *compiler) {
	tk := com.next()
	in := inst.New(inst.IN_CTX_GET, tk.Value(), tk, tk)
	com.Put(in)
}

func literal(com *compiler) {

	var val interface{}
	tk := com.next()

	switch tk.SubType() {
	case SU_BOOL:
		val = tk.Value() == "true"

	case SU_NUMBER:
		val = number.New(tk.Value())

	case SU_STRING:
		val = tk.Value()
	}

	in := inst.New(inst.IN_VAL_PUSH, val, tk, tk)
	com.Put(in)
}
