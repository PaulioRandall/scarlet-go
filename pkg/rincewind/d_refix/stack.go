package refix

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/rincewind/token"
)

type stack struct {
	top  *node
	size int
}

type node struct {
	tk   Token
	next *node
}

func (s *stack) empty() bool {
	return s.size == 0
}

func (s *stack) push(tk Token) {

	s.top = &node{
		tk:   tk,
		next: s.top,
	}

	s.size++
}

func (s *stack) peek() Token {

	if s.size == 0 {
		failNow("Nothing to peek")
	}

	return s.top.tk
}

func (s *stack) pop() Token {

	if s.size == 0 {
		failNow("Nothing to pop")
	}

	tk := s.top.tk
	s.top = s.top.next
	s.size--

	return tk
}
