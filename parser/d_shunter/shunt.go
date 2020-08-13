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

	statement(shy)
	return shy.out.AsContainer()
}

func statement(shy *shuntingYard) {

	for shy.more() {

		switch {
		case shy.inQueue(lexeme.SPELL):
			call(shy)

		case shy.inQueue(lexeme.IDENTIFIER):
			shy.push() // First ID

			if shy.inQueue(lexeme.SEPARATOR) || shy.inQueue(lexeme.ASSIGNMENT) {
				assignment(shy)
				break
			}

			shy.pop()
			fallthrough

		case shy.queueTok().IsTerm():
			expression(shy)

		default:
			panic("Unexpected token: " + shy.queue.Head().String())
		}

		if !shy.queueTok().IsTerminator() {
			panic("Unexpected token: " + shy.queue.Head().String())
		}

		shy.output()
	}
}

func expressions(shy *shuntingYard) {

	for first := true; first || shy.inQueue(lexeme.SEPARATOR); first = false {

		if !first {
			//shy.output()
			shy.discard()
		}

		expression(shy)
	}
}

func expression(shy *shuntingYard) {

	mark := shy.stackSize()

	for !shy.inQueue(lexeme.SEPARATOR) &&
		!shy.inQueue(lexeme.RIGHT_PAREN) &&
		!shy.queueTok().IsTerminator() {

		switch {
		case shy.queueTok().IsTerm():
			shy.output()

		case shy.queueTok().IsOperator():
			for shy.stackTok().Precedence() >= shy.queueTok().Precedence() {
				shy.pop()
			}

			shy.push()

		// case:
		// TODO: Incorporate spells calls and assignments once returns from spell
		// TODO: calls have been implemented.

		default:
			panic("Unexpected token: " + shy.queue.Head().String())
		}
	}

	for mark < shy.stackSize() {
		shy.pop()
	}
}

func call(shy *shuntingYard) {

	shy.push() // @Spell
	shy.push() // "("
	shy.emit(*shy.stack.Top(), lexeme.CALLABLE)

	for !shy.inQueue(lexeme.RIGHT_PAREN) {
		expressions(shy)
	}

	shy.discard() // ")"
	shy.eject()   // "("
	shy.pop()     // @Spell
}

func assignment(shy *shuntingYard) {

	count := 1
	for shy.inQueue(lexeme.SEPARATOR) {
		shy.discard()
		shy.push() // Other IDs
		count++
	}

	shy.push() // :=
	shy.emit(*shy.stack.Top(), lexeme.ASSIGNMENT)

	expressions(shy)
	shy.pop() // :=

	for ; count > 0; count-- {
		shy.pop() // IDs
	}
}
