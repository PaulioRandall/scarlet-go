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
		case shy.inQueue(lexeme.IDENT):
			shy.push() // First ID

			if shy.inQueue(lexeme.DELIM) || shy.inQueue(lexeme.ASSIGN) {
				assignment(shy)
				break
			}

			shy.pop()
			fallthrough

		case shy.inQueue(lexeme.SPELL), shy.queueTok().IsTerm():
			expressions(shy)

		default:
			panic("Unexpected token: " + shy.queue.Head().String())
		}

		if !shy.queueTok().IsTerminator() {
			panic("Unexpected token: " + shy.queue.Head().String())
		}

		shy.output()
	}
}

func assignment(shy *shuntingYard) {

	// First has already been pushed
	mark := shy.stackSize() - 1

	for shy.inQueue(lexeme.DELIM) {
		shy.push()
		shy.push() // Other IDs
	}

	shy.push() // :=
	shy.emit(*shy.stack.Top(), lexeme.ASSIGN)

	expressions(shy)
	shy.pop() // :=

	for mark < shy.stackSize() {
		shy.pop() // IDs
	}
}

func expressions(shy *shuntingYard) {

	mark := shy.stackSize()

	for !shy.queueTok().IsTerminator() {

		switch {
		case shy.inQueue(lexeme.SPELL):
			shy.push()
			shy.emit(*shy.stack.Top(), lexeme.SPELL)

		case shy.inQueue(lexeme.L_PAREN):
			shy.push()

		case shy.queueTok().IsTerm():
			shy.output()

		case shy.queueTok().IsOperator():
			for !shy.inStack(lexeme.L_PAREN) &&
				shy.stackTok().Precedence() >= shy.queueTok().Precedence() {

				shy.pop()
			}

			shy.push()

		case shy.inQueue(lexeme.DELIM):
			for shy.stackTok().IsOperator() {
				shy.pop()
			}

			shy.output()

		case shy.inQueue(lexeme.R_PAREN):
			for !shy.inStack(lexeme.L_PAREN) {
				shy.pop()
			}

			shy.discard()
			shy.eject()

			if shy.inStack(lexeme.SPELL) {
				shy.pop()
			}

		default:
			goto FINISH
		}
	}

FINISH:
	for mark < shy.stackSize() {
		shy.pop()
	}
}
