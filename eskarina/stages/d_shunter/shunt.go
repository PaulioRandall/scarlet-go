package shunter

import (
	"github.com/PaulioRandall/scarlet-go/eskarina/shared/lexeme"
)

func ShuntAll(head *lexeme.Lexeme) *lexeme.Lexeme {

	shy := &shuntingYard{
		queue: lexeme.NewContainer(head),
		stack: &lexeme.Container{},
		out:   &lexeme.Container{},
	}

	shunt(shy)
	return shy.out.Head()
}

func shunt(shy *shuntingYard) {

	for shy.more() {

		switch {
		case shy.queueTok().IsTerm():
			shy.output()

		case shy.inQueue(lexeme.SEPARATOR):
			if shy.inStack(lexeme.LEFT_PAREN) {
				shy.discard() // ","
				break
			}

			shy.pop()

		case shy.inQueue(lexeme.SPELL):
			shy.push() // @Spell
			shy.push() // "("
			shy.emit(*shy.stack.Top(), lexeme.CALLABLE, lexeme.PR_CALLABLE)

		case shy.inQueue(lexeme.RIGHT_PAREN):
			if shy.inStack(lexeme.LEFT_PAREN) {
				shy.discard() // ")"
				shy.eject()   // "("
			}

			shy.pop() // @Spell

		case shy.queueTok().IsTerminator():
			if shy.stack.Empty() {
				shy.output()
				break
			}

			fallthrough

		default:
			panic("Unexpected token: " + shy.queue.Head().String())
		}
	}
}
