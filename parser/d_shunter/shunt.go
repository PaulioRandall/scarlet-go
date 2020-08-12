package shunter

import (
	"github.com/PaulioRandall/scarlet-go/shared/lexeme"
)

func ShuntAll(con *lexeme.Container) *lexeme.Container {

	shy := &shuntingYard{
		queue: con,
		stack: &lexeme.Container{},
		out:   &lexeme.Container{},
	}

	shunt(shy)
	return shy.out.AsContainer()
}

func shunt(shy *shuntingYard) {

	for shy.more() {

		switch {
		case shy.inQueue(lexeme.SPELL):
			call(shy)

		default:
			panic("Unexpected token: " + shy.queue.Head().String())
		}

		if shy.queueTok().IsTerminator() {
			shy.output()
			continue
		}

		panic("Unexpected token: " + shy.queue.Head().String())
	}
}

func call(shy *shuntingYard) {

	shy.push() // @Spell
	shy.push() // "("
	shy.emit(*shy.stack.Top(), lexeme.CALLABLE)

	if !shy.inQueue(lexeme.RIGHT_PAREN) {
		arguments(shy)
	}

	shy.discard() // ")"
	shy.eject()   // "("
	shy.pop()     // @Spell
}

func arguments(shy *shuntingYard) {

	for !shy.inQueue(lexeme.RIGHT_PAREN) {

		if shy.inQueue(lexeme.SEPARATOR) {
			//shy.output()
			shy.discard()
		}

		argument(shy)
	}
}

func argument(shy *shuntingYard) {

	for !shy.inQueue(lexeme.SEPARATOR) && !shy.inQueue(lexeme.RIGHT_PAREN) {

		switch {
		case shy.queueTok().IsTerm():
			shy.output()

		default:
			panic("Unexpected token: " + shy.queue.Head().String())
		}
	}
}
