package series

import (
	"github.com/PaulioRandall/scarlet-go/token/lexeme"
)

type mark struct {
	prev *node
	curr *node
	next *node
}

// More returns true if there are more Lexemes to iterate.
func (m mark) More() bool {
	return m.next != nil
}

func (m *mark) jumpTo(n *node) {
	m.prev = n.prev
	m.curr = n
	m.next = n.next
}

func (m *mark) jumpToStart(head *node) {
	m.prev = nil
	m.curr = nil
	m.next = head
}

func (m *mark) jumpToEnd(tail *node) {
	m.prev = tail
	m.curr = nil
	m.next = nil
}

func (m *mark) nextLex() lexeme.Lexeme {
	if m.next == nil {
		panic("Can't move beyond the end of the series")
	}
	m.jumpTo(m.next)
	return m.curr.data
}

func (m *mark) currLex() lexeme.Lexeme {
	if m.curr == nil {
		return lexeme.Lexeme{}
	}
	return m.curr.data
}

func (m *mark) prevLex() lexeme.Lexeme {
	if m.prev == nil {
		panic("Can't move beyond the start of the series")
	}
	m.jumpTo(m.prev)
	return m.curr.data
}

func (m *mark) insertAfter(n *node) {

	if m.curr == nil {
		panic("Current node missing, can't insert after it")
	}

	var next *node
	if m.curr != nil {
		next = m.curr.next
		unlink(m.curr, next)
	}

	chain(m.curr, n, next)
	m.jumpTo(m.curr)
}

func (m *mark) insertBefore(n *node) {

	if m.curr == nil {
		panic("Current node missing, can't insert before it")
	}

	var prev *node
	if m.curr != nil {
		prev = m.curr.prev
		unlink(prev, m.curr)
	}

	chain(prev, n, m.curr)
	m.jumpTo(m.curr)
}
