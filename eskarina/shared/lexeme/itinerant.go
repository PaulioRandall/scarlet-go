package lexeme

import (
	"strings"
)

type Range interface {
	ToContainer() *Container
	HasPrev() bool
	HasNext() bool
}

type Iterator interface {
	Range
	Prev() bool
	Next() bool
	Curr() *Lexeme
	Remove() *Lexeme
}

type Window interface {
	Iterator
	Behind() *Lexeme
	Ahead() *Lexeme
}

type Itinerant struct {
	behind *Lexeme
	curr   *Lexeme
	ahead  *Lexeme
}

func NewItinerant(head *Lexeme) *Itinerant {
	return &Itinerant{
		ahead: head,
	}
}

func (it *Itinerant) vacate() *Lexeme {

	var head *Lexeme

	switch {
	case it.behind != nil:
		for lex := it.behind; lex != nil; lex = lex.prev {
			head = lex
		}

	case it.curr != nil:
		head = it.curr

	default:
		head = it.ahead
	}

	it.curr, it.behind, it.ahead = nil, nil, nil
	return head
}

func (it *Itinerant) ToContainer() *Container {
	return NewContainer(it.vacate())
}

func (it *Itinerant) HasPrev() bool {
	return it.behind != nil
}

func (it *Itinerant) HasNext() bool {
	return it.ahead != nil
}

func (it *Itinerant) Prev() bool {

	if it.curr == nil && it.behind == nil {
		return false
	}

	if it.behind == nil {
		it.ahead, it.curr = it.curr, nil
		return false
	}

	it.curr = it.behind
	it.ahead = it.curr.next
	it.behind = it.curr.prev
	return true
}

func (it *Itinerant) Next() bool {

	if it.curr == nil && it.ahead == nil {
		return false
	}

	if it.ahead == nil {
		it.behind, it.curr = it.curr, nil
		return false
	}

	it.curr = it.ahead
	it.behind = it.curr.prev
	it.ahead = it.curr.next

	return true
}

func (it *Itinerant) Curr() *Lexeme {
	return it.curr
}

func (it *Itinerant) Behind() *Lexeme {
	return it.behind
}

func (it *Itinerant) Ahead() *Lexeme {
	return it.ahead
}

func (it *Itinerant) Remove() *Lexeme {

	if it.curr == nil {
		return nil
	}

	r := it.curr
	it.curr = nil
	r.remove()

	return r
}

func (it *Itinerant) String() string {

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

	write("Behind: ", it.behind)
	write("Curr  : ", it.curr)
	write("Ahead : ", it.ahead)

	return sb.String()
}
