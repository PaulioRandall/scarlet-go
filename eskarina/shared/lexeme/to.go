package lexeme

type box interface {
	vacate() *Lexeme
}

type To struct {
	b box
}

func NewTo(b box) *To {
	return &To{
		b: b,
	}
}

func (to *To) Container() *Container2 {

	if c, ok := to.b.(*Container2); ok {
		return c
	}

	head := to.b.vacate()
	c := newContainer(head)
	to.b = nil
	return c
}

func (to *To) Queue() Queue2 {
	return to.Container()
}

func (to *To) Stack() Stack2 {
	return to.Container()
}

func (to *To) Itinerant() *Itinerant2 {

	if it, ok := to.b.(*Itinerant2); ok {
		return it
	}

	head := to.b.vacate()
	it := newItinerant(head)
	to.b = nil
	return it
}

func (to *To) Iterator() Iterator2 {
	return to.Itinerant()
}

func (to *To) Window() Window2 {
	return to.Itinerant()
}
