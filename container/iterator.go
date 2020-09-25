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
	curr *node
}

func (it *Iterator) HasNext() bool {
	return it.curr != nil && it.curr.next != nil
}

func (it *Iterator) HasPrev() bool {
	return it.curr != nil && it.curr.prev != nil
}

func (it *Iterator) Next() token.Lexeme {

	if it.curr == nil || it.curr.next == nil {
		return token.Lexeme{}
	}

	it.curr = it.curr.next
	return it.curr.data
}

func (it *Iterator) Curr() token.Lexeme {
	return it.curr.data
}

func (it *Iterator) Prev() token.Lexeme {

	if it.curr == nil || it.curr.prev == nil {
		return token.Lexeme{}
	}

	it.curr = it.curr.prev
	return it.curr.data
}

func (it *Iterator) PeekNext() token.Lexeme {
	if it.curr == nil || it.curr.next == nil {
		return token.Lexeme{}
	}
	return it.curr.next.data
}

func (it *Iterator) PeekPrev() token.Lexeme {

	if it.curr == nil || it.curr.prev == nil {
		return token.Lexeme{}
	}
	return it.curr.prev.data
}

/*
func (it *Iterator) Remove() *Lexeme {

	if it.curr == nil {
		return nil
	}

	r := it.curr
	it.curr = nil
	it.con.remove(r)

	return r
}

func (it *Iterator) Prepend(lex *Lexeme) {

	if it.curr == nil {
		panic("Can't prepend to nil, curr is nil")
	}

	it.con.insertBefore(it.curr, lex)
	it.refresh()
}

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
