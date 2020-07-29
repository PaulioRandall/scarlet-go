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
	Prepend(*Lexeme)
	Append(*Lexeme)
	Split() *Container
	Before() *Lexeme
	After() *Lexeme
	Restart()
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

func (it *Iterator) ToContainer() *Container {
	it.Restart()
	head := it.after
	it.curr, it.before, it.after = nil, nil, nil
	return NewContainer(head)
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
	it.refresh()
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
	it.refresh()
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

func (it *Iterator) Prepend(lex *Lexeme) {

	if it.curr == nil {
		panic("Can't prepend to nil, curr is nil")
	}

	prepend(it.curr, lex)
	it.refresh()
}

func (it *Iterator) Append(lex *Lexeme) {

	if it.curr == nil {
		panic("Can't append to nil, curr is nil")
	}

	append(it.curr, lex)
	it.refresh()
}

func (it *Iterator) Split() *Container {

	if it.before == nil {
		return &Container{}
	}

	it.before.next = nil
	it.curr.prev = nil

	head := it.before
	it.refresh()

	for lex := head; lex != nil; lex = lex.prev {
		head = lex
	}

	return NewContainer(head)
}

func (it *Iterator) Restart() {

	var head *Lexeme

	switch {
	case it.before != nil:
		head = it.before
	case it.curr != nil:
		head = it.curr
	case it.after != nil:
		head = it.after
	default:
		return
	}

	for lex := head; lex != nil; lex = lex.prev {
		head = lex
	}

	it.before = nil
	it.curr = nil
	it.after = head
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
