package compile

import (
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/number"

	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/inst"
	. "github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/inst/codes"
	. "github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/prop"
)

func (com *compiler) println() {

	println("*************************************")

	com.Queue.Descend(func(data inst.Instruction) {
		println(data.String())
	})
}

func next(com *compiler) error {

	defer com.next() // GEN_TERMINATOR, now redundant

	if com.match(PR_PARAMETERS) {
		return call(com)
	}

	return errorUnexpectedToken(com.buff)
}

func call(com *compiler) error {

	com.next() // GEN_PARAMS, now redundant
	argCount := 0

	for !com.match(PR_SPELL) {
		argCount++

		e := expression(com)
		if e != nil {
			return e
		}
	}

	tk := com.next()
	com.Put(inst.Inst{
		InstCode: IN_SPELL,
		InstData: []interface{}{argCount, tk.Value()},
		Opener:   tk,
		Closer:   tk,
	})

	return nil
}

func expression(com *compiler) error {

	switch {
	case com.match(PR_IDENTIFIER):
		identifier(com)

	case com.match(PR_LITERAL):
		literal(com)

	default:
		return errorUnexpectedToken(com.buff)
	}

	return nil
}

func identifier(com *compiler) {
	tk := com.next()
	com.Put(inst.Inst{
		InstCode: IN_CTX_GET,
		InstData: tk.Value(),
		Opener:   tk,
		Closer:   tk,
	})
}

func literal(com *compiler) {

	var val interface{}
	tk := com.peek()

	switch {
	case tk.Is(PR_BOOL):
		val = tk.Value() == "true"

	case tk.Is(PR_NUMBER):
		val = number.New(tk.Value())

	case tk.Is(PR_STRING):
		val = tk.Value()
	}

	com.next()
	com.Put(inst.Inst{
		InstCode: IN_VAL_PUSH,
		InstData: val,
		Opener:   tk,
		Closer:   tk,
	})
}
