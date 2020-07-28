package lexeme

import (
	"fmt"
)

type PropToken interface {
	Is(Prop) bool
	Has(...Prop) bool
	Any(...Prop) bool
}

type Snippet interface {
	At() (line, start, end int)
}

type Node interface {
	NextNode() *Lexeme
	PrevNode() *Lexeme
	ShiftUp()
	ShiftDown()
	Prepend(*Lexeme)
	Append(*Lexeme)
	Remove()
	String() string
}

type Lexeme struct {
	Props []Prop
	Tok   Token
	Raw   string
	Line  int
	Col   int
	Next  *Lexeme
	Prev  *Lexeme
}

func (lex Lexeme) Is(o Prop) bool {

	for _, p := range lex.Props {
		if p == o {
			return true
		}
	}

	return false
}

func (lex Lexeme) Has(others ...Prop) bool {

	for _, o := range others {
		if !lex.Is(o) {
			return false
		}
	}

	return true
}

func (lex Lexeme) Any(others ...Prop) bool {

	for _, o := range others {
		if lex.Is(o) {
			return true
		}
	}

	return false
}

func (lex Lexeme) At() (line, start, end int) {
	return lex.Line, lex.Col, len(lex.Raw)
}

func (lex Lexeme) NextNode() *Lexeme {
	return lex.Next
}

func (lex Lexeme) PrevNode() *Lexeme {
	return lex.Prev
}

func (lex *Lexeme) ShiftUp() {

	if lex.Prev == nil {
		return
	}

	prev := lex.Prev

	if prev.Prev != nil {
		prev.Prev.Next = lex
	}

	if lex.Next != nil {
		lex.Next.Prev = prev
	}

	lex.Prev, prev.Next = prev.Prev, lex.Next
	lex.Next, prev.Prev = prev, lex
}

func (lex *Lexeme) ShiftDown() {

	if lex.Next == nil {
		return
	}

	lex.Next.ShiftUp()
}

func (lex *Lexeme) Prepend(other *Lexeme) {

	if lex.Prev != nil {
		lex.Prev.Next = other
		other.Prev = lex.Prev
	}

	other.Next = lex
	lex.Prev = other
}

func (lex *Lexeme) Append(other *Lexeme) {

	if lex.Next != nil {
		lex.Next.Prepend(other)
		return
	}

	lex.Next = other
	other.Prev = lex
}

func (lex *Lexeme) Remove() {

	if lex.Next != nil {
		lex.Next.Prev = lex.Prev
	}

	if lex.Prev != nil {
		lex.Prev.Next = lex.Next
	}

	lex.Next, lex.Prev = nil, nil
}

func (lex Lexeme) String() string {
	return fmt.Sprintf("%d:%d %s %s %q",
		lex.Line,
		lex.Col,
		lex.Tok.String(),
		"["+JoinProps(",", lex.Props...)+"]",
		lex.Raw,
	)
}
