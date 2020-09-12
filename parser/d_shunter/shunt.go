package shunter

import (
	"github.com/PaulioRandall/scarlet-go/lexeme"
)

func ShuntAll(con *lexeme.Container) *lexeme.Container {

	shy := &shuntingYard{
		queue: con,
		stack: &lexeme.Container{},
		out:   &lexeme.Container{},
	}

	statements(shy)
	return shy.out.AsContainer()
}

func statements(shy *shuntingYard) {
	for shy.more() {
		statement(shy)

		if !shy.queueTok().IsTerminator() {
			panic("Unexpected token: " + shy.queue.Head().String())
		}

		shy.discard()
	}
}

func statement(shy *shuntingYard) {

	switch {
	case shy.queueTok().IsAssignee():
		shy.push() // First IDENT or VOID

		if shy.inQueue(lexeme.DELIM) || shy.inQueue(lexeme.ASSIGN) {
			assignment(shy)
			break
		}

		shy.pop()
		fallthrough

	case shy.inQueue(lexeme.SPELL),
		shy.inQueue(lexeme.L_PAREN),
		shy.queueTok().IsTerm():

		expressions(shy)

	case shy.inQueue(lexeme.L_SQUARE):
		guard(shy)

	case shy.inQueue(lexeme.LOOP):
		loop(shy)

	default:
		panic("Unexpected token: " + shy.queue.Head().String())
	}
}

func assignment(shy *shuntingYard) {

	// First has already been pushed
	mark := shy.stackSize() - 1

	for shy.inQueue(lexeme.DELIM) {
		shy.push()
		shy.push() // Other IDENTs or VOIDs
	}

	shy.push() // :=
	shy.emit(*shy.stack.Top(), lexeme.ASSIGN)

	expressions(shy)
	shy.pop() // :=

	for mark < shy.stackSize() {
		shy.pop() // IDs
	}
}

func guard(shy *shuntingYard) {
	expressions(shy)
	shy.emit(*shy.queue.Head(), lexeme.GUARD)
	block(shy)
}

func loop(shy *shuntingYard) {
	shy.emit(*shy.discard(), lexeme.LOOP)
	expressions(shy)
	block(shy)
}

func block(shy *shuntingYard) {

	shy.output() // {

	if shy.inQueue(lexeme.R_CURLY) {
		shy.output() // }
		return
	}

	for statement(shy); !shy.inQueue(lexeme.R_CURLY); statement(shy) {

		if !shy.queueTok().IsTerminator() {
			panic("Unexpected token: " + shy.queue.Head().String())
		}

		shy.discard() // terminator
	}

	shy.output() // }
}

func expressions(shy *shuntingYard) {

	mark := shy.stackSize()

	for !shy.queueTok().IsTerminator() {

		switch {
		case shy.inQueue(lexeme.SPELL):
			shy.push()
			shy.emit(*shy.stack.Top(), lexeme.SPELL)

		case shy.inQueue(lexeme.L_PAREN), shy.inQueue(lexeme.L_SQUARE):
			shy.push()

		case shy.queueTok().IsTerm():
			shy.output()

		case shy.inQueue(lexeme.VOID):
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
			braceClose(shy, lexeme.L_PAREN)

		case shy.inQueue(lexeme.R_SQUARE):
			braceClose(shy, lexeme.L_SQUARE)

		default:
			goto FINISH
		}
	}

FINISH:
	for mark < shy.stackSize() {
		shy.pop()
	}
}

func braceClose(shy *shuntingYard, opener lexeme.Token) {

	for !shy.inStack(opener) {
		shy.pop()
	}

	shy.discard()
	shy.eject()

	if shy.inStack(lexeme.SPELL) {
		shy.pop()
	}
}
