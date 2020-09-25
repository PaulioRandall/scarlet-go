package container

import (
	"github.com/PaulioRandall/scarlet-go/token"
)

type View interface {
	HasNext() bool
	HasPrev() bool
	Next() token.Lexeme
	Curr() token.Lexeme
	Prev() token.Lexeme
	LookAhead() token.Lexeme
	LookBehind() token.Lexeme
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

func (it *Iterator) HasNext() bool {
	return it.next != nil
}

func (it *Iterator) HasPrev() bool {
	return it.prev != nil
}

func (it *Iterator) Next() token.Lexeme {

	if it.next == nil {
		panic("Can't move beyond the end of a lexeme iterator")
	}

	it.jumpTo(it.next)
	return it.curr.data
}

func (it *Iterator) Curr() token.Lexeme {
	return it.curr.data
}

func (it *Iterator) Prev() token.Lexeme {

	if it.prev == nil {
		panic("Can't move beyond the start of a lexeme iterator")
	}

	it.jumpTo(it.prev)
	return it.curr.data
}

func (it *Iterator) LookAhead() token.Lexeme {
	if it.next == nil {
		return token.Lexeme{}
	}
	return it.curr.next.data
}

func (it *Iterator) LookBehind() token.Lexeme {
	if it.prev == nil {
		return token.Lexeme{}
	}
	return it.curr.prev.data
}

func (it *Iterator) Remove() token.Lexeme {

	if it.curr == nil {
		return token.Lexeme{}
	}

	n := it.curr
	it.curr = nil
	it.con.remove(n)

	return n.data
}

func (it *Iterator) InsertBefore(l token.Lexeme) {

	n := &node{
		data: l,
	}

	it.con.insertBefore(it.curr, n)
	it.jumpTo(it.curr)
}

/*
func (it *Iterator) Append(lex *Lexeme) {

	if it.curr == nil {
		panic("Can't append to nil, curr is nil")
	}

	it.con.insertAfter(it.curr, lex)
	it.refresh()
}

func (it *Iterator) JumpToPrev(f func(View) bool) bool {
	for it.Prev() && !f(it) {
	}

	return !it.SOF()
}

func (it *Iterator) JumpToNext(f func(View) bool) bool {
	for it.Next() && !f(it) {
	}

	return !it.EOF()
}

func (it *Iterator) Restart() {
	it.before = nil
	it.curr = nil
	it.after = it.con.head
}

func (it *Iterator) String() string {

	sb := strings.Builder{}
	write := func(pre string, lex *Lexeme) {
		sb.WriteString(pre)

		if lex == nil {
			sb.WriteString("---")
		} else {
			sb.WriteString(lex.String())
		}

		sb.WriteRune('\n')
	}

	write("Behind: ", it.before)
	write("Curr  : ", it.curr)
	write("After : ", it.after)

	return sb.String()
}

func (it *Iterator) refresh() {
	if it.curr != nil {
		it.before = it.curr.prev
		it.after = it.curr.next
	}
}
*/

func (it *Iterator) jumpTo(n *node) {
	it.prev = n.prev
	it.curr = n
	it.next = n.next
}
