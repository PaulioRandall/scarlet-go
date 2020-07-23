package shunter

import (
	"github.com/PaulioRandall/scarlet-go/pkg/eskarina/lexeme"
	"github.com/PaulioRandall/scarlet-go/pkg/eskarina/prop"
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

func (shy *shuntingYard) inQueue(props ...prop.Prop) bool {
	return shy.queue.Head().Has(props...)
}

func (shy *shuntingYard) inStack(props ...prop.Prop) bool {

	if shy.stack.More() {
		return shy.stack.Top().Has(props...)
	}

	return false
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

func (shy *shuntingYard) emit(ref lexeme.Lexeme, props ...prop.Prop) {

	lex := &lexeme.Lexeme{
		Props: props,
		Line:  ref.Line,
		Col:   ref.Col,
	}

	shy.out.Put(lex)
}
