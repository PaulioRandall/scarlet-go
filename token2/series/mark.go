package series

import (
	"github.com/PaulioRandall/scarlet-go/token2/lexeme"
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

// Next moves the iterator mark onto the next item and returns it. A panic will
// ensue if the end of the iterator has already been reached so Series.More
// should be called before hand.
func (m *mark) Next() lexeme.Lexeme {
	if m.next == nil {
		panic("Can't move beyond the end of the series")
	}
	m.jumpTo(m.next)
	return m.curr.data
}

// Get returns the item at the current iterator mark or the Lexeme zero value if
// there is no item at the mark, Iie. before the first item, after the last
// item, and immediately after an item has been removed.
func (m mark) Get() lexeme.Lexeme {
	if m.curr == nil {
		return lexeme.Lexeme{}
	}
	return m.curr.data
}

// Prev moves the iterator mark onto the previous item and returns it. A panic
// will ensue if the start of the iterator has already been reached.
func (m *mark) Prev() lexeme.Lexeme {
	if m.prev == nil {
		panic("Can't move beyond the start of the series")
	}
	m.jumpTo(m.prev)
	return m.curr.data
}

// LookAhead returns the Lexeme next in the iteration without incrementing the
// iterator mark. An empty Lexeme is returned if there is no item ahead.
func (m mark) LookAhead() lexeme.Lexeme {
	if m.next == nil {
		return lexeme.Lexeme{}
	}
	return m.next.data
}

// Lookback returns the Lexeme previous in the iteration without decrementing
// the iterator mark. An empty Lexeme is returned if there is no item behind.
func (m mark) LookBack() lexeme.Lexeme {
	if m.prev == nil {
		return lexeme.Lexeme{}
	}
	return m.prev.data
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
