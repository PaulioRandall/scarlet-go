package shunter

import (
	"github.com/PaulioRandall/scarlet-go/eskarina/shared/lexeme"
)

type shuntingYard struct {
	queue lexeme.Queue
	stack lexeme.Stack
	out   lexeme.Queue
}

func (shy *shuntingYard) empty() bool {
	return shy.queue.Empty() && shy.stack.Empty()
}

func (shy *shuntingYard) more() bool {
	return shy.queue.More() || shy.stack.More()
}

func (shy *shuntingYard) inQueue(tk lexeme.Token) bool {
	return shy.queue.Head().Tok == tk
}

func (shy *shuntingYard) inStack(tk lexeme.Token) bool {

	if shy.stack.More() {
		return shy.stack.Top().Tok == tk
	}

	return false
}

func (shy *shuntingYard) queueTok() lexeme.Token {

	if shy.queue.More() {
		return shy.queue.Head().Tok
	}

	return lexeme.UNDEFINED
}

func (shy *shuntingYard) stackTok() lexeme.Token {

	if shy.stack.More() {
		return shy.stack.Top().Tok
	}

	return lexeme.UNDEFINED
}

func (shy *shuntingYard) push() {
	shy.stack.Push(shy.queue.Take())
}

func (shy *shuntingYard) discard() {
	shy.queue.Take()
}

func (shy *shuntingYard) pop() {
	shy.out.Put(shy.stack.Pop())
}

func (shy *shuntingYard) eject() {
	shy.stack.Pop()
}

func (shy *shuntingYard) output() {
	shy.out.Put(shy.queue.Take())
}

func (shy *shuntingYard) emit(ref lexeme.Lexeme, tk lexeme.Token, props ...lexeme.Prop) {

	lex := &lexeme.Lexeme{
		Props: props,
		Tok:   tk,
		Line:  ref.Line,
		Col:   ref.Col,
	}

	shy.out.Put(lex)
}
