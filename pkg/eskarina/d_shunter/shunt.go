package shunter

import (
	"github.com/PaulioRandall/scarlet-go/pkg/eskarina/lexeme"
	"github.com/PaulioRandall/scarlet-go/pkg/eskarina/prop"
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
		case shy.inQueue(prop.PR_TERM):
			shy.output()

		case shy.inQueue(prop.PR_SEPARATOR):
			if shy.inStack(prop.PR_PARENTHESIS, prop.PR_OPENER) {
				shy.output()
				break
			}

			shy.pop()

		case shy.inQueue(prop.PR_SPELL):
			shy.output()
			shy.push()
			shy.emit(*shy.stack.Top(), prop.PR_CALLABLE)

		case shy.inQueue(prop.PR_PARENTHESIS, prop.PR_CLOSER):
			if shy.inStack(prop.PR_PARENTHESIS, prop.PR_OPENER) {
				shy.discard()
				shy.eject()
			}

			shy.pop() // Currently this will be a spell, but this will change

		case shy.inQueue(prop.PR_TERMINATOR):
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
