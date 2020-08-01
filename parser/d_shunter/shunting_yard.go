package shunter

import (
	"github.com/PaulioRandall/scarlet-go/shared/lexeme"
)

type Stack interface {
	Empty() bool
	More() bool
	Top() *lexeme.Lexeme
	Push(*lexeme.Lexeme)
	Pop() *lexeme.Lexeme
}

type Queue interface {
	AsContainer() *lexeme.Container
	Empty() bool
	More() bool
	Head() *lexeme.Lexeme
	Put(*lexeme.Lexeme)
	Take() *lexeme.Lexeme
}

type shuntingYard struct {
	queue Queue
	stack Stack
	out   Queue
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

func (shy *shuntingYard) emit(ref lexeme.Lexeme, tk lexeme.Token) {

	lex := &lexeme.Lexeme{
		Tok:  tk,
		Line: ref.Line,
		Col:  ref.Col,
	}

	shy.out.Put(lex)
}
