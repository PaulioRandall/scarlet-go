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

func (lex Lexeme) IsSingle() bool {
	return lex.next == nil && lex.prev == nil
}

func (lex Lexeme) At() (line, start, end int) {
	return lex.Line, lex.Col, len(lex.Raw)
}

func (lex Lexeme) Next() *Lexeme {
	return lex.next
}

func (lex Lexeme) Prev() *Lexeme {
	return lex.prev
}

func (lex Lexeme) String() string {
	return fmt.Sprintf("%d:%d %s %q",
		lex.Line,
		lex.Col,
		lex.Tok.String(),
		lex.Raw,
	)
}

func prepend(base, lex *Lexeme) {

	if base.prev != nil {
		base.prev.next = lex
		lex.prev = base.prev
	}

	lex.next = base
	base.prev = lex
}

func append(base, lex *Lexeme) {

	if base.next != nil {
		base.next.prev = lex
		lex.next = base.next
	}

	base.next = lex
	lex.prev = base
}

func remove(lex *Lexeme) {

	if lex.next != nil {
		lex.next.prev = lex.prev
	}

	if lex.prev != nil {
		lex.prev.next = lex.next
	}

	lex.next, lex.prev = nil, nil
}

func splitNext(lex *Lexeme) (*Lexeme, *Lexeme) {

	next := lex.next

	if next != nil {
		next.prev = nil
		lex.next = nil
	}

	return lex, next
}

func splitPrev(lex *Lexeme) (*Lexeme, *Lexeme) {

	prev := lex.prev

	if prev != nil {
		prev.next = nil
		lex.prev = nil
	}

	return lex, prev
}
