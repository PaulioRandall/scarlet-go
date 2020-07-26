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
		case shy.inQueue(lexeme.PR_TERM):
			shy.output()

		case shy.inQueue(lexeme.PR_SEPARATOR):
			if shy.inStack(lexeme.PR_PARENTHESIS, lexeme.PR_OPENER) {
				shy.discard() // ","
				break
			}

			shy.pop()

		case shy.inQueue(lexeme.PR_SPELL):
			shy.push() // @Spell
			shy.push() // "("
			shy.emit(*shy.stack.Top(), lexeme.PR_CALLABLE)

		case shy.inQueue(lexeme.PR_PARENTHESIS, lexeme.PR_CLOSER):
			if shy.inStack(lexeme.PR_PARENTHESIS, lexeme.PR_OPENER) {
				shy.discard() // ")"
				shy.eject()   // "("
			}

			shy.pop() // @Spell

		case shy.inQueue(lexeme.PR_TERMINATOR):
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
