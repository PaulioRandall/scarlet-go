package container

import (
	//"strings"

	"github.com/PaulioRandall/scarlet-go/token"
)

type View interface {
	HasNext() bool
	HasPrev() bool
	Next() token.Lexeme
	Curr() token.Lexeme
	Prev() token.Lexeme
	PeekPrev() token.Lexeme
	PeekNext() token.Lexeme
	String() string
}

type Iterator struct {
	prev *node
	curr *node
	next *node
}

func (it *Iterator) HasNext() bool {
	return it.next != nil
}

func (it *Iterator) HasPrev() bool {
	return it.prev != nil
}

func (it *Iterator) Next() token.Lexeme {

	if it.next == nil {
		panic("End of lexeme iterator breached")
	}

	it.prev = it.curr
	it.curr = it.next

	if it.curr == nil {
		it.next = nil
	} else {
		it.next = it.next.next
	}

	return it.curr.data
}

func (it *Iterator) Curr() token.Lexeme {
	return it.curr.data
}

func (it *Iterator) Prev() token.Lexeme {

	if it.prev == nil {
		panic("Start of lexeme iterator breached")
	}

	it.next = it.curr
	it.curr = it.prev

	if it.prev == nil {
		it.prev = nil
	} else {
		it.prev = it.prev.prev
	}

	return it.curr.data
}

func (it *Iterator) PeekNext() token.Lexeme {
	if it.next == nil {
		return token.Lexeme{}
	}
	return it.curr.next.data
}

func (it *Iterator) PeekPrev() token.Lexeme {
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
	it.curr.remove()
	it.curr = nil

	return n.data
}

/*
func (it *Iterator) InsertBefore(l token.Lexeme) {

	n := &node{
		data: l,
	}

	if it.curr != nil {
		unlink(it.prev, it.curr)
	}

	link(it.prev, n)
	link(n, it.curr)
	refresh(n)
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