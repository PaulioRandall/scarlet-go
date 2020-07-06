package compile

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/rincewind/instru"
	. "github.com/PaulioRandall/scarlet-go/pkg/rincewind/token"
)

type instruction struct {
	code   Code
	data   interface{}
	opener Token
	closer Token
}

func (ins instruction) Code() Code {
	return ins.code
}

func (ins instruction) Data() interface{} {
	return ins.data
}

func (ins instruction) Begin() (int, int) {
	return ins.opener.Begin()
}

func (ins instruction) End() (int, int) {
	return ins.closer.End()
}

type stack struct {
	top  *stackNode
	size int
}

type stackNode struct {
	in   instruction
	next *stackNode
}

func (s *stack) empty() bool {
	return s.size == 0
}

func (s *stack) push(in instruction) {

	s.top = &stackNode{
		in:   in,
		next: s.top,
	}

	s.size++
}

func (s *stack) peek() instruction {

	if s.size == 0 {
		failNow("Nothing to peek")
	}

	return s.top.in
}

func (s *stack) pop() instruction {

	if s.size == 0 {
		failNow("Nothing to pop")
	}

	in := s.top.in
	s.top = s.top.next
	s.size--

	return in
}
