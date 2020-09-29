package container

import (
	"strings"

	"github.com/PaulioRandall/scarlet-go/token/lexeme"
)

type View interface {
	More() bool
	IsFirst() bool
	Item() lexeme.Lexeme
	LookAhead() lexeme.Lexeme
	LookBack() lexeme.Lexeme
	String() string
}

type Iterator struct { // TODO: Should be private?
	con  *Container
	prev *node
	curr *node
	next *node
}

func (it *Iterator) Container() *Container {
	return it.con
}

func (it *Iterator) More() bool {
	return it.next != nil
}

func (it *Iterator) IsFirst() bool {
	return it.prev == nil
}

func (it *Iterator) Next() lexeme.Lexeme {

	if it.next == nil {
		panic("Can't move beyond the end of a lexeme iterator")
	}

	it.jumpTo(it.next)
	return it.curr.data
}

func (it *Iterator) Item() lexeme.Lexeme {
	if it.curr == nil {
		return lexeme.Lexeme{}
	}
	return it.curr.data
}

func (it *Iterator) Prev() lexeme.Lexeme {

	if it.prev == nil {
		panic("Can't move beyond the start of a lexeme iterator")
	}

	it.jumpTo(it.prev)
	return it.curr.data
}

func (it *Iterator) LookAhead() lexeme.Lexeme {
	if it.next == nil {
		return lexeme.Lexeme{}
	}
	return it.next.data
}

func (it *Iterator) LookBack() lexeme.Lexeme {
	if it.prev == nil {
		return lexeme.Lexeme{}
	}
	return it.prev.data
}

func (it *Iterator) Remove() lexeme.Lexeme {

	if it.curr == nil {
		return lexeme.Lexeme{}
	}

	n := it.curr
	it.curr = nil
	it.con.remove(n)

	return n.data
}

func (it *Iterator) InsertBefore(l lexeme.Lexeme) {

	n := &node{
		data: l,
	}

	if it.curr != nil {
		it.con.insertBefore(it.curr, n)
		it.jumpTo(it.curr)

	} else if it.prev != nil {
		it.con.insertAfter(it.prev, n)
		it.jumpToStart()

	} else {
		it.con.prepend(n)
	}
}

func (it *Iterator) InsertAfter(l lexeme.Lexeme) {

	n := &node{
		data: l,
	}

	if it.curr != nil {
		it.con.insertAfter(it.curr, n)
		it.jumpTo(it.curr)

	} else if it.next != nil {
		it.con.insertBefore(it.next, n)
		it.jumpToEnd()

	} else {
		it.con.append(n)
	}
}

func (it *Iterator) JumpToPrev(f func(View) bool) bool {

	for n := it.prev; n != nil; n = n.prev {
		it.jumpTo(n)
		if f(it) {
			return true
		}
	}

	it.jumpToStart()
	return false
}

func (it *Iterator) JumpToNext(f func(View) bool) bool {

	for n := it.next; n != nil; n = n.next {
		it.jumpTo(n)
		if f(it) {
			return true
		}
	}

	it.jumpToEnd()
	return false
}

func (it *Iterator) String() string {

	var sb strings.Builder
	write := func(pre string, n *node) {

		sb.WriteString(pre)
		if n == nil {
			sb.WriteString("---")
		} else {
			sb.WriteString(n.data.String())
		}

		sb.WriteRune('\n')
	}

	write("Prev: ", it.prev)
	write("Curr: ", it.curr)
	write("Next: ", it.next)

	return sb.String()
}

func (it *Iterator) jumpToStart() {
	it.prev = nil
	it.curr = nil
	it.next = it.con.head
}

func (it *Iterator) jumpToEnd() {
	it.prev = it.con.tail
	it.curr = nil
	it.next = nil
}

func (it *Iterator) jumpTo(n *node) {
	it.prev = n.prev
	it.curr = n
	it.next = n.next
}
