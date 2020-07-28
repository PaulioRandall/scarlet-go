package lexeme

type Range2 interface {
	HasPrev() bool
	HasNext() bool
	ToContainer() *Container2
	AsItinerant() *Itinerant2
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
	Behind() *Lexeme
	Ahead() *Lexeme
}

type Itinerant2 struct {
	behind *Lexeme
	curr   *Lexeme
	ahead  *Lexeme
}

func toItinerant(c *Container2) *Itinerant2 {

	it := &Itinerant2{
		ahead: c.head,
	}

	c.head, c.tail, c.size = nil, nil, 0
	return it
}

func (it *Itinerant2) HasPrev() bool {
	return it.behind != nil
}

func (it *Itinerant2) HasNext() bool {
	return it.ahead != nil
}

func (it *Itinerant2) ToContainer() *Container2 {

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
	return toContainer(head)
}

func (it *Itinerant2) AsItinerant() *Itinerant2 {
	return it
}

func (it *Itinerant2) AsIterator() Iterator2 {
	return it
}

func (it *Itinerant2) AsWindow() Window2 {
	return it
}

func (it *Itinerant2) Prev() bool {

	if it.curr == nil && it.behind == nil {
		return false
	}

	if it.behind == nil {
		it.ahead, it.curr = nil, it.behind
		return false
	}

	it.curr = it.behind
	it.ahead = it.curr.next
	it.behind = it.curr.prev
	return true
}

func (it *Itinerant2) Next() bool {

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

func (it *Itinerant2) Curr() *Lexeme {
	return it.curr
}

func (it *Itinerant2) Behind() *Lexeme {
	return it.behind
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
