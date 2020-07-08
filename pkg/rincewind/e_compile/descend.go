package compile

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/rincewind/inst"
	"github.com/PaulioRandall/scarlet-go/pkg/rincewind/number"
	. "github.com/PaulioRandall/scarlet-go/pkg/rincewind/token"
)

// Example: @Set("x", "Scarlet")
// 1: IN_VAL_PUSH		"x"
// 2: IN_VAL_PUSH		"Scarlet"
// 3: IN_SPELL   		@Set

func next(com *compiler) error {

	if com.match(GE_PARAMS) {
		return call(com)
	}

	return errorUnexpectedToken(com.buff)
}

func call(com *compiler) error {

	com.discard()
	argCount := 0

	for !com.match(GE_SPELL) {
		argCount++

		e := expression(com)
		if e != nil {
			return e
		}
	}

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

	switch {
	case com.match(SU_BOOL):
		val = tk.Value() == "true"

	case com.match(SU_NUMBER):
		val = number.New(tk.Value())

	case com.match(SU_STRING):
		val = tk.Value()
	}

	com.Put(instruction{
		code:   IN_VAL_PUSH,
		data:   val,
		opener: tk,
		closer: tk,
	})
}
