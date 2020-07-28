package lexeme

import (
	"fmt"
)

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
	Tok  Token
	Raw  string
	Line int
	Col  int
	Next *Lexeme
	Prev *Lexeme
	next *Lexeme
	prev *Lexeme
}

func (lex Lexeme) IsSingle() bool {
	return lex.Next == nil && lex.Prev == nil
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

// @Deprecated
func (lex *Lexeme) Prepend(other *Lexeme) {

	if lex.Prev != nil {
		lex.Prev.Next = other
		other.Prev = lex.Prev
	}

	other.Next = lex
	lex.Prev = other
}

func (lex *Lexeme) prepend(other *Lexeme) {

	if lex.prev != nil {
		lex.prev.next = other
		other.prev = lex.prev
	}

	other.next = lex
	lex.prev = other
}

// @Deprecated
func (lex *Lexeme) Append(other *Lexeme) {

	if lex.Next != nil {
		lex.Next.Prepend(other)
		return
	}

	lex.Next = other
	other.Prev = lex
}

func (lex *Lexeme) append(other *Lexeme) {

	if lex.next != nil {
		lex.next.append(other)
		return
	}

	lex.next = other
	other.prev = lex
}

// @Deprecated
func (lex *Lexeme) Remove() {

	if lex.Next != nil {
		lex.Next.Prev = lex.Prev
	}

	if lex.Prev != nil {
		lex.Prev.Next = lex.Next
	}

	lex.Next, lex.Prev = nil, nil
}

func (lex *Lexeme) remove() {

	if lex.next != nil {
		lex.next.prev = lex.prev
	}

	if lex.prev != nil {
		lex.prev.next = lex.next
	}

	lex.next, lex.prev = nil, nil
}

func (lex *Lexeme) SplitBelow() {

	if lex.Next == nil {
		return
	}

	lex.Next.Prev = nil
	lex.Next = nil
}

func (lex Lexeme) String() string {
	return fmt.Sprintf("%d:%d %s %q",
		lex.Line,
		lex.Col,
		lex.Tok.String(),
		lex.Raw,
	)
}
