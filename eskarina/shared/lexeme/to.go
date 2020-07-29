package lexeme

type box interface {
	vacate() *Lexeme
}

type To struct {
	b box
}

func (to *To) Slice() []Lexeme {

	r := []Lexeme{}

	for lex := to.b.vacate(); lex != nil; lex = lex.next {
		lex.remove()
		r = append(r, *lex)
	}

	return r
}

func (to *To) SlicePtr() []*Lexeme {

	r := []*Lexeme{}

	for lex := to.b.vacate(); lex != nil; lex = lex.next {
		lex.remove()
		r = append(r, lex)
	}

	return r
}

func (to *To) Container() *Container {

	if c, ok := to.b.(*Container); ok {
		return c
	}

	head := to.b.vacate()
	c := NewContainer(head)
	to.b = nil
	return c
}

func (to *To) Queue() Queue {
	return to.Container()
}

func (to *To) Stack() Stack {
	return to.Container()
}

func (to *To) Itinerant() *Itinerant {

	if it, ok := to.b.(*Itinerant); ok {
		return it
	}

	head := to.b.vacate()
	it := NewItinerant(head)
	to.b = nil

	return it
}

func (to *To) Iterator() Iterator {
	return to.Itinerant()
}

func (to *To) Window() Window {
	return to.Itinerant()
}
