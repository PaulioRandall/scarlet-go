package compile

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/rincewind/inst"
	"github.com/PaulioRandall/scarlet-go/pkg/rincewind/number"
	. "github.com/PaulioRandall/scarlet-go/pkg/rincewind/token"
)

func (com *compiler) println() {

	println("*************************************")

	com.Queue.Descend(func(data interface{}) {
		println(data.(Instruction).String())
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

	com.Put(instruction{
		code: IN_SPELL,
		data: []interface{}{
			argCount,
			tk.Value(),
		},
		opener: tk,
		closer: tk,
	})

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

	com.Put(instruction{
		code:   IN_CTX_GET,
		data:   tk.Value(),
		opener: tk,
		closer: tk,
	})
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

	com.Put(instruction{
		code:   IN_VAL_PUSH,
		data:   val,
		opener: tk,
		closer: tk,
	})
}
