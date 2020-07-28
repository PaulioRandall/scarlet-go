package lexeme

import (
	"strings"
)

type Range2 interface {
	To() *To
	HasPrev() bool
	HasNext() bool
}

type Iterator2 interface {
	Range2
	Prev() bool
	Next() bool
	Curr() *Lexeme
	Remove() *Lexeme
}

type Window2 interface {
	Iterator2
	Prior() *Lexeme
	Ahead() *Lexeme
}

type Itinerant2 struct {
	prior *Lexeme
	curr  *Lexeme
	ahead *Lexeme
}

func newItinerant(head *Lexeme) *Itinerant2 {
	return &Itinerant2{
		ahead: head,
	}
}

func (it *Itinerant2) vacate() *Lexeme {

	var head *Lexeme

	switch {
	case it.prior != nil:
		for lex := it.prior; lex != nil; lex = lex.prev {
			head = lex
		}

	case it.curr != nil:
		head = it.curr

	default:
		head = it.ahead
	}

	it.curr, it.prior, it.ahead = nil, nil, nil
	return head
}

func (it *Itinerant2) To() *To {
	return &To{
		b: it,
	}
}

func (it *Itinerant2) HasPrev() bool {
	return it.prior != nil
}

func (it *Itinerant2) HasNext() bool {
	return it.ahead != nil
}

func (it *Itinerant2) Prev() bool {

	if it.curr == nil && it.prior == nil {
		return false
	}

	if it.prior == nil {
		it.ahead, it.curr = it.curr, nil
		return false
	}

	it.curr = it.prior
	it.ahead = it.curr.next
	it.prior = it.curr.prev
	return true
}

func (it *Itinerant2) Next() bool {

	if it.curr == nil && it.ahead == nil {
		return false
	}

	if it.ahead == nil {
		it.prior, it.curr = it.curr, nil
		return false
	}

	it.curr = it.ahead
	it.prior = it.curr.prev
	it.ahead = it.curr.next

	return true
}

func (it *Itinerant2) Curr() *Lexeme {
	return it.curr
}

func (it *Itinerant2) Prior() *Lexeme {
	return it.prior
}

func (it *Itinerant2) Ahead() *Lexeme {
	return it.ahead
}

func (it *Itinerant2) Remove() *Lexeme {

	if it.curr == nil {
		return nil
	}

	r := it.curr
	it.curr = nil
	r.remove()

	return r
}

func (it *Itinerant2) String() string {

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

	write("Prior: ", it.prior)
	write("Curr : ", it.curr)
	write("Ahead: ", it.ahead)

	return sb.String()
}
