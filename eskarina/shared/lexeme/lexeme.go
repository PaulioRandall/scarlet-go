package lexeme

import (
	"fmt"
)

type Snippet interface {
	At() (line, start, end int)
}

type Lexeme struct {
	Tok  Token
	Raw  string
	Line int
	Col  int
	next *Lexeme
	prev *Lexeme
}

func (lex Lexeme) IsSingle2() bool {
	return lex.next == nil && lex.prev == nil
}

func (lex Lexeme) At() (line, start, end int) {
	return lex.Line, lex.Col, len(lex.Raw)
}

func (lex Lexeme) Next2() *Lexeme {
	return lex.next
}

func (lex Lexeme) Prev2() *Lexeme {
	return lex.prev
}

func (lex *Lexeme) prepend(other *Lexeme) {

	if lex.prev != nil {
		lex.prev.next = other
		other.prev = lex.prev
	}

	other.next = lex
	lex.prev = other
}

func (lex *Lexeme) append(other *Lexeme) {

	if lex.next != nil {
		lex.next.prev = other
		other.next = lex.next
	}

	lex.next = other
	other.prev = lex
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

func (lex *Lexeme) splitNext() (*Lexeme, *Lexeme) {

	next := lex.next

	if next != nil {
		next.prev = nil
		lex.next = nil
	}

	return lex, next
}

func (lex *Lexeme) splitPrev() (*Lexeme, *Lexeme) {

	prev := lex.prev

	if prev != nil {
		prev.next = nil
		lex.prev = nil
	}

	return lex, prev
}

func (lex Lexeme) String() string {
	return fmt.Sprintf("%d:%d %s %q",
		lex.Line,
		lex.Col,
		lex.Tok.String(),
		lex.Raw,
	)
}
