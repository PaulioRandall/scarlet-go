package lexeme

import (
	"strings"
)

type Iterator2 struct {
	con    *Container
	before *Lexeme
	curr   *Lexeme
	after  *Lexeme
}

func (it *Iterator2) EOF() bool {
	return it.curr == nil && it.after == nil
}

func (it *Iterator2) HasPrev() bool {
	return it.before != nil
}

func (it *Iterator2) HasNext() bool {
	return it.after != nil
}

func (it *Iterator2) Prev() bool {

	if it.curr == nil && it.before == nil {
		return false
	}

	if it.before == nil {
		it.after, it.curr = it.curr, nil
		return false
	}

	it.curr = it.before
	it.refresh()
	return true
}

func (it *Iterator2) Next() bool {

	if it.curr == nil && it.after == nil {
		return false
	}

	if it.after == nil {
		it.before, it.curr = it.curr, nil
		return false
	}

	it.curr = it.after
	it.refresh()
	return true
}

func (it *Iterator2) Curr() *Lexeme {
	return it.curr
}

func (it *Iterator2) Before() *Lexeme {
	return it.before
}

func (it *Iterator2) After() *Lexeme {
	return it.after
}

func (it *Iterator2) Remove() *Lexeme {

	if it.curr == nil {
		return nil
	}

	r := it.curr
	it.curr = nil
	it.con.remove(r)

	return r
}

func (it *Iterator2) Prepend(lex *Lexeme) {

	if it.curr == nil {
		panic("Can't prepend to nil, curr is nil")
	}

	it.con.insertBefore(it.curr, lex)
	it.refresh()
}

func (it *Iterator2) Append(lex *Lexeme) {

	if it.curr == nil {
		panic("Can't append to nil, curr is nil")
	}

	it.con.insertAfter(it.curr, lex)
	it.refresh()
}

func (it *Iterator2) JumpToPrev(f func(Iterator2) bool) bool {
	for it.Prev() && !f(*it) {
	}

	return !it.EOF()
}

func (it *Iterator2) JumpToNext(f func(Iterator2) bool) bool {
	for it.Next() && !f(*it) {
	}

	return !it.EOF()
}

func (it *Iterator2) Restart() {
	it.before = nil
	it.curr = nil
	it.after = it.con.head
}

func (it *Iterator2) String() string {

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

func (it *Iterator2) refresh() {
	if it.curr != nil {
		it.before = it.curr.prev
		it.after = it.curr.next
	}
}
