package lexeme

import (
	"strings"
)

type Range interface {
	AsIterator() *Iterator
	ToContainer() *Container
	HasPrev() bool
	HasNext() bool
	Prev() bool
	Next() bool
	Curr() *Lexeme
	Remove() *Lexeme
	Before() *Lexeme
	After() *Lexeme
	String() string
}

type Iterator struct {
	before *Lexeme
	curr   *Lexeme
	after  *Lexeme
}

func NewIterator(head *Lexeme) *Iterator {
	return &Iterator{
		after: head,
	}
}

func (it *Iterator) vacate() *Lexeme {

	var head *Lexeme

	switch {
	case it.before != nil:
		for lex := it.before; lex != nil; lex = lex.prev {
			head = lex
		}

	case it.curr != nil:
		head = it.curr

	default:
		head = it.after
	}

	it.curr, it.before, it.after = nil, nil, nil
	return head
}

func (it *Iterator) ToContainer() *Container {
	return NewContainer(it.vacate())
}

func (it *Iterator) AsIterator() *Iterator {
	return it
}

func (it *Iterator) HasPrev() bool {
	return it.before != nil
}

func (it *Iterator) HasNext() bool {
	return it.after != nil
}

func (it *Iterator) Prev() bool {

	if it.curr == nil && it.before == nil {
		return false
	}

	if it.before == nil {
		it.after, it.curr = it.curr, nil
		return false
	}

	it.curr = it.before
	it.after = it.curr.next
	it.before = it.curr.prev
	return true
}

func (it *Iterator) Next() bool {

	if it.curr == nil && it.after == nil {
		return false
	}

	if it.after == nil {
		it.before, it.curr = it.curr, nil
		return false
	}

	it.curr = it.after
	it.before = it.curr.prev
	it.after = it.curr.next

	return true
}

func (it *Iterator) Curr() *Lexeme {
	return it.curr
}

func (it *Iterator) Before() *Lexeme {
	return it.before
}

func (it *Iterator) After() *Lexeme {
	return it.after
}

func (it *Iterator) Remove() *Lexeme {

	if it.curr == nil {
		return nil
	}

	r := it.curr
	it.curr = nil
	remove(r)

	return r
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
