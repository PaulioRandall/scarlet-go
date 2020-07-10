package compile

import (
	"github.com/PaulioRandall/scarlet-go/pkg/rincewind/shared/inst"
	"github.com/PaulioRandall/scarlet-go/pkg/rincewind/shared/number"

	. "github.com/PaulioRandall/scarlet-go/pkg/rincewind/shared/token/types"
)

func (com *compiler) println() {

	println("*************************************")

	com.Queue.Descend(func(data interface{}) {
		println(data.(inst.Inst).String())
	})
}

func next(com *compiler) error {

	defer com.discard() // GEN_TERMINATOR, now redundant

	if com.match(GEN_PARAMS) {
		return call(com)
	}

	return errorUnexpectedToken(com.buff)
}

func call(com *compiler) error {

	com.discard() // GEN_PARAMS, now redundant
	argCount := 0

	for !com.match(GEN_SPELL) {
		argCount++

		e := expression(com)
		if e != nil {
			return e
		}
	}

	tk := com.next()
	com.Put(inst.Inst{
		InstCode: inst.IN_SPELL,
		InstData: []interface{}{argCount, tk.Value()},
		Opener:   tk,
		Closer:   tk,
	})

	return nil
}

func expression(com *compiler) error {

	switch {
	case com.match(SUB_IDENTIFIER):
		identifier(com)

	case com.match(GEN_LITERAL):
		literal(com)

	default:
		return errorUnexpectedToken(com.buff)
	}

	return nil
}

func identifier(com *compiler) {
	tk := com.next()
	com.Put(inst.Inst{
		InstCode: inst.IN_CTX_GET,
		InstData: tk.Value(),
		Opener:   tk,
		Closer:   tk,
	})
}

func literal(com *compiler) {

	var val interface{}
	tk := com.next()

	switch tk.SubType() {
	case SUB_BOOL:
		val = tk.Value() == "true"

	case SUB_NUMBER:
		val = number.New(tk.Value())

	case SUB_STRING:
		val = tk.Value()
	}

	com.Put(inst.Inst{
		InstCode: inst.IN_VAL_PUSH,
		InstData: val,
		Opener:   tk,
		Closer:   tk,
	})
}
